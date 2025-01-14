---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_younium Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceYounium DataSource
---

# airbyte_source_younium (Data Source)

SourceYounium DataSource

## Example Usage

```terraform
data "airbyte_source_younium" "my_source_younium" {
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

- `legal_entity` (String) Legal Entity that data should be pulled from
- `password` (String) Account password for younium account API key
- `playground` (Boolean) Property defining if connector is used against playground or production environment
- `source_type` (String) must be one of ["younium"]
- `username` (String) Username for Younium account


