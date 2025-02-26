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

	"github.com/CiscoDevNet/terraform-provider-sdwan/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type CiscoThousandEyes struct {
	Id                  types.String                           `tfsdk:"id"`
	Version             types.Int64                            `tfsdk:"version"`
	TemplateType        types.String                           `tfsdk:"template_type"`
	Name                types.String                           `tfsdk:"name"`
	Description         types.String                           `tfsdk:"description"`
	DeviceTypes         types.List                             `tfsdk:"device_types"`
	VirtualApplications []CiscoThousandEyesVirtualApplications `tfsdk:"virtual_applications"`
}

type CiscoThousandEyesVirtualApplications struct {
	Optional                    types.Bool   `tfsdk:"optional"`
	InstanceId                  types.String `tfsdk:"instance_id"`
	ApplicationType             types.String `tfsdk:"application_type"`
	TeAccountGroupToken         types.String `tfsdk:"te_account_group_token"`
	TeAccountGroupTokenVariable types.String `tfsdk:"te_account_group_token_variable"`
	TeVpn                       types.Int64  `tfsdk:"te_vpn"`
	TeVpnVariable               types.String `tfsdk:"te_vpn_variable"`
	TeAgentIp                   types.String `tfsdk:"te_agent_ip"`
	TeAgentIpVariable           types.String `tfsdk:"te_agent_ip_variable"`
	TeDefaultGateway            types.String `tfsdk:"te_default_gateway"`
	TeDefaultGatewayVariable    types.String `tfsdk:"te_default_gateway_variable"`
	TeNameServer                types.String `tfsdk:"te_name_server"`
	TeNameServerVariable        types.String `tfsdk:"te_name_server_variable"`
	TeHostname                  types.String `tfsdk:"te_hostname"`
	TeHostnameVariable          types.String `tfsdk:"te_hostname_variable"`
	TeWebProxyType              types.String `tfsdk:"te_web_proxy_type"`
	TeProxyHost                 types.String `tfsdk:"te_proxy_host"`
	TeProxyHostVariable         types.String `tfsdk:"te_proxy_host_variable"`
	TeProxyPort                 types.Int64  `tfsdk:"te_proxy_port"`
	TeProxyPortVariable         types.String `tfsdk:"te_proxy_port_variable"`
	TePacUrl                    types.String `tfsdk:"te_pac_url"`
	TePacUrlVariable            types.String `tfsdk:"te_pac_url_variable"`
}

func (data CiscoThousandEyes) getModel() string {
	return "cisco_thousandeyes"
}

