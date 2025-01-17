package describers

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListAccessApplications(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		processAccessApps(ctx, handler, account, cloudFlareChan, &wg)
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

func GetAccessApplication(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	app, err := processAccessApp(ctx, handler, account, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   app.ID,
		Name: app.Name,
		Description: model.AccessApplicationDescription{
			ID:                     app.ID,
			Name:                   app.Name,
			AccountID:              account.ID,
			AccountName:            account.Name,
			Domain:                 app.Domain,
			CreatedAt:              app.CreatedAt,
			Aud:                    app.AUD,
			AutoRedirectToIdentity: app.AutoRedirectToIdentity,
			CustomDenyMessage:      app.CustomDenyMessage,
			CustomDenyURL:          app.CustomDenyURL,
			EnableBindingCookie:    app.EnableBindingCookie,
			SessionDuration:        app.SessionDuration,
			UpdatedAt:              app.UpdatedAt,
			AllowedIDPs:            app.AllowedIdps,
			CORSHeaders:            app.CorsHeaders,
		},
	}
	return &value, nil
}

func processAccessApps(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var accessApps []cloudflare.AccessApplication
	var pageAccessApps []cloudflare.AccessApplication
	var pageData cloudflare.ResultInfo
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		opts := cloudflare.PaginationOptions{
			PerPage: perPage,
			Page:    page,
		}
		for {
			pageAccessApps, pageData, e = handler.Conn.AccessApplications(ctx, account.ID, opts)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			accessApps = append(accessApps, pageAccessApps...)
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
	for _, app := range accessApps {
		wg.Add(1)
		go func(app cloudflare.AccessApplication) {
			defer wg.Done()
			value := models.Resource{
				ID:   app.ID,
				Name: app.Name,
				Description: model.AccessApplicationDescription{
					ID:                     app.ID,
					Name:                   app.Name,
					AccountID:              account.ID,
					AccountName:            account.Name,
					Domain:                 app.Domain,
					CreatedAt:              app.CreatedAt,
					Aud:                    app.AUD,
					AutoRedirectToIdentity: app.AutoRedirectToIdentity,
					CustomDenyMessage:      app.CustomDenyMessage,
					CustomDenyURL:          app.CustomDenyURL,
					EnableBindingCookie:    app.EnableBindingCookie,
					SessionDuration:        app.SessionDuration,
					UpdatedAt:              app.UpdatedAt,
					AllowedIDPs:            app.AllowedIdps,
					CORSHeaders:            app.CorsHeaders,
				},
			}
			cloudFlareChan <- value
		}(app)
	}
}

func processAccessApp(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, resourceID string) (*cloudflare.AccessApplication, error) {
	var accessApp cloudflare.AccessApplication
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		accessApp, e = handler.Conn.AccessApplication(ctx, account.ID, resourceID)
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
	return &accessApp, nil
}
