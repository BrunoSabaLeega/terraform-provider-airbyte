// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type SourceGooglePagespeedInsightsCategories string

const (
	SourceGooglePagespeedInsightsCategoriesAccessibility SourceGooglePagespeedInsightsCategories = "accessibility"
	SourceGooglePagespeedInsightsCategoriesBestPractices SourceGooglePagespeedInsightsCategories = "best-practices"
	SourceGooglePagespeedInsightsCategoriesPerformance   SourceGooglePagespeedInsightsCategories = "performance"
	SourceGooglePagespeedInsightsCategoriesPwa           SourceGooglePagespeedInsightsCategories = "pwa"
	SourceGooglePagespeedInsightsCategoriesSeo           SourceGooglePagespeedInsightsCategories = "seo"
)

func (e SourceGooglePagespeedInsightsCategories) ToPointer() *SourceGooglePagespeedInsightsCategories {
	return &e
}

func (e *SourceGooglePagespeedInsightsCategories) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "accessibility":
		fallthrough
	case "best-practices":
		fallthrough
	case "performance":
		fallthrough
	case "pwa":
		fallthrough
	case "seo":
		*e = SourceGooglePagespeedInsightsCategories(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGooglePagespeedInsightsCategories: %v", v)
	}
}

type SourceGooglePagespeedInsightsGooglePagespeedInsights string

const (
	SourceGooglePagespeedInsightsGooglePagespeedInsightsGooglePagespeedInsights SourceGooglePagespeedInsightsGooglePagespeedInsights = "google-pagespeed-insights"
)

func (e SourceGooglePagespeedInsightsGooglePagespeedInsights) ToPointer() *SourceGooglePagespeedInsightsGooglePagespeedInsights {
	return &e
}

func (e *SourceGooglePagespeedInsightsGooglePagespeedInsights) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "google-pagespeed-insights":
		*e = SourceGooglePagespeedInsightsGooglePagespeedInsights(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGooglePagespeedInsightsGooglePagespeedInsights: %v", v)
	}
}

type SourceGooglePagespeedInsightsStrategies string

const (
	SourceGooglePagespeedInsightsStrategiesDesktop SourceGooglePagespeedInsightsStrategies = "desktop"
	SourceGooglePagespeedInsightsStrategiesMobile  SourceGooglePagespeedInsightsStrategies = "mobile"
)

func (e SourceGooglePagespeedInsightsStrategies) ToPointer() *SourceGooglePagespeedInsightsStrategies {
	return &e
}

func (e *SourceGooglePagespeedInsightsStrategies) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "desktop":
		fallthrough
	case "mobile":
		*e = SourceGooglePagespeedInsightsStrategies(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceGooglePagespeedInsightsStrategies: %v", v)
	}
}

type SourceGooglePagespeedInsights struct {
	// Google PageSpeed API Key. See <a href="https://developers.google.com/speed/docs/insights/v5/get-started#APIKey">here</a>. The key is optional - however the API is heavily rate limited when using without API Key. Creating and using the API key therefore is recommended. The key is case sensitive.
	APIKey *string `json:"api_key,omitempty"`
	// Defines which Lighthouse category to run. One or many of: "accessibility", "best-practices", "performance", "pwa", "seo".
	Categories []SourceGooglePagespeedInsightsCategories            `json:"categories"`
	SourceType SourceGooglePagespeedInsightsGooglePagespeedInsights `json:"sourceType"`
	// The analyses strategy to use. Either "desktop" or "mobile".
	Strategies []SourceGooglePagespeedInsightsStrategies `json:"strategies"`
	// The URLs to retrieve pagespeed information from. The connector will attempt to sync PageSpeed reports for all the defined URLs. Format: https://(www.)url.domain
	Urls []string `json:"urls"`
}
