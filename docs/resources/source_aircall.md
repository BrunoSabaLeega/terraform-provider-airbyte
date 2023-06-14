---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_aircall Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAircall Resource
---

# airbyte_source_aircall (Resource)

SourceAircall Resource



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

- `api_id` (String)
- `api_token` (String)
- `source_type` (String)
- `start_date` (String)

