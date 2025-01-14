// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"airbyte/internal/sdk/pkg/types"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

// SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory - Choose a category to pivot your analytics report around. This selection will organize your data based on the chosen attribute, allowing you to analyze trends and performance from different perspectives.
type SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory string

const (
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryCompany                     SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "COMPANY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryAccount                     SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "ACCOUNT"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryShare                       SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "SHARE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryCampaign                    SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "CAMPAIGN"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryCreative                    SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "CREATIVE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryCampaignGroup               SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "CAMPAIGN_GROUP"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryConversion                  SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "CONVERSION"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryConversationNode            SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "CONVERSATION_NODE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryConversationNodeOptionIndex SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "CONVERSATION_NODE_OPTION_INDEX"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryServingLocation             SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "SERVING_LOCATION"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryCardIndex                   SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "CARD_INDEX"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberCompanySize           SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_COMPANY_SIZE"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberIndustry              SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_INDUSTRY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberSeniority             SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_SENIORITY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberJobTitle              SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_JOB_TITLE "
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberJobFunction           SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_JOB_FUNCTION "
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberCountryV2             SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_COUNTRY_V2 "
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberRegionV2              SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_REGION_V2"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryMemberCompany               SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "MEMBER_COMPANY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryPlacementName               SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "PLACEMENT_NAME"
	SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategoryImpressionDeviceType        SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory = "IMPRESSION_DEVICE_TYPE"
)

func (e SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory) ToPointer() *SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory {
	return &e
}

func (e *SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "COMPANY":
		fallthrough
	case "ACCOUNT":
		fallthrough
	case "SHARE":
		fallthrough
	case "CAMPAIGN":
		fallthrough
	case "CREATIVE":
		fallthrough
	case "CAMPAIGN_GROUP":
		fallthrough
	case "CONVERSION":
		fallthrough
	case "CONVERSATION_NODE":
		fallthrough
	case "CONVERSATION_NODE_OPTION_INDEX":
		fallthrough
	case "SERVING_LOCATION":
		fallthrough
	case "CARD_INDEX":
		fallthrough
	case "MEMBER_COMPANY_SIZE":
		fallthrough
	case "MEMBER_INDUSTRY":
		fallthrough
	case "MEMBER_SENIORITY":
		fallthrough
	case "MEMBER_JOB_TITLE ":
		fallthrough
	case "MEMBER_JOB_FUNCTION ":
		fallthrough
	case "MEMBER_COUNTRY_V2 ":
		fallthrough
	case "MEMBER_REGION_V2":
		fallthrough
	case "MEMBER_COMPANY":
		fallthrough
	case "PLACEMENT_NAME":
		fallthrough
	case "IMPRESSION_DEVICE_TYPE":
		*e = SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory: %v", v)
	}
}

// SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity - Choose how to group the data in your report by time. The options are:<br>- 'ALL': A single result summarizing the entire time range.<br>- 'DAILY': Group results by each day.<br>- 'MONTHLY': Group results by each month.<br>- 'YEARLY': Group results by each year.<br>Selecting a time grouping helps you analyze trends and patterns over different time periods.
type SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity string

const (
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityAll     SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "ALL"
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityDaily   SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "DAILY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityMonthly SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "MONTHLY"
	SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularityYearly  SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity = "YEARLY"
)

func (e SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity) ToPointer() *SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity {
	return &e
}

func (e *SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "ALL":
		fallthrough
	case "DAILY":
		fallthrough
	case "MONTHLY":
		fallthrough
	case "YEARLY":
		*e = SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity: %v", v)
	}
}

// SourceLinkedinAdsAdAnalyticsReportConfiguration - Config for custom ad Analytics Report
type SourceLinkedinAdsAdAnalyticsReportConfiguration struct {
	// The name for the custom report.
	Name string `json:"name"`
	// Choose a category to pivot your analytics report around. This selection will organize your data based on the chosen attribute, allowing you to analyze trends and performance from different perspectives.
	PivotBy SourceLinkedinAdsAdAnalyticsReportConfigurationPivotCategory `json:"pivot_by"`
	// Choose how to group the data in your report by time. The options are:<br>- 'ALL': A single result summarizing the entire time range.<br>- 'DAILY': Group results by each day.<br>- 'MONTHLY': Group results by each month.<br>- 'YEARLY': Group results by each year.<br>Selecting a time grouping helps you analyze trends and patterns over different time periods.
	TimeGranularity SourceLinkedinAdsAdAnalyticsReportConfigurationTimeGranularity `json:"time_granularity"`
}

type SourceLinkedinAdsAuthenticationAccessTokenAuthMethod string

const (
	SourceLinkedinAdsAuthenticationAccessTokenAuthMethodAccessToken SourceLinkedinAdsAuthenticationAccessTokenAuthMethod = "access_token"
)

func (e SourceLinkedinAdsAuthenticationAccessTokenAuthMethod) ToPointer() *SourceLinkedinAdsAuthenticationAccessTokenAuthMethod {
	return &e
}

func (e *SourceLinkedinAdsAuthenticationAccessTokenAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "access_token":
		*e = SourceLinkedinAdsAuthenticationAccessTokenAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAuthenticationAccessTokenAuthMethod: %v", v)
	}
}

type SourceLinkedinAdsAuthenticationAccessToken struct {
	// The access token generated for your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	AccessToken string                                                `json:"access_token"`
	AuthMethod  *SourceLinkedinAdsAuthenticationAccessTokenAuthMethod `json:"auth_method,omitempty"`
}

type SourceLinkedinAdsAuthenticationOAuth20AuthMethod string

