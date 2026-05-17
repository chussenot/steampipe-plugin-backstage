package backstage

import (
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageDomain() *plugin.Table {
	return newCatalogEntityTable(
		"backstage_catalog_domain",
		"Domains in the Backstage catalog.",
		"Domain",
		[]*plugin.Column{
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the domain.", Transform: specFieldTransform("owner")},
			{Name: "subdomain_of", Type: proto.ColumnType_STRING, Description: "Parent domain this domain is part of.", Transform: specFieldTransform("subdomainOf")},
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the domain.", Transform: specFieldTransform("type")},
		},
		plugin.KeyColumnSlice{
			{Name: "owner", Require: plugin.Optional},
			{Name: "subdomain_of", Require: plugin.Optional},
			{Name: "type", Require: plugin.Optional},
		},
	)
}
