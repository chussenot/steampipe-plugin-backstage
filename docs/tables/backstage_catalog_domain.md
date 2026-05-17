# Table: backstage_catalog_domain

A Domain represents the highest level of organization in your software ecosystem. Domains typically group together systems and components that serve a similar business function.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the domain. |
| namespace | string | The namespace the domain belongs to. |
| kind | string | The kind of the entity (always "Domain"). |
| metadata | json | The full metadata of the domain. |
| spec | json | The specification data of the domain. |
| relations | json | The relations of the domain to other entities. |
| title | string | A display name of the domain. |
| description | string | A description of the domain. |
| labels | json | Labels attached to the domain. |
| annotations | json | Annotations attached to the domain. |
| tags | json | A list of tags attached to the domain. |
| links | json | A list of external hyperlinks related to the domain. |
| owner | string | Owner of the domain. |
| subdomain_of | string | Parent domain this domain is part of. |
| type | string | Type of the domain. |

## Examples

### Basic info

```sql
select
  name,
  owner,
  description
from
  backstage_catalog_domain;
```

### List domains with their systems

```sql
select
  d.name as domain_name,
  count(s.name) as system_count
from
  backstage_catalog_domain as d
  left join backstage_catalog_system as s on s.domain = d.name
group by
  d.name
order by
  system_count desc;
```

### Find domains by owner

```sql
select
  name,
  description,
  owner
from
  backstage_catalog_domain
where
  owner = 'team-a';
``` 
