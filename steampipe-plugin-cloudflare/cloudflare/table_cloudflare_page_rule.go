package cloudflare

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCloudflarePageRule(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_page_rule",
		Description: "Page Rules gives the ability to control how Cloudflare works on a URL or subdomain basis.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListPageRule,
		},
		Get: &plugin.GetConfig{
			KeyColumns: plugin.AllColumns([]string{"zone_id", "id"}),
			Hydrate:    opengovernance.GetPageRule,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "Specifies the Page Rule identifier."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Status"), Description: "Specifies the status of the page rule."},

			// Other columns
			{Name: "zone_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ZoneID"), Description: "Specifies the zone identifier."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedOn"), Description: "The time when the page rule is created."},
			{Name: "modified_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ModifiedOn"), Description: "The time when the page rule was last modified."},
			{Name: "priority", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Priority"), Description: "A number that indicates the preference for a page rule over another."},
			{Name: "title", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Title"), Description: "Title of the resource."},

			// JSON columns
			{Name: "actions", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Actions"), Description: "A list of actions to perform if the targets of this rule match the request. Actions can redirect the url to another url or override settings (but not both)."},
			{Name: "targets", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Targets"), Description: "A list of targets to evaluate on a request."},
		}),
	}
}
