---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_mixpanel Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceMixpanel DataSource
---

# airbyte_source_mixpanel (Data Source)

SourceMixpanel DataSource

## Example Usage

```terraform
data "airbyte_source_mixpanel" "my_source_mixpanel" {
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

- `attribution_window` (Number) A period of time for attributing results to ads and the lookback period after those actions occur during which ad results are counted. Default attribution window is 5 days.
- `credentials` (Attributes) Choose how to authenticate to Mixpanel (see [below for nested schema](#nestedatt--configuration--credentials))
- `date_window_size` (Number) Defines window size in days, that used to slice through data. You can reduce it, if amount of data in each window is too big for your environment.
- `end_date` (String) The date in the format YYYY-MM-DD. Any data after this date will not be replicated. Left empty to always sync to most recent date
- `project_id` (Number) Your project ID number. See the <a href="https://help.mixpanel.com/hc/en-us/articles/115004490503-Project-Settings#project-id">docs</a> for more information on how to obtain this.
- `project_timezone` (String) Time zone in which integer date times are stored. The project timezone may be found in the project settings in the <a href="https://help.mixpanel.com/hc/en-us/articles/115004547203-Manage-Timezones-for-Projects-in-Mixpanel">Mixpanel console</a>.
- `region` (String) must be one of ["US", "EU"]
The region of mixpanel domain instance either US or EU.
- `select_properties_by_default` (Boolean) Setting this config parameter to TRUE ensures that new properties on events and engage records are captured. Otherwise new properties will be ignored.
- `source_type` (String) must be one of ["mixpanel"]
- `start_date` (String) The date in the format YYYY-MM-DD. Any data before this date will not be replicated. If this option is not set, the connector will replicate data from up to one year ago by default.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_mixpanel_authentication_wildcard_project_secret` (Attributes) Choose how to authenticate to Mixpanel (see [below for nested schema](#nestedatt--configuration--credentials--source_mixpanel_authentication_wildcard_project_secret))
- `source_mixpanel_authentication_wildcard_service_account` (Attributes) Choose how to authenticate to Mixpanel (see [below for nested schema](#nestedatt--configuration--credentials--source_mixpanel_authentication_wildcard_service_account))
- `source_mixpanel_update_authentication_wildcard_project_secret` (Attributes) Choose how to authenticate to Mixpanel (see [below for nested schema](#nestedatt--configuration--credentials--source_mixpanel_update_authentication_wildcard_project_secret))
- `source_mixpanel_update_authentication_wildcard_service_account` (Attributes) Choose how to authenticate to Mixpanel (see [below for nested schema](#nestedatt--configuration--credentials--source_mixpanel_update_authentication_wildcard_service_account))

<a id="nestedatt--configuration--credentials--source_mixpanel_authentication_wildcard_project_secret"></a>
### Nested Schema for `configuration.credentials.source_mixpanel_authentication_wildcard_project_secret`

Read-Only:

- `api_secret` (String) Mixpanel project secret. See the <a href="https://developer.mixpanel.com/reference/project-secret#managing-a-projects-secret">docs</a> for more information on how to obtain this.
- `option_title` (String) must be one of ["Project Secret"]


<a id="nestedatt--configuration--credentials--source_mixpanel_authentication_wildcard_service_account"></a>
### Nested Schema for `configuration.credentials.source_mixpanel_authentication_wildcard_service_account`

Read-Only:

- `option_title` (String) must be one of ["Service Account"]
- `secret` (String) Mixpanel Service Account Secret. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.
- `username` (String) Mixpanel Service Account Username. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.


<a id="nestedatt--configuration--credentials--source_mixpanel_update_authentication_wildcard_project_secret"></a>
### Nested Schema for `configuration.credentials.source_mixpanel_update_authentication_wildcard_project_secret`

Read-Only:

- `api_secret` (String) Mixpanel project secret. See the <a href="https://developer.mixpanel.com/reference/project-secret#managing-a-projects-secret">docs</a> for more information on how to obtain this.
- `option_title` (String) must be one of ["Project Secret"]


<a id="nestedatt--configuration--credentials--source_mixpanel_update_authentication_wildcard_service_account"></a>
### Nested Schema for `configuration.credentials.source_mixpanel_update_authentication_wildcard_service_account`

Read-Only:

- `option_title` (String) must be one of ["Service Account"]
- `secret` (String) Mixpanel Service Account Secret. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.
- `username` (String) Mixpanel Service Account Username. See the <a href="https://developer.mixpanel.com/reference/service-accounts">docs</a> for more information on how to obtain this.


