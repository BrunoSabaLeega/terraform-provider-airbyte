// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type DestinationMeilisearchUpdate struct {
	// MeiliSearch API Key. See the <a href="https://docs.airbyte.com/integrations/destinations/meilisearch">docs</a> for more information on how to obtain this key.
	APIKey *string `json:"api_key,omitempty"`
	// Hostname of the MeiliSearch instance.
	Host string `json:"host"`
}