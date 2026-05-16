package backstage

import (
	"context"
	"fmt"

	"github.com/datolabs-io/go-backstage/v3"
	"github.com/turbot/steampipe-plugin-sdk/v5/plugin"
)

// connect establishes a connection to the Backstage API using the provided token.
func connect(ctx context.Context, d *plugin.QueryData) (*backstage.Client, error) {
	logger := plugin.Logger(ctx)

	cacheKey := "backstage"
	if cached, ok := d.ConnectionManager.Cache.Get(cacheKey); ok {
		if client, cOk := cached.(*backstage.Client); cOk {
			return client, nil
		}
	}

	config := GetConfig(d.Connection)

	var host, token string
	if config.Host != nil {
		host = *config.Host
	}
	if config.Token != nil {
		token = *config.Token
	}

	logger.Debug("backstage.connect", "host", host)

	client, err := getClient(host, token)
	if err != nil {
		logger.Error("backstage.connect", "connection_error", err)
		return nil, fmt.Errorf("error creating backstage client: %v", err)
	}

	d.ConnectionManager.Cache.Set(cacheKey, client)
	logger.Debug("backstage.connect", "status", "connection successful")
	return client, nil
}
