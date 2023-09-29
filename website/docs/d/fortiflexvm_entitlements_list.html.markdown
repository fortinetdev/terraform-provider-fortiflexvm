---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_list"
description: |-
  Get list of existing entitlements for a configuration.
---

# Data Source: fortiflexvm_entitlements_list

Get list of existing entitlements for a configuration.

Either config_id or (account_id + serial_number) should be provided.

## Example Usage

```hcl
data "fortiflexvm_entitlements_list" "example" {
    config_id = 42
    # either config_id or (account_id + serial_number) should be provided
    # account_id = 12345
    # serial_number = "ELAVMR0000000101"
}
output "my_entitlements_list"{
    value = data.fortiflexvm_entitlements_list.example
}
```

## Argument Reference

The following argument is required:

Either config_id or (account_id + serial_number) should be provided.

* `account_id` - (Optional/Number) Account ID.
* `config_id` - (Optional/Number) The ID of the configuration.
* `program_serial_number` - (Optional/String) The unique serial number of the Program.

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
