package describers

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListLoadBalancerPools(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	go func() {
		processLoadBalancerPools(ctx, handler, cloudFlareChan, &wg)
		wg.Wait()
		close(cloudFlareChan)
	}()
	var values []models.Resource
	for value := range cloudFlareChan {
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

func GetLoadBalancerPool(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	loadBalancerPool, err := processLoadBalancerPool(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   loadBalancerPool.ID,
		Name: loadBalancerPool.Name,
		Description: model.LoadBalancerPoolDescription{
			ID:                loadBalancerPool.ID,
			Name:              loadBalancerPool.Name,
			Enabled:           loadBalancerPool.Enabled,
			Monitor:           loadBalancerPool.Monitor,
			CreatedOn:         loadBalancerPool.CreatedOn,
			Description:       loadBalancerPool.Description,
			Latitude:          loadBalancerPool.Latitude,
			Longitude:         loadBalancerPool.Longitude,
			MinimumOrigins:    loadBalancerPool.MinimumOrigins,
			ModifiedOn:        loadBalancerPool.ModifiedOn,
			NotificationEmail: loadBalancerPool.NotificationEmail,
			CheckRegions:      loadBalancerPool.CheckRegions,
			LoadShedding:      loadBalancerPool.LoadShedding,
			Origins:           loadBalancerPool.Origins,
		},
	}
	return &value, nil
}

func processLoadBalancerPools(ctx context.Context, handler *model.CloudFlareAPIHandler, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var loadBalancerPools []cloudflare.LoadBalancerPool
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		loadBalancerPools, e = handler.Conn.ListLoadBalancerPools(ctx)
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
	for _, loadBalancerPool := range loadBalancerPools {
		wg.Add(1)
		go func(loadBalancerPool cloudflare.LoadBalancerPool) {
			defer wg.Done()
			value := models.Resource{
				ID:   loadBalancerPool.ID,
				Name: loadBalancerPool.Name,
				Description: model.LoadBalancerPoolDescription{
					ID:                loadBalancerPool.ID,
					Name:              loadBalancerPool.Name,
					Enabled:           loadBalancerPool.Enabled,
					Monitor:           loadBalancerPool.Monitor,
					CreatedOn:         loadBalancerPool.CreatedOn,
					Description:       loadBalancerPool.Description,
					Latitude:          loadBalancerPool.Latitude,
					Longitude:         loadBalancerPool.Longitude,
					MinimumOrigins:    loadBalancerPool.MinimumOrigins,
					ModifiedOn:        loadBalancerPool.ModifiedOn,
					NotificationEmail: loadBalancerPool.NotificationEmail,
					CheckRegions:      loadBalancerPool.CheckRegions,
					LoadShedding:      loadBalancerPool.LoadShedding,
					Origins:           loadBalancerPool.Origins,
				},
			}
			cloudFlareChan <- value
		}(loadBalancerPool)
	}
}

func processLoadBalancerPool(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*cloudflare.LoadBalancerPool, error) {
	var loadBalancerPool cloudflare.LoadBalancerPool
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		loadBalancerPool, e = handler.Conn.LoadBalancerPoolDetails(ctx, resourceID)
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
	return &loadBalancerPool, nil
}
