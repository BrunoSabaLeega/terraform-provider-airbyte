---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_braze Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceBraze Resource
---

# airbyte_source_braze (Resource)

SourceBraze Resource

## Example Usage

```terraform
resource "airbyte_source_braze" "my_source_braze" {
  configuration = {
    api_key     = "...my_api_key..."
    source_type = "braze"
    start_date  = "2022-09-06"
    url         = "...my_url..."
  }
  name         = "Rosie Glover"
  secret_id    = "...my_secret_id..."
  workspace_id = "efc5fde1-0a0c-4e21-a9e5-10019c6dc5e3"
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

- `api_key` (String) Braze REST API key
- `source_type` (String) must be one of ["braze"]
- `start_date` (String) Rows after this date will be synced
- `url` (String) Braze REST API endpoint


