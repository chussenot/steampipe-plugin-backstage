# Backstage Plugin for Steampipe

Use SQL to query the Backstage software catalog (entities, components, APIs, systems, resources, users, groups, templates, locations, and domains).

- Getting started: [docs/index.md](docs/index.md)
- Table reference: [docs/tables.md](docs/tables.md)
- Issues: [github.com/chussenot/steampipe-plugin-backstage/issues](https://github.com/chussenot/steampipe-plugin-backstage/issues)

## What this plugin does

This plugin connects to a Backstage instance and exposes catalog data as relational tables:

- `backstage_catalog_entity` (all entities)
- `backstage_catalog_entity_kind` (discovered apiVersion/kind pairs)
- `backstage_catalog_component`, `backstage_catalog_api`, `backstage_catalog_resource`
- `backstage_catalog_system`, `backstage_catalog_domain`
- `backstage_catalog_user`, `backstage_catalog_group`
- `backstage_catalog_template`, `backstage_catalog_location`

## Install

```sh
steampipe plugin install chussenot/backstage
```

## Configure

Create `~/.steampipe/config/backstage.spc`:

```hcl
connection "backstage" {
  plugin = "chussenot/backstage"
  host   = "https://demo.backstage.io"
  # token = "your-token-here"
}
```

You can also configure through environment variables:

```sh
export BACKSTAGE_HOST="https://demo.backstage.io"
export BACKSTAGE_TOKEN="your-token-here"
```

`host` is required either in connection config or `BACKSTAGE_HOST`. `token` is optional for public/anonymous Backstage catalogs.

## Example queries

List entities by kind:

```sql
select kind, count(*) as total
from backstage_catalog_entity
group by kind
order by total desc;
```

List components and owners:

```sql
select name, owner, system
from backstage_catalog_component
order by name;
```

List APIs and their lifecycle:

```sql
select name, type, lifecycle, owner
from backstage_catalog_api
order by name;
```

## Development workflow

Prerequisites:

- Go 1.26+
- Steampipe CLI

Local development:

```sh
git clone https://github.com/chussenot/steampipe-plugin-backstage.git
cd steampipe-plugin-backstage
make test
go vet ./...
make build
make install
cp config/backstage.spc ~/.steampipe/config/backstage.spc
```

Then run:

```sh
steampipe query
```

## Release workflow (Cog + GitHub Actions)

1. Merge conventional commits into `master`.
2. The master release-bump workflow runs `cog bump --auto`.
3. Cog creates version commit + tag (`v*`) and pushes them.
4. Tag workflows build/push the plugin image and verify with a Steampipe smoke query.
5. Validate install:

```sh
steampipe plugin install chussenot/backstage
steampipe query "select count(*) from backstage_catalog_entity;"
```

## Steampipe Hub publication note

Automated build/push and release tagging are handled in this repository. If an additional Steampipe Hub submission/review step is required for indexing, perform that manual submission after the GitHub release/tag is created.

## License

Apache-2.0. See [LICENSE](LICENSE).
