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
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the component.", Transform: specFieldTransform("type")},
			{Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "Lifecycle state of the component.", Transform: specFieldTransform("lifecycle")},
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the component.", Transform: specFieldTransform("owner")},
			{Name: "system", Type: proto.ColumnType_STRING, Description: "System the component belongs to.", Transform: specFieldTransform("system")},
			{Name: "subcomponent_of", Type: proto.ColumnType_STRING, Description: "Parent component this component is part of.", Transform: specFieldTransform("subcomponentOf")},
			{Name: "provides_apis", Type: proto.ColumnType_JSON, Description: "APIs provided by this component.", Transform: specFieldTransform("providesApis")},
			{Name: "consumes_apis", Type: proto.ColumnType_JSON, Description: "APIs consumed by this component.", Transform: specFieldTransform("consumesApis")},
			{Name: "depends_on", Type: proto.ColumnType_JSON, Description: "Entities this component depends on.", Transform: specFieldTransform("dependsOn")},
			{Name: "dependency_of", Type: proto.ColumnType_JSON, Description: "Entities that depend on this component.", Transform: specFieldTransform("dependencyOf")},
		},
		plugin.KeyColumnSlice{
			{Name: "type", Require: plugin.Optional},
			{Name: "lifecycle", Require: plugin.Optional},
			{Name: "owner", Require: plugin.Optional},
			{Name: "system", Require: plugin.Optional},
			{Name: "subcomponent_of", Require: plugin.Optional},
		},
	)
}
