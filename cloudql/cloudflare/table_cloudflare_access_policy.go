package cloudflare

import (
	"context"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

		opengovernance "github.com/opengovern/og-describer-cloudflare/discovery/pkg/es"
)

func tableCloudflareAccessPolicy(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_access_policy",
		Description: "Access Policies define the users or groups who can, or cannot, reach the Application Resource.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListAccessPolicy,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "Access policy unique API identifier."},
			{Name: "name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Name"), Description: "The name of the policy. Only used in the UI."},
			{Name: "application_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ApplicationID"), Description: "The id of application to which policy belongs."},
			{Name: "application_name", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ApplicationName"), Description: "The name of application to which policy belongs."},
			{Name: "account_id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.AccountID"), Description: "The ID of account where application belongs."},

			// Other columns
			{Name: "created_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedAt"), Description: "Timestamp when access policy was created."},
			{Name: "decision", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Decision"), Description: "Defines the action Access will take if the policy matches the user. Allowed values: allow, deny, non_identity, bypass"},
			{Name: "precedence", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Precedence"), Description: "The unique precedence for policies on a single application.Precedence"},
			{Name: "purpose_justification_prompt", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.PurposeJustificationPrompt"), Description: "The text the user will be prompted with when a purpose justification is required."},
			{Name: "purpose_justification_required", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.PurposeJustificationRequired"), Description: "Defines whether or not the user is prompted for a justification when this policy is applied."},
			{Name: "updated_at", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.UpdatedAt"), Description: "Timestamp when access policy was last modified."},

			// JSON columns
			{Name: "approval_groups", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.ApprovalGroups"), Description: "The list of approval groups that must approve the access request."},
			{Name: "exclude", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Exclude"), Description: "The exclude policy works like a NOT logical operator. The user must not satisfy all of the rules in exclude."},
			{Name: "include", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Include"), Description: "The include policy works like an OR logical operator. The user must satisfy one of the rules in includes."},
			{Name: "require", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Require"), Description: "The require policy works like a AND logical operator. The user must satisfy all of the rules in require."},
		}),
	}
}
