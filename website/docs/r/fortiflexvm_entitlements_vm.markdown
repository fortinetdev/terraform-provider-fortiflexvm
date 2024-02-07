---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_vm"
description: |-
  Create and update one VM entitlement based on a configuration.
---

# fortiflexvm_entitlements_vm

Create and update one VM entitlement based on a configuration.

!> The status of newly created VMs is `PENDING`. You can't update a VM entitlement if its current status is `PENDING`. After you [use VM token to activate a virtual machine](https://docs.fortinet.com/document/flex-vm/latest/administration-guide/256339/injecting-the-flex-vm-license), its status will be changed to "ACTIVE" and you can update this resource.


~> By updating this resource, you can change the data in the FortiFlex Platform immediately. Yet it may take several hours for the VMs to update their licenses automatically.


## Example Usage

```hcl
# Create one VM entitlement
# If import, use: terraform import fortiflexvm_entitlements_vm.labelname <serial_number>.<config_id>
# For example: terraform import fortiflexvm_entitlements_vm.example FGVMMLTM23001273.3196
resource "fortiflexvm_entitlements_vm" "example" {
  config_id     = 3196
  description   = "Your description"    # Optional.
  end_date      = "2024-11-12T00:00:00" # Optional. If not set, it will use the program end date automatically.
  # folder_path = "My Assets"           # Optional. If not set, new VM will be in "My Assets"
  # status      = "ACTIVE"              # "ACTIVE" or "STOPPED". Optional. It has many restrictions. Not recommended to set it manually.
}
output "new_entitlement" {
  value = fortiflexvm_entitlements_vm.example
}
output "new_entitlement_token" {
  value = fortiflexvm_entitlements_vm.example.token
}

# Update Entitlement
# The status of newly created VMs is PENDING. You can't update a VM entitlement if its current status is PENDING.
# After you use VM token to activate a VM, its status will be changed to "ACTIVE" and you can use following examples.

# Update entitlement information (Please use this token to activate a VM first)
# If import, use: terraform import fortiflexvm_entitlements_vm.labelname <serial_number>.<config_id>
# After you create or import a fortiflexvm_entitlements_vm resource, you can update it:
resource "fortiflexvm_entitlements_vm" "example" {
  config_id   = 3196                  # New config_id value or unchanged
  description = "Your description"    # Optional.
  end_date    = "2024-11-12T00:00:00" # Optional. If not set, it will use the program end date automatically.
}

# Stop or reactivate a VM (Please use this token to activate a VM first)
resource "fortiflexvm_entitlements_vm" "example" {
  config_id = 3196      # Previous config_id
  status    = "STOPPED" # "ACTIVE" or "STOPPED". Optional. It has many restrictions. Not recommended to set it manually.
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional/Number) Account ID.
* `config_id` - (Required/Number) The ID of a FortiFlex Configuration.
* `description` - (Optional/String) The description of VM entitlement.
* `end_date` - (Optional/String) VM entitlement end date. It can not be before today's date or after the program's end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DDThh:mm:ss`. If not specify, it will use the program end date automatically.
* `folder_path` - (Optional/String) The folder path of the VM. If not set, the new VM will be in "My Assets"
* `status` - (Optional/String) "ACTIVE" or "STOPPED". Use "STOPPED" if you want to stop the VM entitlement. Use "ACTIVE" if you want to reactivate it. It has many restrictions. Not recommended to set it manually.

## Attribute Reference

The following attribute is exported:

* `account_id` - (Number) Account ID.
* `id` - (String) The ID of the resource. Its value will be {serial_number}.{config_id}. For example: "FGVMMLTM23001273.3196"
* `serial_number` - (String) The ID of the VM entitlement.
* `start_date` - (String) Start date. Its format is `YYYY-MM-DDThh:mm:ss.sss`. For example: "2024-07-07T14:32:09.873".
* `status` (String) Four possible values: "PENDING", "ACTIVE", "EXPIRED" and "STOPPED". This attribute can be set as "ACTIVE" or "STOPPED" manually.
* `token` - (String) The token of the VM entitlement.
* `token_status` - (String) The status of the token. Possible value: "NOTUSED" or "USED"

## Import

```
terraform import fortiflexvm_entitlements_vm.labelname {{serial_number}}.{{config_id}}
# For example: terraform import fortiflexvm_entitlements_vm.example FGVMMLTM23001273.3196
```
