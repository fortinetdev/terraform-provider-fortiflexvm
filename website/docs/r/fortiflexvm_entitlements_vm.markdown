---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_vm"
description: |-
  Create and update one VM entitlement based on a configuration.
---

# fortiflexvm_entitlements_vm

Create and update one VM entitlement based on a configuration.

!> Due to the properties of Fortiflex, after you apply `terraform destroy` the status of the entitlement will change to `STOPPED` and stop being charged, rather than being destroyed. To reuse one STOPPED entitlement, please specify `serial_number` in `fortiflexvm_entitlements_vm`. To reuse a group of STOPPED entitlements, please use `fortiflexvm_retrieve_vm_group`.

~> The status of newly created VMs is `PENDING`. After you [use VM token to activate a virtual machine](https://docs.fortinet.com/document/flex-vm/latest/administration-guide/256339/injecting-the-flex-vm-license), its status will be changed to "ACTIVE".


## Example Usage

Create one VM entitlement
```hcl
# If you don't specify serial_number, it will create a new entitlement.
resource "fortiflexvm_entitlements_vm" "example" {
  config_id     = 42
  description   = "Your description"      # Optional.
  # end_date    = "2024-11-12T00:00:00"   # Optional. If not set, it will use the program end date automatically.
  # folder_path = "My Assets"             # Optional. If not set, new VM will be in "My Assets"
  # status      = "ACTIVE"                # "ACTIVE" or "STOPPED". Optional.
  # refresh_token_when_destroy = True     # Optional. Refresh the token when you destroy this resource
}
output "new_entitlement" {
  value = fortiflexvm_entitlements_vm.example
}
output "new_entitlement_token" {
  value = fortiflexvm_entitlements_vm.example.token
}
```

Import & update existing entitlement
```hcl
# If specify both serial_number and config_id, it will import the existing entitlement.
resource "fortiflexvm_entitlements_vm" "example" {
  config_id     = 42                    # New config_id value or unchanged
  serial_number = "FGVMXXXX00000000"
  # description = "Your description"    # Optional.
  # end_date    = "2024-11-12T00:00:00" # Optional. If not set, it will use the program end date automatically.
  status      = "ACTIVE"                # "ACTIVE" or "STOPPED". Optional.
}
# You can also import by using: terraform import fortiflexvm_entitlements_vm.labelname <serial_number>.<config_id>
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional/Number) Account ID.
* `config_id` - (Required/Number) The ID of a FortiFlex Configuration.
* `description` - (Optional/String) The description of VM entitlement.
* `end_date` - (Optional/String) VM entitlement end date. It can not be before today's date or after the program's end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DDThh:mm:ss`. If not specify, it will use the program end date automatically.
* `folder_path` - (Optional/String) The folder path of the VM. If not set, the new VM will be in "My Assets"
* `refresh_token_when_destroy` - (Optional/Boolean) Default value is false. If set it as true, the token of this entitlement will be refreshed when you use `terraform destroy`.
* `serial_number` - (Optional/String) If you specify serial_number, terraform will import the existing entitlement. If you don't specify it, terraform will create a new entitlement.
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

Method 1: Specify `config_id`
```hcl
# If specify both serial_number and config_id, it will import the existing entitlement.
resource "fortiflexvm_entitlements_vm" "example" {
  config_id     = 42                    # New config_id value or unchanged
  serial_number = "FGVMXXXX00000000"
  # description = "Your description"    # Optional.
  # end_date    = "2024-11-12T00:00:00" # Optional. If not set, it will use the program end date automatically.
  status      = "ACTIVE"                # "ACTIVE" or "STOPPED". Optional.
}
# You can also import by using: terraform import fortiflexvm_entitlements_vm.labelname <serial_number>.<config_id>
```

Method 2: Use `terraform import`
```
terraform import fortiflexvm_entitlements_vm.labelname {{serial_number}}.{{config_id}}
# For example: terraform import fortiflexvm_entitlements_vm.example FGVMMLTM23001273.3196
```
