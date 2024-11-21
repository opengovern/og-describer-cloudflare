package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
)

func ListDNSRecords(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	zones, err := getZones(ctx, conn)
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, zone := range zones {
		zoneValues, err := getZoneDNSRecord(ctx, conn, stream, zone)
		if err != nil {
			return nil, err
		}
		values = append(values, zoneValues...)
	}
	return values, nil
}

func GetDNSRecord(ctx context.Context, conn *cloudflare.API, resourceID string) (*models.Resource, error) {
	var zoneID string
	var zoneName string
	zones, err := conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	for _, zone := range zones {
		items, err := conn.DNSRecords(ctx, zone.ID, cloudflare.DNSRecord{})
		if err != nil {
			return nil, err
		}
		for _, item := range items {
			if item.ID == resourceID {
				zoneID = zone.ID
				zoneName = zone.Name
				break
			}
		}
		if zoneID != "" {
			break
		}
	}
	dnsRecord, err := conn.DNSRecord(ctx, zoneID, resourceID)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   dnsRecord.ID,
		Name: dnsRecord.Name,
		Description: JSONAllFieldsMarshaller{
			Value: model.DNSRecordDescription{
				ZoneID:     zoneID,
				ZoneName:   zoneName,
				ID:         dnsRecord.ID,
				Type:       dnsRecord.Type,
				Name:       dnsRecord.Name,
				Content:    dnsRecord.Content,
				TTL:        dnsRecord.TTL,
				CreatedOn:  dnsRecord.CreatedOn,
				Locked:     dnsRecord.Locked,
				ModifiedOn: dnsRecord.ModifiedOn,
				Priority:   dnsRecord.Priority,
				Proxiable:  dnsRecord.Proxiable,
				Proxied:    dnsRecord.Proxied,
				Data:       dnsRecord.Data,
				Meta:       dnsRecord.Meta,
			},
		},
	}
	return &value, nil
}

func getZoneDNSRecord(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender, zone cloudflare.Zone) ([]models.Resource, error) {
	zoneID := zone.ID
	items, err := conn.DNSRecords(ctx, zoneID, cloudflare.DNSRecord{})
	if err != nil {
		return nil, err
	}
	var values []models.Resource
	for _, item := range items {
		value := models.Resource{
			ID:   item.ID,
			Name: item.Name,
			Description: JSONAllFieldsMarshaller{
				Value: model.DNSRecordDescription{
					ZoneID:     zoneID,
					ZoneName:   zone.Name,
					ID:         item.ID,
					Type:       item.Type,
					Name:       item.Name,
					Content:    item.Content,
					TTL:        item.TTL,
					CreatedOn:  item.CreatedOn,
					Locked:     item.Locked,
					ModifiedOn: item.ModifiedOn,
					Priority:   item.Priority,
					Proxiable:  item.Proxiable,
					Proxied:    item.Proxied,
					Data:       item.Data,
					Meta:       item.Meta,
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
