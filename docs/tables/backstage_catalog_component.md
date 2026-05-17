# Table: backstage_catalog_component

Query software components in your Backstage catalog.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the component |
| namespace | string | The namespace the component belongs to |
| kind | string | The kind of the entity (always "Component") |
| metadata | json | The full metadata of the component |
| spec | json | The specification data of the component |
| relations | json | The relations of the component to other entities |
| title | string | A display name of the component |
| description | string | A description of the component |
| labels | json | Labels attached to the component |
| annotations | json | Annotations attached to the component |
| tags | json | A list of tags attached to the component |
| links | json | A list of external hyperlinks related to the component |
| type | string | Type of the component |
| lifecycle | string | Lifecycle state of the component |
| owner | string | Owner of the component |
| system | string | System the component belongs to |
| subcomponent_of | string | Parent component this component is part of |
| provides_apis | json | APIs provided by this component |
| consumes_apis | json | APIs consumed by this component |
| depends_on | json | Entities this component depends on |
| dependency_of | json | Entities that depend on this component |

## Examples

### List all components with their systems

```sql
select
  name,
  system,
  owner
from
  backstage_catalog_component
order by
  system;
```

### Find components without owners

```sql
select
  name,
  system,
  description
from
  backstage_catalog_component
where
  owner is null;
```

### Components grouped by system

```sql
select
  system,
  count(*) as component_count
from
  backstage_catalog_component
group by
  system
order by
  component_count desc;
``` 
