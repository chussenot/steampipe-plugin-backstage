# Table: backstage_catalog_entity

Query all entities in the Backstage catalog regardless of kind.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| kind | string | The kind of the entity (e.g., Component, API, Group, User, System, Domain, Resource, Template, Location). |
| name | string | The name of the entity. |
| namespace | string | The namespace the entity belongs to. |
| metadata | json | The full metadata of the entity, including name, namespace, labels, annotations, tags, and links. |
| spec | json | The specification data of the entity. |
| relations | json | The relations of the entity to other entities. |
| title | string | A display name of the entity. |
| description | string | A description of the entity. |
| labels | json | Labels attached to the entity. |
| annotations | json | Annotations attached to the entity. |
| tags | json | Tags attached to the entity. |
| links | json | Links attached to the entity. |

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
