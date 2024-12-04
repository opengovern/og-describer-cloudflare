package provider

import (
	"github.com/cloudflare/cloudflare-go"
	model "github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/configs"
	"github.com/opengovern/og-describer-cloudflare/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
	"golang.org/x/time/rate"
	"time"
)

// DescribeListByCloudFlare A wrapper to pass cloudflare authorization to describer functions
func DescribeListByCloudFlare(describe func(context.Context, *describer.CloudFlareAPIHandler, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Create cloudflare client using token or (email, api key)
		var conn *cloudflare.API
		var err error
		// Check for the token
		if cfg.Token != "" {
			conn, err = cloudflare.NewWithAPIToken(cfg.Token)
			if err != nil {
				return nil, err
			}
		}

		cloudflareAPIHandler := describer.NewCloudFlareAPIHandler(conn, cfg.AccountID, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, cloudflareAPIHandler, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeSingleByCloudFlare A wrapper to pass cloudflare authorization to describer functions
func DescribeSingleByCloudFlare(describe func(context.Context, *describer.CloudFlareAPIHandler, string) (*model.Resource, error)) model.SingleResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, resourceID string) (*model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Create cloudflare client using token or (email, api key)
		var conn *cloudflare.API
		var err error
		// Check for the token
		if cfg.Token != "" {
			conn, err = cloudflare.NewWithAPIToken(cfg.Token)
			if err != nil {
				return nil, err
			}
		}

		cloudflareAPIHandler := describer.NewCloudFlareAPIHandler(conn, cfg.AccountID, rate.Every(time.Second/4), 1, 10, 5, 5*time.Minute)

		// Get value from describer
		value, err := describe(ctx, cloudflareAPIHandler, resourceID)
		if err != nil {
			return nil, err
		}
		return value, nil
	}
}
