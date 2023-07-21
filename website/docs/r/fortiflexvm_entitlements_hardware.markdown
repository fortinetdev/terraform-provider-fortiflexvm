---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_hardware"
description: |-
  Create and update one hardware entitlement based on a configuration.
---

# fortiflexvm_entitlements_hardware

Create and update one hardware entitlement based on a configuration.


## Example Usage

```hcl
# If import, use: terraform import fortiflexvm_entitlements_hardware.labelname <serial_number>.<config_id>
# For example: terraform import fortiflexvm_entitlements_hardware.example FGT70FTK22000001.5010
resource "fortiflexvm_entitlements_hardware" "example"{ 
  config_id = 5010
  end_date = "2023-11-12T00:00:00" # Optional. If not set, it will use the program end date automatically.
  # status = "ACTIVE" # "ACTIVE" or "STOPPED". Optional. It has many restrictions. Not recommended to set it manually.
}
output "new_entitlement_hw"{
    value = fortiflexvm_entitlements_hardware.example
}
```

## Argument Reference

The following arguments are supported:

* `config_id` - (Required/Number) The ID of a configuration.
* `description` - (Optional/String) The description of hardware entitlement.
* `end_date` - (Optional/String) Hardware entitlement end date. It can not be before today's date or after the program's end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DDThh:mm:ss`. If not specify, it will use the program end date automatically.
* `status` - (Optional/String) "ACTIVE" or "STOPPED". Use "STOPPED" if you want to stop the VM entitlement. Use "ACTIVE" if you want to reactivate it. It has many restrictions. Not recommended to set it manually.

## Attribute Reference

The following attribute is exported:

* `id` - (String) The ID of the resource. Its value will be {serial_number}.{config_id}. For example: "FGT70FTK22000001.5010"
* `serial_number` - (String) The ID of the hardware entitlement.
* `start_date` - (String) Start date. Its format is `YYYY-MM-DDThh:mm:ss.sss`. For example: "2023-07-07T14:32:09.873".
* `status` (String) Four possible values: "PENDING", "ACTIVE", "EXPIRED" and "STOPPED". This attribute can be set as "ACTIVE" or "STOPPED" manually.

## Import

```
terraform import fortiflexvm_entitlements_hardware.labelname {{serial_number}}.{{config_id}}
# For example: terraform import fortiflexvm_entitlements_hardware.example FGT70FTK22000001.5010
```
