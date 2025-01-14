---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_convex Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceConvex Resource
---

# airbyte_source_convex (Resource)

SourceConvex Resource

## Example Usage

```terraform
resource "airbyte_source_convex" "my_source_convex" {
  configuration = {
    access_key     = "...my_access_key..."
    deployment_url = "https://murky-swan-635.convex.cloud"
    source_type    = "convex"
  }
  name         = "Guy Kovacek"
  secret_id    = "...my_secret_id..."
  workspace_id = "a8581a58-208c-454f-afa9-c95f2eac5565"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Optional

- `secret_id` (String) Optional secretID obtained through the public API OAuth redirect flow.

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `access_key` (String) API access key used to retrieve data from Convex.
- `deployment_url` (String)
- `source_type` (String) must be one of ["convex"]


