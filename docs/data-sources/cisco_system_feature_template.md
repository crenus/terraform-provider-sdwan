---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "sdwan_cisco_system_feature_template Data Source - terraform-provider-sdwan"
subcategory: "Feature Templates"
description: |-
  This data source can read the Cisco System feature template.
---

# sdwan_cisco_system_feature_template (Data Source)

This data source can read the Cisco System feature template.

## Example Usage

```terraform
data "sdwan_cisco_system_feature_template" "example" {
  id = "f6b2c44c-693c-4763-b010-895aa3d236bd"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `id` (String) The id of the feature template
- `name` (String) The name of the feature template

### Read-Only

- `admin_tech_on_failure` (Boolean) Collect admin-tech before reboot due to daemon failure
- `admin_tech_on_failure_variable` (String) Variable name
- `affinity_group_number` (Number) Set the affinity group number for router
- `affinity_group_number_variable` (String) Variable name
- `affinity_group_preference` (List of String) Set the affinity group preference
- `affinity_group_preference_variable` (String) Variable name
- `console_baud_rate` (String) Set the console baud rate
- `console_baud_rate_variable` (String) Variable name
- `control_session_pps` (Number) Set the policer rate for control sessions
- `control_session_pps_variable` (String) Variable name
- `controller_group_list` (List of String) Configure a list of comma-separated device groups
- `controller_group_list_variable` (String) Variable name
- `description` (String) The description of the feature template
- `device_groups` (List of String) Device groups (Use comma(,) for multiple groups)
- `device_groups_variable` (String) Variable name
- `device_types` (List of String) List of supported device types
- `enable_mrf_migration` (String) Enable migration mode to Multi-Region Fabric
- `geo_fencing` (Boolean) Enable Geo fencing
- `geo_fencing_range` (Number) Set the device’s geo fencing range
- `geo_fencing_range_variable` (String) Variable name
- `geo_fencing_sms` (Boolean) Enable Geo fencing
- `geo_fencing_sms_phone_numbers` (Attributes List) Set device’s geo fencing SMS phone number (see [below for nested schema](#nestedatt--geo_fencing_sms_phone_numbers))
- `hostname` (String) Set the hostname
- `hostname_variable` (String) Variable name
- `idle_timeout` (Number) Idle CLI timeout in minutes
- `idle_timeout_variable` (String) Variable name
- `latitude` (Number) Set the device’s physical latitude
- `latitude_variable` (String) Variable name
- `location` (String) Set the location of the device
- `location_variable` (String) Variable name
- `longitude` (Number) Set the device’s physical longitude
- `longitude_variable` (String) Variable name
- `max_omp_sessions` (Number) Set the maximum number of OMP sessions <1..100> the device can have
- `max_omp_sessions_variable` (String) Variable name
- `migration_bgp_community` (Number) Set BGP community during migration from BGP-core based network
- `multi_tenant` (Boolean) Device is multi-tenant
- `multi_tenant_variable` (String) Variable name
- `object_trackers` (Attributes List) Object Track configuration (see [below for nested schema](#nestedatt--object_trackers))
- `on_demand_tunnel` (Boolean) Enable or disable On-demand Tunnel
- `on_demand_tunnel_idle_timeout` (Number) Idle CLI timeout in minutes
- `on_demand_tunnel_idle_timeout_variable` (String) Variable name
- `on_demand_tunnel_variable` (String) Variable name
- `overlay_id` (Number) Set the Overlay ID
- `overlay_id_variable` (String) Variable name
- `port_hopping` (Boolean) Enable port hopping
- `port_hopping_variable` (String) Variable name
- `port_offset` (Number) Set the TLOC port offset when multiple devices are behind a NAT
- `port_offset_variable` (String) Variable name
- `region_id` (Number) Set region ID
- `region_id_variable` (String) Variable name
- `role` (String) Set the role for router
- `role_variable` (String) Variable name
- `secondary_region_id` (Number) Set secondary region ID
- `secondary_region_id_variable` (String) Variable name
- `site_id` (Number) Set the site identifier
- `site_id_variable` (String) Variable name
- `system_description` (String) Set a text description of the device
- `system_description_variable` (String) Variable name
- `system_ip` (String) Set the system IP address
- `system_ip_variable` (String) Variable name
- `template_type` (String) The template type
- `timezone` (String) Set the timezone
- `timezone_variable` (String) Variable name
- `track_default_gateway` (Boolean) Enable or disable default gateway tracking
- `track_default_gateway_variable` (String) Variable name
- `track_interface_tag` (Number) OMP Tag attached to routes based on interface tracking
- `track_interface_tag_variable` (String) Variable name
- `track_transport` (Boolean) Configure tracking of transport
- `track_transport_variable` (String) Variable name
- `trackers` (Attributes List) Tracker configuration (see [below for nested schema](#nestedatt--trackers))
- `transport_gateway` (Boolean) Enable transport gateway
- `transport_gateway_variable` (String) Variable name
- `version` (Number) The version of the feature template

<a id="nestedatt--geo_fencing_sms_phone_numbers"></a>
### Nested Schema for `geo_fencing_sms_phone_numbers`

Read-Only:

- `number` (String) Mobile number, ex: +1231234414
- `number_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.


