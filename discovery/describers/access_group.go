package describer

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListAccessGroups(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		processAccessGroups(ctx, handler, account, cloudFlareChan, &wg)
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

func GetAccessGroup(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	group, err := processAccessGroup(ctx, handler, account, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   group.ID,
		Name: group.Name,
		Description: model.AccessGroupDescription{
			ID:          group.ID,
			Name:        group.Name,
			AccountID:   account.ID,
			AccountName: account.Name,
			CreatedAt:   group.CreatedAt,
			UpdatedAt:   group.UpdatedAt,
			Exclude:     group.Exclude,
			Include:     group.Include,
			Require:     group.Require,
		},
	}
	return &value, nil
}

func processAccessGroups(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var accessGroups []cloudflare.AccessGroup
	var pageAccessGroups []cloudflare.AccessGroup
	var pageData cloudflare.ResultInfo
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		opts := cloudflare.PaginationOptions{
			PerPage: perPage,
			Page:    page,
		}
		for {
			pageAccessGroups, pageData, e = handler.Conn.AccessGroups(ctx, account.ID, opts)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			accessGroups = append(accessGroups, pageAccessGroups...)
			if pageData.Page >= pageData.TotalPages {
				break
			}
			opts.Page = opts.Page + 1
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, group := range accessGroups {
		wg.Add(1)
		go func(group cloudflare.AccessGroup) {
			defer wg.Done()
			value := models.Resource{
				ID:   group.ID,
				Name: group.Name,
				Description: model.AccessGroupDescription{
					ID:          group.ID,
					Name:        group.Name,
					AccountID:   account.ID,
					AccountName: account.Name,
					CreatedAt:   group.CreatedAt,
					UpdatedAt:   group.UpdatedAt,
					Exclude:     group.Exclude,
					Include:     group.Include,
					Require:     group.Require,
				},
			}
			cloudFlareChan <- value
		}(group)
	}
}

func processAccessGroup(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, resourceID string) (*cloudflare.AccessGroup, error) {
	var accessGroup cloudflare.AccessGroup
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		accessGroup, e = handler.Conn.AccessGroup(ctx, account.ID, resourceID)
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
	return &accessGroup, nil
}
