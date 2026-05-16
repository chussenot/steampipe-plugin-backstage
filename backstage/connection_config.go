package backstage

import (
	"os"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/schema"
)

type BackstageConfig struct {
	Host  *string `cty:"host"`
	Token *string `cty:"token"`
}

var ConfigSchema = map[string]*schema.Attribute{
	"host": {
		Type:     schema.TypeString,
		Required: true,
	},
	"token": {
		Type: schema.TypeString,
	},
}

func ConfigInstance() interface{} {
	return &BackstageConfig{}
}

func GetConfig(connection *plugin.Connection) BackstageConfig {
	if connection == nil || connection.Config == nil {
		return BackstageConfig{}
	}
	config, _ := connection.Config.(BackstageConfig)

	// Environment variables override connection config
	if host := os.Getenv("BACKSTAGE_HOST"); host != "" {
		config.Host = &host
	}
	if token := os.Getenv("BACKSTAGE_TOKEN"); token != "" {
		config.Token = &token
	}

	return config
}
