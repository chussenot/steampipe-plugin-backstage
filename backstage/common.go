package backstage

import (
	"context"
	"fmt"
	"strings"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

var commonColumns = []*plugin.Column{
	{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the entity.", Transform: transform.FromField("Metadata.Name")},
	{Name: "namespace", Type: proto.ColumnType_STRING, Description: "The namespace the entity belongs to.", Transform: transform.FromField("Metadata.Namespace")},
	{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of the entity."},
	{Name: "metadata", Type: proto.ColumnType_JSON, Description: "The full metadata of the entity."},
	{Name: "spec", Type: proto.ColumnType_JSON, Description: "The specification data of the entity."},
	{Name: "relations", Type: proto.ColumnType_JSON, Description: "The relations of the entity to other entities."},
	{Name: "title", Type: proto.ColumnType_STRING, Description: "A display name of the entity.", Transform: transform.FromField("Metadata.Title")},
	{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the entity.", Transform: transform.FromField("Metadata.Description")},
	{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity.", Transform: transform.FromField("Metadata.Labels")},
	{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations attached to the entity.", Transform: transform.FromField("Metadata.Annotations")},
	{Name: "tags", Type: proto.ColumnType_JSON, Description: "A list of tags attached to the entity.", Transform: transform.FromField("Metadata.Tags")},
	{Name: "links", Type: proto.ColumnType_JSON, Description: "A list of external hyperlinks related to the entity.", Transform: transform.FromField("Metadata.Links")},
}

var commonKeyColumns = plugin.KeyColumnSlice{
	{Name: "name", Require: plugin.Optional},
	{Name: "namespace", Require: plugin.Optional},
	{Name: "kind", Require: plugin.Optional},
}

func specFieldTransform(field string) *transform.ColumnTransforms {
	return transform.FromField(fmt.Sprintf("Spec.%s", field))
}

func newCatalogEntityTable(name, description, kind string, extraColumns []*plugin.Column, extraKeyColumns plugin.KeyColumnSlice) *plugin.Table {
	return &plugin.Table{
		Name:        name,
		Description: description,
		List: &plugin.ListConfig{
			Hydrate:    listCatalogEntitiesByKind(kind),
			KeyColumns: append(append(plugin.KeyColumnSlice{}, commonKeyColumns...), extraKeyColumns...),
		},
		Columns: append(append([]*plugin.Column{}, commonColumns...), extraColumns...),
	}
}

func listCatalogEntitiesByKind(kind string) plugin.HydrateFunc {
	return func(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
		client, err := connect(ctx, d)
		if err != nil {
			return nil, err
		}

		opts := &backstage.ListEntityOptions{
			Filters: buildCatalogFilters(d, kind),
		}

		entities, _, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("listCatalogEntitiesByKind", "query_error", err, "kind", kind, "filters", opts.Filters)
			return nil, fmt.Errorf("error listing entities: %w", err)
		}

		for _, entity := range entities {
			d.StreamListItem(ctx, entity)
			if d.RowsRemaining(ctx) == 0 {
				return nil, nil
			}
		}

		return nil, nil
	}
}

func buildCatalogFilters(d *plugin.QueryData, tableKind string) []string {
	filters := map[string]string{}
	if tableKind != "" {
		filters["kind"] = tableKind
	}

	if kind := d.EqualsQualString("kind"); kind != "" && tableKind == "" {
		filters["kind"] = kind
	}

	for qual, filterKey := range map[string]string{
		"name":      "metadata.name",
		"namespace": "metadata.namespace",
		"owner":     "spec.owner",
		"system":    "spec.system",
		"domain":    "spec.domain",
		"type":      "spec.type",
		"lifecycle": "spec.lifecycle",
	} {
		if value := strings.TrimSpace(d.EqualsQualString(qual)); value != "" {
			filters[filterKey] = value
		}
	}

	if len(filters) == 0 {
		return nil
	}

	var filterParts []string
	for _, key := range []string{"kind", "metadata.name", "metadata.namespace", "spec.owner", "spec.system", "spec.domain", "spec.type", "spec.lifecycle"} {
		if value, ok := filters[key]; ok && value != "" {
			filterParts = append(filterParts, fmt.Sprintf("%s=%s", key, value))
		}
	}

	if len(filterParts) == 0 {
		return nil
	}

	return []string{strings.Join(filterParts, ",")}
}
