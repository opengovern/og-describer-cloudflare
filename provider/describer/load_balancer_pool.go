package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListLoadBalancerPools(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	loadBalancersPools, err := conn.ListLoadBalancerPools(ctx)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, loadBalancersPool := range loadBalancersPools {
		value := models.Resource{
			ID:   loadBalancersPool.ID,
			Name: loadBalancersPool.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.LoadBalancerPoolDescription{
					ID:                loadBalancersPool.ID,
					Name:              loadBalancersPool.Name,
					Enabled:           loadBalancersPool.Enabled,
					Monitor:           loadBalancersPool.Monitor,
					CreatedOn:         loadBalancersPool.CreatedOn,
					Description:       loadBalancersPool.Description,
					Latitude:          loadBalancersPool.Latitude,
					Longitude:         loadBalancersPool.Longitude,
					MinimumOrigins:    loadBalancersPool.MinimumOrigins,
					ModifiedOn:        loadBalancersPool.ModifiedOn,
					NotificationEmail: loadBalancersPool.NotificationEmail,
					CheckRegions:      loadBalancersPool.CheckRegions,
					LoadShedding:      loadBalancersPool.LoadShedding,
					Origins:           loadBalancersPool.Origins,
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

func GetLoadBalancerPool(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	loadBalancersPool, err := conn.LoadBalancerPoolDetails(ctx, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   loadBalancersPool.ID,
		Name: loadBalancersPool.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.LoadBalancerPoolDescription{
				ID:                loadBalancersPool.ID,
				Name:              loadBalancersPool.Name,
				Enabled:           loadBalancersPool.Enabled,
				Monitor:           loadBalancersPool.Monitor,
				CreatedOn:         loadBalancersPool.CreatedOn,
				Description:       loadBalancersPool.Description,
				Latitude:          loadBalancersPool.Latitude,
				Longitude:         loadBalancersPool.Longitude,
				MinimumOrigins:    loadBalancersPool.MinimumOrigins,
				ModifiedOn:        loadBalancersPool.ModifiedOn,
				NotificationEmail: loadBalancersPool.NotificationEmail,
				CheckRegions:      loadBalancersPool.CheckRegions,
				LoadShedding:      loadBalancersPool.LoadShedding,
				Origins:           loadBalancersPool.Origins,
			},
		},
	}
	return &value, nil
}
