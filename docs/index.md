# Backstage Plugin

## Overview

The Backstage plugin for Steampipe lets you query software components, APIs, users, groups, systems, domains, and more from your [Backstage](https://backstage.io/) software catalog using SQL.

## Getting Started

### Installation

Install the plugin with [Steampipe](https://steampipe.io/downloads):

```shell
steampipe plugin install chussenot/backstage
```

### Configuration

Configure the plugin by creating a `~/.steampipe/config/backstage.spc` file:

```hcl
connection "backstage" {
  plugin = "chussenot/backstage"

  # Backstage instance URL (required).
  # Can also be set with the BACKSTAGE_HOST environment variable.
  host = "https://demo.backstage.io"

  # Backstage API token for authentication (optional).
  # Required for Backstage instances with authentication enabled.
  # Can also be set with the BACKSTAGE_TOKEN environment variable.
  # token = "your-token-here"
}
```

### Environment Variables

Credentials can be set via environment variables:

```sh
export BACKSTAGE_HOST="https://demo.backstage.io"
export BACKSTAGE_TOKEN="your-token-here"
```

### Verify the Connection

Run a query to verify everything is working:

```sql
select
  kind,
  metadata ->> 'name' as name,
  metadata ->> 'namespace' as namespace
from
  backstage_catalog_entity
limit 5;
```

If you get results, you're all set! If you get an error, check your configuration.

## Get involved

* Open source: [GitHub Repository](https://github.com/chussenot/steampipe-plugin-backstage)
* Community: [Join #steampipe on Slack →](https://turbot.com/community/join)

## Authentication

The Backstage plugin supports two authentication methods:

### API Token (recommended for authenticated instances)

Set the `token` configuration argument or the `BACKSTAGE_TOKEN` environment variable.

To generate an API token:
1. Log in to your Backstage instance
2. Navigate to your user settings
3. Generate a new API token

For more information about Backstage authentication, see:
- [Backstage Authentication](https://backstage.io/docs/auth/)
- [Backstage Tokens](https://backstage.io/docs/auth/tokens)

### Anonymous Access

If your Backstage instance allows anonymous access, you can omit the `token` configuration argument:

```hcl
connection "backstage" {
  plugin = "chussenot/backstage"
  host   = "https://demo.backstage.io"
}
```

### Required Permissions

When using token authentication, the token needs the following permissions to query the catalog:

- `catalog.entity.read`
- `catalog.location.read`

For more details about Backstage permissions, see:
- [Backstage Permissions](https://backstage.io/docs/permissions/overview)

## Troubleshooting

### Common Issues

* **Connection Failed**: Verify your `host` (or `BACKSTAGE_HOST`) is accessible and includes the protocol (e.g., `https://`)
* **Authentication Failed**: Ensure your token is valid and has the required permissions
* **No Results**: Check that your Backstage instance has entities in its catalog

For more help, join our [Slack community](https://turbot.com/community/join).

