package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageTemplate() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_template",
		"Templates in the Backstage catalog.",
		"Template",
		[]*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the template.", Transform: specFieldTransform("type")},
			{Name: "parameters", Type: proto.ColumnType_JSON, Description: "Parameters defined in the template.", Transform: specFieldTransform("parameters")},
			{Name: "steps", Type: proto.ColumnType_JSON, Description: "Steps defined in the template.", Transform: specFieldTransform("steps")},
		},
		plugin.KeyColumnSlice{
			{Name: "type", Require: plugin.Optional},
		},
	)
}
