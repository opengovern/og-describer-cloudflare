package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListLoadBalancerMonitors(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	loadBalancersMonitors, err := conn.ListLoadBalancerMonitors(ctx)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, loadBalancersMonitor := range loadBalancersMonitors {
		value := models.Resource{
			ID:   loadBalancersMonitor.ID,
			Name: loadBalancersMonitor.ID,
			Description: JSONAllFieldsMarshaller{
				Value: model.LoadBalancerMonitorDescription{
					ID:              loadBalancersMonitor.ID,
					CreatedOn:       loadBalancersMonitor.CreatedOn,
					ModifiedOn:      loadBalancersMonitor.ModifiedOn,
					Type:            loadBalancersMonitor.Type,
					Description:     loadBalancersMonitor.Description,
					Method:          loadBalancersMonitor.Method,
					Path:            loadBalancersMonitor.Path,
					Header:          loadBalancersMonitor.Header,
					Timeout:         loadBalancersMonitor.Timeout,
					Retries:         loadBalancersMonitor.Retries,
					Interval:        loadBalancersMonitor.Interval,
					Port:            loadBalancersMonitor.Port,
					ExpectedBody:    loadBalancersMonitor.ExpectedBody,
					ExpectedCodes:   loadBalancersMonitor.ExpectedCodes,
					FollowRedirects: loadBalancersMonitor.FollowRedirects,
					AllowInsecure:   loadBalancersMonitor.AllowInsecure,
					ProbeZone:       loadBalancersMonitor.ProbeZone,
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

func GetLoadBalancerMonitor(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	loadBalancersMonitor, err := conn.LoadBalancerMonitorDetails(ctx, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   loadBalancersMonitor.ID,
		Name: loadBalancersMonitor.ID,
		Description: JSONAllFieldsMarshaller{
			Value: model.LoadBalancerMonitorDescription{
				ID:              loadBalancersMonitor.ID,
				CreatedOn:       loadBalancersMonitor.CreatedOn,
				ModifiedOn:      loadBalancersMonitor.ModifiedOn,
				Type:            loadBalancersMonitor.Type,
				Description:     loadBalancersMonitor.Description,
				Method:          loadBalancersMonitor.Method,
				Path:            loadBalancersMonitor.Path,
				Header:          loadBalancersMonitor.Header,
				Timeout:         loadBalancersMonitor.Timeout,
				Retries:         loadBalancersMonitor.Retries,
				Interval:        loadBalancersMonitor.Interval,
				Port:            loadBalancersMonitor.Port,
				ExpectedBody:    loadBalancersMonitor.ExpectedBody,
				ExpectedCodes:   loadBalancersMonitor.ExpectedCodes,
				FollowRedirects: loadBalancersMonitor.FollowRedirects,
				AllowInsecure:   loadBalancersMonitor.AllowInsecure,
				ProbeZone:       loadBalancersMonitor.ProbeZone,
			},
		},
	}
	return &value, nil
}
