package describer

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/cloudflare/cloudflare-go"
	"github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
	model "github.com/opengovern/og-describer-cloudflare/discovery/provider"
)

func ListUserAuditLogs(ctx context.Context, handler *model.CloudFlareAPIHandler, stream *models.StreamSender) ([]models.Resource, error) {
	var wg sync.WaitGroup
	cloudFlareChan := make(chan models.Resource)
	go func() {
		processUserAuditLogs(ctx, handler, cloudFlareChan, &wg)
		wg.Wait()
		close(cloudFlareChan)
	}()
	var values []models.Resource
	for value := range cloudFlareChan {
		if stream != nil {
			if err := (*stream)(value); err != nil {
				return nil, err
			}
		} else {
			values = append(values, value)
		}
	}
	return values, nil
}

func processUserAuditLogs(ctx context.Context, handler *model.CloudFlareAPIHandler, cloudFlareChan chan<- models.Resource, wg *sync.WaitGroup) {
	var resp cloudflare.AuditLogResponse
	var auditLogs []cloudflare.AuditLog
	var statusCode *int
	requestFunc := func() (*int, error) {
		var e error
		opts := cloudflare.AuditLogFilter{
			Page:    page,
			PerPage: 10 * perPage,
		}
		for {
			resp, e = handler.Conn.GetUserAuditLogs(ctx, opts)
			if e != nil {
				var httpErr *cloudflare.APIRequestError
				if errors.As(e, &httpErr) {
					statusCode = &httpErr.StatusCode
				}
			}
			if len(resp.Result) == 0 {
				break
			}
			auditLogs = append(auditLogs, resp.Result...)
			opts.Page = opts.Page + 1
		}
		return statusCode, e
	}
	err := handler.DoRequest(ctx, requestFunc)
	if err != nil {
		return
	}
	for _, auditLog := range auditLogs {
		wg.Add(1)
		go func(auditLog cloudflare.AuditLog) {
			defer wg.Done()
			when := convertAuditLogTimeToRFC3339Timestamp(auditLog)
			value := models.Resource{
				ID:   auditLog.ID,
				Name: auditLog.ID,
				Description: model.UserAuditLogDescription{
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
			}
			cloudFlareChan <- value
		}(auditLog)
	}
}

func convertAuditLogTimeToRFC3339Timestamp(auditLog cloudflare.AuditLog) string {
	return auditLog.When.Format(time.RFC3339)
}
