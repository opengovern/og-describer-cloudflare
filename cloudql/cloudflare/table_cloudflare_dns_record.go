package cloudflare

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

		opengovernance "github.com/opengovern/og-describer-cloudflare/discovery/pkg/es"
)

func tableCloudflareDNSRecord(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_dns_record",
		Description: "DNS records for a zone.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListDNSRecord,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"zone_id", "id"}),
			Hydrate:    opengovernance.GetDNSRecord,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "zone_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ZoneID"), Description: "Zone where the record is defined."},
			{Name: "zone_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ZoneName"), Description: "Name of the zone where the record is defined."},
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "ID of the record."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "Type of the record (e.g. A, MX, CNAME)."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "Domain name for the record (e.g. steampipe.io)."},
			{Name: "content", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Content"), Description: "Content or value of the record. Changes by type, including IP address for A records and domain for CNAME records."},
			{Name: "ttl", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.TTL"), Description: "Time to live in seconds of the record."},

			// Other columns
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedOn"), Description: "When the record was created."},
			{Name: "locked", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Locked"), Description: "True if the record is locked."},
			{Name: "modified_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ModifiedOn"), Description: "When the record was last modified."},
			{Name: "priority", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Priority"), Description: "Priority for this record, primarily used for MX records."},
			{Name: "proxiable", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Proxiable"), Description: "True if the record is eligible for Cloudflare's origin protection."},
			{Name: "proxied", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Proxied"), Description: "True if the record has Cloudflare's origin protection."},

			// JSON columns
			{Name: "data", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Data"), Description: "Map of attributes that constitute the record value. Primarily used for LOC and SRV record types."},
			{Name: "meta", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Meta"), Description: "Cloudflare metadata for this record."},
		}),
	}
}
