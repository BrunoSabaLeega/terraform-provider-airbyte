package provider

// import (
// 	"airbyte/internal/sdk/pkg/models/shared"
// 	"github.com/hashicorp/terraform-plugin-framework/types"
// )

// func (r *SourceKafkaResourceModel) ToCreateSDKType() *shared.SourceKafkaCreateRequest {
// 	password := r.Configuration.Password.ValueString()
// 	playground := r.Configuration.Playground.ValueBool()
// 	region := shared.SourceKafkaRegion(r.Configuration.Region.ValueString())
// 	sourceType := shared.SourceKafkaKafka(r.Configuration.SourceType.ValueString())
// 	username := r.Configuration.Username.ValueString()
// 	configuration := shared.SourceKafka{
// 		Password:   password,
// 		Playground: playground,
// 		Region:     region,
// 		SourceType: sourceType,
// 		Username:   username,
// 	}
// 	name := r.Name.ValueString()
// 	secretID := new(string)
// 	if !r.SecretID.IsUnknown() && !r.SecretID.IsNull() {
// 		*secretID = r.SecretID.ValueString()
// 	} else {
// 		secretID = nil
// 	}
// 	workspaceID := r.WorkspaceID.ValueString()
// 	out := shared.SourceKafkaCreateRequest{
// 		Configuration: configuration,
// 		Name:          name,
// 		SecretID:      secretID,
// 		WorkspaceID:   workspaceID,
// 	}
// 	return &out
// }

// func (r *SourceKafkaResourceModel) ToGetSDKType() *shared.SourceKafkaCreateRequest {
// 	out := r.ToCreateSDKType()
// 	return out
// }

// func (r *SourceKafkaResourceModel) ToUpdateSDKType() *shared.SourceKafkaPutRequest {
// 	password := r.Configuration.Password.ValueString()
// 	playground := r.Configuration.Playground.ValueBool()
// 	region := shared.SourceKafkaUpdateRegion(r.Configuration.Region.ValueString())
// 	username := r.Configuration.Username.ValueString()
// 	configuration := shared.SourceKafkaUpdate{
// 		Password:   password,
// 		Playground: playground,
// 		Region:     region,
// 		Username:   username,
// 	}
// 	name := r.Name.ValueString()
// 	workspaceID := r.WorkspaceID.ValueString()
// 	out := shared.SourceKafkaPutRequest{
// 		Configuration: configuration,
// 		Name:          name,
// 		WorkspaceID:   workspaceID,
// 	}
// 	return &out
// }

// func (r *SourceKafkaResourceModel) ToDeleteSDKType() *shared.SourceKafkaCreateRequest {
// 	out := r.ToCreateSDKType()
// 	return out
// }

// func (r *SourceKafkaResourceModel) RefreshFromGetResponse(resp *shared.SourceResponse) {
// 	r.Name = types.StringValue(resp.Name)
// 	r.SourceID = types.StringValue(resp.SourceID)
// 	r.SourceType = types.StringValue(resp.SourceType)
// 	r.WorkspaceID = types.StringValue(resp.WorkspaceID)
// }

// func (r *SourceKafkaResourceModel) RefreshFromCreateResponse(resp *shared.SourceResponse) {
// 	r.RefreshFromGetResponse(resp)
// }