---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_square Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSquare DataSource
---

# airbyte_source_square (Data Source)

SourceSquare DataSource

## Example Usage

```terraform
data "airbyte_source_square" "my_source_square" {
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

- `credentials` (Attributes) Choose how to authenticate to Square. (see [below for nested schema](#nestedatt--configuration--credentials))
- `include_deleted_objects` (Boolean) In some streams there is an option to include deleted objects (Items, Categories, Discounts, Taxes)
- `is_sandbox` (Boolean) Determines whether to use the sandbox or production environment.
- `source_type` (String) must be one of ["square"]
- `start_date` (String) UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated. If not set, all data will be replicated.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_square_authentication_api_key` (Attributes) Choose how to authenticate to Square. (see [below for nested schema](#nestedatt--configuration--credentials--source_square_authentication_api_key))
- `source_square_authentication_oauth_authentication` (Attributes) Choose how to authenticate to Square. (see [below for nested schema](#nestedatt--configuration--credentials--source_square_authentication_oauth_authentication))
- `source_square_update_authentication_api_key` (Attributes) Choose how to authenticate to Square. (see [below for nested schema](#nestedatt--configuration--credentials--source_square_update_authentication_api_key))
- `source_square_update_authentication_oauth_authentication` (Attributes) Choose how to authenticate to Square. (see [below for nested schema](#nestedatt--configuration--credentials--source_square_update_authentication_oauth_authentication))

<a id="nestedatt--configuration--credentials--source_square_authentication_api_key"></a>
### Nested Schema for `configuration.credentials.source_square_authentication_api_key`

Read-Only:

- `api_key` (String) The API key for a Square application
- `auth_type` (String) must be one of ["API Key"]


<a id="nestedatt--configuration--credentials--source_square_authentication_oauth_authentication"></a>
### Nested Schema for `configuration.credentials.source_square_authentication_oauth_authentication`

Read-Only:

- `auth_type` (String) must be one of ["OAuth"]
- `client_id` (String) The Square-issued ID of your application
- `client_secret` (String) The Square-issued application secret for your application
- `refresh_token` (String) A refresh token generated using the above client ID and secret


<a id="nestedatt--configuration--credentials--source_square_update_authentication_api_key"></a>
### Nested Schema for `configuration.credentials.source_square_update_authentication_api_key`

Read-Only:

- `api_key` (String) The API key for a Square application
- `auth_type` (String) must be one of ["API Key"]


<a id="nestedatt--configuration--credentials--source_square_update_authentication_oauth_authentication"></a>
### Nested Schema for `configuration.credentials.source_square_update_authentication_oauth_authentication`

Read-Only:

- `auth_type` (String) must be one of ["OAuth"]
- `client_id` (String) The Square-issued ID of your application
- `client_secret` (String) The Square-issued application secret for your application
- `refresh_token` (String) A refresh token generated using the above client ID and secret

