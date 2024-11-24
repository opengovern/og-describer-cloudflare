package cloudflare

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"

	"github.com/cloudflare/cloudflare-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCloudflareZone(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_zone",
		Description: "A Zone is a domain name along with its subdomains and other identities.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListZone,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.SingleColumn("id"),
			Hydrate:    opengovernance.GetZone,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "Zone identifier tag."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The domain name."},

			// Other columns
			// TODO - do we need this here {Name: "account", Type: proto.ColumnType_JSON, Description: "TODO"},
			{Name: "betas", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Betas"), Description: "Beta feature flags associated with the zone."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedOn"), Description: "When the zone was created."},
			{Name: "deactivation_reason", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.DeactivationReason"), Description: "TODO"},
			{Name: "development_mode", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.DevelopmentMode"), Description: "The interval (in seconds) from when development mode expires (positive integer) or last expired (negative integer) for the domain. If development mode has never been enabled, this value is 0."},
			{Name: "dnssec", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.DNSSEC"), Description: "DNSSEC settings for the zone."},
			{Name: "host", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Host"), Description: "TODO"},
			{Name: "meta", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Meta"), Description: "Metadata associated with the zone."},
			{Name: "modified_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ModifiedOn"), Description: "When the zone was last modified."},
			{Name: "name_servers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.NameServers"), Description: "Cloudflare-assigned name servers. This is only populated for zones that use Cloudflare DNS."},
			{Name: "original_dnshost", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.OriginalDNSHost"), Description: "DNS host at the time of switching to Cloudflare."},
			{Name: "original_name_servers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.OriginalNameServers"), Description: "Original name servers before moving to Cloudflare."},
			{Name: "original_registrar", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.OriginalRegistrar"), Description: "Registrar for the domain at the time of switching to Cloudflare."},
			{Name: "owner", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Owner"), Description: "Information about the user or organization that owns the zone."},
			{Name: "paused", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Paused"), Description: "Indicates if the zone is only using Cloudflare DNS services. A true value means the zone will not receive security or performance benefits."},
			{Name: "permissions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Permissions"), Description: "Available permissions on the zone for the current user requesting the item."},
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Settings").Transform(settingsToStandard), Description: "Simple key value map of zone settings like advanced_ddos = on. Full settings details are in settings_src."},
			//{Name: "settings_src", Type: proto.ColumnType_JSON, Hydrate: getZoneSettings, Transform: transform.FromValue(), Description: "Original source form of zone settings for caching, security and other features of Cloudflare."},
			{Name: "plan", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Plan"), Description: "Current plan associated with the zone."},
			{Name: "plan_pending", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.PlanPending"), Description: "Pending plan change associated with the zone."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Status"), Description: "Status of the zone."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "A full zone implies that DNS is hosted with Cloudflare. A partial zone is typically a partner-hosted zone or a CNAME setup."},
			//{Name: "universal_ssl_settings", Type: proto.ColumnType_JSON, Hydrate: getZoneUniversalSSLSettings, Transform: transform.FromValue(), Description: "Universal SSL settings for a zone."},
			{Name: "vanity_name_servers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.VanityNameServers"), Description: "Custom name servers for the zone."},
			// TODO - It's unclear when this is set {Name: "verification_key", Type: proto.ColumnType_STRING, Description: "TODO"},
		}),
	}
}

func settingsToStandard(ctx context.Context, d *transform.TransformData) (interface{}, error) {
	settings := d.HydrateItem.([]cloudflare.ZoneSetting)
	// Convert the settings into a map, which makes them a lot easier to query by name
	settingsMap := map[string]interface{}{}
	for _, i := range settings {
		settingsMap[i.ID] = i.Value
	}
	return settingsMap, nil
}
