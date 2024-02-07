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
  # account_id = 12345 # optional
}

output "my_groups_list" {
  value = data.fortiflexvm_groups_list.example
}
```

## Argument Reference

The following argument is required:

* `account_id` - (Optional/Number) The account ID.

## Attribute Reference

The following attributes are exported:

* `groups` - (List of Object) List of groups. The structure of [`groups` block](#nestedatt--groups) is documented below.
* `id` - (String) An ID for the resource. Its value is the string "GroupsList".

<a id="nestedatt--groups"></a>
The `groups` block contains:

* `account_id` - (Number) The account ID of this group. If not specified argument `account_id`, the value of attribute `groups->account_id` will be 0.
* `available_tokens` - (Number) The number of available tokens in the group.
* `folder_path` - (String) The folder path of the group.
* `used_tokens` - (Number) The number of tokens used in the group.


