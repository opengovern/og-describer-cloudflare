package maps
import (
	"github.com/opengovern/og-describer-cloudflare/discovery/describers"
	"github.com/opengovern/og-describer-cloudflare/discovery/provider"
	"github.com/opengovern/og-describer-cloudflare/platform/constants"
	"github.com/opengovern/og-util/pkg/integration/interfaces"
	model "github.com/opengovern/og-describer-cloudflare/discovery/pkg/models"
)
var ResourceTypes = map[string]model.ResourceType{

	"CloudFlare/Access/Application": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Access/Application",
		Tags:                 map[string][]string{
            "category": {"Access"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListAccessApplications),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetAccessApplication),
	},

	"CloudFlare/Access/Group": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Access/Group",
		Tags:                 map[string][]string{
            "category": {"Access"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListAccessGroups),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetAccessGroup),
	},

	"CloudFlare/Access/Policy": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Access/Policy",
		Tags:                 map[string][]string{
            "category": {"Access"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListAccessPolicies),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetAccessPolicy),
	},

	"CloudFlare/Account": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Account",
		Tags:                 map[string][]string{
            "category": {"Account"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListAccounts),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetAccount),
	},

	"CloudFlare/Account/Member": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Account/Member",
		Tags:                 map[string][]string{
            "category": {"Account"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListAccountMembers),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetAccountMember),
	},

	"CloudFlare/Account/Role": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Account/Role",
		Tags:                 map[string][]string{
            "category": {"Account"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListAccountRoles),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetAccountRole),
	},

	"CloudFlare/ApiToken": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/ApiToken",
		Tags:                 map[string][]string{
            "category": {"ApiToken"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListAPITokens),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetAPIToken),
	},

	"CloudFlare/DNSRecord": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/DNSRecord",
		Tags:                 map[string][]string{
            "category": {"DNSRecord"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListDNSRecords),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetDNSRecord),
	},

	"CloudFlare/Firewall/Rule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Firewall/Rule",
		Tags:                 map[string][]string{
            "category": {"Firewall"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListFireWallRules),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetFireWallRule),
	},

	"CloudFlare/LoadBalancer": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/LoadBalancer",
		Tags:                 map[string][]string{
            "category": {"LoadBalancer"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListLoadBalancers),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetLoadBalancer),
	},

	"CloudFlare/LoadBalancer/Monitor": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/LoadBalancer/Monitor",
		Tags:                 map[string][]string{
            "category": {"LoadBalancer"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListLoadBalancerMonitors),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetLoadBalancerMonitor),
	},

	"CloudFlare/LoadBalancer/Pool": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/LoadBalancer/Pool",
		Tags:                 map[string][]string{
            "category": {"LoadBalancer"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListLoadBalancerPools),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetLoadBalancerPool),
	},

	"CloudFlare/PageRule": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/PageRule",
		Tags:                 map[string][]string{
            "category": {"PageRule"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListPageRules),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetPageRule),
	},

	"CloudFlare/User": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/User",
		Tags:                 map[string][]string{
            "category": {"User"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListUsers),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetUser),
	},

	"CloudFlare/User/AuditLog": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/User/AuditLog",
		Tags:                 map[string][]string{
            "category": {"User"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListUserAuditLogs),
		GetDescriber:         nil,
	},

	"CloudFlare/WorkerRoute": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/WorkerRoute",
		Tags:                 map[string][]string{
            "category": {"WorkerRoute"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListWorkerRoutes),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetWorkerRoute),
	},

	"CloudFlare/Zone": {
		IntegrationType:      constants.IntegrationName,
		ResourceName:         "CloudFlare/Zone",
		Tags:                 map[string][]string{
            "category": {"Zone"},
        },
		Labels:               map[string]string{
        },
		Annotations:          map[string]string{
        },
		ListDescriber:        provider.DescribeListByCloudFlare(describers.ListZones),
		GetDescriber:         provider.DescribeSingleByCloudFlare(describers.GetZone),
	},
}


var ResourceTypeConfigs = map[string]*interfaces.ResourceTypeConfiguration{

	"CloudFlare/Access/Application": {
		Name:         "CloudFlare/Access/Application",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/Access/Group": {
		Name:         "CloudFlare/Access/Group",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/Access/Policy": {
		Name:         "CloudFlare/Access/Policy",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/Account": {
		Name:         "CloudFlare/Account",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/Account/Member": {
		Name:         "CloudFlare/Account/Member",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/Account/Role": {
		Name:         "CloudFlare/Account/Role",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/ApiToken": {
		Name:         "CloudFlare/ApiToken",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/DNSRecord": {
		Name:         "CloudFlare/DNSRecord",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/Firewall/Rule": {
		Name:         "CloudFlare/Firewall/Rule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/LoadBalancer": {
		Name:         "CloudFlare/LoadBalancer",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/LoadBalancer/Monitor": {
		Name:         "CloudFlare/LoadBalancer/Monitor",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/LoadBalancer/Pool": {
		Name:         "CloudFlare/LoadBalancer/Pool",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/PageRule": {
		Name:         "CloudFlare/PageRule",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/User": {
		Name:         "CloudFlare/User",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/User/AuditLog": {
		Name:         "CloudFlare/User/AuditLog",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/WorkerRoute": {
		Name:         "CloudFlare/WorkerRoute",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},

	"CloudFlare/Zone": {
		Name:         "CloudFlare/Zone",
		IntegrationType:      constants.IntegrationName,
		Description:                 "",
		
	},
}


var ResourceTypesList = []string{
  "CloudFlare/Access/Application",
  "CloudFlare/Access/Group",
  "CloudFlare/Access/Policy",
  "CloudFlare/Account",
  "CloudFlare/Account/Member",
  "CloudFlare/Account/Role",
  "CloudFlare/ApiToken",
  "CloudFlare/DNSRecord",
  "CloudFlare/Firewall/Rule",
  "CloudFlare/LoadBalancer",
  "CloudFlare/LoadBalancer/Monitor",
  "CloudFlare/LoadBalancer/Pool",
  "CloudFlare/PageRule",
  "CloudFlare/User",
  "CloudFlare/User/AuditLog",
  "CloudFlare/WorkerRoute",
  "CloudFlare/Zone",
}