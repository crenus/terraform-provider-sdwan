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

func TestAccSdwanCiscoSIGCredentialsFeatureTemplate(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_minimum(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_organization", "org1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_base_uri", "abc"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_username", "user1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_password", "password123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_cloud_name", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_username", "partner1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_password", "password123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_api_key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_secret", "secret123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_organization_id", "org1"),
				),
			},
			{
				Config: testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_all(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_organization", "org1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_base_uri", "abc"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_username", "user1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_password", "password123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_cloud_name", "1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_username", "partner1"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_password", "password123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "zscaler_partner_api_key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_key", "key123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_api_secret", "secret123"),
					resource.TestCheckResourceAttr("sdwan_cisco_sig_credentials_feature_template.test", "umbrella_organization_id", "org1"),
				),
			},
		},
	})
}

func testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_minimum() string {
	return `
	resource "sdwan_cisco_sig_credentials_feature_template" "test" {
		name = "TF_TEST_MIN"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		zscaler_organization = "org1"
		zscaler_partner_base_uri = "abc"
		zscaler_username = "user1"
		zscaler_password = "password123"
		zscaler_cloud_name = 1
		zscaler_partner_username = "partner1"
		zscaler_partner_password = "password123"
		zscaler_partner_api_key = "key123"
		umbrella_api_key = "key123"
		umbrella_api_secret = "secret123"
		umbrella_organization_id = "org1"
	}
	`
}

func testAccSdwanCiscoSIGCredentialsFeatureTemplateConfig_all() string {
	return `
	resource "sdwan_cisco_sig_credentials_feature_template" "test" {
		name = "TF_TEST_ALL"
		description = "Terraform integration test"
		device_types = ["vedge-C8000V"]
		zscaler_organization = "org1"
		zscaler_partner_base_uri = "abc"
		zscaler_username = "user1"
		zscaler_password = "password123"
		zscaler_cloud_name = 1
		zscaler_partner_username = "partner1"
		zscaler_partner_password = "password123"
		zscaler_partner_api_key = "key123"
		umbrella_api_key = "key123"
		umbrella_api_secret = "secret123"
		umbrella_organization_id = "org1"
	}
	`
}
