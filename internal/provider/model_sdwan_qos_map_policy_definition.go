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

	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

type QoSMapPolicyDefinition struct {
	Id            types.String                          `tfsdk:"id"`
	Version       types.Int64                           `tfsdk:"version"`
	Name          types.String                          `tfsdk:"name"`
	Description   types.String                          `tfsdk:"description"`
	QosSchedulers []QoSMapPolicyDefinitionQosSchedulers `tfsdk:"qos_schedulers"`
}

type QoSMapPolicyDefinitionQosSchedulers struct {
	Queue            types.Int64  `tfsdk:"queue"`
	ClassMapId       types.String `tfsdk:"class_map_id"`
	ClassMapVersion  types.Int64  `tfsdk:"class_map_version"`
	BandwidthPercent types.Int64  `tfsdk:"bandwidth_percent"`
	BufferPercent    types.Int64  `tfsdk:"buffer_percent"`
	Burst            types.Int64  `tfsdk:"burst"`
	DropType         types.String `tfsdk:"drop_type"`
	SchedulingType   types.String `tfsdk:"scheduling_type"`
}

func (data QoSMapPolicyDefinition) toBody(ctx context.Context) string {
	body := ""
	body, _ = sjson.Set(body, "type", "qosMap")
	if !data.Name.IsNull() {
		body, _ = sjson.Set(body, "name", data.Name.ValueString())
	}
	if !data.Description.IsNull() {
		body, _ = sjson.Set(body, "description", data.Description.ValueString())
	}
	if len(data.QosSchedulers) > 0 {
		body, _ = sjson.Set(body, "definition.qosSchedulers", []interface{}{})
		for _, item := range data.QosSchedulers {
			itemBody := ""
			if !item.Queue.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "queue", fmt.Sprint(item.Queue.ValueInt64()))
			}
			if !item.ClassMapId.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "classMapRef", item.ClassMapId.ValueString())
			}
			if !item.BandwidthPercent.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bandwidthPercent", fmt.Sprint(item.BandwidthPercent.ValueInt64()))
			}
			if !item.BufferPercent.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "bufferPercent", fmt.Sprint(item.BufferPercent.ValueInt64()))
			}
			if !item.Burst.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "burst", fmt.Sprint(item.Burst.ValueInt64()))
			}
			if !item.DropType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "drops", item.DropType.ValueString())
			}
			if !item.SchedulingType.IsNull() {
				itemBody, _ = sjson.Set(itemBody, "scheduling", item.SchedulingType.ValueString())
			}
			body, _ = sjson.SetRaw(body, "definition.qosSchedulers.-1", itemBody)
		}
	}
	return body
}

func (data *QoSMapPolicyDefinition) fromBody(ctx context.Context, res gjson.Result) {
	state := *data
	if value := res.Get("name"); value.Exists() {
		data.Name = types.StringValue(value.String())
	} else {
		data.Name = types.StringNull()
	}
	if value := res.Get("description"); value.Exists() {
		data.Description = types.StringValue(value.String())
	} else {
		data.Description = types.StringNull()
	}
	if value := res.Get("definition.qosSchedulers"); value.Exists() {
		data.QosSchedulers = make([]QoSMapPolicyDefinitionQosSchedulers, 0)
		value.ForEach(func(k, v gjson.Result) bool {
			item := QoSMapPolicyDefinitionQosSchedulers{}
			if cValue := v.Get("queue"); cValue.Exists() {
				item.Queue = types.Int64Value(cValue.Int())
			} else {
				item.Queue = types.Int64Null()
			}
			if cValue := v.Get("classMapRef"); cValue.Exists() {
				item.ClassMapId = types.StringValue(cValue.String())
			} else {
				item.ClassMapId = types.StringNull()
			}
			if cValue := v.Get("bandwidthPercent"); cValue.Exists() {
				item.BandwidthPercent = types.Int64Value(cValue.Int())
			} else {
				item.BandwidthPercent = types.Int64Null()
			}
			if cValue := v.Get("bufferPercent"); cValue.Exists() {
				item.BufferPercent = types.Int64Value(cValue.Int())
			} else {
				item.BufferPercent = types.Int64Null()
			}
			if cValue := v.Get("burst"); cValue.Exists() {
				item.Burst = types.Int64Value(cValue.Int())
			} else {
				item.Burst = types.Int64Null()
			}
			if cValue := v.Get("drops"); cValue.Exists() {
				item.DropType = types.StringValue(cValue.String())
			} else {
				item.DropType = types.StringNull()
			}
			if cValue := v.Get("scheduling"); cValue.Exists() {
				item.SchedulingType = types.StringValue(cValue.String())
			} else {
				item.SchedulingType = types.StringNull()
			}
			data.QosSchedulers = append(data.QosSchedulers, item)
			return true
		})
	}
	data.updateVersions(ctx, &state)
}

func (data *QoSMapPolicyDefinition) hasChanges(ctx context.Context, state *QoSMapPolicyDefinition) bool {
	hasChanges := false
	if !data.Name.Equal(state.Name) {
		hasChanges = true
	}
	if !data.Description.Equal(state.Description) {
		hasChanges = true
	}
	if len(data.QosSchedulers) != len(state.QosSchedulers) {
		hasChanges = true
	} else {
		for i := range data.QosSchedulers {
			if !data.QosSchedulers[i].Queue.Equal(state.QosSchedulers[i].Queue) {
				hasChanges = true
			}
			if !data.QosSchedulers[i].ClassMapId.Equal(state.QosSchedulers[i].ClassMapId) {
				hasChanges = true
			}
			if !data.QosSchedulers[i].BandwidthPercent.Equal(state.QosSchedulers[i].BandwidthPercent) {
				hasChanges = true
			}
			if !data.QosSchedulers[i].BufferPercent.Equal(state.QosSchedulers[i].BufferPercent) {
				hasChanges = true
			}
			if !data.QosSchedulers[i].Burst.Equal(state.QosSchedulers[i].Burst) {
				hasChanges = true
			}
			if !data.QosSchedulers[i].DropType.Equal(state.QosSchedulers[i].DropType) {
				hasChanges = true
			}
			if !data.QosSchedulers[i].SchedulingType.Equal(state.QosSchedulers[i].SchedulingType) {
				hasChanges = true
			}
		}
	}
	return hasChanges
}

func (data *QoSMapPolicyDefinition) updateVersions(ctx context.Context, state *QoSMapPolicyDefinition) {
	for i := range data.QosSchedulers {
		dataKeys := [...]string{fmt.Sprintf("%v", data.QosSchedulers[i].Queue.ValueInt64())}
		stateIndex := -1
		for j := range state.QosSchedulers {
			stateKeys := [...]string{fmt.Sprintf("%v", state.QosSchedulers[j].Queue.ValueInt64())}
			if dataKeys == stateKeys {
				stateIndex = j
				break
			}
		}
		if stateIndex > -1 {
			data.QosSchedulers[i].ClassMapVersion = state.QosSchedulers[stateIndex].ClassMapVersion
		} else {
			data.QosSchedulers[i].ClassMapVersion = types.Int64Null()
		}
	}
}
