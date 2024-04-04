---
subcategory: "Special"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_retrieve_vm_group"
description: |-
  Retrieve a group of STOPPED VM entitlements and change them to ACTIVE.
---

# fortiflexvm_retrieve_vm_group

This resource is used to retrieve a group of existing `STOPPED` (or `STOPPED` and `PENDING`) VM entitlements (with empty description) and change them to `ACTIVE`.

When you destroy this resource, the retrieved entitlements will refresh their token, change their status back to `STOPPED` and change their description to empty.

If you want to create new entitlements, please use `fortiflexvm_entitlements_vm`.

## How this resource works

~> This resource is special. It is important to know how this resource works before you use it.

To determine whether this resource can hold one entitlement, this resource queries the information of this entitlement first. If the status of this entitlement is "STOPPED" and the description is an empty string, the resource will update the entitlement's description to its `task_name` (to pre-hold this entitlement), sleep for `preempt_interval` (default is 1) second, and query the information of this entitlement again. If the description is still the same, then the resource owns this entitlement and changes entitlement status to "ACTIVE". If the description has changed, it means other tasks pre-hold this entitlement at the same time. Then the resource gives up this entitlement. 

By doing the above steps, this resource can do its best effort to avoid entitlement overlap when more than 2 retrieving entitlement requests are running at the same time.


## Example Usage

Retrieve "STOPPED" entitlements
```hcl
resource "fortiflexvm_retrieve_vm_group" "task1" {
  task_name = "task1" # Unique task name
  config_id = 1234    # Your config ID
  count_num = 3
  require_exact_count = true   # If retrieve less than 3 (count_num) entitlements, release retrieved entitlements and report an error
}
resource "fortiflexvm_retrieve_vm_group" "task2" {
  task_name = "task2" # Unique task name
  config_id = 1234    # Your config ID
  count_num = 3
  require_exact_count = true   # If retrieve less than 3 (count_num) entitlements, release retrieved entitlements and report an error
}
output "task1_tokens" {
  value = { for key, vm in fortiflexvm_retrieve_vm_group.task1.entitlements : vm.serial_number => vm.token }
}
output "task2_tokens" {
  value = { for key, vm in fortiflexvm_retrieve_vm_group.task2.entitlements : vm.serial_number => vm.token }
}
```

Retrieve both "STOPPED" and "PENDING" entitlements
```hcl
resource "fortiflexvm_retrieve_vm_group" "task1" {
  task_name           = "task1" # Unique task name
  config_id           = 1234    # Your config ID
  count_num           = 3
  retrieve_status     = ["STOPPED", "PENDING"] # Default value is ["STOPPED"]
  require_exact_count = true   # If retrieve less than 3 (count_num) entitlements, release retrieved entitlements and report an error
}
```

## Argument Reference

* `task_name` - (Required/String) Name of your task. It should be unqiue. This argument cannot be modified after the resource is created.
* `count_num` - (Required/Number) Number of entitlements you want.
* `config_id` - (Required/Number) The ID of the configuration. This argument cannot be modified after the resource is created.
* `preempt_interval` - (Optinal/Number) Default is 1. The second wait to preempt each entitlement. The larger this value, the longer time you need to wait, and the less probability you get entitlement overlap. Normally, 1 second is good enough.
* `refresh_token_when_destroy` - (Optinal/Boolean) Default value is true. If it is true, the token of all entitlements will be refreshed when you use `terraform destroy`.
* `refresh_token_when_create` - (Optinal/Boolean) Default value is false. If it is true, the token of all entitlements will be refreshed when you use `terraform apply` and create the resource.
* `retrieve_status` - (Optinal/List of string) The entitlements with what status you want to retrieve. The default value is ["STOPPED"]. You can set it as ["STOPPED", "PENDING"] if you want to retrieve both "STOPPED" and "PENDING" entitlements.
* `require_exact_count`- (Optinal/Boolean) Default value is false. If it is true and the resource retrieves less than (count_num) entitlements, it will release retrieved entitlements and report an error.


## Attribute Reference

The following attributes are exported:

* `id` - (String) The ID of the configuration. Its value is variable `config_id`.
* `entitlements` - (List of Object) List of existing entitlements using the specified configuration. The structure of [`entitlements` block](#nestedatt--entitlements) is documented below.

<a id="nestedatt--entitlements"></a>
The `entitlements` block contains:

* `account_id` - (Number) Account ID.
* `config_id` - (Number) The ID of the configuration this entitlement used.
* `description` - (String) The description of entitlement.
* `end_date` - (String) Entitlement end date.
* `serial_number` - (String) The unique serial number of the entitlement.
* `start_date` - (String) Entitlement creation date.
* `status` - (String) Entitlement status. Possible values: `PENDING`, `ACTIVE`, `STOPPED` or `EXPIRED`.
* `token` - (String) Entitlement token. Empty for hardware entitlements.
* `token_status` - (String) The status of the Entitlement token. Possible values: `NOTUSED` or `USED`. Empty for hardware entitlements.

## Import

This resource does not support import.