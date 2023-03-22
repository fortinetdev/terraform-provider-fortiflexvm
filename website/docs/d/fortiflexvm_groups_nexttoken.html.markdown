---
subcategory: "Groups"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_groups_nexttoken"
description: |-
  Get next available (unused) token.
---

# Data Source: fortiflexvm_groups_nexttoken
Get next available (unused) token.

Returns first available token by asset folder or Configuration id (or both can be specified in the request).

## Example Usage

```hcl
data "fortiflexvm_groups_nexttoken" "example" {
    config_id = 42
    folder_path = "My Assets"
}

output "output1"{
    value = data.fortiflexvm_groups_nexttoken.example
}
```

## Argument Reference

The following arguments are supported:

**Althrough both arguments are optional, please at least specify one argument to get the result.**

* `config_id` (Optional/Number) The ID of a Flex VM Configuration.
* `folder_path` (Optional/String) Folder path.

## Attribute Reference

The following attributes are exported:

* `id` - (String) An ID for the resource.
* `vms` - (List of Object) **One** Virtual Machine with unused token. The structure of [`vms` block](#nestedatt--vms)is documented below.

<a id="nestedatt--vms"></a>
The `vms` block contains:

* `config_id` - (Number) The ID of the Flex VM Configuration this VM used.
* `description` - (String) The description of the VM.
* `end_date` - (String) VM end date.
* `serial_number` - (String) The unique serial number of the VM.
* `start_date` - (String) VM creation date.
* `status` - (String) VM status. Possible values: `PENDING`, `ACTIVE`, `STOPPED` or `EXPIRED`.
* `token` - (String) VM token.
* `token_status` - (String) The status of the VM token. Possible values: `NOTUSED` or `USED`.


