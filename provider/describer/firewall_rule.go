package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
)

func ListFireWallRules(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	zones, err := getZones(ctx, conn)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, zone := range zones {
		zoneValues, err := GetZoneFirewallRules(ctx, conn, stream, zone)
		if err != nil {
			return nil, err
		}
		values = append(values, zoneValues...)
	}
	return values, nil
}

func GetFireWallRule(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	var zoneID string
	zones, err := conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		fireWallRules, err := conn.FirewallRules(ctx, zoneID, cloudflare.PaginationOptions{})
		if err != nil {
			return nil, err
		}
		for _, fireWallRule := range fireWallRules {
			if fireWallRule.ID == resourceID {
				zoneID = zone.ID
				break
			}
		}
		if zoneID != "" {
			break
		}
	}
	fireWallRule, err := conn.FirewallRule(ctx, zoneID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   fireWallRule.ID,
		Name: fireWallRule.ID,
		Description: JSONAllFieldsMarshaller{
			Value: model.FireWallRuleDescription{
				ID:          fireWallRule.ID,
				Paused:      fireWallRule.Paused,
				Description: fireWallRule.Description,
				Action:      fireWallRule.Action,
				Priority:    fireWallRule.Priority,
				Title:       fireWallRule.ID,
				Filter:      fireWallRule.Filter,
				Products:    fireWallRule.Products,
				CreatedOn:   fireWallRule.CreatedOn,
				ModifiedOn:  fireWallRule.ModifiedOn,
				ZoneID:      zoneID,
			},
		},
	}
	return &value, nil
}

func GetZoneFirewallRules(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender, zone cloudflare.Zone) ([]models.Resource, error) {
	zoneID := zone.ID
	fireWallRules, err := conn.FirewallRules(ctx, zoneID, cloudflare.PaginationOptions{})
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, fireWallRule := range fireWallRules {
		value := models.Resource{
			ID:   fireWallRule.ID,
			Name: fireWallRule.ID,
			Description: JSONAllFieldsMarshaller{
				Value: model.FireWallRuleDescription{
					ID:          fireWallRule.ID,
					Paused:      fireWallRule.Paused,
					Description: fireWallRule.Description,
					Action:      fireWallRule.Action,
					Priority:    fireWallRule.Priority,
					Title:       fireWallRule.ID,
					Filter:      fireWallRule.Filter,
					Products:    fireWallRule.Products,
					CreatedOn:   fireWallRule.CreatedOn,
					ModifiedOn:  fireWallRule.ModifiedOn,
					ZoneID:      zoneID,
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
