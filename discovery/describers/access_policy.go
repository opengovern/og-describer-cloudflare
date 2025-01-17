package describer

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListAccessPolicies(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	apps, err := getApplications(ctx, handler, account.ID)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, app := range apps {
			processAccessPolicies(ctx, handler, account, app, cloudFlareChan, &wg)
		}
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

func GetAccessPolicy(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	var appID *string
	var appName *string
	account, err := getAccount(ctx, handler)
	if err != nil {
		return nil, err
	}
	apps, err := getApplications(ctx, handler, account.ID)
	if err != nil {
		return nil, err
	}
	policy, err := processAccessPolicy(ctx, handler, account, apps, resourceID, appID, appName)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   policy.ID,
		Name: policy.Name,
		Description: model.AccessPolicyDescription{
			ID:                           policy.ID,
			Name:                         policy.Name,
			AccountID:                    account.ID,
			ApplicationID:                *appID,
			ApplicationName:              *appName,
			CreatedAt:                    policy.CreatedAt,
			Decision:                     policy.Decision,
			Precedence:                   policy.Precedence,
			PurposeJustificationPrompt:   policy.PurposeJustificationPrompt,
			PurposeJustificationRequired: policy.PurposeJustificationRequired,
			UpdatedAt:                    policy.UpdatedAt,
			ApprovalGroups:               policy.ApprovalGroups,
			Exclude:                      policy.Exclude,
			Include:                      policy.Include,
			Require:                      policy.Require,
		},
	}
	return &value, nil
}

func processAccessPolicies(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, app cloudflare.AccessApplication, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var accessPolicies []cloudflare.AccessPolicy
	var pageAccessPolicies []cloudflare.AccessPolicy
	var pageData cloudflare.ResultInfo
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		opts := cloudflare.PaginationOptions{
			PerPage: perPage,
			Page:    page,
		}
		for {
			pageAccessPolicies, pageData, e = handler.Conn.AccessPolicies(ctx, account.ID, app.ID, opts)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			accessPolicies = append(accessPolicies, pageAccessPolicies...)
			if pageData.Page >= pageData.TotalPages {
				break
			}
			opts.Page = opts.Page + 1
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, policy := range accessPolicies {
		wg.Add(1)
		go func(policy cloudflare.AccessPolicy) {
			defer wg.Done()
			value := models.Resource{
				ID:   policy.ID,
				Name: policy.Name,
				Description: model.AccessPolicyDescription{
					ID:                           policy.ID,
					Name:                         policy.Name,
					AccountID:                    account.ID,
					ApplicationID:                app.ID,
					ApplicationName:              app.Name,
					CreatedAt:                    policy.CreatedAt,
					Decision:                     policy.Decision,
					Precedence:                   policy.Precedence,
					PurposeJustificationPrompt:   policy.PurposeJustificationPrompt,
					PurposeJustificationRequired: policy.PurposeJustificationRequired,
					UpdatedAt:                    policy.UpdatedAt,
					ApprovalGroups:               policy.ApprovalGroups,
					Exclude:                      policy.Exclude,
					Include:                      policy.Include,
					Require:                      policy.Require,
				},
			}
			cloudFlareChan <- value
		}(policy)
	}
}

func processAccessPolicy(ctx context.Context, handler *model.CloudFlareAPIHandler, account *cloudflare.Account, apps []cloudflare.AccessApplication, resourceID string, appID, appName *string) (*cloudflare.AccessPolicy, error) {
	var accessPolicy cloudflare.AccessPolicy
	var accessPolicies []cloudflare.AccessPolicy
	var pageData cloudflare.ResultInfo
	var statusCode *int
	for _, app := range apps {
		requestFunc := func() (*int, error) {
			var e error
			opts := cloudflare.PaginationOptions{
				PerPage: perPage,
				Page:    page,
			}
			found := false
			for {
				accessPolicies, pageData, e = handler.Conn.AccessPolicies(ctx, account.ID, app.ID, opts)
				if e != nil {
					var httpErr *cloudflare.APIRequestError
					if errors.As(e, &httpErr) {
						statusCode = &httpErr.StatusCode
					}
				}
				for _, policy := range accessPolicies {
					if policy.ID == resourceID {
						accessPolicy = policy
						appID = &app.ID
						appName = &app.Name
						found = true
						break
					}
				}
				if found {
					break
				}
				if pageData.Page >= pageData.TotalPages {
					break
				}
				opts.Page = opts.Page + 1
			}
			return statusCode, e
		}
		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return nil, err
		}
		if accessPolicy.ID != "" {
			return &accessPolicy, nil
		}
	}
	return nil, errors.New("access policy with this ID doesn't exist")
}
