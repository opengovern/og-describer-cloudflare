package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListAccessPolicies(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, nil
	}
	apps, err := getApplications(ctx, conn, account.ID)
	if err != nil {
		return nil, nil
	}
	var values []models.Resource
	for _, app := range apps {
		accountValues, err := getAppAccessPolicies(ctx, conn, stream, account.ID, app)
		if err != nil {
			return nil, err
		}
		values = append(values, accountValues...)
	}
	return values, nil
}

func GetAccessPolicy(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	var appID string
	var appName string
	opts := cloudflare.PaginationOptions{
		PerPage: perPage,
		Page:    page,
	}
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, nil
	}
	apps, err := getApplications(ctx, conn, account.ID)
	if err != nil {
		return nil, nil
	}
	for _, app := range apps {
		accessPolicies, _, err := conn.AccessPolicies(ctx, account.ID, app.ID, opts)
		if err != nil {
			return nil, err
		}
		for _, accessPolicy := range accessPolicies {
			if accessPolicy.ID == resourceID {
				appID = app.ID
				appName = app.Name
				break
			}
		}
		if appID != "" {
			break
		}
	}
	item, err := conn.AccessPolicy(ctx, account.ID, appID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   item.ID,
		Name: item.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.AccessPolicyDescription{
				ID:                           item.ID,
				Name:                         item.Name,
				AccountID:                    account.ID,
				ApplicationID:                appID,
				ApplicationName:              appName,
				CreatedAt:                    item.CreatedAt,
				Decision:                     item.Decision,
				Precedence:                   item.Precedence,
				PurposeJustificationPrompt:   item.PurposeJustificationPrompt,
				PurposeJustificationRequired: item.PurposeJustificationRequired,
				UpdatedAt:                    item.UpdatedAt,
				ApprovalGroups:               item.ApprovalGroups,
				Exclude:                      item.Exclude,
				Include:                      item.Include,
				Require:                      item.Require,
			},
		},
	}
	return &value, nil
}

func getAppAccessPolicies(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender, accountID string, app cloudflare.AccessApplication) ([]models.Resource, error) {
	appID := app.ID
	opts := cloudflare.PaginationOptions{
		PerPage: perPage,
		Page:    page,
	}
	var values []models.Resource
	for {
		items, resultInfo, err := conn.AccessPolicies(ctx, accountID, appID, opts)
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			value := models.Resource{
				ID:   item.ID,
				Name: item.Name,
				Description: JSONAllFieldsMarshaller{
					Value: model.AccessPolicyDescription{
						ID:                           item.ID,
						Name:                         item.Name,
						AccountID:                    accountID,
						ApplicationID:                app.ID,
						ApplicationName:              app.Name,
						CreatedAt:                    item.CreatedAt,
						Decision:                     item.Decision,
						Precedence:                   item.Precedence,
						PurposeJustificationPrompt:   item.PurposeJustificationPrompt,
						PurposeJustificationRequired: item.PurposeJustificationRequired,
						UpdatedAt:                    item.UpdatedAt,
						ApprovalGroups:               item.ApprovalGroups,
						Exclude:                      item.Exclude,
						Include:                      item.Include,
						Require:                      item.Require,
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
		if resultInfo.Page >= resultInfo.TotalPages {
			break
		}
		opts.Page = opts.Page + 1
	}
	return values, nil
}
