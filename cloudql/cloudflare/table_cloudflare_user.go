package cloudflare

import (
	"context"
		opengovernance "github.com/opengovern/og-describer-cloudflare/discovery/pkg/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCloudflareUser(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_user",
		Description: "Information about the current user making the request.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUser,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "ID of the user."},
			{Name: "email", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Email"), Description: "Email of the user."},
			{Name: "username", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Username"), Description: "Username (actually often in ID style) of the user."},

			// Other columns
			{Name: "telephone", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Telephone"), Description: "Telephone number of the user."},
			{Name: "first_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.FirstName"), Description: "First name of the user."},
			{Name: "last_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.LastName"), Description: "Last name of the user."},
			{Name: "country", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Country"), Description: "Country of the user."},
			{Name: "zipcode", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Zipcode"), Description: "Zipcode of the user."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedOn"), Description: "When the user was created."},
			{Name: "modified_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ModifiedOn"), Description: "When the user was last modified."},
			{Name: "api_key", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.APIKey"), Description: "API Key for the user."},
			{Name: "two_factor_authentication_enabled", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.TwoFactorAuthenticationEnabled"), Description: "True if two factor authentication is enabled for this user."},

			// JSON columns
			{Name: "betas", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Betas"), Description: "Beta feature flags associated with the user."},
			{Name: "organizations", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Organizations"), Description: "Organizations the user is a member of."},
		}),
	}
}
