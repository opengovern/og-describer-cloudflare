package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"sync"
)

func ListWorkerRoutes(ctx context.Context, handler *CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	zones, err := getZones(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, zone := range zones {
			processWorkerRoutes(ctx, handler, zone, cloudFlareChan, &wg)
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

func GetWorkerRoute(ctx context.Context, handler *CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	var zoneID *string
	var zoneName *string
	zones, err := handler.Conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	route, err := processWorkerRoute(ctx, handler, zones, resourceID, zoneID, zoneName)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   route.ID,
		Name: route.ID,
		Description: model.WorkerRouteDescription{
			ID:       route.ID,
			ZoneName: *zoneName,
			Pattern:  route.Pattern,
			Script:   route.Script,
			ZoneID:   *zoneID,
		},
	}
	return &value, nil
}

func processWorkerRoutes(ctx context.Context, handler *CloudFlareAPIHandler, zone cloudflare.Zone, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var resp cloudflare.WorkerRoutesResponse
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		resp, e = handler.Conn.ListWorkerRoutes(ctx, zone.ID)
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
	for _, route := range resp.Routes {
		wg.Add(1)
		go func(route cloudflare.WorkerRoute) {
			defer wg.Done()
			value := models.Resource{
				ID:   route.ID,
				Name: route.ID,
				Description: model.WorkerRouteDescription{
					ID:       route.ID,
					ZoneName: zone.Name,
					Pattern:  route.Pattern,
					Script:   route.Script,
					ZoneID:   zone.ID,
				},
			}
			cloudFlareChan <- value
		}(route)
	}
}

func processWorkerRoute(ctx context.Context, handler *CloudFlareAPIHandler, zones []cloudflare.Zone, resourceID string, zoneID, zoneName *string) (*cloudflare.WorkerRoute, error) {
	var resp cloudflare.WorkerRoutesResponse
	var workerRoute cloudflare.WorkerRoute
	var statusCode *int
	for _, zone := range zones {
		requestFunc := func() (*int, error) {
			var e error
			resp, e = handler.Conn.ListWorkerRoutes(ctx, zone.ID)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			for _, route := range resp.Routes {
				if route.ID == resourceID {
					workerRoute = route
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
		if workerRoute.ID != "" {
			return &workerRoute, nil
		}
	}
	return nil, errors.New("DNS record with this ID doesn't exist")
}