func (data CiscoThousandEyes) toBody(ctx context.Context) string {
	body := ""

	var device_types []string
	data.DeviceTypes.ElementsAs(ctx, &device_types, false)
	body, _ = sjson.Set(body, "deviceType", device_types)
	body, _ = sjson.Set(body, "factoryDefault", false)
	body, _ = sjson.Set(body, "templateDescription", data.Description.ValueString())
	body, _ = sjson.Set(body, "templateMinVersion", "15.0.0")
	body, _ = sjson.Set(body, "templateName", data.Name.ValueString())
	body, _ = sjson.Set(body, "templateType", "cisco_thousandeyes")
	body, _ = sjson.Set(body, "templateDefinition", map[string]interface{}{})

	path := "templateDefinition."
	if len(data.VirtualApplications) > 0 {
		body, _ = sjson.Set(body, path+"virtual-application."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipType", "constant")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipPrimaryKey", []string{"application-type"})
		body, _ = sjson.Set(body, path+"virtual-application."+"vipValue", []interface{}{})
	} else {
		body, _ = sjson.Set(body, path+"virtual-application."+"vipObjectType", "tree")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipType", "ignore")
		body, _ = sjson.Set(body, path+"virtual-application."+"vipPrimaryKey", []string{"application-type"})
		body, _ = sjson.Set(body, path+"virtual-application."+"vipValue", []interface{}{})
	}
	for _, item := range data.VirtualApplications {
		itemBody := ""
		itemAttributes := make([]string, 0)
		itemAttributes = append(itemAttributes, "instance-id")
		if item.InstanceId.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "instance-id."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "instance-id."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "instance-id."+"vipValue", item.InstanceId.ValueString())
		}
		itemAttributes = append(itemAttributes, "application-type")
		if item.ApplicationType.IsNull() {
		} else {
			itemBody, _ = sjson.Set(itemBody, "application-type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "application-type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "application-type."+"vipValue", item.ApplicationType.ValueString())
		}
		itemAttributes = append(itemAttributes, "token")

		if !item.TeAccountGroupTokenVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.token."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.token."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.token."+"vipVariableName", item.TeAccountGroupTokenVariable.ValueString())
		} else if item.TeAccountGroupToken.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te", map[string]interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.token."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.token."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.token."+"vipValue", item.TeAccountGroupToken.ValueString())
		}
		itemAttributes = append(itemAttributes, "vpn")

		if !item.TeVpnVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipVariableName", item.TeVpnVariable.ValueString())
		} else if item.TeVpn.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.vpn."+"vipValue", item.TeVpn.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "te-mgmt-ip")

		if !item.TeAgentIpVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.te-mgmt-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.te-mgmt-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.te-mgmt-ip."+"vipVariableName", item.TeAgentIpVariable.ValueString())
		} else if item.TeAgentIp.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te", map[string]interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.te-mgmt-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.te-mgmt-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.te-mgmt-ip."+"vipValue", item.TeAgentIp.ValueString())
		}
		itemAttributes = append(itemAttributes, "te-vpg-ip")

		if !item.TeDefaultGatewayVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.te-vpg-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.te-vpg-ip."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.te-vpg-ip."+"vipVariableName", item.TeDefaultGatewayVariable.ValueString())
		} else if item.TeDefaultGateway.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te", map[string]interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.te-vpg-ip."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.te-vpg-ip."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.te-vpg-ip."+"vipValue", item.TeDefaultGateway.ValueString())
		}
		itemAttributes = append(itemAttributes, "name-server")

		if !item.TeNameServerVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipVariableName", item.TeNameServerVariable.ValueString())
		} else if item.TeNameServer.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.name-server."+"vipValue", item.TeNameServer.ValueString())
		}
		itemAttributes = append(itemAttributes, "hostname")

		if !item.TeHostnameVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipVariableName", item.TeHostnameVariable.ValueString())
		} else if item.TeHostname.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipType", "ignore")
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.hostname."+"vipValue", item.TeHostname.ValueString())
		}
		itemAttributes = append(itemAttributes, "proxy_type")
		if item.TeWebProxyType.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te", map[string]interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_type."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_type."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_type."+"vipValue", item.TeWebProxyType.ValueString())
		}
		itemAttributes = append(itemAttributes, "proxy_host")

		if !item.TeProxyHostVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_host."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_host."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_host."+"vipVariableName", item.TeProxyHostVariable.ValueString())
		} else if item.TeProxyHost.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static", map[string]interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_host."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_host."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_host."+"vipValue", item.TeProxyHost.ValueString())
		}
		itemAttributes = append(itemAttributes, "proxy_port")

		if !item.TeProxyPortVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_port."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_port."+"vipVariableName", item.TeProxyPortVariable.ValueString())
		} else if item.TeProxyPort.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static", map[string]interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_port."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_port."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_static.proxy_port."+"vipValue", item.TeProxyPort.ValueInt64())
		}
		itemAttributes = append(itemAttributes, "pac_url")

		if !item.TePacUrlVariable.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_pac.pac_url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_pac.pac_url."+"vipType", "variableName")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_pac.pac_url."+"vipVariableName", item.TePacUrlVariable.ValueString())
		} else if item.TePacUrl.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_pac", map[string]interface{}{})
		} else {
			itemBody, _ = sjson.Set(itemBody, "te.proxy_pac.pac_url."+"vipObjectType", "object")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_pac.pac_url."+"vipType", "constant")
			itemBody, _ = sjson.Set(itemBody, "te.proxy_pac.pac_url."+"vipValue", item.TePacUrl.ValueString())
		}
		if !item.Optional.IsNull() {
			itemBody, _ = sjson.Set(itemBody, "vipOptional", item.Optional.ValueBool())
			itemBody, _ = sjson.Set(itemBody, "priority-order", itemAttributes)
		}
		body, _ = sjson.SetRaw(body, path+"virtual-application."+"vipValue.-1", itemBody)
	}
	return body
}

