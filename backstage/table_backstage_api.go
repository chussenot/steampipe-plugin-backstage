package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define additional columns specific to the API table
var apiSpecificColumns = []*plugin.Column{
	{Name: "owner", Type: proto.ColumnType_STRING, Description: "Owner of the API"},
	{Name: "definition", Type: proto.ColumnType_JSON, Description: "API definition"},
	{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the API"},
	{Name: "lifecycle", Type: proto.ColumnType_STRING, Description: "Lifecycle state of the API"},
}

func tableBackstageAPI() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_api",
		Description: "APIs in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listAPIs,
		},
		Columns: append(commonColumns, apiSpecificColumns...), // Union of commonColumns and apiSpecificColumns
	}
}

func listAPIs(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=API"},
		Fields:  []string{},
	}

	var cursor string
	for {
		apis, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_api.listAPIs", "query_error", err)
			return nil, fmt.Errorf("error listing APIs: %v", err)
		}

		for _, api := range apis {
			d.StreamListItem(ctx, api)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
