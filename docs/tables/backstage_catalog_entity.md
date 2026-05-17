# Table: backstage_catalog_entity

Query all entities in the Backstage catalog regardless of kind.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| api_version | string | The API version of the entity schema. |
| kind | string | The kind of the entity (e.g., Component, API, Group, User, System, Domain, Resource, Template, Location). |
| name | string | The name of the entity. |
| namespace | string | The namespace the entity belongs to. |
| metadata | json | The full metadata of the entity, including name, namespace, labels, annotations, tags, and links. |
| spec | json | The specification data of the entity. |
| relations | json | The relations of the entity to other entities. |
| status | json | Entity processing and health status. |
| status_items | json | `status.items` entries attached to the entity. |
| title | string | A display name of the entity. |
| description | string | A description of the entity. |
| labels | json | Labels attached to the entity. |
| annotations | json | Annotations attached to the entity. |
| labels_kv | json | Flattened labels in key/value format. |
| annotations_kv | json | Flattened annotations in key/value format. |
| annotation_github_project_slug | string | Value of `metadata.annotations['github.com/project-slug']`. |
| annotation_techdocs_ref | string | Value of `metadata.annotations['backstage.io/techdocs-ref']`. |
| annotation_source_location | string | Value of `metadata.annotations['backstage.io/source-location']`. |
| annotation_managed_by_location | string | Value of `metadata.annotations['backstage.io/managed-by-location']`. |
| tags | json | Tags attached to the entity. |
| links | json | Links attached to the entity. |
| relation_owned_by | json | Target refs from `ownedBy` relations. |
| relation_part_of | json | Target refs from `partOf` relations. |
| relation_depends_on | json | Target refs from `dependsOn` relations. |
| relation_has_part | json | Target refs from `hasPart` relations. |

## Examples

### List all entities

```sql
select
  kind,
  metadata ->> 'name' as name,
  metadata ->> 'namespace' as namespace
from
  backstage_catalog_entity;
```

### Count entities by kind

```sql
select
  kind,
  count(*) as entity_count
from
  backstage_catalog_entity
group by
  kind
order by
  entity_count desc;
```

### Query custom kinds with apiVersion and kind filters

```sql
select
  api_version,
  kind,
  name
from
  backstage_catalog_entity
where
  api_version = 'my-company.net/v1'
  and kind = 'Capability'
order by
  name;
```

### Use helper annotation columns

```sql
select
  name,
  annotation_github_project_slug,
  annotation_techdocs_ref,
  annotation_source_location
from
  backstage_catalog_entity
where
  annotation_github_project_slug is not null;
```

### Inspect authoritative ownership and dependency relations

```sql
select
  kind,
  name,
  relation_owned_by,
  relation_part_of,
  relation_depends_on
from
  backstage_catalog_entity
where
  relation_owned_by is not null
  or relation_depends_on is not null;
```

### Detect missing owners for kinds that should have one

```sql
select
  kind,
  name,
  namespace
from
  backstage_catalog_entity
where
  kind in ('Component', 'API', 'Resource', 'System', 'Domain')
  and coalesce(spec ->> 'owner', '') = ''
order by
  kind,
  name;
```

### Detect non-standard lifecycle values by kind

```sql
select
  kind,
  spec ->> 'lifecycle' as lifecycle,
  count(*) as entity_count
from
  backstage_catalog_entity
where
  kind in ('Component', 'API')
  and spec ->> 'lifecycle' is not null
  and spec ->> 'lifecycle' not in ('experimental', 'production', 'deprecated')
group by
  kind,
  spec ->> 'lifecycle'
order by
  entity_count desc;
```

### Find entities in a specific namespace

```sql
select
  kind,
  metadata ->> 'name' as name
from
  backstage_catalog_entity
where
  metadata ->> 'namespace' = 'default'
order by
  kind,
  metadata ->> 'name';
```

### List entities with a specific tag

```sql
select
  kind,
  metadata ->> 'name' as name,
  metadata -> 'tags' as tags
from
  backstage_catalog_entity
where
  metadata -> 'tags' ? 'java';
```
