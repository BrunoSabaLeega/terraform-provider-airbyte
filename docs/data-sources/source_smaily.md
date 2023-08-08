---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_smaily Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSmaily DataSource
---

# airbyte_source_smaily (Data Source)

SourceSmaily DataSource

## Example Usage

```terraform
data "airbyte_source_smaily" "my_source_smaily" {
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

- `api_password` (String) API user password. See https://smaily.com/help/api/general/create-api-user/
- `api_subdomain` (String) API Subdomain. See https://smaily.com/help/api/general/create-api-user/
- `api_username` (String) API user username. See https://smaily.com/help/api/general/create-api-user/
- `source_type` (String) must be one of ["smaily"]

