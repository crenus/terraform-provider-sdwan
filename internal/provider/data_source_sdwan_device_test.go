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

func TestAccDataSourceSdwanDevice(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccDataSourceSdwanDeviceConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.sdwan_device.test", "device_id", "2081c2f4-3f9f-4fee-8078-dcc8904e368d"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "uuid", "2081c2f4-3f9f-4fee-8078-dcc8904e368d"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "site_id", "400"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "serial_number", "2AJC9DJ"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "hostname", "vEdge-123"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "reachability", "reachable"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "status", "normal"),
					resource.TestCheckResourceAttr("data.sdwan_device.test", "state", "green"),
				),
			},
		},
	})
}

const testAccDataSourceSdwanDeviceConfig = `


resource "sdwan_device" "test" {
  device_id = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
  uuid = "2081c2f4-3f9f-4fee-8078-dcc8904e368d"
  site_id = "400"
  serial_number = "2AJC9DJ"
  hostname = "vEdge-123"
  reachability = "reachable"
  status = "normal"
  state = "green"
}

data "sdwan_device" "test" {
  id = sdwan_device.test.id
}
`
