# Table: backstage_catalog_location

A Location in Backstage represents a source from which entity definitions are ingested into the catalog. Locations can be URLs, files, or other sources containing entity definitions.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the location. |
| namespace | string | The namespace the location belongs to. |
| kind | string | The kind of the entity (always "Location"). |
| metadata | json | The full metadata of the location. |
| spec | json | The specification data of the location. |
| relations | json | The relations of the location to other entities. |
| title | string | A display name of the location. |
| description | string | A description of the location. |
| labels | json | Labels attached to the location. |
| annotations | json | Annotations attached to the location. |
| tags | json | A list of tags attached to the location. |
| links | json | A list of external hyperlinks related to the location. |
| type | string | Type of the location (url, file, etc.). |
| target | string | Target of the location (URL or file path). |
| targets | json | Targets of the location. |
| presence | string | Whether the target is required or optional. |

## Examples

### Basic info

```sql
select
  name,
  type,
  target
from
  backstage_catalog_location;
```

### List locations by type

```sql
select
  type,
  count(*) as location_count
from
  backstage_catalog_location
group by
  type
order by
  location_count desc;
```

### Find locations with specific annotations

```sql
select
  name,
  target,
  annotations->>'backstage.io/managed-by' as managed_by
from
  backstage_catalog_location
where
  annotations->>'backstage.io/managed-by' is not null;
```
