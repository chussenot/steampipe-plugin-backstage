# Table: backstage_catalog_api

A Backstage API entity represents an interface that can be exposed by a component. APIs are defined by a specification in a machine-readable format, such as OpenAPI, GraphQL, gRPC, or AsyncAPI.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the API. |
| namespace | string | The namespace the API belongs to. |
| kind | string | The kind of the entity (always "API"). |
| metadata | json | The full metadata of the API. |
| spec | json | The specification data of the API. |
| relations | json | The relations of the API to other entities. |
| title | string | A display name of the API. |
| description | string | A description of the API. |
| labels | json | Labels attached to the API. |
| annotations | json | Annotations attached to the API. |
| tags | json | A list of tags attached to the API. |
| links | json | A list of external hyperlinks related to the API. |
| owner | string | Owner of the API. |
| system | string | System the API belongs to. |
| definition | json | API definition. |
| type | string | Type of the API. |
| lifecycle | string | Lifecycle state of the API. |

## Examples

### Basic info

```sql
select
  name,
  type,
  lifecycle,
  owner
from
  backstage_catalog_api;
```

### List APIs by type

```sql
select
  type,
  count(*) as api_count
from
  backstage_catalog_api
group by
  type
order by
  api_count desc;
```

### List APIs without owners

```sql
select
  name,
  type,
  description
from
  backstage_catalog_api
where
  owner is null;
```

### Get APIs with specific annotation

```sql
select
  name,
  annotations->>'backstage.io/techdocs-ref' as techdocs_ref
from
  backstage_catalog_api
where
  annotations->>'backstage.io/techdocs-ref' is not null;
``` 
