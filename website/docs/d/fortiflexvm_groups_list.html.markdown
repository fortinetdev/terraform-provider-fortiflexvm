---
subcategory: "Groups"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_groups_list"
description: |-
  Get list of FortiFlex groups (asset folders).
---

# Data Source: fortiflexvm_groups_list
Get list of FortiFlex groups (asset folders).

Returns list of FortiFlex groups (asset folders that have FortiFlex products in them).


## Example Usage

```hcl
data "fortiflexvm_groups_list" "example" {
}

output "my_groups_list"{
    value = data.fortiflexvm_groups_list.example
}
```

## Argument Reference

No argument is required.

## Attribute Reference

The following attributes are exported:

* `groups` - (List of Object) List of groups. The structure of [`groups` block](#nestedatt--groups) is documented below.
* `id` - (String) An ID for the resource. Its value is the string "GroupsList".

<a id="nestedatt--groups"></a>
The `groups` block contains:

* `available_tokens` - (Number) The number of available tokens in the group.
* `folder_path` - (String) The folder path of the group.
* `used_tokens` - (Number) The number of tokens used in the group.


