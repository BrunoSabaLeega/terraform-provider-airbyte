---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_workspace Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  Workspace Resource
---

# airbyte_workspace (Resource)

Workspace Resource

## Example Usage

```terraform
resource "airbyte_workspace" "my_workspace" {
  name = "Glenda Schiller DDS"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `name` (String) Name of the workspace

### Read-Only

- `data_residency` (String) must be one of ["auto", "us", "eu"]
- `workspace_id` (String)


