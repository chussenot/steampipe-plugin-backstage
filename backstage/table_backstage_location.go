package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageLocation() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_location",
		"Locations in the Backstage catalog.",
		"Location",
		[]*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the location.", Transform: specFieldTransform("type")},
			{Name: "target", Type: proto.ColumnType_STRING, Description: "Target of the location.", Transform: specFieldTransform("target")},
			{Name: "targets", Type: proto.ColumnType_JSON, Description: "Targets of the location.", Transform: specFieldTransform("targets")},
			{Name: "presence", Type: proto.ColumnType_STRING, Description: "Whether the location target is required or optional.", Transform: specFieldTransform("presence")},
		},
		plugin.KeyColumnSlice{
			{Name: "type", Require: plugin.Optional},
			{Name: "presence", Require: plugin.Optional},
		},
	)
}
