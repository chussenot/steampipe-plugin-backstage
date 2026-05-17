# Table: backstage_catalog_system

A Backstage System represents a collection of resources and components that work together to deliver value to customers. Systems help you understand how your software ecosystem fits together.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the system. |
| namespace | string | The namespace the system belongs to. |
| kind | string | The kind of the entity (always "System"). |
| metadata | json | The full metadata of the system. |
| spec | json | The specification data of the system. |
| relations | json | The relations of the system to other entities. |
| title | string | A display name of the system. |
| description | string | A description of the system. |
| labels | json | Labels attached to the system. |
| annotations | json | Annotations attached to the system. |
| tags | json | A list of tags attached to the system. |
| links | json | A list of external hyperlinks related to the system. |
| owner | string | Owner of the system. |
| domain | string | Domain the system belongs to. |
| type | string | Type of the system. |

## Examples

### Basic info

```sql
select
  name,
  domain,
  description
from
  backstage_catalog_system;
```

### List systems with their components

```sql
select
  s.name as system_name,
  count(c.name) as component_count
from
  backstage_catalog_system as s
  left join backstage_catalog_component as c on c.system = s.name
group by
  s.name
order by
  component_count desc;
```

### Find systems without domains

```sql
select
  name,
  description
from
  backstage_catalog_system
where
  domain is null;
``` 
