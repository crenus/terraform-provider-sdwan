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

package provider

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccSdwanActivateCentralizedPolicy(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanActivateCentralizedPolicyConfig,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_activate_centralized_policy.test", "version", "0"),
				),
			},
		},
	})
}

const testAccSdwanActivateCentralizedPolicyConfig = `
resource "sdwan_site_list_policy_object" "sites1" {
  name = "TF_TEST"
  entries = [
    {
      site_id = "100-200"
    }
  ]
}

resource "sdwan_vpn_list_policy_object" "vpns1" {
  name = "TF_TEST"
  entries = [
    {
      vpn_id = "100-200"
    }
  ]
}

resource "sdwan_traffic_data_policy_definition" "data1" {
  name           = "TF_TEST"
  description    = "My description"
  default_action = "drop"
  sequences = [
    {
      id      = 1
      name    = "Seq1"
      type    = "applicationFirewall"
      ip_type = "ipv4"
      action_entries = [
        {
          type = "log"
          log  = true
        }
      ]
    }
  ]
}

resource "sdwan_centralized_policy" "test" {
	name = "Example"
	description = "My description"
	definitions = [{
		id = sdwan_traffic_data_policy_definition.data1.id
		type = "data"
		entries = [{
			site_list_ids = [sdwan_site_list_policy_object.sites1.id]
			vpn_list_ids = [sdwan_vpn_list_policy_object.vpns1.id]
			direction = "service"
		}]
	}]
}

resource "sdwan_activate_centralized_policy" "test" {
	id      = sdwan_centralized_policy.test.id
	version = sdwan_centralized_policy.test.version
}
`
