package describer

import (
	"context"
	"errors"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"sync"
)

func ListDNSRecords(ctx context.Context, handler *CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	zones, err := getZones(ctx, handler)
	if err != nil {
		return nil, err
	}
	go func() {
		for _, zone := range zones {
			processDNSRecords(ctx, handler, zone, cloudFlareChan, &wg)
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

func GetDNSRecord(ctx context.Context, handler *CloudFlareAPIHandler, resourceID string) (*models.Resource, error) {
	var zoneID *string
	var zoneName *string
	zones, err := handler.Conn.ListZones(ctx)
	if err != nil {
		return nil, err
	}
	dnsRecord, err := processDNSRecord(ctx, handler, zones, resourceID, zoneID, zoneName)
	if err != nil {
		return nil, err
	}
	value := models.Resource{
		ID:   dnsRecord.ID,
		Name: dnsRecord.Name,
		Description: model.DNSRecordDescription{
			ZoneID:     *zoneID,
			ZoneName:   *zoneName,
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
	}
	return &value, nil
}

func processDNSRecords(ctx context.Context, handler *CloudFlareAPIHandler, zone cloudflare.Zone, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var dnsRecords []cloudflare.DNSRecord
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		dnsRecords, e = handler.Conn.DNSRecords(ctx, zone.ID, cloudflare.DNSRecord{})
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
	for _, dnsRecord := range dnsRecords {
		wg.Add(1)
		go func(dnsRecord cloudflare.DNSRecord) {
			defer wg.Done()
			value := models.Resource{
				ID:   dnsRecord.ID,
				Name: dnsRecord.Name,
				Description: model.DNSRecordDescription{
					ZoneID:     zone.ID,
					ZoneName:   zone.Name,
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
			}
			cloudFlareChan <- value
		}(dnsRecord)
	}
}

func processDNSRecord(ctx context.Context, handler *CloudFlareAPIHandler, zones []cloudflare.Zone, resourceID string, zoneID, zoneName *string) (*cloudflare.DNSRecord, error) {
	var dnsRecords []cloudflare.DNSRecord
	var dnsRecord cloudflare.DNSRecord
	var statusCode *int
	for _, zone := range zones {
		requestFunc := func() (*int, error) {
			var e error
			dnsRecords, e = handler.Conn.DNSRecords(ctx, zone.ID, cloudflare.DNSRecord{})
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			for _, record := range dnsRecords {
				if record.ID == resourceID {
					dnsRecord = record
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
		if dnsRecord.ID != "" {
			return &dnsRecord, nil
		}
	}
	return nil, errors.New("DNS record with this ID doesn't exist")
}