const (
	SourceLinkedinAdsAuthenticationOAuth20AuthMethodOAuth20 SourceLinkedinAdsAuthenticationOAuth20AuthMethod = "oAuth2.0"
)

func (e SourceLinkedinAdsAuthenticationOAuth20AuthMethod) ToPointer() *SourceLinkedinAdsAuthenticationOAuth20AuthMethod {
	return &e
}

func (e *SourceLinkedinAdsAuthenticationOAuth20AuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "oAuth2.0":
		*e = SourceLinkedinAdsAuthenticationOAuth20AuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsAuthenticationOAuth20AuthMethod: %v", v)
	}
}

type SourceLinkedinAdsAuthenticationOAuth20 struct {
	AuthMethod *SourceLinkedinAdsAuthenticationOAuth20AuthMethod `json:"auth_method,omitempty"`
	// The client ID of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	ClientID string `json:"client_id"`
	// The client secret of your developer application. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	ClientSecret string `json:"client_secret"`
	// The key to refresh the expired access token. Refer to our <a href='https://docs.airbyte.com/integrations/sources/linkedin-ads#setup-guide'>documentation</a> for more information.
	RefreshToken string `json:"refresh_token"`
}

type SourceLinkedinAdsAuthenticationType string

const (
	SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationOAuth20     SourceLinkedinAdsAuthenticationType = "source-linkedin-ads_Authentication_OAuth2.0"
	SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationAccessToken SourceLinkedinAdsAuthenticationType = "source-linkedin-ads_Authentication_Access Token"
)

type SourceLinkedinAdsAuthentication struct {
	SourceLinkedinAdsAuthenticationOAuth20     *SourceLinkedinAdsAuthenticationOAuth20
	SourceLinkedinAdsAuthenticationAccessToken *SourceLinkedinAdsAuthenticationAccessToken

	Type SourceLinkedinAdsAuthenticationType
}

func CreateSourceLinkedinAdsAuthenticationSourceLinkedinAdsAuthenticationOAuth20(sourceLinkedinAdsAuthenticationOAuth20 SourceLinkedinAdsAuthenticationOAuth20) SourceLinkedinAdsAuthentication {
	typ := SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationOAuth20

	return SourceLinkedinAdsAuthentication{
		SourceLinkedinAdsAuthenticationOAuth20: &sourceLinkedinAdsAuthenticationOAuth20,
		Type:                                   typ,
	}
}

func CreateSourceLinkedinAdsAuthenticationSourceLinkedinAdsAuthenticationAccessToken(sourceLinkedinAdsAuthenticationAccessToken SourceLinkedinAdsAuthenticationAccessToken) SourceLinkedinAdsAuthentication {
	typ := SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationAccessToken

	return SourceLinkedinAdsAuthentication{
		SourceLinkedinAdsAuthenticationAccessToken: &sourceLinkedinAdsAuthenticationAccessToken,
		Type: typ,
	}
}

func (u *SourceLinkedinAdsAuthentication) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	sourceLinkedinAdsAuthenticationAccessToken := new(SourceLinkedinAdsAuthenticationAccessToken)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinAdsAuthenticationAccessToken); err == nil {
		u.SourceLinkedinAdsAuthenticationAccessToken = sourceLinkedinAdsAuthenticationAccessToken
		u.Type = SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationAccessToken
		return nil
	}

	sourceLinkedinAdsAuthenticationOAuth20 := new(SourceLinkedinAdsAuthenticationOAuth20)
	d = json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err := d.Decode(&sourceLinkedinAdsAuthenticationOAuth20); err == nil {
		u.SourceLinkedinAdsAuthenticationOAuth20 = sourceLinkedinAdsAuthenticationOAuth20
		u.Type = SourceLinkedinAdsAuthenticationTypeSourceLinkedinAdsAuthenticationOAuth20
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u SourceLinkedinAdsAuthentication) MarshalJSON() ([]byte, error) {
	if u.SourceLinkedinAdsAuthenticationAccessToken != nil {
		return json.Marshal(u.SourceLinkedinAdsAuthenticationAccessToken)
	}

	if u.SourceLinkedinAdsAuthenticationOAuth20 != nil {
		return json.Marshal(u.SourceLinkedinAdsAuthenticationOAuth20)
	}

	return nil, nil
}

type SourceLinkedinAdsLinkedinAds string

const (
	SourceLinkedinAdsLinkedinAdsLinkedinAds SourceLinkedinAdsLinkedinAds = "linkedin-ads"
)

func (e SourceLinkedinAdsLinkedinAds) ToPointer() *SourceLinkedinAdsLinkedinAds {
	return &e
}

func (e *SourceLinkedinAdsLinkedinAds) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "linkedin-ads":
		*e = SourceLinkedinAdsLinkedinAds(v)
		return nil
	default:
		return fmt.Errorf("invalid value for SourceLinkedinAdsLinkedinAds: %v", v)
	}
}

type SourceLinkedinAds struct {
	// Specify the account IDs to pull data from, separated by a space. Leave this field empty if you want to pull the data from all accounts accessible by the authenticated user. See the <a href="https://www.linkedin.com/help/linkedin/answer/a424270/find-linkedin-ads-account-details?lang=en">LinkedIn docs</a> to locate these IDs.
	AccountIds         []int64                                           `json:"account_ids,omitempty"`
	AdAnalyticsReports []SourceLinkedinAdsAdAnalyticsReportConfiguration `json:"ad_analytics_reports,omitempty"`
	Credentials        *SourceLinkedinAdsAuthentication                  `json:"credentials,omitempty"`
	SourceType         SourceLinkedinAdsLinkedinAds                      `json:"sourceType"`
	// UTC date in the format YYYY-MM-DD. Any data before this date will not be replicated.
	StartDate types.Date `json:"start_date"`
}
