---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_auth0 Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceAuth0 Resource
---

# airbyte_source_auth0 (Resource)

SourceAuth0 Resource

## Example Usage

```terraform
resource "airbyte_source_auth0" "my_source_auth0" {
  configuration = {
    base_url = "https://dev-yourOrg.us.auth0.com/"
    credentials = {
      source_auth0_authentication_method_o_auth2_access_token = {
        access_token = "...my_access_token..."
        auth_type    = "oauth2_access_token"
      }
    }
    source_type = "auth0"
    start_date  = "2023-08-05T00:43:59.244Z"
  }
  name         = "Willard McLaughlin"
  secret_id    = "...my_secret_id..."
  workspace_id = "75dad636-c600-4503-98bb-31180f739ae9"
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

- `base_url` (String) The Authentication API is served over HTTPS. All URLs referenced in the documentation have the following base `https://YOUR_DOMAIN`
- `credentials` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials))
- `source_type` (String) must be one of ["auth0"]

Optional:

- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Optional:

- `source_auth0_authentication_method_o_auth2_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_auth0_authentication_method_o_auth2_access_token))
- `source_auth0_authentication_method_o_auth2_confidential_application` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_auth0_authentication_method_o_auth2_confidential_application))
- `source_auth0_update_authentication_method_o_auth2_access_token` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_auth0_update_authentication_method_o_auth2_access_token))
- `source_auth0_update_authentication_method_o_auth2_confidential_application` (Attributes) (see [below for nested schema](#nestedatt--configuration--credentials--source_auth0_update_authentication_method_o_auth2_confidential_application))

<a id="nestedatt--configuration--credentials--source_auth0_authentication_method_o_auth2_access_token"></a>
### Nested Schema for `configuration.credentials.source_auth0_authentication_method_o_auth2_access_token`

Required:

- `access_token` (String) Also called <a href="https://auth0.com/docs/secure/tokens/access-tokens/get-management-api-access-tokens-for-testing">API Access Token </a> The access token used to call the Auth0 Management API Token. It's a JWT that contains specific grant permissions knowns as scopes.
- `auth_type` (String) must be one of ["oauth2_access_token"]


<a id="nestedatt--configuration--credentials--source_auth0_authentication_method_o_auth2_confidential_application"></a>
### Nested Schema for `configuration.credentials.source_auth0_authentication_method_o_auth2_confidential_application`

Required:

- `audience` (String) The audience for the token, which is your API. You can find this in the Identifier field on your  <a href="https://manage.auth0.com/#/apis">API's settings tab</a>
- `auth_type` (String) must be one of ["oauth2_confidential_application"]
- `client_id` (String) Your application's Client ID. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
- `client_secret` (String) Your application's Client Secret. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.


<a id="nestedatt--configuration--credentials--source_auth0_update_authentication_method_o_auth2_access_token"></a>
### Nested Schema for `configuration.credentials.source_auth0_update_authentication_method_o_auth2_access_token`

Required:

- `access_token` (String) Also called <a href="https://auth0.com/docs/secure/tokens/access-tokens/get-management-api-access-tokens-for-testing">API Access Token </a> The access token used to call the Auth0 Management API Token. It's a JWT that contains specific grant permissions knowns as scopes.
- `auth_type` (String) must be one of ["oauth2_access_token"]


<a id="nestedatt--configuration--credentials--source_auth0_update_authentication_method_o_auth2_confidential_application"></a>
### Nested Schema for `configuration.credentials.source_auth0_update_authentication_method_o_auth2_confidential_application`

Required:

- `audience` (String) The audience for the token, which is your API. You can find this in the Identifier field on your  <a href="https://manage.auth0.com/#/apis">API's settings tab</a>
- `auth_type` (String) must be one of ["oauth2_confidential_application"]
- `client_id` (String) Your application's Client ID. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.
- `client_secret` (String) Your application's Client Secret. You can find this value on the <a href="https://manage.auth0.com/#/applications">application's settings tab</a> after you login the admin portal.


