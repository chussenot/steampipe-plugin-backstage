package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageSystem() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_system",
		"Systems in the Backstage catalog.",
		"System",
		[]*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain the system belongs to.", Transform: specFieldTransform("domain")},
		},
		plugin.KeyColumnSlice{
			{Name: "domain", Require: plugin.Optional},
		},
	)
}
