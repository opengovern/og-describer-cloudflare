package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	"github.com/turbot/go-kit/helpers"
)

func ListAccessApplications(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, nil
	}
	opts := cloudflare.PaginationOptions{
		PerPage: perPage,
		Page:    page,
	}
	type ListPageResponse struct {
		Applications []cloudflare.AccessApplication
		resp         cloudflare.ResultInfo
	}
	listPage := func(ctx context.Context) (interface{}, error) {
		applications, resp, err := conn.AccessApplications(ctx, account.ID, opts)
		return ListPageResponse{
			Applications: applications,
			resp:         resp,
		}, err
	}
	var values []models.Resource
	for {
		listPageResponse, err := retry(
			ctx,
			func() (interface{}, error) {
				return listPage(ctx)
			},
			shouldRetryError,
		)
		if err != nil {
			var cloudFlareErr *cloudflare.APIRequestError
			if errors.As(err, &cloudFlareErr) {
				if helpers.StringSliceContains(cloudFlareErr.ErrorMessages(), "Access is not enabled. Visit the Access dashboard at https://dash.cloudflare.com/ and click the 'Enable Access' button.") {
					return nil, nil
				}
			}
			return nil, err
		}
		listResponse := listPageResponse.(ListPageResponse)
		apps := listResponse.Applications
		resp := listResponse.resp
		for _, app := range apps {
			value := models.Resource{
				ID:   app.ID,
				Name: app.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.AccessApplicationDescription{
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
				},
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		}
		if resp.Page >= resp.TotalPages {
			break
		}
		opts.Page = opts.Page + 1
	}
	return values, nil
}
