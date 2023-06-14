// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"time"
)

func (r *SourceFacebookMarketingResourceModel) ToCreateSDKType() *shared.SourceFacebookMarketingCreateRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	accountID := r.Configuration.AccountID.ValueString()
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	customInsights := make([]shared.SourceFacebookMarketingInsightConfig, 0)
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		actionBreakdowns := make([]shared.SourceFacebookMarketingInsightConfigValidActionBreakdowns, 0)
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.SourceFacebookMarketingInsightConfigValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		breakdowns := make([]shared.SourceFacebookMarketingInsightConfigValidBreakdowns, 0)
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.SourceFacebookMarketingInsightConfigValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		fields := make([]shared.SourceFacebookMarketingInsightConfigValidEnums, 0)
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingInsightConfigValidEnums(fieldsItem.ValueString()))
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.SourceFacebookMarketingInsightConfigLevel)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.SourceFacebookMarketingInsightConfigLevel(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.SourceFacebookMarketingInsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	maxBatchSize := new(int64)
	if !r.Configuration.MaxBatchSize.IsUnknown() && !r.Configuration.MaxBatchSize.IsNull() {
		*maxBatchSize = r.Configuration.MaxBatchSize.ValueInt64()
	} else {
		maxBatchSize = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	sourceType := shared.SourceFacebookMarketingFacebookMarketing(r.Configuration.SourceType.ValueString())
	startDate1, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceFacebookMarketing{
		AccessToken:                accessToken,
		AccountID:                  accountID,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		IncludeDeleted:             includeDeleted,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		MaxBatchSize:               maxBatchSize,
		PageSize:                   pageSize,
		SourceType:                 sourceType,
		StartDate:                  startDate1,
	}
	name1 := r.Name.ValueString()
	secretID := new(string)
	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
		*secretID = r.SecretID.ValueString()
	} else {
		secretID = nil
	}
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingCreateRequest{
		Configuration: configuration,
		Name:          name1,
		SecretID:      secretID,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFacebookMarketingResourceModel) ToGetSDKType() *shared.SourceFacebookMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFacebookMarketingResourceModel) ToUpdateSDKType() *shared.SourceFacebookMarketingPutRequest {
	accessToken := r.Configuration.AccessToken.ValueString()
	accountID := r.Configuration.AccountID.ValueString()
	actionBreakdownsAllowEmpty := new(bool)
	if !r.Configuration.ActionBreakdownsAllowEmpty.IsUnknown() && !r.Configuration.ActionBreakdownsAllowEmpty.IsNull() {
		*actionBreakdownsAllowEmpty = r.Configuration.ActionBreakdownsAllowEmpty.ValueBool()
	} else {
		actionBreakdownsAllowEmpty = nil
	}
	customInsights := make([]shared.SourceFacebookMarketingUpdateInsightConfig, 0)
	for _, customInsightsItem := range r.Configuration.CustomInsights {
		actionBreakdowns := make([]shared.SourceFacebookMarketingUpdateInsightConfigValidActionBreakdowns, 0)
		for _, actionBreakdownsItem := range customInsightsItem.ActionBreakdowns {
			actionBreakdowns = append(actionBreakdowns, shared.SourceFacebookMarketingUpdateInsightConfigValidActionBreakdowns(actionBreakdownsItem.ValueString()))
		}
		breakdowns := make([]shared.SourceFacebookMarketingUpdateInsightConfigValidBreakdowns, 0)
		for _, breakdownsItem := range customInsightsItem.Breakdowns {
			breakdowns = append(breakdowns, shared.SourceFacebookMarketingUpdateInsightConfigValidBreakdowns(breakdownsItem.ValueString()))
		}
		endDate := new(time.Time)
		if !customInsightsItem.EndDate.IsUnknown() && !customInsightsItem.EndDate.IsNull() {
			*endDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.EndDate.ValueString())
		} else {
			endDate = nil
		}
		fields := make([]shared.SourceFacebookMarketingUpdateInsightConfigValidEnums, 0)
		for _, fieldsItem := range customInsightsItem.Fields {
			fields = append(fields, shared.SourceFacebookMarketingUpdateInsightConfigValidEnums(fieldsItem.ValueString()))
		}
		insightsLookbackWindow := new(int64)
		if !customInsightsItem.InsightsLookbackWindow.IsUnknown() && !customInsightsItem.InsightsLookbackWindow.IsNull() {
			*insightsLookbackWindow = customInsightsItem.InsightsLookbackWindow.ValueInt64()
		} else {
			insightsLookbackWindow = nil
		}
		level := new(shared.SourceFacebookMarketingUpdateInsightConfigLevel)
		if !customInsightsItem.Level.IsUnknown() && !customInsightsItem.Level.IsNull() {
			*level = shared.SourceFacebookMarketingUpdateInsightConfigLevel(customInsightsItem.Level.ValueString())
		} else {
			level = nil
		}
		name := customInsightsItem.Name.ValueString()
		startDate := new(time.Time)
		if !customInsightsItem.StartDate.IsUnknown() && !customInsightsItem.StartDate.IsNull() {
			*startDate, _ = time.Parse(time.RFC3339Nano, customInsightsItem.StartDate.ValueString())
		} else {
			startDate = nil
		}
		timeIncrement := new(int64)
		if !customInsightsItem.TimeIncrement.IsUnknown() && !customInsightsItem.TimeIncrement.IsNull() {
			*timeIncrement = customInsightsItem.TimeIncrement.ValueInt64()
		} else {
			timeIncrement = nil
		}
		customInsights = append(customInsights, shared.SourceFacebookMarketingUpdateInsightConfig{
			ActionBreakdowns:       actionBreakdowns,
			Breakdowns:             breakdowns,
			EndDate:                endDate,
			Fields:                 fields,
			InsightsLookbackWindow: insightsLookbackWindow,
			Level:                  level,
			Name:                   name,
			StartDate:              startDate,
			TimeIncrement:          timeIncrement,
		})
	}
	endDate1 := new(time.Time)
	if !r.Configuration.EndDate.IsUnknown() && !r.Configuration.EndDate.IsNull() {
		*endDate1, _ = time.Parse(time.RFC3339Nano, r.Configuration.EndDate.ValueString())
	} else {
		endDate1 = nil
	}
	fetchThumbnailImages := new(bool)
	if !r.Configuration.FetchThumbnailImages.IsUnknown() && !r.Configuration.FetchThumbnailImages.IsNull() {
		*fetchThumbnailImages = r.Configuration.FetchThumbnailImages.ValueBool()
	} else {
		fetchThumbnailImages = nil
	}
	includeDeleted := new(bool)
	if !r.Configuration.IncludeDeleted.IsUnknown() && !r.Configuration.IncludeDeleted.IsNull() {
		*includeDeleted = r.Configuration.IncludeDeleted.ValueBool()
	} else {
		includeDeleted = nil
	}
	insightsLookbackWindow1 := new(int64)
	if !r.Configuration.InsightsLookbackWindow.IsUnknown() && !r.Configuration.InsightsLookbackWindow.IsNull() {
		*insightsLookbackWindow1 = r.Configuration.InsightsLookbackWindow.ValueInt64()
	} else {
		insightsLookbackWindow1 = nil
	}
	maxBatchSize := new(int64)
	if !r.Configuration.MaxBatchSize.IsUnknown() && !r.Configuration.MaxBatchSize.IsNull() {
		*maxBatchSize = r.Configuration.MaxBatchSize.ValueInt64()
	} else {
		maxBatchSize = nil
	}
	pageSize := new(int64)
	if !r.Configuration.PageSize.IsUnknown() && !r.Configuration.PageSize.IsNull() {
		*pageSize = r.Configuration.PageSize.ValueInt64()
	} else {
		pageSize = nil
	}
	startDate1, _ := time.Parse(time.RFC3339Nano, r.Configuration.StartDate.ValueString())
	configuration := shared.SourceFacebookMarketingUpdate{
		AccessToken:                accessToken,
		AccountID:                  accountID,
		ActionBreakdownsAllowEmpty: actionBreakdownsAllowEmpty,
		CustomInsights:             customInsights,
		EndDate:                    endDate1,
		FetchThumbnailImages:       fetchThumbnailImages,
		IncludeDeleted:             includeDeleted,
		InsightsLookbackWindow:     insightsLookbackWindow1,
		MaxBatchSize:               maxBatchSize,
		PageSize:                   pageSize,
		StartDate:                  startDate1,
	}
	name1 := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.SourceFacebookMarketingPutRequest{
		Configuration: configuration,
		Name:          name1,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *SourceFacebookMarketingResourceModel) ToDeleteSDKType() *shared.SourceFacebookMarketingCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *SourceFacebookMarketingResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
	r.Name = types.StringValue(resp.Name)
	r.SourceID = types.StringValue(resp.SourceID)
	r.SourceType = types.StringValue(resp.SourceType)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *SourceFacebookMarketingResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
	r.RefreshFromGetResponse(resp)
}