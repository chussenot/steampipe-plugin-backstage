package backstage

import "github.com/turbot/steampipe-plugin-sdk/v5/plugin"

func tableBackstageResource() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_resource",
		"Resources in the Backstage catalog.",
		"Resource",
		nil,
		nil,
	)
}
