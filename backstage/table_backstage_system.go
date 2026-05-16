package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageSystem() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_system",
		Description: "Systems in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listSystems,
		},
		Columns: append(commonColumns, []*plugin.Column{
			{Name: "domain", Type: proto.ColumnType_STRING, Description: "Domain the system belongs to"},
		}...), // Union of commonColumns and specific columns for systems
	}
}

func listSystems(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=System"},
		Fields:  []string{},
	}

	var cursor string
	for {
		systems, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_system.listSystems", "query_error", err)
			return nil, fmt.Errorf("error listing systems: %v", err)
		}

		for _, system := range systems {
			d.StreamListItem(ctx, system)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
