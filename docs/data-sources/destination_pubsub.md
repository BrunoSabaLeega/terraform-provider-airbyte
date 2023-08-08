---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_pubsub Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationPubsub DataSource
---

# airbyte_destination_pubsub (Data Source)

DestinationPubsub DataSource

## Example Usage

```terraform
data "airbyte_destination_pubsub" "my_destination_pubsub" {
  destination_id = "...my_destination_id..."
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `destination_id` (String)

### Read-Only

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Read-Only:

- `batching_delay_threshold` (Number) Number of ms before the buffer is flushed
- `batching_element_count_threshold` (Number) Number of messages before the buffer is flushed
- `batching_enabled` (Boolean) If TRUE messages will be buffered instead of sending them one by one
- `batching_request_bytes_threshold` (Number) Number of bytes before the buffer is flushed
- `credentials_json` (String) The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.com/integrations/destinations/pubsub">docs</a> if you need help generating this key.
- `destination_type` (String) must be one of ["pubsub"]
- `ordering_enabled` (Boolean) If TRUE PubSub publisher will have <a href="https://cloud.google.com/pubsub/docs/ordering">message ordering</a> enabled. Every message will have an ordering key of stream
- `project_id` (String) The GCP project ID for the project containing the target PubSub.
- `topic_id` (String) The PubSub topic ID in the given GCP project ID.

