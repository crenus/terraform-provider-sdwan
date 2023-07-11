---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_tloc_list_policy_object Resource - terraform-provider-sdwan"
subcategory: "Policy Objects"
description: |-
  This resource can manage a TLOC List policy object.
---

# sdwan_tloc_list_policy_object (Resource)

This resource can manage a TLOC List policy object.

## Example Usage

```terraform
resource "sdwan_tloc_list_policy_object" "example" {
  name = "Example"
  entries = [
    {
      tloc_ip       = "1.1.1.2"
      color         = "blue"
      encapsulation = "gre"
      preference    = 10
    }
  ]
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `entries` (Attributes List) List of entries (see [below for nested schema](#nestedatt--entries))
- `name` (String) The name of the policy object

### Read-Only

- `id` (String) The id of the policy object
- `version` (Number) The version of the feature template

<a id="nestedatt--entries"></a>
### Nested Schema for `entries`

Required:

- `color` (String) Color
  - Choices: `default`, `3g`, `biz-internet`, `blue`, `bronze`, `custom1`, `custom2`, `custom3`, `gold`, `green`, `lte`, `metro-ethernet`, `mpls`, `private1`, `private2`, `private3`, `private4`, `private5`, `private6`, `public-internet`, `red`, `silver`
- `encapsulation` (String) Encapsulation
  - Choices: `ipsec`, `gre`
- `tloc_ip` (String) TLOC IP

Optional:

- `preference` (Number) Preference
  - Range: `0`-`4294967295`

## Import

Import is supported using the following syntax:

```shell
terraform import sdwan_tloc_list_policy_object.example "f6b2c44c-693c-4763-b010-895aa3d236bd"
```