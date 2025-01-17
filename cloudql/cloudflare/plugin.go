package cloudflare

import (
	"context"
	essdk "github.com/opengovern/og-util/pkg/opengovernance-es-sdk"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name: "steampipe-plugin-cloudflare",
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: essdk.ConfigInstance,
			Schema:      essdk.ConfigSchema(),
		},
		DefaultTransform: transform.FromCamel(),
		TableMap: map[string]*plugin.Table{
			"cloudflare_access_application":    tableCloudflareAccessApplication(ctx),
			"cloudflare_access_group":          tableCloudflareAccessGroup(ctx),
			"cloudflare_access_policy":         tableCloudflareAccessPolicy(ctx),
			"cloudflare_account":               tableCloudflareAccount(ctx),
			"cloudflare_account_member":        tableCloudflareAccountMember(ctx),
			"cloudflare_account_role":          tableCloudflareAccountRole(ctx),
			"cloudflare_api_token":             tableCloudflareAPIToken(ctx),
			"cloudflare_dns_record":            tableCloudflareDNSRecord(ctx),
			"cloudflare_firewall_rule":         tableCloudflareFirewallRule(ctx),
			"cloudflare_load_balancer":         tableCloudflareLoadBalancer(ctx),
			"cloudflare_load_balancer_monitor": tableCloudflareLoadBalancerMonitor(ctx),
			"cloudflare_load_balancer_pool":    tableCloudflareLoadBalancerPool(ctx),
			"cloudflare_page_rule":             tableCloudflarePageRule(ctx),
			"cloudflare_user":                  tableCloudflareUser(ctx),
			"cloudflare_user_audit_log":        tableCloudflareUserAuditLog(ctx),
			"cloudflare_worker_route":          tableCloudflareWorkerRoute(ctx),
			"cloudflare_zone":                  tableCloudflareZone(ctx),
		},
	}
	for key, table := range p.TableMap {
		if table == nil {
			continue
		}
		if table.Get != nil && table.Get.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}
		if table.List != nil && table.List.Hydrate == nil {
			delete(p.TableMap, key)
			continue
		}

		opengovernanceTable := false
		for _, col := range table.Columns {
			if col != nil && col.Name == "platform_account_id" {
				opengovernanceTable = true
			}
		}

		if opengovernanceTable {
			if table.Get != nil {
				table.Get.KeyColumns = append(table.Get.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}

			if table.List != nil {
				table.List.KeyColumns = append(table.List.KeyColumns, plugin.OptionalColumns([]string{"platform_account_id", "platform_resource_id"})...)
			}
		}
	}
	return p
}
