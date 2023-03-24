---
subcategory: "VMS"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_vms_update"
description: |-
  Update a VM, stop or reactivate a VM, or regenerate a VM token.
---

# fortiflexvm_vms_update
Update a VM, stop or reactivate a VM, or regenerate a VM token.

!> You can use [fortiflexvm_vms_create](./fortiflexvm_vms_create.html.markdown) to create VMs. The status of newly created VMs is `PENDING`. You can't update a VM if its current status is `PENDING`. Please [use VM token to activate a virtual machine](https://docs.fortinet.com/document/flex-vm/latest/administration-guide/256339/injecting-the-flex-vm-license) before using this API.

~> By using this resource, you can change the data in the FortiFlex VM Platform immediately. Yet it may take several hours for the VMs to update their licenses automatically.

## Example Usage

```hcl
resource "fortiflexvm_vms_update" "example"{
  serial_number = "FGVMMLTM0000XXXX"
  config_id = 42
  description = "Modify through Terraform"
  end_date = "2023-11-11T00:00:00"
  regenerate_token = false
  status = "DISABLE"
}
```

-> If you want to update the `description` or `end_date`, please remember to specify the `config_id`, even if it is unchanged.

## Argument Reference

The following arguments are supported:

* `serial_number` - (Required/String) The unique serial number of the VM to be updated.
* `config_id` - (Required if you want to update the `config_id`, `description` or `end_date`/Number) Set a new Flex VM Configuration.
* `description` - (Optional/String) Set a new description.
* `end_date` - (Optional/String) Set a new end date. Any format that satisfies [ISO 8601](https://www.w3.org/TR/NOTE-datetime-970915.html) is accepted. Recommended format: `YYYY-MM-DDThh:mm:ss`.
* `regenerate_token` - (Optional/Boolean) Whether to regenerate a new token. If this argument is `true`, every time you run `terraform apply`, the system will generate a new token for your VM. Please remember to set it as `false` if you don't want to regenerate the token anymore.
* `status` - (Optional/String) Valid values: `ACTIVE` or `STOPPED`. If you want to use this argument, it is highly recommended to also specify the argument `config_id`.

## Attribute Reference

The following attribute is exported:

* `id` - (String) An ID for the resource.

## Import

~> Currently, our Provider only supports importing one VM profile.

VM profile can be imported by using the following steps:

First, specify the `config_id` when you configure the provider.
```
provider "fortiflexvm" {
  username = "ABCDEFG"
  password = "HIJKLMN"
  import_options= toset(["config_id=42"])
}
```

Then, use the following command to import the VM profile.
```
terraform import fortiflexvm_vms_update.labelname {{serial_number}}
```