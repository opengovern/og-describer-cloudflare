package cloudflare

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

		opengovernance "github.com/opengovern/og-describer-cloudflare/discovery/pkg/es"
)

func tableCloudflareFirewallRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_firewall_rule",
		Description: "Cloudflare Firewall Rules is a flexible and intuitive framework for filtering HTTP requests.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListFireWallRule,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"zone_id", "id"}),
			Hydrate:    opengovernance.GetFireWallRule,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "Specifies the Firewall Rule identifier."},
			{Name: "zone_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ZoneID"), Description: "Specifies the zone identifier."},
			{Name: "paused", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.Paused"), Description: "Indicates whether the firewall rule is currently paused."},
			{Name: "priority", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Priority"), Description: "The priority of the rule to allow control of processing order. A lower number indicates high priority. If not provided, any rules with a priority will be sequenced before those without."},
			{Name: "action", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Action"), Description: "The action to apply to a matched request."},

			// Other columns
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedOn"), Description: "The time when the firewall rule is created."},
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Description"), Description: "A description of the rule to help identify it."},
			{Name: "modified_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ModifiedOn"), Description: "The time when the firewall rule is updated."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Title"), Description: "Title of the resource."},

			// JSON columns
			{Name: "filter", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Filter"), Description: "A set of firewall properties."},
			{Name: "products", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Products"), Description: "A list of products to bypass for a request when the bypass action is used."},
		}),
	}
}
