package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// Define your Backstage catalog entity table
func tableBackstageCatalogEntities() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_entity",
		Description: "Retrieve entities from the Backstage catalog.",
		List: &plugin.ListConfig{
			Hydrate: listBackstageCatalogEntities,
		},
		Columns: []*plugin.Column{
			{Name: "name", Type: plugin.TypeString, Description: "The name of the entity."},
			{Name: "namespace", Type: plugin.TypeString, Description: "The namespace of the entity."},
			{Name: "kind", Type: plugin.TypeString, Description: "The type of the entity (Component, API, etc.)."},
			{Name: "metadata", Type: plugin.TypeJSON, Description: "Metadata of the entity."},
		},
	}
}

func listBackstageCatalogEntities(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	config := GetConfig(d.Connection)
	client := &http.Client{}

	// Construct the request
	req, err := http.NewRequest("GET", config.BaseURL+"/api/catalog/entities", nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+config.ApiToken)

	// Execute the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse the response
	var entities []map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&entities)
	if err != nil {
		return nil, err
	}

	// Stream each entity to Steampipe
	for _, entity := range entities {
		d.StreamListItem(ctx, entity)
	}

	return nil, nil
}
