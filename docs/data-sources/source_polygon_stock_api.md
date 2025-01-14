---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_polygon_stock_api Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourcePolygonStockAPI DataSource
---

# airbyte_source_polygon_stock_api (Data Source)

SourcePolygonStockAPI DataSource

## Example Usage

```terraform
data "airbyte_source_polygon_stock_api" "my_source_polygonstockapi" {
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

- `adjusted` (String) Determines whether or not the results are adjusted for splits. By default, results are adjusted and set to true. Set this to false to get results that are NOT adjusted for splits.
- `api_key` (String) Your API ACCESS Key
- `end_date` (String) The target date for the aggregate window.
- `limit` (Number) The target date for the aggregate window.
- `multiplier` (Number) The size of the timespan multiplier.
- `sort` (String) Sort the results by timestamp. asc will return results in ascending order (oldest at the top), desc will return results in descending order (newest at the top).
- `source_type` (String) must be one of ["polygon-stock-api"]
- `start_date` (String) The beginning date for the aggregate window.
- `stocks_ticker` (String) The exchange symbol that this item is traded under.
- `timespan` (String) The size of the time window.


