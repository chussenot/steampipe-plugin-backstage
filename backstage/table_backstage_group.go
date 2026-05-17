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
			{Name: "parent", Type: proto.ColumnType_STRING, Description: "Parent group of this group.", Transform: specFieldTransform("parent")},
			{Name: "children", Type: proto.ColumnType_JSON, Description: "Child groups of this group.", Transform: specFieldTransform("children")},
			{Name: "members", Type: proto.ColumnType_JSON, Description: "Members belonging to this group.", Transform: specFieldTransform("members")},
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display name of the group profile.", Transform: specFieldTransform("profile.displayName")},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email address from the group profile.", Transform: specFieldTransform("profile.email")},
			{Name: "picture", Type: proto.ColumnType_STRING, Description: "Picture URL from the group profile.", Transform: specFieldTransform("profile.picture")},
		},
		plugin.KeyColumnSlice{
			{Name: "type", Require: plugin.Optional},
			{Name: "parent", Require: plugin.Optional},
		},
	)
}
