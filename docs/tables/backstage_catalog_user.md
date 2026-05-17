# Table: backstage_catalog_user

A User in Backstage represents an individual person in the organization. Users can be members of groups and can own various entities in the catalog.

## Columns

| Column | Type | Description |
|--------|------|-------------|
| name | string | The name of the user. |
| namespace | string | The namespace the user belongs to. |
| kind | string | The kind of the entity (always "User"). |
| metadata | json | The full metadata of the user. |
| spec | json | The specification data of the user. |
| relations | json | The relations of the user to other entities. |
| title | string | A display name of the user. |
| description | string | A description of the user. |
| labels | json | Labels attached to the user. |
| annotations | json | Annotations attached to the user. |
| tags | json | A list of tags attached to the user. |
| links | json | A list of external hyperlinks related to the user. |
| display_name | string | Display name of the user profile. |
| email | string | Email address of the user. |
| picture | string | Picture URL of the user. |
| member_of | json | Groups the user belongs to. |

## Examples

### Basic info
```sql
select
  name,
  email,
  title
from
  backstage_catalog_user;
```

### List users and their groups
```sql
select
  name as user_name,
  member_of
from
  backstage_catalog_user
where
  member_of is not null;
```

### Find users with specific roles
```sql
select
  name,
  email
from
  backstage_catalog_user
where
  spec->>'role' = 'admin';
``` 
