---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_survey_sparrow Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSurveySparrow DataSource
---

# airbyte_source_survey_sparrow (Data Source)

SourceSurveySparrow DataSource

## Example Usage

```terraform
data "airbyte_source_survey_sparrow" "my_source_surveysparrow" {
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

- `access_token` (String) Your access token. See <a href="https://developers.surveysparrow.com/rest-apis#authentication">here</a>. The key is case sensitive.
- `region` (Attributes) Is your account location is EU based? If yes, the base url to retrieve data will be different. (see [below for nested schema](#nestedatt--configuration--region))
- `source_type` (String) must be one of ["survey-sparrow"]
- `survey_id` (List of String) A List of your survey ids for survey-specific stream

<a id="nestedatt--configuration--region"></a>
### Nested Schema for `configuration.region`

Read-Only:

- `source_survey_sparrow_base_url_eu_based_account` (Attributes) Is your account location is EU based? If yes, the base url to retrieve data will be different. (see [below for nested schema](#nestedatt--configuration--region--source_survey_sparrow_base_url_eu_based_account))
- `source_survey_sparrow_base_url_global_account` (Attributes) Is your account location is EU based? If yes, the base url to retrieve data will be different. (see [below for nested schema](#nestedatt--configuration--region--source_survey_sparrow_base_url_global_account))
- `source_survey_sparrow_update_base_url_eu_based_account` (Attributes) Is your account location is EU based? If yes, the base url to retrieve data will be different. (see [below for nested schema](#nestedatt--configuration--region--source_survey_sparrow_update_base_url_eu_based_account))
- `source_survey_sparrow_update_base_url_global_account` (Attributes) Is your account location is EU based? If yes, the base url to retrieve data will be different. (see [below for nested schema](#nestedatt--configuration--region--source_survey_sparrow_update_base_url_global_account))

<a id="nestedatt--configuration--region--source_survey_sparrow_base_url_eu_based_account"></a>
### Nested Schema for `configuration.region.source_survey_sparrow_base_url_eu_based_account`

Read-Only:

- `url_base` (String) must be one of ["https://eu-api.surveysparrow.com/v3"]


<a id="nestedatt--configuration--region--source_survey_sparrow_base_url_global_account"></a>
### Nested Schema for `configuration.region.source_survey_sparrow_base_url_global_account`

Read-Only:

- `url_base` (String) must be one of ["https://api.surveysparrow.com/v3"]


<a id="nestedatt--configuration--region--source_survey_sparrow_update_base_url_eu_based_account"></a>
### Nested Schema for `configuration.region.source_survey_sparrow_update_base_url_eu_based_account`

Read-Only:

- `url_base` (String) must be one of ["https://eu-api.surveysparrow.com/v3"]


<a id="nestedatt--configuration--region--source_survey_sparrow_update_base_url_global_account"></a>
### Nested Schema for `configuration.region.source_survey_sparrow_update_base_url_global_account`

Read-Only:

- `url_base` (String) must be one of ["https://api.surveysparrow.com/v3"]


