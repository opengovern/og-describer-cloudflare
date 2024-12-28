package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"sync"
)

func ListLoadBalancers(ctx context.Context, handler *CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	zones, err := getZones(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, zone := range zones {
			processLoadBalancers(ctx, handler, zone, cloudFlareChan, &wg)
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

func GetLoadBalancer(ctx context.Context, handler *CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	var zoneID *string
	var zoneName *string
	zones, err := handler.Conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	loadBalancer, err := processLoadBalancer(ctx, handler, zones, resourceID, zoneID, zoneName)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   loadBalancer.ID,
		Name: loadBalancer.Name,
		Description: model.LoadBalancerDescription{
			ID:                        loadBalancer.ID,
			Name:                      loadBalancer.Name,
			ZoneName:                  *zoneName,
			ZoneID:                    *zoneID,
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
	}
	return &value, nil
}

func processLoadBalancers(ctx context.Context, handler *CloudFlareAPIHandler, zone cloudflare.Zone, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var loadBalancers []cloudflare.LoadBalancer
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		loadBalancers, e = handler.Conn.ListLoadBalancers(ctx, zone.ID)
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
	for _, loadBalancer := range loadBalancers {
		wg.Add(1)
		go func(loadBalancer cloudflare.LoadBalancer) {
			defer wg.Done()
			value := models.Resource{
				ID:   loadBalancer.ID,
				Name: loadBalancer.Name,
				Description: model.LoadBalancerDescription{
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
			}
			cloudFlareChan <- value
		}(loadBalancer)
	}
}

func processLoadBalancer(ctx context.Context, handler *CloudFlareAPIHandler, zones []cloudflare.Zone, resourceID string, zoneID, zoneName *string) (*cloudflare.LoadBalancer, error) {
	var loadBalancers []cloudflare.LoadBalancer
	var loadBalancer cloudflare.LoadBalancer
	var statusCode *int
	for _, zone := range zones {
		requestFunc := func() (*int, error) {
			var e error
			loadBalancers, e = handler.Conn.ListLoadBalancers(ctx, zone.ID)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			for _, lb := range loadBalancers {
				if lb.ID == resourceID {
					loadBalancer = lb
					zoneID = &zone.ID
					zoneName = &zone.Name
					break
				}
			}
			return statusCode, e
		}
		err := handler.DoRequest(ctx, requestFunc)
		if err != nil {
			return nil, err
		}
		if loadBalancer.ID != "" {
			return &loadBalancer, nil
		}
	}
	return nil, errors.New("DNS record with this ID doesn't exist")
}
