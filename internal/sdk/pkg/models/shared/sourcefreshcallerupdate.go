// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type SourceFreshcallerUpdate struct {
	// Freshcaller API Key. See the <a href="https://docs.airbyte.com/integrations/sources/freshcaller">docs</a> for more information on how to obtain this key.
	APIKey string `json:"api_key"`
	// Used to construct Base URL for the Freshcaller APIs
	Domain string `json:"domain"`
	// The number of requests per minute that this source allowed to use. There is a rate limit of 50 requests per minute per app per account.
	RequestsPerMinute *int64 `json:"requests_per_minute,omitempty"`
	// UTC date and time. Any data created after this date will be replicated.
	StartDate time.Time `json:"start_date"`
	// Lag in minutes for each sync, i.e., at time T, data for the time range [prev_sync_time, T-30] will be fetched
	SyncLagMinutes *int64 `json:"sync_lag_minutes,omitempty"`
}