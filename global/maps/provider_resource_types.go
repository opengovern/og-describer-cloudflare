package maps
import (
	
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{}

var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{
}


var ResourceTypesList = []string{
  "Github/Actions/Artifact",
  "Github/Actions/Runner",
  "Github/Actions/Secret",
  "Github/Actions/WorkflowRun",
  "Github/Branch",
  "Github/Branch/Protection",
  "Github/Commit",
  "Github/Issue",
  "Github/License",
  "Github/Organization",
  "Github/Organization/Collaborator",
  "Github/Organization/Dependabot/Alert",
  "Github/Organization/Member",
  "Github/Organization/Team",
  "Github/PullRequest",
  "Github/Release",
  "Github/Repository",
  "Github/Repository/Collaborator",
  "Github/Repository/DependabotAlert",
  "Github/Repository/Deployment",
  "Github/Repository/Environment",
  "Github/Repository/Ruleset",
  "Github/Repository/SBOM",
  "Github/Repository/VulnerabilityAlert",
  "Github/Tag",
  "Github/Team/Member",
  "Github/User",
  "Github/Workflow",
  "Github/Container/Package",
  "Github/Package/Maven",
  "Github/NPM/Package",
  "Github/Nuget/Package",
  "Github/Artifact/DockerFile",
}