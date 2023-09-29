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
	_ datasource.DataSource              = &DevicesDataSource{}
	_ datasource.DataSourceWithConfigure = &DevicesDataSource{}
)

func NewDevicesDataSource() datasource.DataSource {
	return &DevicesDataSource{}
}

type DevicesDataSource struct {
	client *sdwan.Client
}

func (d *DevicesDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *DevicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		// This description is used by the documentation generator and the language server.
		MarkdownDescription: "This data source can read the Devices .",

		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				MarkdownDescription: "The id of the object",
				Required:            true,
			},
			"version": schema.Int64Attribute{
				MarkdownDescription: "The version of the object",
				Computed:            true,
			},
			"device_id": schema.StringAttribute{
				MarkdownDescription: "",
				Computed:            true,
			},
			"uuid": schema.StringAttribute{
				MarkdownDescription: "Unique identifier for device",
				Computed:            true,
			},
			"site_id": schema.StringAttribute{
				MarkdownDescription: "Site id for respective device",
				Computed:            true,
			},
			"serial_number": schema.StringAttribute{
				MarkdownDescription: "Serial number for device. Could be board or virtual identifier",
				Computed:            true,
			},
			"hostname": schema.StringAttribute{
				MarkdownDescription: "Hostname for respective device",
				Computed:            true,
			},
			"reachability": schema.StringAttribute{
				MarkdownDescription: "Reachability of device",
				Computed:            true,
			},
			"status": schema.StringAttribute{
				MarkdownDescription: "Status for respective device",
				Computed:            true,
			},
			"state": schema.StringAttribute{
				MarkdownDescription: "state for respective device",
				Computed:            true,
			},
		},
	}
}

func (d *DevicesDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, _ *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	d.client = req.ProviderData.(*SdwanProviderData).Client
}

func (d *DevicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var config Devices

	// Read config
	diags := req.Config.Get(ctx, &config)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	tflog.Debug(ctx, fmt.Sprintf("%s: Beginning Read", config.Id.String()))

	res, err := d.client.Get("/device" + config.Id.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Failed to retrieve object, got error: %s", err))
		return
	}

	config.fromBody(ctx, res)

	tflog.Debug(ctx, fmt.Sprintf("%s: Read finished successfully", config.Id.ValueString()))

	diags = resp.State.Set(ctx, &config)
	resp.Diagnostics.Append(diags...)
}
