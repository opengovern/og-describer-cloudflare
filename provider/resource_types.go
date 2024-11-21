package provider

import (
	model "github.com/opengovern/og-describer-cloudflare/pkg/sdk/models"
	"github.com/opengovern/og-describer-cloudflare/provider/configs"
	"github.com/opengovern/og-describer-cloudflare/provider/describer"
)

var ResourceTypes = map[string]model.ResourceType{

	"CloudFlare/Access/Application": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Access/Application",
		Tags: map[string][]string{
			"category": {"Access"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListAccessApplications),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetAccessApplication),
	},

	"CloudFlare/Access/Group": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Access/Group",
		Tags: map[string][]string{
			"category": {"Access"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListAccessGroups),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetAccessGroup),
	},

	"CloudFlare/Access/Policy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Access/Policy",
		Tags: map[string][]string{
			"category": {"Access"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListAccessPolicies),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetAccessPolicy),
	},

	"CloudFlare/Account": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Account",
		Tags: map[string][]string{
			"category": {"Account"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListAccounts),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetAccount),
	},

	"CloudFlare/Account/Member": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Account/Member",
		Tags: map[string][]string{
			"category": {"Account"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListAccountMembers),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetAccountMember),
	},

	"CloudFlare/Account/Role": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Account/Role",
		Tags: map[string][]string{
			"category": {"Account"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListAccountRoles),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetAccountRole),
	},

	"CloudFlare/ApiToken": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/ApiToken",
		Tags: map[string][]string{
			"category": {"ApiToken"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: nil,
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetApiToken),
	},

	"CloudFlare/DNSRecord": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/DNSRecord",
		Tags: map[string][]string{
			"category": {"DNSRecord"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListDNSRecords),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetDNSRecord),
	},

	"CloudFlare/Firewall/Rule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Firewall/Rule",
		Tags: map[string][]string{
			"category": {"Firewall"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListFireWallRules),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetFireWallRule),
	},

	"CloudFlare/LoadBalancer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/LoadBalancer",
		Tags: map[string][]string{
			"category": {"LoadBalancer"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListLoadBalancers),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetLoadBalancer),
	},

	"CloudFlare/LoadBalancer/Monitor": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/LoadBalancer/Monitor",
		Tags: map[string][]string{
			"category": {"LoadBalancer"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListLoadBalancerMonitors),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetLoadBalancerMonitor),
	},

	"CloudFlare/LoadBalancer/Pool": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/LoadBalancer/Pool",
		Tags: map[string][]string{
			"category": {"LoadBalancer"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListLoadBalancerPools),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetLoadBalancerPool),
	},

	"CloudFlare/PageRule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/PageRule",
		Tags: map[string][]string{
			"category": {"PageRule"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListPageRules),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetPageRule),
	},

	"CloudFlare/User": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/User",
		Tags: map[string][]string{
			"category": {"User"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListUsers),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetUser),
	},

	"CloudFlare/User/AuditLog": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/User/AuditLog",
		Tags: map[string][]string{
			"category": {"User"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListUserAuditLogs),
		GetDescriber:  nil,
	},

	"CloudFlare/WorkerRoute": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/WorkerRoute",
		Tags: map[string][]string{
			"category": {"WorkerRoute"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListWorkerRoutes),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetWorkerRoute),
	},

	"CloudFlare/Zone": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Zone",
		Tags: map[string][]string{
			"category": {"Zone"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeListByCloudFlare(describer.ListZones),
		GetDescriber:  DescribeSingleByCloudFlare(describer.GetZone),
	},
}
