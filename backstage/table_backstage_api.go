package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define additional columns specific to the API table
var apiSpecificColumns = []*plugin.Column{
	{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the API.", Transform: specFieldTransform("owner")},
	{Name: "system", Type: proto.ColumnType_STRING, Description: "System the API belongs to.", Transform: specFieldTransform("system")},
	{Name: "definition", Type: proto.ColumnType_JSON, Description: "API definition.", Transform: specFieldTransform("definition")},
	{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the API.", Transform: specFieldTransform("type")},
	{Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "Lifecycle state of the API.", Transform: specFieldTransform("lifecycle")},
}

func tableBackstageAPI() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_api",
		"APIs in the Backstage catalog.",
		"API",
		apiSpecificColumns,
		plugin.KeyColumnSlice{
			{Name: "owner", Require: plugin.Optional},
			{Name: "system", Require: plugin.Optional},
			{Name: "type", Require: plugin.Optional},
			{Name: "lifecycle", Require: plugin.Optional},
		},
	)
}
