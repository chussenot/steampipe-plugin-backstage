package backstage

import "github.com/turbot/steampipe-plugin-sdk/v5/plugin"

func tableBackstageEntity() *plugin.Table {
	// Define the table schema and other properties here
	return &plugin.Table{
		Name:        "backstage_entity",
		Description: "Backstage entity table",
		// Add other table properties here
	}
}
