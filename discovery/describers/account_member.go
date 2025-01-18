package describers

import (
	"context"
	"errors"
	"strings"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListAccountMembers(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		processAccountMembers(ctx, handler, account, cloudFlareChan, &wg)
		wg.Wait()
		close(cloudFlareChan)
	}()
	var values []models.Resource
	for value := range cloudFlareChan {
		if stream != nil {
			if err = (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func GetAccountMember(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	accountMember, err := processAccountMember(ctx, handler, account, resourceID)
	if err != nil {
		return nil, err
	}
	title := accountMemberTitle(*accountMember)
	value := models.Resource{
		ID:   accountMember.ID,
		Name: title,
		Description: model.AccountMemberDescription{
			UserEmail: accountMember.User.Email,
			ID:        accountMember.ID,
			Status:    accountMember.Status,
			AccountID: account.ID,
			Code:      accountMember.Code,
			User:      accountMember.User,
			Roles:     accountMember.Roles,
			Title:     title,
		},
	}
	return &value, nil
}

func processAccountMembers(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var accountMembers []cloudflare.AccountMember
	var pageAccountMembers []cloudflare.AccountMember
	var pageData cloudflare.ResultInfo
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		var pages cloudflare.PaginationOptions
		for {
			pageAccountMembers, pageData, e = handler.Conn.AccountMembers(ctx, account.ID, pages)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			accountMembers = append(accountMembers, pageAccountMembers...)
			if pageData.Page == pageData.TotalPages {
				break
			}
			pages.Page = pageData.Page + 1
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, accountMember := range accountMembers {
		wg.Add(1)
		go func(accountMember cloudflare.AccountMember) {
			defer wg.Done()
			title := accountMemberTitle(accountMember)
			value := models.Resource{
				ID:   accountMember.ID,
				Name: title,
				Description: model.AccountMemberDescription{
					UserEmail: accountMember.User.Email,
					ID:        accountMember.ID,
					Status:    accountMember.Status,
					AccountID: account.ID,
					Code:      accountMember.Code,
					User:      accountMember.User,
					Roles:     accountMember.Roles,
					Title:     title,
				},
			}
			cloudFlareChan <- value
		}(accountMember)
	}
}

func processAccountMember(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, resourceID string) (*cloudflare.AccountMember, error) {
	var accountMember cloudflare.AccountMember
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		accountMember, e = handler.Conn.AccountMember(ctx, account.ID, resourceID)
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
	return &accountMember, nil
}

func accountMemberTitle(accountMember cloudflare.AccountMember) string {
	if len(accountMember.User.FirstName) > 0 && len(accountMember.User.LastName) > 0 {
		return accountMember.User.FirstName + " " + accountMember.User.LastName
	}
	return strings.Split(accountMember.User.Email, "@")[0]
}
