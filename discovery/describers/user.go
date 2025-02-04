package describers

import (
	"context"
	"errors"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListUsers(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	user, err := processUser(ctx, handler)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	value := models.Resource{
		ID:   user.ID,
		Name: user.Username,
		Description: model.UserDescription{
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
	}
	if stream != nil {
		if err = (*stream)(value); err != nil {
			return nil, err
		}
	} else {
		values = append(values, value)
	}
	return values, nil
}

func GetUser(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	user, err := processUser(ctx, handler)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   user.ID,
		Name: user.Username,
		Description: model.UserDescription{
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
	}
	return &value, nil
}

func processUser(ctx context.Context, handler *model.CloudFlareAPIHandler) (*cloudflare.User, error) {
	var user cloudflare.User
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		user, e = handler.Conn.UserDetails(ctx)
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
	return &user, nil
}
