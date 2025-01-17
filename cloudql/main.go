package main

import (
	"github.com/opengovern/og-describer-cloudflare/cloudql/cloudflare"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: cloudflare.Plugin})
}
