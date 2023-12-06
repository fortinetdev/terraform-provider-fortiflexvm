---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_cloud"
description: |-
  Create and update one cloud entitlement based on a configuration.
---

# fortiflexvm_entitlements_cloud

Create and update one cloud entitlement based on a configuration.

~> Each account can create at most one entitlement for each cloud product.

## Example Usage

```hcl
resource "fortiflexvm_config" "example"{
  product_type = "FC_EMS_CLOUD"
  program_serial_number = "ELAVMS0000xxxxxx"
  name = "FC_EMS_OP_template"
  fc_ems_cloud {
    ztna_num = 225                # Value should be divisible by 25. Number between 0 and 25000 (inclusive)
    ztna_fgf_num = 225            # Value should be divisible by 25. Number between 0 and 25000 (inclusive)
    epp_ztna_num = 125            # Value should be divisible by 25. Number between 0 and 25000 (inclusive)
    epp_ztna_fgf_num = 125        # Value should be divisible by 25. Number between 0 and 25000 (inclusive)
    chromebook = 100              # Value should be divisible by 25. Number between 0 and 25000 (inclusive) 
    addons = ["BPS"]              # [] or ["BPS"]
  }
  status = "ACTIVE"
}

# Each account can create at most one entitlement for each cloud product.
resource "fortiflexvm_entitlements_cloud" "example"{ 
  config_id = fortiflexvm_config.example.id
  description = "Use v2 terraform" # Optional.
  end_date = "2023-12-12T00:00:00" # Optional. If not set or empty "", it will use the program's end date automatically.
  folder_path = "My Assets/v2" # Optional. If not set, new VM will be in "My Assets"
  # status = "ACTIVE" # Optional, It has many restrictions. Not recommended to set it manually.
}
output "new_entitlement"{
    value = fortiflexvm_entitlements_cloud.example
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional/Number) Account ID.
* `config_id` - (Required/Number) The ID of a FortiFlex Configuration.
* `description` - (Optional/String) The description of the entitlement.
* `end_date` - (Optional/String) Cloud entitlement end date. It can not be before today's date or after the program's end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DDThh:mm:ss`. If not specify, it will use the program end date automatically.
* `folder_path` - (Optional/String) The folder path of the cloud entitlement. If not set, the new cloud entitlement will be in "My Assets"
* `status` - (Optional/String) "ACTIVE" or "STOPPED". Use "STOPPED" if you want to stop the cloud entitlement. Use "ACTIVE" if you want to reactivate it. It has many restrictions. Not recommended to set it manually.

## Attribute Reference

The following attribute is exported:

* `account_id` - (Number) Account ID.
* `id` - (String) The ID of the resource. Its value will be {serial_number}.{config_id}. For example: "FEMSPO8823000143.3196"
* `serial_number` - (String) The ID of the cloud entitlement.
* `start_date` - (String) Start date. Its format is `YYYY-MM-DDThh:mm:ss.sss`. For example: "2023-07-07T14:32:09.873".
* `status` (String) Four possible values: "PENDING", "ACTIVE", "EXPIRED" and "STOPPED". This attribute can be set as "ACTIVE" or "STOPPED" manually.

## Import

```
terraform import fortiflexvm_entitlements_cloud.labelname {{serial_number}}.{{config_id}}
# For example: terraform import fortiflexvm_entitlements_cloud.example FEMSPO8823000143.3196
```
