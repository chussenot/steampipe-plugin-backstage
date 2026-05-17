package backstage

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin/transform"
)

var commonColumns = []*plugin.Column{
	{Name: "api_version", Type: proto.ColumnType_STRING, Description: "The API version of the entity schema.", Transform: transform.FromField("ApiVersion")},
	{Name: "name", Type: proto.ColumnType_STRING, Description: "The name of the entity.", Transform: transform.FromField("Metadata.Name")},
	{Name: "namespace", Type: proto.ColumnType_STRING, Description: "The namespace the entity belongs to.", Transform: transform.FromField("Metadata.Namespace")},
	{Name: "kind", Type: proto.ColumnType_STRING, Description: "The kind of the entity."},
	{Name: "metadata", Type: proto.ColumnType_JSON, Description: "The full metadata of the entity."},
	{Name: "spec", Type: proto.ColumnType_JSON, Description: "The specification data of the entity."},
	{Name: "relations", Type: proto.ColumnType_JSON, Description: "The relations of the entity to other entities."},
	{Name: "status", Type: proto.ColumnType_JSON, Description: "Entity processing and health status.", Transform: transform.FromField("Status")},
	{Name: "status_items", Type: proto.ColumnType_JSON, Description: "Status items extracted from status.items.", Transform: transform.FromField("Status.Items")},
	{Name: "title", Type: proto.ColumnType_STRING, Description: "A display name of the entity.", Transform: transform.FromField("Metadata.Title")},
	{Name: "description", Type: proto.ColumnType_STRING, Description: "A description of the entity.", Transform: transform.FromField("Metadata.Description")},
	{Name: "labels", Type: proto.ColumnType_JSON, Description: "Labels attached to the entity.", Transform: transform.FromField("Metadata.Labels")},
	{Name: "annotations", Type: proto.ColumnType_JSON, Description: "Annotations attached to the entity.", Transform: transform.FromField("Metadata.Annotations")},
	{Name: "labels_kv", Type: proto.ColumnType_JSON, Description: "Flattened labels as key/value objects.", Transform: transform.FromP(entityMapToKeyValueListTransform, "labels")},
	{Name: "annotations_kv", Type: proto.ColumnType_JSON, Description: "Flattened annotations as key/value objects.", Transform: transform.FromP(entityMapToKeyValueListTransform, "annotations")},
	{Name: "annotation_github_project_slug", Type: proto.ColumnType_STRING, Description: "Value of metadata.annotations['github.com/project-slug'].", Transform: transform.FromP(entityAnnotationValueTransform, "github.com/project-slug")},
	{Name: "annotation_techdocs_ref", Type: proto.ColumnType_STRING, Description: "Value of metadata.annotations['backstage.io/techdocs-ref'].", Transform: transform.FromP(entityAnnotationValueTransform, "backstage.io/techdocs-ref")},
	{Name: "annotation_source_location", Type: proto.ColumnType_STRING, Description: "Value of metadata.annotations['backstage.io/source-location'].", Transform: transform.FromP(entityAnnotationValueTransform, "backstage.io/source-location")},
	{Name: "annotation_managed_by_location", Type: proto.ColumnType_STRING, Description: "Value of metadata.annotations['backstage.io/managed-by-location'].", Transform: transform.FromP(entityAnnotationValueTransform, "backstage.io/managed-by-location")},
	{Name: "tags", Type: proto.ColumnType_JSON, Description: "A list of tags attached to the entity.", Transform: transform.FromField("Metadata.Tags")},
	{Name: "links", Type: proto.ColumnType_JSON, Description: "A list of external hyperlinks related to the entity.", Transform: transform.FromField("Metadata.Links")},
	{Name: "relation_owned_by", Type: proto.ColumnType_JSON, Description: "Target refs from ownedBy relations.", Transform: transform.FromP(relationTargetsByTypeTransform, "ownedBy")},
	{Name: "relation_part_of", Type: proto.ColumnType_JSON, Description: "Target refs from partOf relations.", Transform: transform.FromP(relationTargetsByTypeTransform, "partOf")},
	{Name: "relation_depends_on", Type: proto.ColumnType_JSON, Description: "Target refs from dependsOn relations.", Transform: transform.FromP(relationTargetsByTypeTransform, "dependsOn")},
	{Name: "relation_has_part", Type: proto.ColumnType_JSON, Description: "Target refs from hasPart relations.", Transform: transform.FromP(relationTargetsByTypeTransform, "hasPart")},
}

