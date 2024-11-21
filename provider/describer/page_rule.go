package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListPageRules(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	zones, err := getZones(ctx, conn)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, zone := range zones {
		zoneValues, err := getZonePageRules(ctx, conn, stream, zone)
		if err != nil {
			return nil, err
		}
		values = append(values, zoneValues...)
	}
	return values, nil
}

func GetPageRule(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	var zoneID string
	zones, err := conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		pageRules, err := conn.ListPageRules(ctx, zone.ID)
		if err != nil {
			return nil, err
		}
		for _, pageRule := range pageRules {
			if pageRule.ID == resourceID {
				zoneID = zone.ID
				break
			}
		}
		if zoneID != "" {
			break
		}
	}
	pageRule, err := conn.PageRule(ctx, zoneID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   pageRule.ID,
		Name: pageRule.ID,
		Description: JSONAllFieldsMarshaller{
			Value: model.PageRuleDescription{
				ID:         pageRule.ID,
				Status:     pageRule.Status,
				ZoneID:     zoneID,
				CreatedOn:  pageRule.CreatedOn,
				ModifiedOn: pageRule.ModifiedOn,
				Priority:   pageRule.Priority,
				Title:      pageRule.ID,
				Actions:    pageRule.Actions,
				Targets:    pageRule.Targets,
			},
		},
	}
	return &value, nil
}

func getZonePageRules(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender, zone cloudflare.Zone) ([]models.Resource, error) {
	pageRules, err := conn.ListPageRules(ctx, zone.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, pageRule := range pageRules {
		value := models.Resource{
			ID:   pageRule.ID,
			Name: pageRule.ID,
			Description: JSONAllFieldsMarshaller{
				Value: model.PageRuleDescription{
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
