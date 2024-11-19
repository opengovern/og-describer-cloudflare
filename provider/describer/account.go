package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
)

func GetAccount(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	account, _, err := conn.Account(ctx, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   account.ID,
		Name: account.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.AccountDescription{
				ID:       account.ID,
				Name:     account.Name,
				Type:     account.Type,
				Settings: account.Settings,
			},
		},
	}
	return &value, nil
}
