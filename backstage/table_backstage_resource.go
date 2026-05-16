package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageResource() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_resource",
		Description: "Resources in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listResources,
		},
		Columns: commonColumns,
	}
}

func listResources(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Resource"},
		Fields:  []string{},
	}

	var cursor string
	for {
		resources, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_resource.listResources", "query_error", err)
			return nil, fmt.Errorf("error listing resources: %v", err)
		}

		for _, resource := range resources {
			d.StreamListItem(ctx, resource)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
