package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"sync"
)

func ListZones(ctx context.Context, handler *CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	go func() {
		processZones(ctx, handler, cloudFlareChan, &wg)
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

func GetZone(ctx context.Context, handler *CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	zone, err := processZone(ctx, handler, resourceID)
	if err != nil {
		return nil, err
	}
	zoneSetting, err := getZoneSettings(ctx, handler, zone)
	if err != nil {
		return nil, err
	}
	settings := settingsToStandard(zoneSetting)
	zoneDNSSEC, err := getZoneDNSSEC(ctx, handler, zone)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   zone.ID,
		Name: zone.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.ZoneDescription{
				ID:              zone.ID,
				Name:            zone.Name,
				Betas:           zone.Betas,
				CreatedOn:       zone.CreatedOn,
				DevelopmentMode: zone.DevMode,
				DNSSEC:          *zoneDNSSEC,
				Host: struct {
					Name    string
					Website string
				}{Name: zone.Host.Name, Website: zone.Host.Website},
				Meta:                zone.Meta,
				ModifiedOn:          zone.ModifiedOn,
				NameServers:         zone.NameServers,
				OriginalDNSHost:     zone.OriginalDNSHost,
				OriginalNameServers: zone.OriginalNS,
				OriginalRegistrar:   zone.OriginalRegistrar,
				Owner:               zone.Owner,
				Paused:              zone.Paused,
				Permissions:         zone.Permissions,
				Settings:            settings,
				Plan:                zone.Plan,
				PlanPending:         zone.PlanPending,
				Status:              zone.Status,
				Type:                zone.Type,
				VanityNameServers:   zone.VanityNS,
			},
		},
	}
	return &value, nil
}

func processZones(ctx context.Context, handler *CloudFlareAPIHandler, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var resp cloudflare.ZonesResponse
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		resp, e = handler.Conn.ListZonesContext(ctx)
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
	for _, zone := range resp.Result {
		wg.Add(1)
		go func(zone cloudflare.Zone) {
			defer wg.Done()
			zoneSetting, err := getZoneSettings(ctx, handler, &zone)
			if err == nil {
				settings := settingsToStandard(zoneSetting)
				zoneDNSSEC, err := getZoneDNSSEC(ctx, handler, &zone)
				if err == nil {
					value := models.Resource{
						ID:   zone.ID,
						Name: zone.Name,
						Description: JSONAllFieldsMarshaller{
							Value: model.ZoneDescription{
								ID:              zone.ID,
								Name:            zone.Name,
								Betas:           zone.Betas,
								CreatedOn:       zone.CreatedOn,
								DevelopmentMode: zone.DevMode,
								DNSSEC:          *zoneDNSSEC,
								Host: struct {
									Name    string
									Website string
								}{Name: zone.Host.Name, Website: zone.Host.Website},
								Meta:                zone.Meta,
								ModifiedOn:          zone.ModifiedOn,
								NameServers:         zone.NameServers,
								OriginalDNSHost:     zone.OriginalDNSHost,
								OriginalNameServers: zone.OriginalNS,
								OriginalRegistrar:   zone.OriginalRegistrar,
								Owner:               zone.Owner,
								Paused:              zone.Paused,
								Permissions:         zone.Permissions,
								Settings:            settings,
								Plan:                zone.Plan,
								PlanPending:         zone.PlanPending,
								Status:              zone.Status,
								Type:                zone.Type,
								VanityNameServers:   zone.VanityNS,
							},
						},
					}
					cloudFlareChan <- value
				}
			}
		}(zone)
	}
}

func processZone(ctx context.Context, handler *CloudFlareAPIHandler, resourceID string) (*cloudflare.Zone, error) {
	var zone cloudflare.Zone
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		zone, e = handler.Conn.ZoneDetails(ctx, resourceID)
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
	return &zone, nil
}

func getZoneSettings(ctx context.Context, handler *CloudFlareAPIHandler, zone *cloudflare.Zone) ([]cloudflare.ZoneSetting, error) {
	var resp *cloudflare.ZoneSettingResponse
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		resp, e = handler.Conn.ZoneSettings(ctx, zone.ID)
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
	return resp.Result, nil
}

func settingsToStandard(zoneSetting []cloudflare.ZoneSetting) map[string]interface{} {
	settingsMap := map[string]interface{}{}
	for _, item := range zoneSetting {
		settingsMap[item.ID] = item.Value
	}
	return settingsMap
}

func getZoneDNSSEC(ctx context.Context, handler *CloudFlareAPIHandler, zone *cloudflare.Zone) (*cloudflare.ZoneDNSSEC, error) {
	var zoneDNSSEC cloudflare.ZoneDNSSEC
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		zoneDNSSEC, e = handler.Conn.ZoneDNSSECSetting(ctx, zone.ID)
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
	return &zoneDNSSEC, nil
}
