package backstage

import (
	"context"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type backstageConfig struct {
	BaseURL  string
	ApiToken string
}

func ConfigInstance() interface{} {
	return &backstageConfig{}
}

func Plugin(ctx context.Context) *plugin.Plugin {
	return &plugin.Plugin{
		Name: "steampipe-plugin-backstage",
		TableMap: map[string]*plugin.Table{
			"backstage_entity": tableBackstageEntity(),
		},
	}
}
