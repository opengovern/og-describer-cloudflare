package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
	"github.com/turbot/go-kit/helpers"
)

func GetAllAccessGroups(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	accountValues, err := GetAccountAccessGroups(ctx, conn, stream, *account)
	if err != nil {
		return nil, err
	}
	values = append(values, accountValues...)
	return values, nil
}

func GetAccountAccessGroups(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender, account cloudflare.Account) ([]models.Resource, error) {
	opts := cloudflare.PaginationOptions{
		PerPage: 100,
		Page:    1,
	}
	type ListPageResponse struct {
		Groups []cloudflare.AccessGroup
		resp   cloudflare.ResultInfo
	}
	listPage := func(ctx context.Context) (interface{}, error) {
		groups, resp, err := conn.AccessGroups(ctx, account.ID, opts)
		return ListPageResponse{
			Groups: groups,
			resp:   resp,
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
		groups := listResponse.Groups
		resp := listResponse.resp
		for _, group := range groups {
			value := models.Resource{
				ID:   group.ID,
				Name: group.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.AccessGroupDescription{
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
