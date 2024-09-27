package main

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

func pluginFunc() *plugin.Plugin {
	return &plugin.Plugin{
		Name: "steampipe-plugin-backstage",
		ConnectionConfigSchema: &schema.ConnectionConfigSchema{
			NewInstance: func() interface{} { return &backstageConfig{} },
			Schema: map[string]*schema.Schema{
				"base_url": {
					Type:     schema.TypeString,
					Required: true,
				},
				"api_token": {
					Type:     schema.TypeString,
					Required: true,
				},
			},
		},
		Tables: []*plugin.Table{
			tableBackstageCatalogEntities(),
		},
	}
}
