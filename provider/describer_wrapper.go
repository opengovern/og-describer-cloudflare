package provider

import (
	"github.com/cloudflare/cloudflare-go"
	model "github.com/opengovern/og-describer-template/pkg/sdk/models"
	"github.com/opengovern/og-describer-template/provider/configs"
	"github.com/opengovern/og-describer-template/provider/describer"
	"github.com/opengovern/og-util/pkg/describe/enums"
	"golang.org/x/net/context"
)

// DescribeByIntegration TODO: implement a wrapper to pass integration authorization to describer functions
func DescribeByCloudFlare(describe func(context.Context, *cloudflare.API, *model.StreamSender) ([]model.Resource, error)) model.ResourceDescriber {
	return func(ctx context.Context, cfg configs.IntegrationCredentials, triggerType enums.DescribeTriggerType, additionalParameters map[string]string, stream *model.StreamSender) ([]model.Resource, error) {
		ctx = describer.WithTriggerType(ctx, triggerType)

		// Create cloudflare client using token or (email, api key)
		var conn *cloudflare.API
		var err error
		// First: check for the token
		if cfg.Token != nil {
			conn, err = cloudflare.NewWithAPIToken(*cfg.Token)
			if err != nil {
				return nil, err
			}
		}
		// Second: Email + API Key
		if cfg.Email != nil && cfg.APIKey != nil {
			conn, err = cloudflare.New(*cfg.APIKey, *cfg.Email)
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
