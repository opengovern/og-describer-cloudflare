package describer

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListFireWallRules(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	zones, err := getZones(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, zone := range zones {
			processFirewallRules(ctx, handler, zone, cloudFlareChan, &wg)
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

func GetFireWallRule(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	var zoneID *string
	zones, err := handler.Conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	firewallRule, err := processFirewallRule(ctx, handler, zones, resourceID, zoneID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   firewallRule.ID,
		Name: firewallRule.ID,
		Description: model.FireWallRuleDescription{
			ID:          firewallRule.ID,
			Paused:      firewallRule.Paused,
			Description: firewallRule.Description,
			Action:      firewallRule.Action,
			Priority:    firewallRule.Priority,
			Title:       firewallRule.ID,
			Filter:      firewallRule.Filter,
			Products:    firewallRule.Products,
			CreatedOn:   firewallRule.CreatedOn,
			ModifiedOn:  firewallRule.ModifiedOn,
			ZoneID:      *zoneID,
		},
	}
	return &value, nil
}

func processFirewallRules(ctx context.Context, handler *model.CloudFlareAPIHandler, zone cloudflare.Zone, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var firewallRules []cloudflare.FirewallRule
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		firewallRules, e = handler.Conn.FirewallRules(ctx, zone.ID, cloudflare.PaginationOptions{})
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
	for _, firewallRule := range firewallRules {
		wg.Add(1)
		go func(firewallRule cloudflare.FirewallRule) {
			defer wg.Done()
			value := models.Resource{
				ID:   firewallRule.ID,
				Name: firewallRule.ID,
				Description: model.FireWallRuleDescription{
					ID:          firewallRule.ID,
					Paused:      firewallRule.Paused,
					Description: firewallRule.Description,
					Action:      firewallRule.Action,
					Priority:    firewallRule.Priority,
					Title:       firewallRule.ID,
					Filter:      firewallRule.Filter,
					Products:    firewallRule.Products,
					CreatedOn:   firewallRule.CreatedOn,
					ModifiedOn:  firewallRule.ModifiedOn,
					ZoneID:      zone.ID,
				},
			}
			cloudFlareChan <- value
		}(firewallRule)
	}
}

func processFirewallRule(ctx context.Context, handler *model.CloudFlareAPIHandler, zones []cloudflare.Zone, resourceID string, zoneID *string) (*cloudflare.FirewallRule, error) {
	var firewallRules []cloudflare.FirewallRule
	var firewallRule cloudflare.FirewallRule
	var statusCode *int
	for _, zone := range zones {
		requestFunc := func() (*int, error) {
			var e error
			firewallRules, e = handler.Conn.FirewallRules(ctx, zone.ID, cloudflare.PaginationOptions{})
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			for _, rule := range firewallRules {
				if rule.ID == resourceID {
					firewallRule = rule
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
		if firewallRule.ID != "" {
			return &firewallRule, nil
		}
	}
	return nil, errors.New("DNS record with this ID doesn't exist")
}
