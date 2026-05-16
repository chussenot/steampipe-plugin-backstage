package backstage

import (
	"context"
	"strings"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

func Plugin(ctx context.Context) *plugin.Plugin {
	p := &plugin.Plugin{
		Name:             "steampipe-plugin-backstage",
		DefaultTransform: transform.FromGo().NullIfZero(),
		DefaultGetConfig: &plugin.GetConfig{
			ShouldIgnoreError: isNotFoundError,
		},
		ConnectionConfigSchema: &plugin.ConnectionConfigSchema{
			NewInstance: ConfigInstance,
			Schema:      ConfigSchema,
		},
		TableMap: map[string]*plugin.Table{
			// Catalog entities
			"backstage_catalog_entity":    tableBackstageEntity(),
			"backstage_catalog_system":    tableBackstageSystem(),
			"backstage_catalog_domain":    tableBackstageDomain(),
			"backstage_catalog_component": tableBackstageComponent(),
			"backstage_catalog_api":       tableBackstageAPI(),
			"backstage_catalog_resource":  tableBackstageResource(),

			// Organizational entities
			"backstage_catalog_group": tableBackstageGroup(),
			"backstage_catalog_user":  tableBackstageUser(),

			// Other entities
			"backstage_catalog_template": tableBackstageTemplate(),
			"backstage_catalog_location": tableBackstageLocation(),
		},
	}

	return p
}

// isNotFoundError returns true if the error is a not found error
func isNotFoundError(err error) bool {
	return err != nil && (strings.Contains(err.Error(), "404") || strings.Contains(err.Error(), "not found"))
}
