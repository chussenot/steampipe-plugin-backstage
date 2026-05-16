package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageGroup() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_group",
		Description: "Groups in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listGroups,
		},
		Columns: append(commonColumns, []*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the entity"},
		}...), // Union of commonColumns and specific columns for groups
	}
}

func listGroups(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	// Get the limit
	limit := d.QueryContext.Limit
	if limit != nil {
		if *limit < 1 {
			return nil, nil
		}
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Group"},
		Fields:  []string{},
	}

	// Handle pagination
	var cursor string
	for {
		groups, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_group.listGroups", "query_error", err)
			return nil, fmt.Errorf("error listing groups: %v", err)
		}

		for _, group := range groups {
			d.StreamListItem(ctx, group)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
