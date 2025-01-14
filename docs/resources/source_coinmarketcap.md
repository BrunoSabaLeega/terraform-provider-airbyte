---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_coinmarketcap Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceCoinmarketcap Resource
---

# airbyte_source_coinmarketcap (Resource)

SourceCoinmarketcap Resource

## Example Usage

```terraform
resource "airbyte_source_coinmarketcap" "my_source_coinmarketcap" {
  configuration = {
    api_key     = "...my_api_key..."
    data_type   = "historical"
    source_type = "coinmarketcap"
    symbols = [
      "...",
    ]
  }
  name         = "Meredith Kassulke"
  secret_id    = "...my_secret_id..."
  workspace_id = "1804e54c-82f1-468a-b63c-8873e484380b"
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

- `api_key` (String) Your API Key. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Authentication">here</a>. The token is case sensitive.
- `data_type` (String) must be one of ["latest", "historical"]
/latest: Latest market ticker quotes and averages for cryptocurrencies and exchanges. /historical: Intervals of historic market data like OHLCV data or data for use in charting libraries. See <a href="https://coinmarketcap.com/api/documentation/v1/#section/Endpoint-Overview">here</a>.
- `source_type` (String) must be one of ["coinmarketcap"]

Optional:

- `symbols` (List of String) Cryptocurrency symbols. (only used for quotes stream)


