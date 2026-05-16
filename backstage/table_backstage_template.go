package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageTemplate() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_template",
		Description: "Templates in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listTemplates,
		},
		Columns: append(commonColumns, []*plugin.Column{
			{Name: "type", Type: proto.ColumnType_STRING, Description: "Type of the template"},
			{Name: "parameters", Type: proto.ColumnType_JSON, Description: "Parameters defined in the template"},
			{Name: "steps", Type: proto.ColumnType_JSON, Description: "Steps defined in the template"},
		}...), // Union of commonColumns and specific columns for templates
	}
}

func listTemplates(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=Template"},
		Fields:  []string{},
	}

	var cursor string
	for {
		templates, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_template.listTemplates", "query_error", err)
			return nil, fmt.Errorf("error listing templates: %v", err)
		}

		for _, template := range templates {
			d.StreamListItem(ctx, template)
		}

		cursor = resp.Header.Get("Link")
		if cursor == "" {
			break
		}
	}

	return nil, nil
}
