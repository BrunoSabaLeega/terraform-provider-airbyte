---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_nytimes Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceNytimes Resource
---

# airbyte_source_nytimes (Resource)

SourceNytimes Resource

## Example Usage

```terraform
resource "airbyte_source_nytimes" "my_source_nytimes" {
  configuration = {
    api_key     = "...my_api_key..."
    end_date    = "1851-01"
    period      = "7"
    share_type  = "facebook"
    source_type = "nytimes"
    start_date  = "2022-08"
  }
  name         = "Mr. Emily Macejkovic"
  secret_id    = "...my_secret_id..."
  workspace_id = "4fe44472-97cd-43b1-9d3b-bce247b7684e"
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

- `api_key` (String) API Key
- `period` (Number) must be one of ["1", "7", "30"]
Period of time (in days)
- `source_type` (String) must be one of ["nytimes"]
- `start_date` (String) Start date to begin the article retrieval (format YYYY-MM)

Optional:

- `end_date` (String) End date to stop the article retrieval (format YYYY-MM)
- `share_type` (String) must be one of ["facebook"]
Share Type


