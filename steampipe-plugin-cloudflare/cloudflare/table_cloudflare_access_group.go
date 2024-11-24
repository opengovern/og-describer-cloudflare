package cloudflare

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
)

func tableCloudflareAccessGroup(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_access_group",
		Description: "Access Groups allows to define a set of users to which an application policy can be applied.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccessGroup,
		},
		// Get Config - Currently SDK is not supporting get call
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "Identifier of the access group."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "Friendly name of the access group."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AccountID"), Description: "ID of the account, access group belongs."},
			{Name: "account_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AccountName"), Description: "Name of the account, access group belongs."},

			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "Timestamp when access group was created."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "TImestamp when access group was last modified."},

			// JSON columns
			{Name: "exclude", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Exclude"), Description: "The exclude policy works like a NOT logical operator. The user must not satisfy all of the rules in exclude."},
			{Name: "include", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Include"), Description: "The include policy works like an OR logical operator. The user must satisfy one of the rules in includes."},
			{Name: "require", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Require"), Description: "The require policy works like a AND logical operator. The user must satisfy all of the rules in require."},
		}),
	}
}
