package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageGroup() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_group",
		"Groups in the Backstage catalog.",
		"Group",
		[]*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the group.", Transform: specFieldTransform("type")},
		},
		plugin.KeyColumnSlice{
			{Name: "type", Require: plugin.Optional},
		},
	)
}
