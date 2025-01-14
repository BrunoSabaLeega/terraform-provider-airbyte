---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_asana Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAsana DataSource
---

# airbyte_source_asana (Data Source)

SourceAsana DataSource

## Example Usage

```terraform
data "airbyte_source_asana" "my_source_asana" {
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

- `credentials` (Attributes) Choose how to authenticate to Github (see [below for nested schema](#nestedatt--configuration--credentials))
- `source_type` (String) must be one of ["asana"]

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `source_asana_authentication_mechanism_authenticate_via_asana_oauth` (Attributes) Choose how to authenticate to Github (see [below for nested schema](#nestedatt--configuration--credentials--source_asana_authentication_mechanism_authenticate_via_asana_oauth))
- `source_asana_authentication_mechanism_authenticate_with_personal_access_token` (Attributes) Choose how to authenticate to Github (see [below for nested schema](#nestedatt--configuration--credentials--source_asana_authentication_mechanism_authenticate_with_personal_access_token))
- `source_asana_update_authentication_mechanism_authenticate_via_asana_oauth` (Attributes) Choose how to authenticate to Github (see [below for nested schema](#nestedatt--configuration--credentials--source_asana_update_authentication_mechanism_authenticate_via_asana_oauth))
- `source_asana_update_authentication_mechanism_authenticate_with_personal_access_token` (Attributes) Choose how to authenticate to Github (see [below for nested schema](#nestedatt--configuration--credentials--source_asana_update_authentication_mechanism_authenticate_with_personal_access_token))

<a id="nestedatt--configuration--credentials--source_asana_authentication_mechanism_authenticate_via_asana_oauth"></a>
### Nested Schema for `configuration.credentials.source_asana_authentication_mechanism_authenticate_via_asana_oauth`

Read-Only:

- `client_id` (String)
- `client_secret` (String)
- `option_title` (String) must be one of ["OAuth Credentials"]
OAuth Credentials
- `refresh_token` (String)


<a id="nestedatt--configuration--credentials--source_asana_authentication_mechanism_authenticate_with_personal_access_token"></a>
### Nested Schema for `configuration.credentials.source_asana_authentication_mechanism_authenticate_with_personal_access_token`

Read-Only:

- `option_title` (String) must be one of ["PAT Credentials"]
PAT Credentials
- `personal_access_token` (String) Asana Personal Access Token (generate yours <a href="https://app.asana.com/0/developer-console">here</a>).


<a id="nestedatt--configuration--credentials--source_asana_update_authentication_mechanism_authenticate_via_asana_oauth"></a>
### Nested Schema for `configuration.credentials.source_asana_update_authentication_mechanism_authenticate_via_asana_oauth`

Read-Only:

- `client_id` (String)
- `client_secret` (String)
- `option_title` (String) must be one of ["OAuth Credentials"]
OAuth Credentials
- `refresh_token` (String)


<a id="nestedatt--configuration--credentials--source_asana_update_authentication_mechanism_authenticate_with_personal_access_token"></a>
### Nested Schema for `configuration.credentials.source_asana_update_authentication_mechanism_authenticate_with_personal_access_token`

Read-Only:

- `option_title` (String) must be one of ["PAT Credentials"]
PAT Credentials
- `personal_access_token` (String) Asana Personal Access Token (generate yours <a href="https://app.asana.com/0/developer-console">here</a>).


