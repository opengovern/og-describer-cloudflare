package describer

import (
	"context"
	"errors"
	"sync"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListLoadBalancerMonitors(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	go func() {
		processLoadBalancerMonitors(ctx, handler, cloudFlareChan, &wg)
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

func GetLoadBalancerMonitor(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	loadBalancerMonitor, err := processLoadBalancerMonitor(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   loadBalancerMonitor.ID,
		Name: loadBalancerMonitor.ID,
		Description: model.LoadBalancerMonitorDescription{
			ID:              loadBalancerMonitor.ID,
			CreatedOn:       loadBalancerMonitor.CreatedOn,
			ModifiedOn:      loadBalancerMonitor.ModifiedOn,
			Type:            loadBalancerMonitor.Type,
			Description:     loadBalancerMonitor.Description,
			Method:          loadBalancerMonitor.Method,
			Path:            loadBalancerMonitor.Path,
			Header:          loadBalancerMonitor.Header,
			Timeout:         loadBalancerMonitor.Timeout,
			Retries:         loadBalancerMonitor.Retries,
			Interval:        loadBalancerMonitor.Interval,
			Port:            loadBalancerMonitor.Port,
			ExpectedBody:    loadBalancerMonitor.ExpectedBody,
			ExpectedCodes:   loadBalancerMonitor.ExpectedCodes,
			FollowRedirects: loadBalancerMonitor.FollowRedirects,
			AllowInsecure:   loadBalancerMonitor.AllowInsecure,
			ProbeZone:       loadBalancerMonitor.ProbeZone,
		},
	}
	return &value, nil
}

func processLoadBalancerMonitors(ctx context.Context, handler *model.CloudFlareAPIHandler, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var loadBalancerMonitors []cloudflare.LoadBalancerMonitor
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		loadBalancerMonitors, e = handler.Conn.ListLoadBalancerMonitors(ctx)
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
	for _, loadBalancerMonitor := range loadBalancerMonitors {
		wg.Add(1)
		go func(loadBalancerMonitor cloudflare.LoadBalancerMonitor) {
			defer wg.Done()
			value := models.Resource{
				ID:   loadBalancerMonitor.ID,
				Name: loadBalancerMonitor.ID,
				Description: model.LoadBalancerMonitorDescription{
					ID:              loadBalancerMonitor.ID,
					CreatedOn:       loadBalancerMonitor.CreatedOn,
					ModifiedOn:      loadBalancerMonitor.ModifiedOn,
					Type:            loadBalancerMonitor.Type,
					Description:     loadBalancerMonitor.Description,
					Method:          loadBalancerMonitor.Method,
					Path:            loadBalancerMonitor.Path,
					Header:          loadBalancerMonitor.Header,
					Timeout:         loadBalancerMonitor.Timeout,
					Retries:         loadBalancerMonitor.Retries,
					Interval:        loadBalancerMonitor.Interval,
					Port:            loadBalancerMonitor.Port,
					ExpectedBody:    loadBalancerMonitor.ExpectedBody,
					ExpectedCodes:   loadBalancerMonitor.ExpectedCodes,
					FollowRedirects: loadBalancerMonitor.FollowRedirects,
					AllowInsecure:   loadBalancerMonitor.AllowInsecure,
					ProbeZone:       loadBalancerMonitor.ProbeZone,
				},
			}
			cloudFlareChan <- value
		}(loadBalancerMonitor)
	}
}

func processLoadBalancerMonitor(ctx context.Context, handler *model.CloudFlareAPIHandler, resourceID string) (*cloudflare.LoadBalancerMonitor, error) {
	var loadBalancerMonitor cloudflare.LoadBalancerMonitor
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		loadBalancerMonitor, e = handler.Conn.LoadBalancerMonitorDetails(ctx, resourceID)
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
	return &loadBalancerMonitor, nil
}
