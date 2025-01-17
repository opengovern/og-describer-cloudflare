package describer

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListPageRules(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	zones, err := getZones(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, zone := range zones {
			processPageRules(ctx, handler, zone, cloudFlareChan, &wg)
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

func GetPageRule(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	var zoneID *string
	zones, err := handler.Conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	pageRule, err := processPageRule(ctx, handler, zones, resourceID, zoneID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   pageRule.ID,
		Name: pageRule.ID,
		Description: model.PageRuleDescription{
			ID:         pageRule.ID,
			Status:     pageRule.Status,
			ZoneID:     *zoneID,
			CreatedOn:  pageRule.CreatedOn,
			ModifiedOn: pageRule.ModifiedOn,
			Priority:   pageRule.Priority,
			Title:      pageRule.ID,
			Actions:    pageRule.Actions,
			Targets:    pageRule.Targets,
		},
	}
	return &value, nil
}

func processPageRules(ctx context.Context, handler *model.CloudFlareAPIHandler, zone cloudflare.Zone, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var pageRules []cloudflare.PageRule
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		pageRules, e = handler.Conn.ListPageRules(ctx, zone.ID)
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
		return
	}
	for _, pageRule := range pageRules {
		wg.Add(1)
		go func(pageRule cloudflare.PageRule) {
			defer wg.Done()
			value := models.Resource{
				ID:   pageRule.ID,
				Name: pageRule.ID,
				Description: model.PageRuleDescription{
					ID:         pageRule.ID,
					Status:     pageRule.Status,
					ZoneID:     zone.ID,
					CreatedOn:  pageRule.CreatedOn,
					ModifiedOn: pageRule.ModifiedOn,
					Priority:   pageRule.Priority,
					Title:      pageRule.ID,
					Actions:    pageRule.Actions,
					Targets:    pageRule.Targets,
				},
			}
			cloudFlareChan <- value
		}(pageRule)
	}
}

func processPageRule(ctx context.Context, handler *model.CloudFlareAPIHandler, zones []cloudflare.Zone, resourceID string, zoneID *string) (*cloudflare.PageRule, error) {
	var pageRules []cloudflare.PageRule
	var pageRule cloudflare.PageRule
	var statusCode *int
	for _, zone := range zones {
		requestFunc := func() (*int, error) {
			var e error
			pageRules, e = handler.Conn.ListPageRules(ctx, zone.ID)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			for _, rule := range pageRules {
				if rule.ID == resourceID {
					pageRule = rule
					zoneID = &zone.ID
					break
				}
			}
			return statusCode, e
		}
		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return nil, err
		}
		if pageRule.ID != "" {
			return &pageRule, nil
		}
	}
	return nil, errors.New("DNS record with this ID doesn't exist")
}
