---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_sonar_cloud Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSonarCloud DataSource
---

# airbyte_source_sonar_cloud (Data Source)

SourceSonarCloud DataSource

## Example Usage

```terraform
data "airbyte_source_sonar_cloud" "my_source_sonarcloud" {
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

- `component_keys` (List of String) Comma-separated list of component keys.
- `end_date` (String) To retrieve issues created before the given date (inclusive).
- `organization` (String) Organization key. See <a href="https://docs.sonarcloud.io/appendices/project-information/#project-and-organization-keys">here</a>.
- `source_type` (String) must be one of ["sonar-cloud"]
- `start_date` (String) To retrieve issues created after the given date (inclusive).
- `user_token` (String) Your User Token. See <a href="https://docs.sonarcloud.io/advanced-setup/user-accounts/">here</a>. The token is case sensitive.


