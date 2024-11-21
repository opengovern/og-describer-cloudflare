package steampipe

import (
	"github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
)

var Map = map[string]string{
  "CloudFlare/Access/Application": "cloudflare_access_application",
  "CloudFlare/Access/Group": "cloudflare_access_group",
  "CloudFlare/Access/Policy": "cloudflare_access_policy",
  "CloudFlare/Account": "cloudflare_account",
  "CloudFlare/Account/Member": "cloudflare_account_member",
  "CloudFlare/Account/Role": "cloudflare_account_role",
  "CloudFlare/ApiToken": "cloudflare_api_token",
  "CloudFlare/DNSRecord": "cloudflare_dns_record",
  "CloudFlare/Firewall/Rule": "cloudflare_firewall_rule",
  "CloudFlare/LoadBalancer": "cloudflare_load_balancer",
  "CloudFlare/LoadBalancer/Monitor": "cloudflare_load_balancer_monitor",
  "CloudFlare/LoadBalancer/Pool": "cloudflare_load_balancer_pool",
  "CloudFlare/PageRule": "cloudflare_page_rule",
  "CloudFlare/User": "cloudflare_user",
  "CloudFlare/User/AuditLog": "cloudflare_user_audit_log",
  "CloudFlare/WorkerRoute": "cloudflare_worker_route",
  "CloudFlare/Zone": "cloudflare_zone",
}

var DescriptionMap = map[string]interface{}{
  "CloudFlare/Access/Application": opengovernance.AccessApplication{},
  "CloudFlare/Access/Group": opengovernance.AccessGroup{},
  "CloudFlare/Access/Policy": opengovernance.AccessPolicy{},
  "CloudFlare/Account": opengovernance.Account{},
  "CloudFlare/Account/Member": opengovernance.AccountMember{},
  "CloudFlare/Account/Role": opengovernance.AccountRole{},
  "CloudFlare/ApiToken": opengovernance.ApiToken{},
  "CloudFlare/DNSRecord": opengovernance.DNSRecord{},
  "CloudFlare/Firewall/Rule": opengovernance.FireWallRule{},
  "CloudFlare/LoadBalancer": opengovernance.LoadBalancer{},
  "CloudFlare/LoadBalancer/Monitor": opengovernance.LoadBalancerMonitor{},
  "CloudFlare/LoadBalancer/Pool": opengovernance.LoadBalancerPool{},
  "CloudFlare/PageRule": opengovernance.PageRule{},
  "CloudFlare/User": opengovernance.User{},
  "CloudFlare/User/AuditLog": opengovernance.UserAuditLog{},
  "CloudFlare/WorkerRoute": opengovernance.WorkerRoute{},
  "CloudFlare/Zone": opengovernance.Zone{},
}

var ReverseMap = map[string]string{
  "cloudflare_access_application": "CloudFlare/Access/Application",
  "cloudflare_access_group": "CloudFlare/Access/Group",
  "cloudflare_access_policy": "CloudFlare/Access/Policy",
  "cloudflare_account": "CloudFlare/Account",
  "cloudflare_account_member": "CloudFlare/Account/Member",
  "cloudflare_account_role": "CloudFlare/Account/Role",
  "cloudflare_api_token": "CloudFlare/ApiToken",
  "cloudflare_dns_record": "CloudFlare/DNSRecord",
  "cloudflare_firewall_rule": "CloudFlare/Firewall/Rule",
  "cloudflare_load_balancer": "CloudFlare/LoadBalancer",
  "cloudflare_load_balancer_monitor": "CloudFlare/LoadBalancer/Monitor",
  "cloudflare_load_balancer_pool": "CloudFlare/LoadBalancer/Pool",
  "cloudflare_page_rule": "CloudFlare/PageRule",
  "cloudflare_user": "CloudFlare/User",
  "cloudflare_user_audit_log": "CloudFlare/User/AuditLog",
  "cloudflare_worker_route": "CloudFlare/WorkerRoute",
  "cloudflare_zone": "CloudFlare/Zone",
}
