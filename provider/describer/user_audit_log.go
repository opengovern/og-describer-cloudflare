package describer

import (
	"context"
	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/model"
	"time"
)

func ListUserAuditLogs(ctx context.Context, conn *cloudflare.API, stream *models.StreamSender) ([]models.Resource, error) {
	opts := cloudflare.AuditLogFilter{
		Page:    page,
		PerPage: 10 * perPage,
	}
	var values []models.Resource
	for {
		resp, err := conn.GetUserAuditLogs(ctx, opts)
		if err != nil {
			return nil, err
		}
		if len(resp.Result) == 0 {
			break
		}
		for _, auditLog := range resp.Result {
			when := convertAuditLogTimeToRFC3339Timestamp(auditLog)
			value := models.Resource{
				ID:   auditLog.ID,
				Name: auditLog.ID,
				Description: JSONAllFieldsMarshaller{
					Value: model.UserAuditLogDescription{
						ActorEmail:   auditLog.Actor.Email,
						ActorID:      auditLog.Actor.ID,
						ActorIP:      auditLog.Actor.IP,
						ActorType:    auditLog.Actor.Type,
						ID:           auditLog.ID,
						NewValue:     auditLog.NewValue,
						OldValue:     auditLog.OldValue,
						OwnerID:      auditLog.Owner.ID,
						When:         when,
						Action:       auditLog.Action,
						Metadata:     auditLog.Metadata,
						NewValueJSON: auditLog.NewValueJSON,
						OldValueJSON: auditLog.OldValueJSON,
						Resource:     auditLog.Resource,
					},
				},
			}
			if stream != nil {
				if err := (*stream)(value); err != nil {
					return nil, err
				}
			} else {
				values = append(values, value)
			}
		}
		opts.Page = opts.Page + 1
	}
	return values, nil
}

func convertAuditLogTimeToRFC3339Timestamp(auditLog cloudflare.AuditLog) string {
	return auditLog.When.Format(time.RFC3339)
}
