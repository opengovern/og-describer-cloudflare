package cloudflare

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCloudflareWorkerRoute(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_worker_route",
		Description: "Routes are basic patterns used to enable or disable workers that match requests.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListWorkerRoute,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "API item identifier tag."},
			{Name: "zone_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ZoneName"), Description: "Specifies the zone name."},
			{Name: "pattern", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Pattern"), Description: "Patterns decide what (if any) script is matched based on the URL of that request."},
			{Name: "script", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Script"), Description: "Name of the script to apply when the route is matched. The route is skipped when this is blank/missing."},
			{Name: "zone_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ZoneID"), Description: "Specifies the zone identifier."},
		}),
	}
}
