package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageLocation() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_location",
		Description: "Locations in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listLocations,
		},
		Columns: append(commonColumns, []*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the location"},
			{Name: "target", Type: proto.ColumnType_STRING, Description: "Target of the location"},
		}...), // Union of commonColumns and specific columns for locations
	}
}

func listLocations(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Location"},
		Fields:  []string{},
	}

	var cursor string
	for {
		locations, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_location.listLocations", "query_error", err)
			return nil, fmt.Errorf("error listing locations: %v", err)
		}

		for _, location := range locations {
			d.StreamListItem(ctx, location)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
