# Backstage Plugin for Steampipe

Use SQL to query namespaces, components, APIs, users, groups and more from [Backstage](https://backstage.io/).

- **[Get started →](docs/index.md)**
- Documentation: [Table definitions & examples](docs/tables.md)
- Community: [Join #steampipe on Slack →](https://turbot.com/community/join)
- Get involved: [Issues](https://github.com/chussenot/steampipe-plugin-backstage/issues)

## Quick start

Install the plugin with [Steampipe](https://steampipe.io/downloads):

```shell
steampipe plugin install chussenot/backstage
```

[Configure the plugin](docs/index.md#configuration) by editing `~/.steampipe/config/backstage.spc`:

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

Or set environment variables:

```shell
export BACKSTAGE_HOST="https://demo.backstage.io"
export BACKSTAGE_TOKEN="your-token-here"
```

Start Steampipe:

```shell
steampipe query
```

Run a query:

```sql
select
  kind,
  metadata ->> 'name' as name,
  metadata ->> 'namespace' as namespace,
  metadata ->> 'description' as description
from
  backstage_catalog_entity
where
  kind = 'Component';
```

## Developing

Prerequisites:

- [Steampipe](https://steampipe.io/downloads)
- [Golang](https://golang.org/doc/install)

Clone:

```sh
git clone https://github.com/chussenot/steampipe-plugin-backstage.git
cd steampipe-plugin-backstage
```

Build, which automatically installs the new version to your `~/.steampipe/plugins` directory:

```sh
make install
```

Configure the plugin:

```sh
cp config/* ~/.steampipe/config
vi ~/.steampipe/config/backstage.spc
```

Try it!

```shell
steampipe query
> .inspect backstage
```

Further reading:

- [Writing plugins](https://steampipe.io/docs/develop/writing-plugins)
- [Writing your first table](https://steampipe.io/docs/develop/writing-your-first-table)

## Resources

- [steampipe](https://steampipe.io)
- [backstage](https://backstage.io/)
- [plugin release checklist](https://steampipe.io/docs/develop/plugin-release-checklist)
- [go-backstage](https://github.com/datolabs-io/go-backstage)
- [steampipe plugin standards](https://steampipe.io/docs/develop/standards#naming)

## License

Apache 2.0 — see [LICENSE](LICENSE).
