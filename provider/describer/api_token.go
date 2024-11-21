package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListAPITokens(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	tokens, err := conn.APITokens(ctx)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, token := range tokens {
		value := models.Resource{
			ID:   token.ID,
			Name: token.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.ApiTokenDescription{
					ID:         token.ID,
					Name:       token.Name,
					Status:     token.Status,
					Condition:  token.Condition,
					ExpiresOn:  token.ExpiresOn,
					IssuedOn:   token.IssuedOn,
					ModifiedOn: token.ModifiedOn,
					NotBefore:  token.NotBefore,
					Policies:   token.Policies,
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

func GetApiToken(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	token, err := conn.GetAPIToken(ctx, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   token.ID,
		Name: token.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.ApiTokenDescription{
				ID:         token.ID,
				Name:       token.Name,
				Status:     token.Status,
				Condition:  token.Condition,
				ExpiresOn:  token.ExpiresOn,
				IssuedOn:   token.IssuedOn,
				ModifiedOn: token.ModifiedOn,
				NotBefore:  token.NotBefore,
				Policies:   token.Policies,
			},
		},
	}
	return &value, nil
}
