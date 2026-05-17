# Table: backstage_catalog_resource

A Resource in Backstage represents infrastructure or services that are used by components. This can include databases, S3 buckets, queues, and other infrastructure components.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the resource. |
| namespace | string | The namespace the resource belongs to. |
| kind | string | The kind of the entity (always "Resource"). |
| metadata | json | The full metadata of the resource. |
| spec | json | The specification data of the resource. |
| relations | json | The relations of the resource to other entities. |
| title | string | A display name of the resource. |
| description | string | A description of the resource. |
| labels | json | Labels attached to the resource. |
| annotations | json | Annotations attached to the resource. |
| tags | json | A list of tags attached to the resource. |
| links | json | A list of external hyperlinks related to the resource. |
| type | string | Type of the resource. |
| owner | string | Owner of the resource. |
| system | string | System the resource belongs to. |
| depends_on | json | Entities this resource depends on. |
| dependency_of | json | Entities that depend on this resource. |

## Examples

### Basic info

```sql
select
  name,
  type,
  owner,
  system
from
  backstage_catalog_resource;
```

### List resources by type

```sql
select
  type,
  count(*) as resource_count
from
  backstage_catalog_resource
group by
  type
order by
  resource_count desc;
```

### Find resources by system

```sql
select
  r.name as resource_name,
  r.type as resource_type,
  s.name as system_name
from
  backstage_catalog_resource as r
  join backstage_catalog_system as s on s.name = r.system;
```