var commonKeyColumns = plugin.KeyColumnSlice{
	{Name: "api_version", Require: plugin.Optional},
	{Name: "name", Require: plugin.Optional},
	{Name: "namespace", Require: plugin.Optional},
	{Name: "kind", Require: plugin.Optional},
}

func specFieldTransform(field string) *transform.ColumnTransforms {
	return transform.FromField(fmt.Sprintf("Spec.%s", field))
}

func entityAnnotationValueTransform(_ context.Context, d *transform.TransformData) (interface{}, error) {
	annotationKey, ok := d.Param.(string)
	if !ok || annotationKey == "" {
		return nil, nil
	}

	entity, ok := d.Value.(backstage.Entity)
	if !ok || len(entity.Metadata.Annotations) == 0 {
		return nil, nil
	}

	value, ok := entity.Metadata.Annotations[annotationKey]
	if !ok || value == "" {
		return nil, nil
	}

	return value, nil
}

func entityMapToKeyValueListTransform(_ context.Context, d *transform.TransformData) (interface{}, error) {
	entity, ok := d.Value.(backstage.Entity)
	if !ok {
		return nil, nil
	}

	mapName, ok := d.Param.(string)
	if !ok || mapName == "" {
		return nil, nil
	}

	var values map[string]string
	switch mapName {
	case "labels":
		values = entity.Metadata.Labels
	case "annotations":
		values = entity.Metadata.Annotations
	default:
		return nil, nil
	}
	if len(values) == 0 {
		return nil, nil
	}

	keys := make([]string, 0, len(values))
	for k := range values {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	flattened := make([]map[string]string, 0, len(values))
	for _, key := range keys {
		flattened = append(flattened, map[string]string{
			"key":   key,
			"value": values[key],
		})
	}

	return flattened, nil
}

func relationTargetsByTypeTransform(_ context.Context, d *transform.TransformData) (interface{}, error) {
	relationType, ok := d.Param.(string)
	if !ok || relationType == "" {
		return nil, nil
	}

	entity, ok := d.Value.(backstage.Entity)
	if !ok || len(entity.Relations) == 0 {
		return nil, nil
	}

	targetRefs := make([]string, 0)
	for _, relation := range entity.Relations {
		if relation.Type == relationType && relation.TargetRef != "" {
			targetRefs = append(targetRefs, relation.TargetRef)
		}
	}

	if len(targetRefs) == 0 {
		return nil, nil
	}

	return targetRefs, nil
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
		"api_version":     "apiVersion",
		"name":            "metadata.name",
		"namespace":       "metadata.namespace",
		"owner":           "spec.owner",
		"system":          "spec.system",
		"domain":          "spec.domain",
		"type":            "spec.type",
		"lifecycle":       "spec.lifecycle",
		"subcomponent_of": "spec.subcomponentOf",
		"subdomain_of":    "spec.subdomainOf",
		"presence":        "spec.presence",
	} {
		if value := strings.TrimSpace(d.EqualsQualString(qual)); value != "" {
			filters[filterKey] = value
		}
	}

	if len(filters) == 0 {
		return nil
	}

	var filterParts []string
	for _, key := range []string{"apiVersion", "kind", "metadata.name", "metadata.namespace", "spec.owner", "spec.system", "spec.domain", "spec.type", "spec.lifecycle", "spec.subcomponentOf", "spec.subdomainOf", "spec.presence"} {
		if value, ok := filters[key]; ok && value != "" {
			filterParts = append(filterParts, fmt.Sprintf("%s=%s", key, value))
		}
	}

	if len(filterParts) == 0 {
		return nil
	}

	return []string{strings.Join(filterParts, ",")}
}
