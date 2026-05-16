package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageDomain() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_domain",
		Description: "Domains in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listDomains,
		},
		Columns: append(commonColumns, []*plugin.Column{
			{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the domain"},
		}...), // Union of commonColumns and specific columns for domains
	}
}

func listDomains(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Domain"},
		Fields:  []string{},
	}

	var cursor string
	for {
		domains, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_domain.listDomains", "query_error", err)
			return nil, fmt.Errorf("error listing domains: %v", err)
		}

		for _, domain := range domains {
			d.StreamListItem(ctx, domain)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
