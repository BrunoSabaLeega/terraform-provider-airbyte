---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_mongodb_internal_poc Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceMongodbInternalPoc DataSource
---

# airbyte_source_mongodb_internal_poc (Data Source)

SourceMongodbInternalPoc DataSource

## Example Usage

```terraform
data "airbyte_source_mongodb_internal_poc" "my_source_mongodbinternalpoc" {
  secret_id = "...my_secret_id..."
  source_id = "...my_source_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `source_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `auth_source` (String) The authentication source where the user information is stored.
- `connection_string` (String) The connection string of the database that you want to replicate..
- `password` (String) The password associated with this username.
- `replica_set` (String) The name of the replica set to be replicated.
- `source_type` (String) must be one of ["mongodb-internal-poc"]
- `user` (String) The username which is used to access the database.


