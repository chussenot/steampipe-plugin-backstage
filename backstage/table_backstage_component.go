package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageComponent() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_component",
		"Components in the Backstage catalog.",
		"Component",
		[]*plugin.Column{
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the component.", Transform: specFieldTransform("owner")},
			{Name: "system", Type: proto.ColumnType_STRING, Description: "System the component belongs to.", Transform: specFieldTransform("system")},
		},
		plugin.KeyColumnSlice{
			{Name: "owner", Require: plugin.Optional},
			{Name: "system", Require: plugin.Optional},
		},
	)
}
