---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_pypi Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePypi DataSource
---

# airbyte_source_pypi (Data Source)

SourcePypi DataSource

## Example Usage

```terraform
data "airbyte_source_pypi" "my_source_pypi" {
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

- `project_name` (String) Name of the project/package. Can only be in lowercase with hyphen. This is the name used using pip command for installing the package.
- `source_type` (String) must be one of ["pypi"]
- `version` (String) Version of the project/package.  Use it to find a particular release instead of all releases.


