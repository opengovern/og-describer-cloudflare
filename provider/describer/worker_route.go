package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListWorkerRoutes(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	zones, err := getZones(ctx, conn)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, zone := range zones {
		zoneValues, err := getZoneWorkerRoutes(ctx, conn, stream, zone)
		if err != nil {
			return nil, err
		}
		values = append(values, zoneValues...)
	}
	return values, nil
}

func GetWorkerRoute(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	var zoneID string
	var zoneName string
	zones, err := conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		resp, err := conn.ListWorkerRoutes(ctx, zone.ID)
		if err != nil {
			return nil, err
		}
		for _, workerRoute := range resp.Routes {
			if workerRoute.ID == resourceID {
				zoneID = zone.ID
				zoneName = zone.Name
				break
			}
		}
		if zoneID != "" {
			break
		}
	}
	route, err := conn.GetWorkerRoute(ctx, zoneID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   route.ID,
		Name: route.ID,
		Description: JSONAllFieldsMarshaller{
			Value: model.WorkerRouteDescription{
				ID:       route.ID,
				ZoneName: zoneName,
				Pattern:  route.Pattern,
				Script:   route.Script,
				ZoneID:   zoneID,
			},
		},
	}
	return &value, nil
}

func getZoneWorkerRoutes(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender, zone cloudflare.Zone) ([]models.Resource, error) {
	resp, err := conn.ListWorkerRoutes(ctx, zone.ID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, route := range resp.Routes {
		value := models.Resource{
			ID:   route.ID,
			Name: route.ID,
			Description: JSONAllFieldsMarshaller{
				Value: model.WorkerRouteDescription{
					ID:       route.ID,
					ZoneName: zone.Name,
					Pattern:  route.Pattern,
					Script:   route.Script,
					ZoneID:   zone.ID,
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
