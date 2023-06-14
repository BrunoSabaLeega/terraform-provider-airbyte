---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_tvmaze_schedule Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceTvmazeSchedule Resource
---

# airbyte_source_tvmaze_schedule (Resource)

SourceTvmazeSchedule Resource



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Optional

- `secret_id` (String)

### Read-Only

- `source_id` (String)
- `source_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `domestic_schedule_country_code` (String)
- `source_type` (String)
- `start_date` (String)

Optional:

- `end_date` (String)
- `web_schedule_country_code` (String)

