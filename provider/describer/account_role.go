package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"sync"
)

func ListAccountRoles(ctx context.Context, handler *CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		processAccountRoles(ctx, handler, account, cloudFlareChan, &wg)
		wg.Wait()
		close(cloudFlareChan)
	}()
	var values []models.Resource
	for value := range cloudFlareChan {
		if stream != nil {
			if err = (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func GetAccountRole(ctx context.Context, handler *CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	accountRole, err := processAccountRole(ctx, handler, account, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   accountRole.ID,
		Name: accountRole.Name,
		Description: model.AccountRoleDescription{
			ID:          accountRole.ID,
			Name:        accountRole.Name,
			Description: accountRole.Description,
			Permissions: accountRole.Permissions,
			AccountID:   account.ID,
			Title:       accountRole.Name,
		},
	}
	return &value, nil
}

func processAccountRoles(ctx context.Context, handler *CloudFlareAPIHandler, account *cloudflare.Account, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var accountRoles []cloudflare.AccountRole
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		accountRoles, e = handler.Conn.AccountRoles(ctx, account.ID)
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
	for _, accountRole := range accountRoles {
		wg.Add(1)
		go func(accountRole cloudflare.AccountRole) {
			defer wg.Done()
			value := models.Resource{
				ID:   accountRole.ID,
				Name: accountRole.Name,
				Description: model.AccountRoleDescription{
					ID:          accountRole.ID,
					Name:        accountRole.Name,
					Description: accountRole.Description,
					Permissions: accountRole.Permissions,
					AccountID:   account.ID,
					Title:       accountRole.Name,
				},
			}
			cloudFlareChan <- value
		}(accountRole)
	}
}

func processAccountRole(ctx context.Context, handler *CloudFlareAPIHandler, account *cloudflare.Account, resourceID string) (*cloudflare.AccountRole, error) {
	var accountRole cloudflare.AccountRole
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		accountRole, e = handler.Conn.AccountRole(ctx, account.ID, resourceID)
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
	return &accountRole, nil
}
