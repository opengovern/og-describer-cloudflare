package cloudflare

import (
	"context"
		opengovernance "github.com/opengovern/og-describer-cloudflare/discovery/pkg/es"
	"time"

	"github.com/cloudflare/cloudflare-go"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func tableCloudflareUserAuditLog(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:             "cloudflare_user_audit_log",
		Description:      "Cloudflare User Audit Logs",
		DefaultTransform: transform.FromCamel(),
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListUserAuditLog,
		},
		Columns: commonColumns([]*plugin.Column{
			{
				Name:        "actor_email",
				Description: "Email of the actor.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActorEmail"),
			},
			{
				Name:        "actor_id",
				Description: "Unique identifier of the actor in Cloudflare’s system.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActorID"),
			},
			{
				Name:        "actor_ip",
				Description: "Physical network address of the actor.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActorIP"),
			},
			{
				Name:        "actor_type",
				Description: "Type of user that started the audit trail.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ActorType"),
			},
			{
				Name:        "id",
				Description: "Unique identifier of an audit log.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.ID"),
			},
			{
				Name:        "new_value",
				Description: "Contains the new value for the audited item.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.NewValue"),
			},
			{
				Name:        "old_value",
				Description: "Contains the old value for the audited item.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OldValue"),
			},
			{
				Name:        "owner_id",
				Description: "The identifier of the user that was acting or was acted on behalf of. If a user did the action themselves, this value will be the same as the ActorID.",
				Type:        proto.ColumnType_STRING,
				Transform:   transform.FromField("Description.OwnerID"),
			},
			{
				Name:        "when",
				Description: "When the change happened.",
				Type:        proto.ColumnType_TIMESTAMP,
				Transform:   transform.FromField("Description.When").Transform(convertAuditLogTimeToRFC3339Timestamp),
			},
			{
				Name:        "action",
				Description: "The action that was taken.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Action"),
			},
			{
				Name:        "metadata",
				Description: "Additional audit log-specific information. Metadata is organized in key:value pairs. Key and Value formats can vary by ResourceType.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Metadata"),
			},
			{
				Name:        "new_value_json",
				Description: "Contains the new value for the audited item in JSON format.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.NewValueJSON"),
			},
			{
				Name:        "old_value_json",
				Description: "Contains the old value for the audited item in JSON format.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.OldValueJSON"),
			},
			{
				Name:        "resource",
				Description: "Target resource the action was performed on.",
				Type:        proto.ColumnType_JSON,
				Transform:   transform.FromField("Description.Resource"),
			},
		}),
	}
}

//// TRANSFORM FUNCTION

func convertAuditLogTimeToRFC3339Timestamp(_ context.Context, d *transform.TransformData) (interface{}, error) {
	data := d.HydrateItem.(cloudflare.AuditLog)
	return data.When.Format(time.RFC3339), nil
}
