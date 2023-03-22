---
subcategory: "VMS"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_vms_create"
description: |-
  Create one or more VMs based on a Flex VM Configuration.
---

# fortiflexvm_vms_create

Create one or more VMs based on a Flex VM Configuration.

-> This API is only used to create one or more VMs. **Once you declare a resource block and use `terraform apply`, any modification to the resource blocks will not make efforts**. (e.g., if you already used the resource block in the `Example Usage` to apply one VM, you can not apply another VM by just replacing the value of `vm_count` with 2. You need to write another resource block to create additional VMs). To modify a VM, please refer to [fortiflexvm_vms_update](./fortiflexvm_vms_update.html.markdown).

## Example Usage

```hcl
resource "fortiflexvm_vms_create" "example"{
  config_id = 42
  description = "Create through Terraform"
  end_date = "2023-11-11T00:00:00"
  folder_path = "My Assets"
  vm_count = 1
}
```

## Argument Reference

The following arguments are supported:

* `config_id` - (Required/Number) The ID of a Flex VM Configuration.
* `description` - (Optional/String) The description of VM(s).
* `end_date` - (Optional/String) VM(s) end date. It can not be before today's date or after the program's end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DDThh:mm:ss`.
* `folder_path` - (Optional/String) The folder path of the VM(s).
* `vm_count` - (Optional/Number) The number of VM(s) to be created. The default value is 1.

## Attribute Reference

The following attribute is exported:

* `id` - (String) An ID for the resource.


