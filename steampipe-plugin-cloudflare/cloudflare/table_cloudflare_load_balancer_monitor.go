package cloudflare

import (
	"context"
	opengovernance "github.com/opengovern/og-describer-cloudflare/pkg/sdk/es"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"

	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableCloudflareLoadBalancerMonitor(ctx context.Context) *plugin.Table {
	return &plugin.Table{
		Name:        "cloudflare_load_balancer_monitor",
		Description: "A monitor issues health checks at regular intervals to evaluate the health of an origin pool.",
		List: &plugin.ListConfig{
			Hydrate: opengovernance.ListLoadBalancerMonitor,
		},
		Columns: commonColumns([]*plugin.Column{
			// Top columns
			{Name: "id", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ID"), Description: "Load balancer monitor API item identifier."},
			{Name: "created_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.CreatedOn"), Description: "Timestamp when the load balancer monitor was created."},
			{Name: "modified_on", Type: proto.ColumnType_TIMESTAMP, Transform: transform.FromField("Description.ModifiedOn"), Description: "Timestamp when the load balancer monitor was last modified."},
			{Name: "type", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Type"), Description: "The protocol to use for the healthcheck. Currently supported protocols are \"HTTP\", \"HTTPS\" and \"TCP\". Default: \"http\"."},

			// Other columns
			{Name: "description", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Description"), Description: "Monitor description."},
			{Name: "method", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Method"), Description: "The method to use for the health check. Valid values are any valid HTTP verb if type is \"http\" or \"https\", or connection_established if type is \"tcp\". Default: \"GET\" if type is \"http\" or \"https\", or \"connection_established\" if type is \"tcp\" ."},
			{Name: "path", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.Path"), Description: "The endpoint path to health check against. Default: \"/\". Only valid if type is \"http\" or \"https\"."},
			{Name: "header", Type: proto.ColumnType_JSON, Transform: transform.FromField("Description.Header"), Description: "The HTTP request headers to send in the health check. It is recommended you set a Host header by default. The User-Agent header cannot be overridden. Fields documented below. Only valid if type is \"http\" or \"https\"."},
			{Name: "timeout", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Timeout"), Description: "The timeout (in seconds) before marking the health check as failed. Default: 5."},
			{Name: "retries", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Retries"), Description: "The number of retries to attempt in case of a timeout before marking the origin as unhealthy. Retries are attempted immediately. Default: 2."},
			{Name: "interval", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Interval"), Description: "The interval between each health check. Shorter intervals may improve failover time, but will increase load on the origins as we check from multiple locations. Default: 60."},
			{Name: "port", Type: proto.ColumnType_INT, Transform: transform.FromField("Description.Port"), Description: "The port number to use for the healthcheck, required when creating a TCP monitor. Valid values are in the range 0-65535"},
			{Name: "expected_body", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ExpectedBody"), Description: "A case-insensitive sub-string to look for in the response body. If this string is not found, the origin will be marked as unhealthy. Only valid if type is \"http\" or \"https\". Default: \"\"."},
			{Name: "expected_codes", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ExpectedCodes"), Description: "The expected HTTP response code or code range of the health check. Eg 2xx. Only valid and required if type is \"http\" or \"https\"."},
			{Name: "follow_redirects", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.FollowRedirects"), Description: "Follow redirects if returned by the origin. Only valid if type is \"http\" or \"https\"."},
			{Name: "allow_insecure", Type: proto.ColumnType_BOOL, Transform: transform.FromField("Description.AllowInsecure"), Description: "Do not validate the certificate when monitor use HTTPS. Only valid if type is \"http\" or \"https\"."},
			{Name: "probe_zone", Type: proto.ColumnType_STRING, Transform: transform.FromField("Description.ProbeZone"), Description: "Assign this monitor to emulate the specified zone while probing. Only valid if type is \"http\" or \"https\"."},
		}),
	}
}