func (data *CiscoThousandEyes) fromBody(ctx context.Context, res gjson.Result) {
	if value := res.Get("deviceType"); value.Exists() {
		data.DeviceTypes = helpers.GetStringList(value.Array())
	} else {
		data.DeviceTypes = types.ListNull(types.StringType)
	}
	if value := res.Get("templateDescription"); value.Exists() && value.String() != "" {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("templateName"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("templateType"); value.Exists() {
		data.TemplateType = types.StringValue(value.String())
	} else {
		data.TemplateType = types.StringNull()
	}

	path := "templateDefinition."
	if value := res.Get(path + "virtual-application.vipValue"); len(value.Array()) > 0 {
		data.VirtualApplications = make([]CiscoThousandEyesVirtualApplications, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := CiscoThousandEyesVirtualApplications{}
			if cValue := v.Get("vipOptional"); cValue.Exists() {
				item.Optional = types.BoolValue(cValue.Bool())
			} else {
				item.Optional = types.BoolNull()
			}
			if cValue := v.Get("instance-id.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.InstanceId = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.InstanceId = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("instance-id.vipValue")
					item.InstanceId = types.StringValue(cv.String())

				}
			} else {
				item.InstanceId = types.StringNull()

			}
			if cValue := v.Get("application-type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.ApplicationType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.ApplicationType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("application-type.vipValue")
					item.ApplicationType = types.StringValue(cv.String())

				}
			} else {
				item.ApplicationType = types.StringNull()

			}
			if cValue := v.Get("te.token.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeAccountGroupToken = types.StringNull()

					cv := v.Get("te.token.vipVariableName")
					item.TeAccountGroupTokenVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeAccountGroupToken = types.StringNull()
					item.TeAccountGroupTokenVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.token.vipValue")
					item.TeAccountGroupToken = types.StringValue(cv.String())
					item.TeAccountGroupTokenVariable = types.StringNull()
				}
			} else {
				item.TeAccountGroupToken = types.StringNull()
				item.TeAccountGroupTokenVariable = types.StringNull()
			}
			if cValue := v.Get("te.vpn.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeVpn = types.Int64Null()

					cv := v.Get("te.vpn.vipVariableName")
					item.TeVpnVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeVpn = types.Int64Null()
					item.TeVpnVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.vpn.vipValue")
					item.TeVpn = types.Int64Value(cv.Int())
					item.TeVpnVariable = types.StringNull()
				}
			} else {
				item.TeVpn = types.Int64Null()
				item.TeVpnVariable = types.StringNull()
			}
			if cValue := v.Get("te.te-mgmt-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeAgentIp = types.StringNull()

					cv := v.Get("te.te-mgmt-ip.vipVariableName")
					item.TeAgentIpVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeAgentIp = types.StringNull()
					item.TeAgentIpVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.te-mgmt-ip.vipValue")
					item.TeAgentIp = types.StringValue(cv.String())
					item.TeAgentIpVariable = types.StringNull()
				}
			} else {
				item.TeAgentIp = types.StringNull()
				item.TeAgentIpVariable = types.StringNull()
			}
			if cValue := v.Get("te.te-vpg-ip.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeDefaultGateway = types.StringNull()

					cv := v.Get("te.te-vpg-ip.vipVariableName")
					item.TeDefaultGatewayVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeDefaultGateway = types.StringNull()
					item.TeDefaultGatewayVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.te-vpg-ip.vipValue")
					item.TeDefaultGateway = types.StringValue(cv.String())
					item.TeDefaultGatewayVariable = types.StringNull()
				}
			} else {
				item.TeDefaultGateway = types.StringNull()
				item.TeDefaultGatewayVariable = types.StringNull()
			}
			if cValue := v.Get("te.name-server.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeNameServer = types.StringNull()

					cv := v.Get("te.name-server.vipVariableName")
					item.TeNameServerVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeNameServer = types.StringNull()
					item.TeNameServerVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.name-server.vipValue")
					item.TeNameServer = types.StringValue(cv.String())
					item.TeNameServerVariable = types.StringNull()
				}
			} else {
				item.TeNameServer = types.StringNull()
				item.TeNameServerVariable = types.StringNull()
			}
			if cValue := v.Get("te.hostname.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeHostname = types.StringNull()

					cv := v.Get("te.hostname.vipVariableName")
					item.TeHostnameVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeHostname = types.StringNull()
					item.TeHostnameVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.hostname.vipValue")
					item.TeHostname = types.StringValue(cv.String())
					item.TeHostnameVariable = types.StringNull()
				}
			} else {
				item.TeHostname = types.StringNull()
				item.TeHostnameVariable = types.StringNull()
			}
			if cValue := v.Get("te.proxy_type.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeWebProxyType = types.StringNull()

				} else if cValue.String() == "ignore" {
					item.TeWebProxyType = types.StringNull()

				} else if cValue.String() == "constant" {
					cv := v.Get("te.proxy_type.vipValue")
					item.TeWebProxyType = types.StringValue(cv.String())

				}
			} else {
				item.TeWebProxyType = types.StringNull()

			}
			if cValue := v.Get("te.proxy_static.proxy_host.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeProxyHost = types.StringNull()

					cv := v.Get("te.proxy_static.proxy_host.vipVariableName")
					item.TeProxyHostVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeProxyHost = types.StringNull()
					item.TeProxyHostVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.proxy_static.proxy_host.vipValue")
					item.TeProxyHost = types.StringValue(cv.String())
					item.TeProxyHostVariable = types.StringNull()
				}
			} else {
				item.TeProxyHost = types.StringNull()
				item.TeProxyHostVariable = types.StringNull()
			}
			if cValue := v.Get("te.proxy_static.proxy_port.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TeProxyPort = types.Int64Null()

					cv := v.Get("te.proxy_static.proxy_port.vipVariableName")
					item.TeProxyPortVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TeProxyPort = types.Int64Null()
					item.TeProxyPortVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.proxy_static.proxy_port.vipValue")
					item.TeProxyPort = types.Int64Value(cv.Int())
					item.TeProxyPortVariable = types.StringNull()
				}
			} else {
				item.TeProxyPort = types.Int64Null()
				item.TeProxyPortVariable = types.StringNull()
			}
			if cValue := v.Get("te.proxy_pac.pac_url.vipType"); cValue.Exists() {
				if cValue.String() == "variableName" {
					item.TePacUrl = types.StringNull()

					cv := v.Get("te.proxy_pac.pac_url.vipVariableName")
					item.TePacUrlVariable = types.StringValue(cv.String())

				} else if cValue.String() == "ignore" {
					item.TePacUrl = types.StringNull()
					item.TePacUrlVariable = types.StringNull()
				} else if cValue.String() == "constant" {
					cv := v.Get("te.proxy_pac.pac_url.vipValue")
					item.TePacUrl = types.StringValue(cv.String())
					item.TePacUrlVariable = types.StringNull()
				}
			} else {
				item.TePacUrl = types.StringNull()
				item.TePacUrlVariable = types.StringNull()
			}
			data.VirtualApplications = append(data.VirtualApplications, item)
			return true
		})
	}
}

