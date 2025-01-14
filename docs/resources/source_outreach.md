---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_outreach Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceOutreach Resource
---

# airbyte_source_outreach (Resource)

SourceOutreach Resource

## Example Usage

```terraform
resource "airbyte_source_outreach" "my_source_outreach" {
  configuration = {
    client_id     = "...my_client_id..."
    client_secret = "...my_client_secret..."
    redirect_uri  = "...my_redirect_uri..."
    refresh_token = "...my_refresh_token..."
    source_type   = "outreach"
    start_date    = "2020-11-16T00:00:00Z"
  }
  name         = "Kim Kirlin"
  secret_id    = "...my_secret_id..."
  workspace_id = "8e0cc885-187e-44de-84af-28c5dddb46aa"
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

- `client_id` (String) The Client ID of your Outreach developer application.
- `client_secret` (String) The Client Secret of your Outreach developer application.
- `redirect_uri` (String) A Redirect URI is the location where the authorization server sends the user once the app has been successfully authorized and granted an authorization code or access token.
- `refresh_token` (String) The token for obtaining the new access token.
- `source_type` (String) must be one of ["outreach"]
- `start_date` (String) The date from which you'd like to replicate data for Outreach API, in the format YYYY-MM-DDT00:00:00Z. All data generated after this date will be replicated.


