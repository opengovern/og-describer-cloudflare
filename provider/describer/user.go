package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListUsers(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	user, err := conn.UserDetails(ctx)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	value := models.Resource{
		ID:   user.ID,
		Name: user.Username,
		Description: JSONAllFieldsMarshaller{
			Value: model.UserDescription{
				ID:                             user.ID,
				Email:                          user.Email,
				Username:                       user.Username,
				Telephone:                      user.Telephone,
				FirstName:                      user.FirstName,
				LastName:                       user.LastName,
				Country:                        user.Country,
				Zipcode:                        user.Zipcode,
				CreatedOn:                      user.CreatedOn,
				ModifiedOn:                     user.ModifiedOn,
				APIKey:                         user.APIKey,
				TwoFactorAuthenticationEnabled: user.TwoFA,
				Betas:                          user.Betas,
				Organizations:                  user.Accounts,
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

func GetUser(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	user, err := conn.UserDetails(ctx)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   user.ID,
		Name: user.Username,
		Description: JSONAllFieldsMarshaller{
			Value: model.UserDescription{
				ID:                             user.ID,
				Email:                          user.Email,
				Username:                       user.Username,
				Telephone:                      user.Telephone,
				FirstName:                      user.FirstName,
				LastName:                       user.LastName,
				Country:                        user.Country,
				Zipcode:                        user.Zipcode,
				CreatedOn:                      user.CreatedOn,
				ModifiedOn:                     user.ModifiedOn,
				APIKey:                         user.APIKey,
				TwoFactorAuthenticationEnabled: user.TwoFA,
				Betas:                          user.Betas,
				Organizations:                  user.Accounts,
			},
		},
	}
	return &value, nil
}
