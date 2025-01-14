---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_destination_firestore Resource - terraform-provider-airbyte"
subcategory: ""
description: |-
  DestinationFirestore Resource
---

# airbyte_destination_firestore (Resource)

DestinationFirestore Resource

## Example Usage

```terraform
resource "airbyte_destination_firestore" "my_destination_firestore" {
  configuration = {
    credentials_json = "...my_credentials_json..."
    destination_type = "firestore"
    project_id       = "...my_project_id..."
  }
  name         = "Paula Jacobs I"
  workspace_id = "f16d9f5f-ce6c-4556-946c-3e250fb008c4"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `configuration` (Attributes) (see [below for nested schema](#nestedatt--configuration))
- `name` (String)
- `workspace_id` (String)

### Read-Only

- `destination_id` (String)
- `destination_type` (String)

<a id="nestedatt--configuration"></a>
### Nested Schema for `configuration`

Required:

- `destination_type` (String) must be one of ["firestore"]
- `project_id` (String) The GCP project ID for the project containing the target BigQuery dataset.

Optional:

- `credentials_json` (String) The contents of the JSON service account key. Check out the <a href="https://docs.airbyte.io/integrations/destinations/firestore">docs</a> if you need help generating this key. Default credentials will be used if this field is left empty.


