// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import "github.com/hashicorp/terraform-plugin-framework/types"

type DestinationAwsDatalakeAuthenticationModeIAMRole struct {
	CredentialsTitle types.String `tfsdk:"credentials_title"`
	RoleArn          types.String `tfsdk:"role_arn"`
}