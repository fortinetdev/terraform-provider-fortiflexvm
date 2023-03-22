---
subcategory: "VMS"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_vms_list"
description: |-
  Get list of existing VMs for a Flex VM Configuration.
---

# Data Source: fortiflexvm_vms_list
Get list of existing VMs for a Flex VM Configuration.

## Example Usage

```hcl
data "fortiflexvm_vms_list" "example" {
    config_id = 42
}

output "my_vms_list"{
    value = data.fortiflexvm_vms_list.example
}
```

## Argument Reference

The following argument is required:

* `config_id` - (Required/Number) The ID of a Flex VM Configuration.

## Attribute Reference

The following attributes are exported:

* `id` - (String) An ID for the resource.
* `vms` - (List of Object) List of existing VMs using the specified Flex VM Configuration. The structure of [`vms` block](#nestedatt--vms) is documented below.

<a id="nestedatt--vms"></a>
The `vms` block contains:

* `config_id` - (Number) The ID of the Flex VM Configuration this VM used.
* `description` - (String) The description of VM.
* `end_date` - (String) VM end date.
* `serial_number` - (String) The unique serial number of the VM.
* `start_date` - (String) VM creation date.
* `status` - (String) VM status. Possible values: `PENDING`, `ACTIVE`, `STOPPED` or `EXPIRED`.
* `token` - (String) VM token.
* `token_status` - (String) The status of the VM token. Possible values: `NOTUSED` or `USED`.
