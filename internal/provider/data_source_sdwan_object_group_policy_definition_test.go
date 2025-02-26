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
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccDataSourceSdwanObjectGroupPolicyDefinition(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanObjectGroupPolicyDefinitionConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_object_group_policy_definition.test", "name", "Example"),
					resource.TestCheckResourceAttr("data.sdwan_object_group_policy_definition.test", "description", "My description"),
					resource.TestCheckResourceAttr("data.sdwan_object_group_policy_definition.test", "ipv4_prefix", "10.1.1.0/24"),
					resource.TestCheckResourceAttr("data.sdwan_object_group_policy_definition.test", "fqdn", "cisco.com"),
					resource.TestCheckResourceAttr("data.sdwan_object_group_policy_definition.test", "port", "80-90"),
					resource.TestCheckResourceAttr("data.sdwan_object_group_policy_definition.test", "geo_location", "AF"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanObjectGroupPolicyDefinitionConfig = `


resource "sdwan_object_group_policy_definition" "test" {
  name = "Example"
  description = "My description"
  ipv4_prefix = "10.1.1.0/24"
  fqdn = "cisco.com"
  port = "80-90"
  geo_location = "AF"
}

data "sdwan_object_group_policy_definition" "test" {
  id = sdwan_object_group_policy_definition.test.id
}
`
