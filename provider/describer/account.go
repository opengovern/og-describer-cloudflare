package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListAccounts(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
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
	if stream != nil {
		if err := (*stream)(value); err != nil {
			return nil, err
		}
	} else {
		values = append(values, value)
	}
	return values, nil
}

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
