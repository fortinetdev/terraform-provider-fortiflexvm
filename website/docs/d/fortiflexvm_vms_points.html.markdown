---
subcategory: "VMS"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_vms_points"
description: |-
  Get point usage for VMs.
---

# Data Source: fortiflexvm_vms_points
Get point usage for VMs.

Returns total points consumed by one or more virtual machines in a date range.

## Example Usage

```hcl
data "fortiflexvm_vms_points" "example" {
    config_id = 42
    start_date = "2022-11-25"
    end_date = "2022-03-07"
}

output "output1"{
    value = data.fortiflexvm_vms_points.example
}

```

## Argument Reference

The following arguments are required:

* `config_id` - (Required/Number) The ID of a Flex VM Configuration.
* `end_date` - (Required/String) Specify an end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DD`.
* `start_date` - (Required/String) Specify a start date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DD`.

## Attribute Reference

The following attributes are exported:

* `id` - (String) An ID for the resource.
* `vms` - (List of Object) List of existing VMs using the specified Flex VM Configuration. The structure of [`vms` block](#nestedatt--vms) is documented below.

<a id="nestedatt--vms"></a>
The `vms` block contains:

* `points` - (Number) The points consumed by this VM in the date range.
* `serial_number` - (String) The unique serial number of the VM.


