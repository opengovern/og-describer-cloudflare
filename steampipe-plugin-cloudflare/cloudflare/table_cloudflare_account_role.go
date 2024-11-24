package cloudflare

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"

	"github.com/cloudflare/cloudflare-go"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type accountRoleInfo = struct {
	ID          string                                      `json:"id"`
	Name        string                                      `json:"name"`
	Description string                                      `json:"description"`
	Permissions map[string]cloudflare.AccountRolePermission `json:"permissions"`
	AccountID   string
}

//// TABLE DEFINITION

func tableCloudflareAccountRole(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_account_role",
		Description: "A Role defines what permissions a Member of an Account has.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccountRole,
		},
		Get: &plugin.GetConfig{
			Hydrate:    opengovernance.GetAccountRole,
			KeyColumns: plugin.AllColumns([]string{"account_id", "id"}),
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{
				Name:        "id",
				Description: "Specifies the Role identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
			},
			{
				Name:        "name",
				Description: "Specifies the name of the role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Name"),
			},
			{
				Name:        "account_id",
				Description: "Specifies the account id where the role is created at.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountID"),
			},

			// Other columns
			{
				Name:        "description",
				Description: "A description of the role.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Description"),
			},
			{
				Name:        "permissions",
				Description: "A list of permissions attached with the role.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Permissions"),
			},
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Title"),
			},
		}),
	}
}
