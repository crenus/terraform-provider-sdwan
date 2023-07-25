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

func TestAccDataSourceSdwanCLIDeviceTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanCLIDeviceTemplateConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "name", "Example"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "description", "My description"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "device_type", "vedge-ISR-4331"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "cli_type", "device"),
					resource.TestCheckResourceAttr("data.sdwan_cli_device_template.test", "cli_configuration", " system\n host-name             R1-ISR4331-1200-1"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanCLIDeviceTemplateConfig = `


resource "sdwan_cli_device_template" "test" {
  name = "Example"
  description = "My description"
  device_type = "vedge-ISR-4331"
  cli_type = "device"
  cli_configuration = " system\n host-name             R1-ISR4331-1200-1"
}

data "sdwan_cli_device_template" "test" {
  id = sdwan_cli_device_template.test.id
}
`
