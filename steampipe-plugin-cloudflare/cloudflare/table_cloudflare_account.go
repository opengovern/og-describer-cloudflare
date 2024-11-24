package cloudflare

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"

	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
)

func tableCloudflareAccount(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_account",
		Description: "Accounts the user has access to.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccount,
		},
		Get: &plugin.GetConfig{
			Hydrate:    opengovernance.GetAccount,
			KeyColumns: plugin.SingleColumn("id"),
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "ID of the account."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "Name of the account."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "Type of the account."},

			// JSON columns
			{Name: "settings", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Settings"), Description: "Settings for the account."},
		}),
	}
}
