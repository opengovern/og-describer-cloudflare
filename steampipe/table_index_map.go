package steampipe

import (
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
)

var Map = map[string]string{
  "CloudFlare/Access/Application": "cloudflare_access_application",
  "CloudFlare/Access/Group": "cloudflare_access_group",
  "CloudFlare/Access/Policy": "cloudflare_access_policy",
  "CloudFlare/Account": "cloudflare_account",
  "CloudFlare/DNSRecord": "cloudflare_dns_record",
  "CloudFlare/Firewall/Rule": "cloudflare_firewall_rule",
  "CloudFlare/LoadBalancer": "cloudflare_load_balancer",
}

var DescriptionMap = map[string]interface{}{
  "CloudFlare/Access/Application": opengovernance.AccessApplication{},
  "CloudFlare/Access/Group": opengovernance.AccessGroup{},
  "CloudFlare/Access/Policy": opengovernance.AccessPolicy{},
  "CloudFlare/Account": opengovernance.Account{},
  "CloudFlare/DNSRecord": opengovernance.DNSRecord{},
  "CloudFlare/Firewall/Rule": opengovernance.FireWallRule{},
  "CloudFlare/LoadBalancer": opengovernance.LoadBalancer{},
}

var ReverseMap = map[string]string{
  "cloudflare_access_application": "CloudFlare/Access/Application",
  "cloudflare_access_group": "CloudFlare/Access/Group",
  "cloudflare_access_policy": "CloudFlare/Access/Policy",
  "cloudflare_account": "CloudFlare/Account",
  "cloudflare_dns_record": "CloudFlare/DNSRecord",
  "cloudflare_firewall_rule": "CloudFlare/Firewall/Rule",
  "cloudflare_load_balancer": "CloudFlare/LoadBalancer",
}
