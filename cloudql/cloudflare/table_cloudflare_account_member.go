package cloudflare

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
		opengovernance "github.com/opengovern/og-describer-cloudflare/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

type accountMemberInfo = struct {
	ID        string                              `json:"id"`
	Code      string                              `json:"code"`
	User      cloudflare.AccountMemberUserDetails `json:"user"`
	Status    string                              `json:"status"`
	Roles     []cloudflare.AccountRole            `json:"roles"`
	AccountID string
}

//// TABLE DEFINITION

func tableCloudflareAccountMember(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_account_member",
		Description: "Cloudflare Account Member",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccountMember,
		},
		Get: &plugin.GetConfig{
			Hydrate:    opengovernance.GetAccountMember,
			KeyColumns: plugin.AllColumns([]string{"account_id", "id"}),
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "user_email",
				Description: "Specifies the user email.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.UserEmail"),
			},
			{
				Name:        "id",
				Description: "Specifies the account membership identifier.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
			},
			{
				Name:        "status",
				Description: "A member's status in the account.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Status"),
			},
			{
				Name:        "account_id",
				Description: "Specifies the account id, the member is associated with.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.AccountID"),
			},
			{
				Name:        "code",
				Description: "The unique activation code for the account membership.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Code"),
			},
			{
				Name:        "user",
				Description: "A set of information about the user.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.User"),
			},
			{
				Name:        "roles",
				Description: "A list of permissions that a Member of an Account has.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Roles"),
			},

			// steampipe standard columns
			{
				Name:        "title",
				Description: "Title of the resource.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.Title"),
			},
		}),
	}
}
