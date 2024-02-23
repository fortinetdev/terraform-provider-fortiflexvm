---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_list"
description: |-
  Get list of existing entitlements for a configuration.
---

# Data Source: fortiflexvm_entitlements_list

Get list of existing entitlements for a configuration.

Either config_id or (account_id + program_serial_number) should be provided.

## Example Usage

```hcl
data "fortiflexvm_entitlements_list" "example" {
  # either config_id or (account_id + program_serial_number) should be provided
  config_id = 42
  # account_id = 12345
  # program_serial_number = "ELAVMR0000000101"

  # Filter options:
  # description   = ""                 # optional
  # serial_number = "FGVMXXXX00000000" # optional
  # status        = "PENDING"          # optional ACTIVE, PENDING, EXPIRED, STOPPED
  # token_status  = "NOTUSED"          # optional USED, NOTUSED
}
output "my_entitlements_list" {
  value = data.fortiflexvm_entitlements_list.example
}
```

## Argument Reference

The following argument is required:

Either config_id or (account_id + serial_number) should be provided.

* `account_id` - (Optional/Number) Account ID.
* `config_id` - (Optional/Number) The ID of the configuration.
* `description` - (Optional/String) Filter option. The retrieved entitlments must have the same description.
* `program_serial_number` - (Optional/String) The unique serial number of the Program.
* `serial_number` - (Optional/String) The retrieved entitlments must have the same serial_number.
* `status` - (Optional/String) Filter option. The retrieved entitlments must have the same status. `ACTIVE`, `STOPPED`, `PENDING` or `EXPIRED`.
* `token_status` - (Optional/String) Filter option. The retrieved entitlments must have the same token_status. `USED` or `NOTUSED`


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