<a id="nestedatt--object_trackers"></a>
### Nested Schema for `object_trackers`

Read-Only:

- `boolean` (String) Type of grouping to be performed for tracker group
- `boolean_variable` (String) Variable name
- `group_tracks_ids` (Attributes List) Tracks id in group configuration (see [below for nested schema](#nestedatt--object_trackers--group_tracks_ids))
- `interface` (String) interface name
- `interface_variable` (String) Variable name
- `ip` (String) IP address of route
- `ip_variable` (String) Variable name
- `mask` (String) Route Ip Mask
- `mask_variable` (String) Variable name
- `object_number` (Number) Object tracker ID
- `object_number_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `sig` (String) service sig
- `sig_variable` (String) Variable name
- `vpn_id` (Number) VPN

<a id="nestedatt--object_trackers--group_tracks_ids"></a>
### Nested Schema for `object_trackers.group_tracks_ids`

Read-Only:

- `optional` (Boolean) Indicates if list item is considered optional.
- `track_id` (Number) Track id
- `track_id_variable` (String) Variable name



<a id="nestedatt--trackers"></a>
### Nested Schema for `trackers`

Read-Only:

- `boolean` (String) Type of grouping to be performed for tracker group
- `boolean_variable` (String) Variable name
- `elements` (List of String) Tracker member names separated by space
- `elements_variable` (String) Variable name
- `endpoint_api_url` (String) API url of endpoint
- `endpoint_api_url_variable` (String) Variable name
- `endpoint_dns_name` (String) DNS name of endpoint
- `endpoint_dns_name_variable` (String) Variable name
- `endpoint_ip` (String) IP address of endpoint
- `endpoint_ip_variable` (String) Variable name
- `interval` (Number) Probe interval <10..600> seconds
- `interval_variable` (String) Variable name
- `multiplier` (Number) Probe failure multiplier <1..10> failed attempts
- `multiplier_variable` (String) Variable name
- `name` (String) Tracker name
- `name_variable` (String) Variable name
- `optional` (Boolean) Indicates if list item is considered optional.
- `threshold` (Number) Probe Timeout threshold <100..1000> milliseconds
- `threshold_variable` (String) Variable name
- `transport_endpoint_ip` (String) IP address of endpoint
- `transport_endpoint_ip_variable` (String) Variable name
- `transport_endpoint_port` (Number) TCP/UDP port pf endpoint
- `transport_endpoint_port_variable` (String) Variable name
- `transport_endpoint_protocol` (String) transport protocol: TCP/UDP
- `transport_endpoint_protocol_variable` (String) Variable name
- `type` (String) Default(Interface)
- `type_variable` (String) Variable name
