package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListAccountRoles(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, err
	}
	accountRoles, err := conn.AccountRoles(ctx, account.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, accountRole := range accountRoles {
		value := models.Resource{
			ID:   accountRole.ID,
			Name: accountRole.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.AccountRoleDescription{
					ID:          accountRole.ID,
					Name:        accountRole.Name,
					Description: accountRole.Description,
					Permissions: accountRole.Permissions,
					AccountID:   account.ID,
					Title:       accountRole.Name,
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
	return values, nil
}

func GetAccountRole(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, err
	}
	accountRole, err := conn.AccountRole(ctx, account.ID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   accountRole.ID,
		Name: accountRole.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.AccountRoleDescription{
				ID:          accountRole.ID,
				Name:        accountRole.Name,
				Description: accountRole.Description,
				Permissions: accountRole.Permissions,
				AccountID:   account.ID,
				Title:       accountRole.Name,
			},
		},
	}
	return &value, nil
}
