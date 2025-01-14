---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_braze Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceBraze DataSource
---

# airbyte_source_braze (Data Source)

SourceBraze DataSource

## Example Usage

```terraform
data "airbyte_source_braze" "my_source_braze" {
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

- `api_key` (String) Braze REST API key
- `source_type` (String) must be one of ["braze"]
- `start_date` (String) Rows after this date will be synced
- `url` (String) Braze REST API endpoint


