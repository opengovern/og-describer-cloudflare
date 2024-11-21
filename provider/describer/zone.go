package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListZones(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	resp, err := conn.ListZonesContext(ctx)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, zone := range resp.Result {
		zoneSetting, err := getZoneSettings(ctx, conn, zone)
		if err != nil {
			return nil, err
		}
		settings := settingsToStandard(zoneSetting)
		zoneDNSSEC, err := getZoneDNSSEC(ctx, conn, zone)
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

func GetZone(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	zone, err := conn.ZoneDetails(ctx, resourceID)
	if err != nil {
		return nil, err
	}
	zoneSetting, err := getZoneSettings(ctx, conn, zone)
	if err != nil {
		return nil, err
	}
	settings := settingsToStandard(zoneSetting)
	zoneDNSSEC, err := getZoneDNSSEC(ctx, conn, zone)
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

func getZoneSettings(ctx context.Context, conn *cloudflare.API, zone cloudflare.Zone) ([]cloudflare.ZoneSetting, error) {
	resp, err := conn.ZoneSettings(ctx, zone.ID)
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

func getZoneDNSSEC(ctx context.Context, conn *cloudflare.API, zone cloudflare.Zone) (*cloudflare.ZoneDNSSEC, error) {
	zoneDNSSEC, err := conn.ZoneDNSSECSetting(ctx, zone.ID)
	if err != nil {
		return nil, err
	}
	return &zoneDNSSEC, nil
}
