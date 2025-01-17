package describer

import (
	"context"
	"errors"

	"github.com/cloudflare/cloudflare-go"
)

const (
	perPage = 100
	page    = 1
)



func getAccount(ctx context.Context, handler *model.CloudFlareAPIHandler) (*cloudflare.Account, error) {
	var account cloudflare.Account
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		account, _, e = handler.Conn.Account(ctx, handler.AccountID)
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

func getApplications(ctx context.Context, handler *model.CloudFlareAPIHandler, accountID string) ([]cloudflare.AccessApplication, error) {
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
			pageAccessApps, pageData, e = handler.Conn.AccessApplications(ctx, accountID, opts)
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
		return nil, err
	}
	return accessApps, nil
}

func getZones(ctx context.Context, handler *model.CloudFlareAPIHandler) ([]cloudflare.Zone, error) {
	var zones []cloudflare.Zone
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		zones, e = handler.Conn.ListZones(ctx)
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
	return zones, nil
}

