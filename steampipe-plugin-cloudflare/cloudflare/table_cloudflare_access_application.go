package cloudflare

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
)

func tableCloudflareAccessApplication(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_access_application",
		Description: "Access Applications are used to restrict access to a whole application using an authorisation gateway managed by Cloudflare.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccessApplication,
		},
		// Get Config - Currently SDK is not supporting get call
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "Application API uuid."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "Friendly name of the access application."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AccountID"), Description: "ID of the account, access application belongs."},
			{Name: "account_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AccountName"), Description: "Name of the account, access application belongs."},
			{Name: "domain", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Domain"), Description: "The domain and path that access will block."},
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "Timestamp when the application was created."},

			// Other columns
			{Name: "aud", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AUD"), Description: "Audience tag."},
			{Name: "auto_redirect_to_identity", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.AutoRedirectToIdentity"), Description: "Option to skip identity provider selection if only one is configured in allowed_idps. Defaults to false (disabled)."},
			{Name: "custom_deny_message", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CustomDenyMessage"), Description: "Option that returns a custom error message when a user is denied access to the application."},
			{Name: "custom_deny_url", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.CustomDenyURL"), Description: "Option that redirects to a custom URL when a user is denied access to the application."},
			{Name: "enable_binding_cookie", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.EnableBindingCookie"), Description: "Option to provide increased security against compromised authorization tokens and CSRF attacks by requiring an additional \"binding\" cookie on requests. Defaults to false."},
			{Name: "session_duration", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.SessionDuration"), Description: "How often a user will be forced to re-authorise. Must be in the format \"48h\" or \"2h45m\". Valid time units are ns, us (or µs), ms, s, m, h. Defaults to 24h."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "Timestamp when the application was last modified."},

			// JSON columns
			{Name: "allowed_idps", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.AllowedIDPs"), Description: "The identity providers selected for the application."},
			{Name: "cors_headers", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.CORSHeaders"), Description: "CORS configuration for the access application. See below for reference structure."},
		}),
	}
}
