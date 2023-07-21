---
subcategory: "Programs"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_programs_list"
description: |-
  Get list of FortiFlex Programs for the account.
---

# Data Source: fortiflexvm_programs_list
Get list of FortiFlex Programs for the account.

## Example Usage

```hcl
data "fortiflexvm_programs_list" "example" {}

output "my_programs_list"{
    value = data.fortiflexvm_programs_list.example
}
```

## Argument Reference

No argument is required.

## Attribute Reference

The following attributes are exported:

* `id` - (String) An ID for the resource. Its value is the string "ProgramsList".
* `programs` - (List of Object) List of FortiFlex Programs for the account. The structure of [`programs` block](#nestedatt--programs) is documented below.

<a id="nestedatt--programs"></a>
The `programs` block contains:

* `account_id` - (Number) Your account ID.
* `end_date` - (String) FortiFlex Program end date.
* `has_support_coverage` - (Boolean) <!-- Whether the current date is between the start_date and end_date. -->
* `serial_number` - (String) The unique serial number of this FortiFlex Program.
* `start_date` - (String) FortiFlex Program creation date.