package backstage

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

type entityKindSummary struct {
	ApiVersion  string
	Kind        string
	EntityCount int64
}

func tableBackstageEntityKind() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_entity_kind",
		Description: "Discovered entity kind and apiVersion combinations in the Backstage catalog.",
		List: &plugin.ListConfig{
			Hydrate: listCatalogEntityKinds,
			KeyColumns: plugin.KeyColumnSlice{
				{Name: "api_version", Require: plugin.Optional},
				{Name: "kind", Require: plugin.Optional},
			},
		},
		Columns: []*plugin.Column{
			{Name: "api_version", Type: proto.ColumnType_STRING, Description: "Entity apiVersion."},
			{Name: "kind", Type: proto.ColumnType_STRING, Description: "Entity kind."},
			{Name: "entity_count", Type: proto.ColumnType_INT, Description: "Count of entities for this apiVersion/kind pair."},
		},
	}
}

func listCatalogEntityKinds(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: buildEntityKindFilters(d),
	}

	entities, _, err := client.Catalog.Entities.List(ctx, opts)
	if err != nil {
		plugin.Logger(ctx).Error("listCatalogEntityKinds", "query_error", err, "filters", opts.Filters)
		return nil, fmt.Errorf("error listing entities for kind summary: %w", err)
	}

	aggregated := map[string]*entityKindSummary{}
	for _, entity := range entities {
		key := fmt.Sprintf("%s|%s", entity.ApiVersion, entity.Kind)
		if _, ok := aggregated[key]; !ok {
			aggregated[key] = &entityKindSummary{
				ApiVersion: entity.ApiVersion,
				Kind:       entity.Kind,
			}
		}
		aggregated[key].EntityCount++
	}

	keys := make([]string, 0, len(aggregated))
	for key := range aggregated {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	for _, key := range keys {
		d.StreamListItem(ctx, aggregated[key])
		if d.RowsRemaining(ctx) == 0 {
			return nil, nil
		}
	}

	return nil, nil
}

func buildEntityKindFilters(d *plugin.QueryData) []string {
	filters := map[string]string{}

	if kind := d.EqualsQualString("kind"); kind != "" {
		filters["kind"] = kind
	}
	if apiVersion := d.EqualsQualString("api_version"); apiVersion != "" {
		filters["apiVersion"] = apiVersion
	}

	if len(filters) == 0 {
		return nil
	}

	filterParts := make([]string, 0, len(filters))
	for _, key := range []string{"apiVersion", "kind"} {
		if value, ok := filters[key]; ok && value != "" {
			filterParts = append(filterParts, fmt.Sprintf("%s=%s", key, value))
		}
	}

	if len(filterParts) == 0 {
		return nil
	}

	return []string{strings.Join(filterParts, ",")}
}
