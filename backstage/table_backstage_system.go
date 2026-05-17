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
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the system.", Transform: specFieldTransform("owner")},
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain the system belongs to.", Transform: specFieldTransform("domain")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the system.", Transform: specFieldTransform("type")},
		},
		plugin.KeyColumnSlice{
			{Name: "owner", Require: plugin.Optional},
			{Name: "domain", Require: plugin.Optional},
			{Name: "type", Require: plugin.Optional},
		},
	)
}
