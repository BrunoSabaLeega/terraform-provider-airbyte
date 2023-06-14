// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type SourceFreshserviceUpdate struct {
	// Freshservice API Key. See <a href="https://api.freshservice.com/#authentication">here</a>. The key is case sensitive.
	APIKey string `json:"api_key"`
	// The name of your Freshservice domain
	DomainName string `json:"domain_name"`
	// UTC date and time in the format 2020-10-01T00:00:00Z. Any data before this date will not be replicated.
	StartDate string `json:"start_date"`
}