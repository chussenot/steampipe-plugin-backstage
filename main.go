package main

import (
	"github.com/chussenot/steampipe-plugin-backstage/backstage"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{PluginFunc: backstage.Plugin})
}
