---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_points"
description: |-
  Get point usage for entitlements.
---

# Data Source: fortiflexvm_entitlements_points
Get point usage for entitlements.

Returns total points consumed by one or more entitlements in a date range.

## Example Usage

```hcl
data "fortiflexvm_entitlements_points" "example" {
  # account_id = 12345 # optional
  config_id  = 42
  start_date = "2023-11-25"
  end_date   = "2024-10-07"
}

output "my_entitlements_points" {
  value = data.fortiflexvm_entitlements_points.example
}
```

## Argument Reference

The following arguments are required:

* `account_id` - (Optional/Number) The account ID.
* `config_id` - (Required/Number) The ID of a configuration.
* `end_date` - (Required/String) Specify an end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DD`.
* `start_date` - (Required/String) Specify a start date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DD`.

## Attribute Reference

The following attributes are exported:

* `id` - (String) An ID for the resource. Its value will be `{config_id}.{start_date}.{end_date}`. For example: "42.2022-11-25.2022-03-07".
* `entitlements` - (List of Object) List of existing entitlements using the specified configuration. The structure of [`entitlements` block](#nestedatt--entitlements) is documented below.

<a id="nestedatt--entitlements"></a>
The `entitlements` block contains:

* `account_id` - (Number) The account ID of this group.
* `points` - (Number) The points consumed by this entitlement in the date range.
* `serial_number` - (String) The unique serial number of the entitlement.


