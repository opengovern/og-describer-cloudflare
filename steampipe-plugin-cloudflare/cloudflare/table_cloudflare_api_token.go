package cloudflare

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCloudflareAPIToken(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_api_token",
		Description: "API tokens for the user.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListApiToken,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "ID of the API token."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "Name of the API token."},
			{Name: "status", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Status"), Description: "Status of the API token."},

			// Other columns
			{Name: "condition", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Condition"), Description: "Conditions (e.g. client IP ranges) associated with the API token."},
			{Name: "expires_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ExpiresOn"), Description: "When the API token expires."},
			{Name: "issued_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.IssuedOn"), Description: "When the API token was issued."},
			{Name: "modified_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ModifiedOn"), Description: "When the API token was last modified."},
			{Name: "not_before", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.NotBefore"), Description: "When the API token becomes valid."},

			// JSON columns
			{Name: "policies", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Policies"), Description: "Policies associated with this API token."},
		}),
	}
}