func (data *CiscoThousandEyes) hasChanges(ctx context.Context, state *CiscoThousandEyes) bool {
	hasChanges := false
	if len(data.VirtualApplications) != len(state.VirtualApplications) {
		hasChanges = true
	} else {
		for i := range data.VirtualApplications {
			if !data.VirtualApplications[i].InstanceId.Equal(state.VirtualApplications[i].InstanceId) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].ApplicationType.Equal(state.VirtualApplications[i].ApplicationType) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeAccountGroupToken.Equal(state.VirtualApplications[i].TeAccountGroupToken) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeVpn.Equal(state.VirtualApplications[i].TeVpn) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeAgentIp.Equal(state.VirtualApplications[i].TeAgentIp) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeDefaultGateway.Equal(state.VirtualApplications[i].TeDefaultGateway) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeNameServer.Equal(state.VirtualApplications[i].TeNameServer) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeHostname.Equal(state.VirtualApplications[i].TeHostname) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeWebProxyType.Equal(state.VirtualApplications[i].TeWebProxyType) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeProxyHost.Equal(state.VirtualApplications[i].TeProxyHost) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TeProxyPort.Equal(state.VirtualApplications[i].TeProxyPort) {
				hasChanges = true
			}
			if !data.VirtualApplications[i].TePacUrl.Equal(state.VirtualApplications[i].TePacUrl) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}
