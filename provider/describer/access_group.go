package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"github.com/turbot/go-kit/helpers"
)

func ListAccessGroups(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, nil
	}
	opts := cloudflare.PaginationOptions{
		PerPage: perPage,
		Page:    page,
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

func GetAccessGroup(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, nil
	}
	group, err := conn.AccessGroup(ctx, account.ID, resourceID)
	if err != nil {
		return nil, err
	}
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
	return &value, nil
}
