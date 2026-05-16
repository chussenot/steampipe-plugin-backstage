# Table: backstage_catalog_entity

Query all entities in the Backstage catalog regardless of kind.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| kind | string | The kind of the entity (e.g., Component, API, Group, User, System, Domain, Resource, Template, Location). |
| metadata | json | The full metadata of the entity, including name, namespace, labels, annotations, tags, and links. |
| spec | json | The specification data of the entity. |

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
