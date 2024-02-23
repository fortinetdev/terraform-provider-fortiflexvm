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
  # account_id = 12345 # optional
  config_id   = 42
  folder_path = "My Assets"           # optional
  status      = ["ACTIVE", "PENDING"] # optional
}

output "my_groups_nexttoken" {
  value = data.fortiflexvm_groups_nexttoken.example
}
```

## Argument Reference

The following arguments are supported:

**Either account_id or config_id is required.**

* `account_id` - (Optional/Number) The account ID.
* `config_id` (Optional/Number) The ID of a configuration.
* `folder_path` (Optional/String) Folder path.
* `status` (Optional/List of String) The status of the entitlement.

## Attribute Reference

The following attributes are exported:

* `id` - (String) An ID for the resource. Its value will be `{config_id}.{folder_path}`. For example: "42.My Assets". If config_id or folder_path is not specified, their value will be "none". For example: "none.My Assets" or "42.none".
* `entitlements` - (List of Object) **One** Virtual Machine with unused token. The structure of [`entitlements` block](#nestedatt--entitlements)is documented below.

<a id="nestedatt--entitlements"></a>
The `entitlements` block contains:

* `account_id` - (Number) The account ID of this entitlements. If not specified argument `account_id`, the value of attribute `entitlements->account_id` will be 0.
* `config_id` - (Number) The ID of the configuration this entitlement used.
* `description` - (String) The description of the entitlement.
* `end_date` - (String) Entitlement end date.
* `serial_number` - (String) The unique serial number of the entitlement.
* `start_date` - (String) Entitlement creation date.
* `status` - (String) Entitlement status. Possible values: `PENDING`, `ACTIVE`, `STOPPED` or `EXPIRED`.
* `token` - (String) Entitlement token. Empty for hardware entitlements.
* `token_status` - (String) The status of the Entitlement token. Possible values: `NOTUSED` or `USED`. Empty for hardware entitlements.



