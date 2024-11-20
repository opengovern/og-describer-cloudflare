package provider

import (
	"github.com/cloudflare/cloudflare-go"
	model "github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/configs"
	"github.com/opengovern/og-describer-cloudflare/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
)

// DescribeByCloudFlareList A wrapper to pass cloudflare authorization to describer functions
func DescribeByCloudFlareList(describe func(context.Context, *cloudflare.API, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Create cloudflare client using token or (email, api key)
		var conn *cloudflare.API
		var err error
		// First: check for the token
		if cfg.Token != "" {
			conn, err = cloudflare.NewWithAPIToken(cfg.Token)
			if err != nil {
				return nil, err
			}
		}
		// Second: Email + API Key
		if cfg.Email != "" && cfg.APIKey != "" {
			conn, err = cloudflare.New(cfg.APIKey, cfg.Email)
			if err != nil {
				return nil, err
			}
		}

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, conn, stream)
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}

// DescribeByCloudFlareGet A wrapper to pass cloudflare authorization to describer functions
func DescribeByCloudFlareGet(describe func(context.Context, *cloudflare.API, string) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Create cloudflare client using token or (email, api key)
		var conn *cloudflare.API
		var err error
		// First: check for the token
		if cfg.Token != "" {
			conn, err = cloudflare.NewWithAPIToken(cfg.Token)
			if err != nil {
				return nil, err
			}
		}
		// Second: Email + API Key
		if cfg.Email != "" && cfg.APIKey != "" {
			conn, err = cloudflare.New(cfg.APIKey, cfg.Email)
			if err != nil {
				return nil, err
			}
		}

		// Get values from describer
		var values []model.Resource
		result, err := describe(ctx, conn, "")
		if err != nil {
			return nil, err
		}
		values = append(values, result...)
		return values, nil
	}
}
