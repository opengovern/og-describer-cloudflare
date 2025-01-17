package describers

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListAPITokens(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	go func() {
		processAPITokens(ctx, handler, cloudFlareChan, &wg)
		wg.Wait()
		close(cloudFlareChan)
	}()
	var values []models.Resource
	for value := range cloudFlareChan {
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func GetAPIToken(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	token, err := processAPIToken(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   token.ID,
		Name: token.Name,
		Description: model.ApiTokenDescription{
			ID:         token.ID,
			Name:       token.Name,
			Status:     token.Status,
			Condition:  token.Condition,
			ExpiresOn:  token.ExpiresOn,
			IssuedOn:   token.IssuedOn,
			ModifiedOn: token.ModifiedOn,
			NotBefore:  token.NotBefore,
			Policies:   token.Policies,
		},
	}
	return &value, nil
}

func processAPITokens(ctx context.Context, handler *model.CloudFlareAPIHandler, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var apiTokens []cloudflare.APIToken
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		apiTokens, e = handler.Conn.APITokens(ctx)
		if e != nil {
			var httpErr *cloudflare.APIRequestError
			if errors.As(e, &httpErr) {
				statusCode = &httpErr.StatusCode
			}
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, token := range apiTokens {
		wg.Add(1)
		go func(token cloudflare.APIToken) {
			defer wg.Done()
			value := models.Resource{
				ID:   token.ID,
				Name: token.Name,
				Description: model.ApiTokenDescription{
					ID:         token.ID,
					Name:       token.Name,
					Status:     token.Status,
					Condition:  token.Condition,
					ExpiresOn:  token.ExpiresOn,
					IssuedOn:   token.IssuedOn,
					ModifiedOn: token.ModifiedOn,
					NotBefore:  token.NotBefore,
					Policies:   token.Policies,
				},
			}
			cloudFlareChan <- value
		}(token)
	}
}

func processAPIToken(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*cloudflare.APIToken, error) {
	var apiToken cloudflare.APIToken
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		apiToken, e = handler.Conn.GetAPIToken(ctx, resourceID)
		if e != nil {
			var httpErr *cloudflare.APIRequestError
			if errors.As(e, &httpErr) {
				statusCode = &httpErr.StatusCode
			}
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return nil, err
	}
	return &apiToken, nil
}
