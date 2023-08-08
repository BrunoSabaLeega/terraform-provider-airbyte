---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_cumulio Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationCumulio DataSource
---

# airbyte_destination_cumulio (Data Source)

DestinationCumulio DataSource

## Example Usage

```terraform
data "airbyte_destination_cumulio" "my_destination_cumulio" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `api_host` (String) URL of the Cumul.io API (e.g. 'https://api.cumul.io', 'https://api.us.cumul.io', or VPC-specific API url). Defaults to 'https://api.cumul.io'.
- `api_key` (String) An API key generated in Cumul.io's platform (can be generated here: https://app.cumul.io/start/profile/integration).
- `api_token` (String) The corresponding API token generated in Cumul.io's platform (can be generated here: https://app.cumul.io/start/profile/integration).
- `destination_type` (String) must be one of ["cumulio"]

