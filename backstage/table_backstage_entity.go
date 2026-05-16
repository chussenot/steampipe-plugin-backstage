package backstage

import "github.com/turbot/steampipe-plugin-sdk/v5/plugin"

func tableBackstageEntity() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_entity",
		"Generic entities in the Backstage catalog.",
		"",
		nil,
		nil,
	)
}
