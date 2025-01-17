package describer

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListAccounts(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	go func() {
		processAccounts(ctx, handler, cloudFlareChan, &wg)
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

func GetAccount(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	account, err := processAccount(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   account.ID,
		Name: account.Name,
		Description: model.AccountDescription{
			ID:       account.ID,
			Name:     account.Name,
			Type:     account.Type,
			Settings: account.Settings,
		},
	}
	return &value, nil
}

func processAccounts(ctx context.Context, handler *model.CloudFlareAPIHandler, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var accounts []cloudflare.Account
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		pageOpts := cloudflare.PaginationOptions{
			PerPage: perPage,
			Page:    page,
		}
		accounts, _, e = handler.Conn.Accounts(ctx, pageOpts)
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
	for _, account := range accounts {
		wg.Add(1)
		go func(account cloudflare.Account) {
			defer wg.Done()
			value := models.Resource{
				ID:   account.ID,
				Name: account.Name,
				Description: model.AccountDescription{
					ID:       account.ID,
					Name:     account.Name,
					Type:     account.Type,
					Settings: account.Settings,
				},
			}
			cloudFlareChan <- value
		}(account)
	}
}

func processAccount(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*cloudflare.Account, error) {
	var account cloudflare.Account
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		account, _, e = handler.Conn.Account(ctx, resourceID)
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
	return &account, nil
}
