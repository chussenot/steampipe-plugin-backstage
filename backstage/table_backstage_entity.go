package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageEntity() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_entity",
		Description: "Generic entities in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listEntities,
		},
		Columns: commonColumns,
	}
}

func listEntities(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	logger := plugin.Logger(ctx)

	client, err := connect(ctx, d)
	if err != nil {
		logger.Error("listEntities", "connection_error", err)
		return nil, err
	}

	logger.Debug("listEntities", "status", "starting entity fetch")

	opts := &backstage.ListEntityOptions{
		Fields: commonFields,
	}

	var cursor string
	var entityCount int
	for {
		entities, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			logger.Error("listEntities", "query_error", err, "cursor", cursor)
			return nil, fmt.Errorf("error listing entities: %v", err)
		}

		for _, entity := range entities {
			entityCount++
			d.StreamListItem(ctx, entity)
		}

		logger.Debug("listEntities", "batch_count", len(entities), "total_count", entityCount)

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
		logger.Debug("listEntities", "pagination_cursor", cursor)
	}
	logger.Info("listEntities", "final_count", entityCount)
	return nil, nil
}
