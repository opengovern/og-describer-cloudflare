package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"strings"
)

func ListAccountMembers(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, err
	}
	var pages cloudflare.PaginationOptions
	var values []models.Resource
	for {
		accountMembers, pageData, err := conn.AccountMembers(ctx, account.ID, pages)
		if err != nil {
			return nil, err
		}
		for _, accountMember := range accountMembers {
			title := accountMemberTitle(accountMember)
			value := models.Resource{
				ID:   accountMember.ID,
				Name: title,
				Description: JSONAllFieldsMarshaller{
					Value: model.AccountMemberDescription{
						UserEmail: accountMember.User.Email,
						ID:        accountMember.ID,
						Status:    accountMember.Status,
						AccountID: account.ID,
						Code:      accountMember.Code,
						User:      accountMember.User,
						Roles:     accountMember.Roles,
						Title:     title,
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
		if pageData.Page == pageData.TotalPages {
			break
		}
		pages.Page = pageData.Page + 1
	}
	return values, nil
}

func GetAccountMember(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	account, err := getAccount(ctx, conn)
	if err != nil {
		return nil, err
	}
	accountMember, err := conn.AccountMember(ctx, account.ID, resourceID)
	if err != nil {
		return nil, err
	}
	title := accountMemberTitle(accountMember)
	value := models.Resource{
		ID:   accountMember.ID,
		Name: title,
		Description: JSONAllFieldsMarshaller{
			Value: model.AccountMemberDescription{
				UserEmail: accountMember.User.Email,
				ID:        accountMember.ID,
				Status:    accountMember.Status,
				AccountID: account.ID,
				Code:      accountMember.Code,
				User:      accountMember.User,
				Roles:     accountMember.Roles,
				Title:     title,
			},
		},
	}
	return &value, nil
}

func accountMemberTitle(accountMember cloudflare.AccountMember) string {
	if len(accountMember.User.FirstName) > 0 && len(accountMember.User.LastName) > 0 {
		return accountMember.User.FirstName + " " + accountMember.User.LastName
	}
	return strings.Split(accountMember.User.Email, "@")[0]
}
