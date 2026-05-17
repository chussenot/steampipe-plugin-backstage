# Table: backstage_catalog_group

A Group in Backstage represents an organizational team or unit. Groups can own components, systems, and other entities, and can have members (users) assigned to them.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the group. |
| namespace | string | The namespace the group belongs to. |
| kind | string | The kind of the entity (always "Group"). |
| metadata | json | The full metadata of the group. |
| spec | json | The specification data of the group. |
| relations | json | The relations of the group to other entities. |
| title | string | A display name of the group. |
| description | string | A description of the group. |
| labels | json | Labels attached to the group. |
| annotations | json | Annotations attached to the group. |
| tags | json | A list of tags attached to the group. |
| links | json | A list of external hyperlinks related to the group. |
| parent | string | Parent group of this group. |
| children | json | Child groups of this group. |
| members | json | Members belonging to this group. |
| display_name | string | Display name of the group profile. |
| email | string | Email address from the group profile. |
| picture | string | Picture URL from the group profile. |

## Examples

### Basic info

```sql
select
  name,
  description,
  parent
from
  backstage_catalog_group;
```

### List groups with their member count

```sql
select
  name as group_name,
  jsonb_array_length(coalesce(members, '[]'::jsonb)) as member_count
from
  backstage_catalog_group
order by
  member_count desc;
```

### Find groups without parent groups

```sql
select
  name,
  description
from
  backstage_catalog_group
where
  parent is null;
```
