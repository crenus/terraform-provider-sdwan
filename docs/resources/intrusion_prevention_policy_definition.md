---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_intrusion_prevention_policy_definition Resource - terraform-provider-sdwan"
subcategory: "Security Policies"
description: |-
  This resource can manage a Intrusion Prevention Policy Definition .
---

# sdwan_intrusion_prevention_policy_definition (Resource)

This resource can manage a Intrusion Prevention Policy Definition .

## Example Usage

```terraform
resource "sdwan_intrusion_prevention_policy_definition" "example" {
  name            = "Example"
  description     = "My description"
  mode            = "security"
  inspection_mode = "protection"
  log_level       = "alert"
  signature_set   = "connectivity"
  target_vpns     = ["1"]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the policy definition
- `name` (String) The name of the policy definition

### Optional

- `inspection_mode` (String) The inspection mode
  - Choices: `protection`, `detection`
- `ips_signature_list_id` (String) IPS signature list ID
- `ips_signature_list_version` (Number) IPS signature list version
- `log_level` (String) Log level
  - Choices: `emergency`, `alert`, `critical`, `error`, `warning`, `notice`, `info`, `debug`
- `mode` (String) The policy mode
  - Choices: `security`, `unified`
- `signature_set` (String) Signature set
  - Choices: `balanced`, `connectivity`, `security`
- `target_vpns` (List of String) List of VPN IDs

### Read-Only

- `id` (String) The id of the object
- `version` (Number) The version of the object

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_intrusion_prevention_policy_definition.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```
