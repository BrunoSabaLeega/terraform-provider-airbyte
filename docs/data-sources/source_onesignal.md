---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_onesignal Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceOnesignal DataSource
---

# airbyte_source_onesignal (Data Source)

SourceOnesignal DataSource

## Example Usage

```terraform
data "airbyte_source_onesignal" "my_source_onesignal" {
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

- `applications` (Attributes List) Applications keys, see the <a href="https://documentation.onesignal.com/docs/accounts-and-keys">docs</a> for more information on how to obtain this data (see [below for nested schema](#nestedatt--configuration--applications))
- `outcome_names` (String) Comma-separated list of names and the value (sum/count) for the returned outcome data. See the <a href="https://documentation.onesignal.com/reference/view-outcomes">docs</a> for more details
- `source_type` (String) must be one of ["onesignal"]
- `start_date` (String) The date from which you'd like to replicate data for OneSignal API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.
- `user_auth_key` (String) OneSignal User Auth Key, see the <a href="https://documentation.onesignal.com/docs/accounts-and-keys#user-auth-key">docs</a> for more information on how to obtain this key.

<a id="nestedatt--configuration--applications"></a>
### Nested Schema for `configuration.applications`

Read-Only:

- `app_api_key` (String)
- `app_id` (String)
- `app_name` (String)


