---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_rule_set_policy_definition Resource - terraform-provider-sdwan"
subcategory: "Security Policies"
description: |-
  This resource can manage a Rule Set Policy Definition .
---

# sdwan_rule_set_policy_definition (Resource)

This resource can manage a Rule Set Policy Definition .

## Example Usage

```terraform
resource "sdwan_rule_set_policy_definition" "example" {
  name        = "Example"
  description = "My description"
  rules = [
    {
      name                     = "Rule1"
      order                    = 1
      source_ipv4_prefix       = "10.1.1.0/24"
      source_fqdn              = "cisco.com"
      source_port              = "80-90"
      source_geo_location      = "AF"
      destination_ipv4_prefix  = "10.1.1.0/24"
      destination_fqdn         = "cisco.com"
      destination_port         = "80-90"
      destination_geo_location = "AF"
      protocol                 = "cifs"
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `description` (String) The description of the policy definition
- `name` (String) The name of the policy definition
- `rules` (Attributes List) List of rules (see [below for nested schema](#nestedatt--rules))

### Read-Only

- `id` (String) The id of the object
- `version` (Number) The version of the object

<a id="nestedatt--rules"></a>
### Nested Schema for `rules`

Required:

- `name` (String) The name of the rule
- `order` (Number) The order of the rule

Optional:

- `destination_data_fqdn_prefix_list_id` (String) Destination data FQDN prefix list ID
- `destination_data_fqdn_prefix_list_version` (Number) Destination data FQDN prefix list version
- `destination_data_ipv4_prefix_list_id` (String) Destination data IPv4 prefix list ID
- `destination_data_ipv4_prefix_list_version` (Number) Destination data IPv4 prefix list version
- `destination_fqdn` (String) Destination fully qualified domain name
- `destination_geo_location` (String) Destination geo location
- `destination_geo_location_list_id` (String) Destination geo location list ID
- `destination_geo_location_list_version` (Number) Destination geo location list version
- `destination_ipv4_prefix` (String) Destination IPv4 prefix
- `destination_ipv4_prefix_variable` (String) Destination IPv4 prefix variable name
- `destination_object_group_id` (String) Destination object group ID
- `destination_object_group_version` (Number) Destination object group version
- `destination_port` (String) Destination port or range of ports
- `destination_port_list_id` (String) Destination port list ID
- `destination_port_list_version` (Number) Destination port list version
- `protocol` (String) Protocol name
- `protocol_list_id` (String) Protocol list ID
- `protocol_list_version` (Number) Protocol list version
- `protocol_number` (Number) Protocol number
  - Range: `0`-`255`
- `source_data_fqdn_prefix_list_id` (String) Source data FQDN prefix list ID
- `source_data_fqdn_prefix_list_version` (Number) Source data FQDN prefix list version
- `source_data_ipv4_prefix_list_id` (String) Source data IPv4 prefix list ID
- `source_data_ipv4_prefix_list_version` (Number) Source data IPv4 prefix list version
- `source_fqdn` (String) Source fully qualified domain name
- `source_geo_location` (String) Source geo location
- `source_geo_location_list_id` (String) Source geo location list ID
- `source_geo_location_list_version` (Number) Source geo location list version
- `source_ipv4_prefix` (String) Source IPv4 prefix
- `source_ipv4_prefix_variable` (String) Source IPv4 prefix variable name
- `source_object_group_id` (String) Source object group ID
- `source_object_group_version` (Number) Source object group version
- `source_port` (String) Source port or range of ports
- `source_port_list_id` (String) Source port list ID
- `source_port_list_version` (Number) Source port list version

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_rule_set_policy_definition.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```