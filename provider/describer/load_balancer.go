package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/model"
)

func ListLoadBalancers(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	zones, err := getZones(ctx, conn)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, zone := range zones {
		zoneValues, err := GetZoneLoadBalancers(ctx, conn, stream, zone)
		if err != nil {
			return nil, err
		}
		values = append(values, zoneValues...)
	}
	return values, nil
}

func GetLoadBalancer(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	var zoneID string
	var zoneName string
	zones, err := conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		loadBalancers, err := conn.ListLoadBalancers(ctx, zone.ID)
		if err != nil {
			return nil, err
		}
		for _, lb := range loadBalancers {
			if lb.ID == resourceID {
				zoneID = zone.ID
				zoneName = zone.Name
				break
			}
		}
		if zoneID != "" {
			break
		}
	}
	loadBalancer, err := conn.LoadBalancerDetails(ctx, zoneID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   loadBalancer.ID,
		Name: loadBalancer.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.LoadBalancerDescription{
				ID:                        loadBalancer.ID,
				Name:                      loadBalancer.Name,
				ZoneName:                  zoneName,
				ZoneID:                    zoneID,
				TTL:                       loadBalancer.TTL,
				Enabled:                   loadBalancer.Enabled,
				CreatedOn:                 loadBalancer.CreatedOn,
				Description:               loadBalancer.Description,
				FallbackPool:              loadBalancer.FallbackPool,
				ModifiedOn:                loadBalancer.ModifiedOn,
				Proxied:                   loadBalancer.Proxied,
				SessionAffinity:           loadBalancer.Persistence,
				SessionAffinityTTL:        loadBalancer.PersistenceTTL,
				SteeringPolicy:            loadBalancer.SteeringPolicy,
				DefaultPools:              loadBalancer.DefaultPools,
				PopPools:                  loadBalancer.PopPools,
				RegionPools:               loadBalancer.RegionPools,
				SessionAffinityAttributes: loadBalancer.SessionAffinityAttributes,
			},
		},
	}
	return &value, nil
}

func GetZoneLoadBalancers(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender, zone cloudflare.Zone) ([]models.Resource, error) {
	zoneID := zone.ID
	loadBalancers, err := conn.ListLoadBalancers(ctx, zoneID)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, loadBalancer := range loadBalancers {
		value := models.Resource{
			ID:   loadBalancer.ID,
			Name: loadBalancer.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.LoadBalancerDescription{
					ID:                        loadBalancer.ID,
					Name:                      loadBalancer.Name,
					ZoneName:                  zone.Name,
					ZoneID:                    zone.ID,
					TTL:                       loadBalancer.TTL,
					Enabled:                   loadBalancer.Enabled,
					CreatedOn:                 loadBalancer.CreatedOn,
					Description:               loadBalancer.Description,
					FallbackPool:              loadBalancer.FallbackPool,
					ModifiedOn:                loadBalancer.ModifiedOn,
					Proxied:                   loadBalancer.Proxied,
					SessionAffinity:           loadBalancer.Persistence,
					SessionAffinityTTL:        loadBalancer.PersistenceTTL,
					SteeringPolicy:            loadBalancer.SteeringPolicy,
					DefaultPools:              loadBalancer.DefaultPools,
					PopPools:                  loadBalancer.PopPools,
					RegionPools:               loadBalancer.RegionPools,
					SessionAffinityAttributes: loadBalancer.SessionAffinityAttributes,
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
