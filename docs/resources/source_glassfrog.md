---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_glassfrog Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceGlassfrog Resource
---

# airbyte_source_glassfrog (Resource)

SourceGlassfrog Resource

## Example Usage

```terraform
resource "airbyte_source_glassfrog" "my_source_glassfrog" {
  configuration = {
    api_key     = "...my_api_key..."
    source_type = "glassfrog"
  }
  name         = "Carl Davis"
  secret_id    = "...my_secret_id..."
  workspace_id = "891f82ce-1157-4172-b053-77dcfa89df97"
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

- `api_key` (String) API key provided by Glassfrog
- `source_type` (String) must be one of ["glassfrog"]


