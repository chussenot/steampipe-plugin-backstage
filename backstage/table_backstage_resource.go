package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageResource() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_resource",
		"Resources in the Backstage catalog.",
		"Resource",
		[]*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the resource.", Transform: specFieldTransform("type")},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the resource.", Transform: specFieldTransform("owner")},
			{Name: "system", Type: proto.ColumnType_STRING, Description: "System the resource belongs to.", Transform: specFieldTransform("system")},
		},
		plugin.KeyColumnSlice{
			{Name: "type", Require: plugin.Optional},
			{Name: "owner", Require: plugin.Optional},
			{Name: "system", Require: plugin.Optional},
		},
	)
}
