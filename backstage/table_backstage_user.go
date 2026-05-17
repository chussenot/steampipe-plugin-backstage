package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageUser() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_user",
		"Users in the Backstage catalog.",
		"User",
		[]*plugin.Column{
			{Name: "display_name", Type: proto.ColumnType_STRING, Description: "Display name of the user profile.", Transform: specFieldTransform("profile.displayName")},
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email address of the user.", Transform: specFieldTransform("profile.email")},
			{Name: "picture", Type: proto.ColumnType_STRING, Description: "Picture URL of the user.", Transform: specFieldTransform("profile.picture")},
			{Name: "member_of", Type: proto.ColumnType_JSON, Description: "Groups the user belongs to.", Transform: specFieldTransform("memberOf")},
		},
		nil,
	)
}
