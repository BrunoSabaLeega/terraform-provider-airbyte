---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_netsuite Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceNetsuite Resource
---

# airbyte_source_netsuite (Resource)

SourceNetsuite Resource

## Example Usage

```terraform
resource "airbyte_source_netsuite" "my_source_netsuite" {
  configuration = {
    consumer_key    = "...my_consumer_key..."
    consumer_secret = "...my_consumer_secret..."
    object_types = [
      "...",
    ]
    realm          = "...my_realm..."
    source_type    = "netsuite"
    start_datetime = "2017-01-25T00:00:00Z"
    token_key      = "...my_token_key..."
    token_secret   = "...my_token_secret..."
    window_in_days = 7
  }
  name         = "Miss Meredith Hand"
  secret_id    = "...my_secret_id..."
  workspace_id = "4bf01bad-8706-4d46-882b-fbdc41ff5d4e"
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

- `consumer_key` (String) Consumer key associated with your integration
- `consumer_secret` (String) Consumer secret associated with your integration
- `realm` (String) Netsuite realm e.g. 2344535, as for `production` or 2344535_SB1, as for the `sandbox`
- `source_type` (String) must be one of ["netsuite"]
- `start_datetime` (String) Starting point for your data replication, in format of "YYYY-MM-DDTHH:mm:ssZ"
- `token_key` (String) Access token key
- `token_secret` (String) Access token secret

Optional:

- `object_types` (List of String) The API names of the Netsuite objects you want to sync. Setting this speeds up the connection setup process by limiting the number of schemas that need to be retrieved from Netsuite.
- `window_in_days` (Number) The amount of days used to query the data with date chunks. Set smaller value, if you have lots of data.


