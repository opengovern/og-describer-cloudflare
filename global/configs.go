package global

import "github.com/opengovern/og-util/pkg/integration"


const (
	IntegrationTypeLower             = "cloudflare"                                    // example: aws, azure
	IntegrationName = integration.Type("cloudflare_account")          // example: AWS_ACCOUNT, AZURE_SUBSCRIPTION
	OGPluginRepoURL                  = "github.com/opengovern/og-describer-cloudflare" // example: github.com/opengovern/og-describer-aws
)


type IntegrationCredentials struct {
	Token     string `json:"token"`
	AccountID string `json:"account_id"`
}
