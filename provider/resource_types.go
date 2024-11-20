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
		ListDescriber: DescribeByCloudFlareList(describer.ListAccessApplications),
		GetDescriber:  nil,
	},

	"CloudFlare/Access/Group": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Access/Group",
		Tags: map[string][]string{
			"category": {"Access"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByCloudFlareList(describer.ListAccessGroups),
		GetDescriber:  nil,
	},

	"CloudFlare/Access/Policy": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Access/Policy",
		Tags: map[string][]string{
			"category": {"Access"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByCloudFlareList(describer.ListAccessPolicies),
		GetDescriber:  nil,
	},

	"CloudFlare/Account": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Account",
		Tags: map[string][]string{
			"category": {"Account"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: nil,
		//GetDescriber:         DescribeByCloudFlareGet(describer.GetAccount),
		GetDescriber: nil,
	},

	"CloudFlare/DNSRecord": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/DNSRecord",
		Tags: map[string][]string{
			"category": {"DNSRecord"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByCloudFlareList(describer.ListDNSRecords),
		//GetDescriber:         DescribeByCloudFlareGet(describer.GetDNSRecord),
		GetDescriber: nil,
	},

	"CloudFlare/Firewall/Rule": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/Firewall/Rule",
		Tags: map[string][]string{
			"category": {"Firewall"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByCloudFlareList(describer.ListFireWallRules),
		//GetDescriber:         DescribeByCloudFlareGet(describer.GetFireWallRule),
		GetDescriber: nil,
	},

	"CloudFlare/LoadBalancer": {
		IntegrationType: configs.IntegrationName,
		ResourceName:    "CloudFlare/LoadBalancer",
		Tags: map[string][]string{
			"category": {"LoadBalancer"},
		},
		Labels:        map[string]string{},
		Annotations:   map[string]string{},
		ListDescriber: DescribeByCloudFlareList(describer.ListLoadBalancers),
		//GetDescriber:         DescribeByCloudFlareGet(describer.GetLoadBalancer),
		GetDescriber: nil,
	},
}
