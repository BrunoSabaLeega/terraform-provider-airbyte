// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"airbyte/internal/sdk/pkg/models/shared"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func (r *DestinationPubsubResourceModel) ToCreateSDKType() *shared.DestinationPubsubCreateRequest {
	batchingDelayThreshold := new(int64)
	if !r.Configuration.BatchingDelayThreshold.IsUnknown() && !r.Configuration.BatchingDelayThreshold.IsNull() {
		*batchingDelayThreshold = r.Configuration.BatchingDelayThreshold.ValueInt64()
	} else {
		batchingDelayThreshold = nil
	}
	batchingElementCountThreshold := new(int64)
	if !r.Configuration.BatchingElementCountThreshold.IsUnknown() && !r.Configuration.BatchingElementCountThreshold.IsNull() {
		*batchingElementCountThreshold = r.Configuration.BatchingElementCountThreshold.ValueInt64()
	} else {
		batchingElementCountThreshold = nil
	}
	batchingEnabled := r.Configuration.BatchingEnabled.ValueBool()
	batchingRequestBytesThreshold := new(int64)
	if !r.Configuration.BatchingRequestBytesThreshold.IsUnknown() && !r.Configuration.BatchingRequestBytesThreshold.IsNull() {
		*batchingRequestBytesThreshold = r.Configuration.BatchingRequestBytesThreshold.ValueInt64()
	} else {
		batchingRequestBytesThreshold = nil
	}
	credentialsJSON := r.Configuration.CredentialsJSON.ValueString()
	destinationType := shared.DestinationPubsubPubsub(r.Configuration.DestinationType.ValueString())
	orderingEnabled := r.Configuration.OrderingEnabled.ValueBool()
	projectID := r.Configuration.ProjectID.ValueString()
	topicID := r.Configuration.TopicID.ValueString()
	configuration := shared.DestinationPubsub{
		BatchingDelayThreshold:        batchingDelayThreshold,
		BatchingElementCountThreshold: batchingElementCountThreshold,
		BatchingEnabled:               batchingEnabled,
		BatchingRequestBytesThreshold: batchingRequestBytesThreshold,
		CredentialsJSON:               credentialsJSON,
		DestinationType:               destinationType,
		OrderingEnabled:               orderingEnabled,
		ProjectID:                     projectID,
		TopicID:                       topicID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPubsubCreateRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationPubsubResourceModel) ToGetSDKType() *shared.DestinationPubsubCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationPubsubResourceModel) ToUpdateSDKType() *shared.DestinationPubsubPutRequest {
	batchingDelayThreshold := new(int64)
	if !r.Configuration.BatchingDelayThreshold.IsUnknown() && !r.Configuration.BatchingDelayThreshold.IsNull() {
		*batchingDelayThreshold = r.Configuration.BatchingDelayThreshold.ValueInt64()
	} else {
		batchingDelayThreshold = nil
	}
	batchingElementCountThreshold := new(int64)
	if !r.Configuration.BatchingElementCountThreshold.IsUnknown() && !r.Configuration.BatchingElementCountThreshold.IsNull() {
		*batchingElementCountThreshold = r.Configuration.BatchingElementCountThreshold.ValueInt64()
	} else {
		batchingElementCountThreshold = nil
	}
	batchingEnabled := r.Configuration.BatchingEnabled.ValueBool()
	batchingRequestBytesThreshold := new(int64)
	if !r.Configuration.BatchingRequestBytesThreshold.IsUnknown() && !r.Configuration.BatchingRequestBytesThreshold.IsNull() {
		*batchingRequestBytesThreshold = r.Configuration.BatchingRequestBytesThreshold.ValueInt64()
	} else {
		batchingRequestBytesThreshold = nil
	}
	credentialsJSON := r.Configuration.CredentialsJSON.ValueString()
	orderingEnabled := r.Configuration.OrderingEnabled.ValueBool()
	projectID := r.Configuration.ProjectID.ValueString()
	topicID := r.Configuration.TopicID.ValueString()
	configuration := shared.DestinationPubsubUpdate{
		BatchingDelayThreshold:        batchingDelayThreshold,
		BatchingElementCountThreshold: batchingElementCountThreshold,
		BatchingEnabled:               batchingEnabled,
		BatchingRequestBytesThreshold: batchingRequestBytesThreshold,
		CredentialsJSON:               credentialsJSON,
		OrderingEnabled:               orderingEnabled,
		ProjectID:                     projectID,
		TopicID:                       topicID,
	}
	name := r.Name.ValueString()
	workspaceID := r.WorkspaceID.ValueString()
	out := shared.DestinationPubsubPutRequest{
		Configuration: configuration,
		Name:          name,
		WorkspaceID:   workspaceID,
	}
	return &out
}

func (r *DestinationPubsubResourceModel) ToDeleteSDKType() *shared.DestinationPubsubCreateRequest {
	out := r.ToCreateSDKType()
	return out
}

func (r *DestinationPubsubResourceModel) RefreshFromGetResponse(resp *shared.DestinationResponse) {
	r.DestinationID = types.StringValue(resp.DestinationID)
	r.DestinationType = types.StringValue(resp.DestinationType)
	r.Name = types.StringValue(resp.Name)
	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
}

func (r *DestinationPubsubResourceModel) RefreshFromCreateResponse(resp *shared.DestinationResponse) {
	r.RefreshFromGetResponse(resp)
}