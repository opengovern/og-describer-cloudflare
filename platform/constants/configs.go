package constants

import "github.com/opengovern/og-util/pkg/integration"
import _ "embed"

//go:embed ui-spec.json
var UISpec []byte

//go:embed manifest.yaml
var Manifest []byte

//go:embed Setup.md
var SetupMd []byte

const (
	IntegrationName = integration.Type("github_account") // example: aws_cloud, azure_subscription
)

const (
	DescriberDeploymentName = "og-describer-github"
	DescriberRunCommand     = "/og-describer-github"
)
