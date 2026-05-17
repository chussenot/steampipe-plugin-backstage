# Table: backstage_catalog_entity_kind

List discovered `apiVersion` + `kind` pairs found in the Backstage catalog, including custom kinds.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| api_version | string | Entity apiVersion. |
| kind | string | Entity kind. |
| entity_count | int | Count of entities for this apiVersion/kind pair. |

## Examples

### List all discovered kinds

```sql
select
  api_version,
  kind,
  entity_count
from
  backstage_catalog_entity_kind
order by
  entity_count desc,
  api_version,
  kind;
```

### Find non-core/custom kinds

```sql
select
  api_version,
  kind,
  entity_count
from
  backstage_catalog_entity_kind
where
  api_version <> 'backstage.io/v1alpha1'
  or kind not in ('Component', 'API', 'Resource', 'System', 'Domain', 'Group', 'User', 'Template', 'Location')
order by
  entity_count desc;
```
