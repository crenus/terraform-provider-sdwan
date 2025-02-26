// Copyright © 2023 Cisco Systems, Inc. and its affiliates.
// All rights reserved.
//
// Licensed under the Mozilla Public License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://mozilla.org/MPL/2.0/
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// SPDX-License-Identifier: MPL-2.0

// Code generated by "gen/generator.go"; DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-log/tflog"
	"github.com/netascode/go-sdwan"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &RuleSetPolicyDefinitionDataSource{}
	_ datasource.DataSourceWithConfigure = &RuleSetPolicyDefinitionDataSource{}
)

func NewRuleSetPolicyDefinitionDataSource() datasource.DataSource {
	return &RuleSetPolicyDefinitionDataSource{}
}

type RuleSetPolicyDefinitionDataSource struct {
	client *sdwan.Client
}

func (d *RuleSetPolicyDefinitionDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rule_set_policy_definition"
}

func (d *RuleSetPolicyDefinitionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Rule Set Policy Definition .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"name": schema.StringAttribute{
				MarkdownDescription: "The name of the policy definition",
				Computed:            true,
			},
			"description": schema.StringAttribute{
				MarkdownDescription: "The description of the policy definition",
				Computed:            true,
			},
			"rules": schema.ListNestedAttribute{
				MarkdownDescription: "List of rules",
				Computed:            true,
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							MarkdownDescription: "The name of the rule",
							Computed:            true,
						},
						"order": schema.Int64Attribute{
							MarkdownDescription: "The order of the rule",
							Computed:            true,
						},
						"source_object_group_id": schema.StringAttribute{
							MarkdownDescription: "Source object group ID",
							Computed:            true,
						},
						"source_object_group_version": schema.Int64Attribute{
							MarkdownDescription: "Source object group version",
							Computed:            true,
						},
						"source_data_ipv4_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Source data IPv4 prefix list ID",
							Computed:            true,
						},
						"source_data_ipv4_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: "Source data IPv4 prefix list version",
							Computed:            true,
						},
						"source_ipv4_prefix": schema.StringAttribute{
							MarkdownDescription: "Source IPv4 prefix",
							Computed:            true,
						},
						"source_ipv4_prefix_variable": schema.StringAttribute{
							MarkdownDescription: "Source IPv4 prefix variable name",
							Computed:            true,
						},
						"source_data_fqdn_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Source data FQDN prefix list ID",
							Computed:            true,
						},
						"source_data_fqdn_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: "Source data FQDN prefix list version",
							Computed:            true,
						},
						"source_fqdn": schema.StringAttribute{
							MarkdownDescription: "Source fully qualified domain name",
							Computed:            true,
						},
						"source_port_list_id": schema.StringAttribute{
							MarkdownDescription: "Source port list ID",
							Computed:            true,
						},
						"source_port_list_version": schema.Int64Attribute{
							MarkdownDescription: "Source port list version",
							Computed:            true,
						},
						"source_port": schema.StringAttribute{
							MarkdownDescription: "Source port or range of ports",
							Computed:            true,
						},
						"source_geo_location_list_id": schema.StringAttribute{
							MarkdownDescription: "Source geo location list ID",
							Computed:            true,
						},
						"source_geo_location_list_version": schema.Int64Attribute{
							MarkdownDescription: "Source geo location list version",
							Computed:            true,
						},
						"source_geo_location": schema.StringAttribute{
							MarkdownDescription: "Source geo location",
							Computed:            true,
						},
						"destination_object_group_id": schema.StringAttribute{
							MarkdownDescription: "Destination object group ID",
							Computed:            true,
						},
						"destination_object_group_version": schema.Int64Attribute{
							MarkdownDescription: "Destination object group version",
							Computed:            true,
						},
						"destination_data_ipv4_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Destination data IPv4 prefix list ID",
							Computed:            true,
						},
						"destination_data_ipv4_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: "Destination data IPv4 prefix list version",
							Computed:            true,
						},
						"destination_ipv4_prefix": schema.StringAttribute{
							MarkdownDescription: "Destination IPv4 prefix",
							Computed:            true,
						},
						"destination_ipv4_prefix_variable": schema.StringAttribute{
							MarkdownDescription: "Destination IPv4 prefix variable name",
							Computed:            true,
						},
						"destination_data_fqdn_prefix_list_id": schema.StringAttribute{
							MarkdownDescription: "Destination data FQDN prefix list ID",
							Computed:            true,
						},
						"destination_data_fqdn_prefix_list_version": schema.Int64Attribute{
							MarkdownDescription: "Destination data FQDN prefix list version",
							Computed:            true,
						},
						"destination_fqdn": schema.StringAttribute{
							MarkdownDescription: "Destination fully qualified domain name",
							Computed:            true,
						},
						"destination_port_list_id": schema.StringAttribute{
							MarkdownDescription: "Destination port list ID",
							Computed:            true,
						},
						"destination_port_list_version": schema.Int64Attribute{
							MarkdownDescription: "Destination port list version",
							Computed:            true,
						},
						"destination_port": schema.StringAttribute{
							MarkdownDescription: "Destination port or range of ports",
							Computed:            true,
						},
						"destination_geo_location_list_id": schema.StringAttribute{
							MarkdownDescription: "Destination geo location list ID",
							Computed:            true,
						},
						"destination_geo_location_list_version": schema.Int64Attribute{
							MarkdownDescription: "Destination geo location list version",
							Computed:            true,
						},
						"destination_geo_location": schema.StringAttribute{
							MarkdownDescription: "Destination geo location",
							Computed:            true,
						},
						"protocol_list_id": schema.StringAttribute{
							MarkdownDescription: "Protocol list ID",
							Computed:            true,
						},
						"protocol_list_version": schema.Int64Attribute{
							MarkdownDescription: "Protocol list version",
							Computed:            true,
						},
						"protocol": schema.StringAttribute{
							MarkdownDescription: "Protocol name",
							Computed:            true,
						},
						"protocol_number": schema.Int64Attribute{
							MarkdownDescription: "Protocol number",
							Computed:            true,
						},
					},
				},
			},
		},
	}
}

func (d *RuleSetPolicyDefinitionDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *RuleSetPolicyDefinitionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config RuleSetPolicyDefinition

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/template/policy/definition/ruleset/" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
