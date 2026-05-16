package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/grpc/proto"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

func tableBackstageUser() *plugin.Table {
	return &plugin.Table{
		Name:        "backstage_catalog_user",
		Description: "Users in the Backstage catalog",
		List: &plugin.ListConfig{
			Hydrate: listUsers,
		},
		Columns: append(commonColumns, []*plugin.Column{
			{Name: "email", Type: proto.ColumnType_STRING, Description: "Email of the user"},
			{Name: "picture", Type: proto.ColumnType_STRING, Description: "Picture URL of the user"},
			{Name: "memberof", Type: proto.ColumnType_JSON, Description: "Groups the user belongs to"},
		}...), // Union of commonColumns and specific columns for users
	}
}

func listUsers(ctx context.Context, d *plugin.QueryData, _ *plugin.HydrateData) (interface{}, error) {
	client, err := connect(ctx, d)
	if err != nil {
		return nil, err
	}

	opts := &backstage.ListEntityOptions{
		Filters: []string{"kind=User"},
		Fields:  []string{},
	}

	for {
		users, resp, err := client.Catalog.Entities.List(ctx, opts)
		if err != nil {
			plugin.Logger(ctx).Error("backstage_catalog_user.listUsers", "query_error", err)
			return nil, fmt.Errorf("error listing users: %v", err)
		}

		for _, user := range users {
			d.StreamListItem(ctx, user)
		}

		if resp.Header.Get("Link") == "" {
			break
		}
	}

	return nil, nil
}
