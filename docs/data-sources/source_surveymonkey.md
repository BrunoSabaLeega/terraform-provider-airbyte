---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "airbyte_source_surveymonkey Data Source - terraform-provider-airbyte"
subcategory: ""
description: |-
  SourceSurveymonkey DataSource
---

# airbyte_source_surveymonkey (Data Source)

SourceSurveymonkey DataSource

## Example Usage

```terraform
data "airbyte_source_surveymonkey" "my_source_surveymonkey" {
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

- `credentials` (Attributes) The authorization method to use to retrieve data from SurveyMonkey (see [below for nested schema](#nestedatt--configuration--credentials))
- `origin` (String) must be one of ["USA", "Europe", "Canada"]
Depending on the originating datacenter of the SurveyMonkey account, the API access URL may be different.
- `source_type` (String) must be one of ["surveymonkey"]
- `start_date` (String) UTC date and time in the format 2017-01-25T00:00:00Z. Any data before this date will not be replicated.
- `survey_ids` (List of String) IDs of the surveys from which you'd like to replicate data. If left empty, data from all boards to which you have access will be replicated.

<a id="nestedatt--configuration--credentials"></a>
### Nested Schema for `configuration.credentials`

Read-Only:

- `access_token` (String) Access Token for making authenticated requests. See the <a href="https://docs.airbyte.io/integrations/sources/surveymonkey">docs</a> for information on how to generate this key.
- `auth_method` (String) must be one of ["oauth2.0"]
- `client_id` (String) The Client ID of the SurveyMonkey developer application.
- `client_secret` (String) The Client Secret of the SurveyMonkey developer application.


