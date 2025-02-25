---
subcategory: "Groups"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_groups_nexttoken"
description: |-
  Terraform ephemeral resource to get next available (unused) token.
---

# Ephemeral: fortiflexvm_groups_nexttoken
Terraform ephemeral resource to get next available (unused) token.

Returns first available token by asset folder or Configuration id (or both can be specified in the request).

## Example Usage

```hcl
provider "fortiflexvm" {
  username = "<USERNAME>"
  password = "<PASSWORD>"
}

ephemeral "fortiflexvm_groups_nexttoken" "example" {
  config_id   = 42
  folder_path = "My Assets"           # optional
  status      = ["ACTIVE", "PENDING"] # optional
}

provider "fortios" {
  hostname = "<HOSTNAME>"
  username = "<USERNAME>"
  password = "<PASSWORD>"
  insecure = "true"
}

resource "fortios_system_license_fortiflex" "test" {
  token_writeonly = ephemeral.fortiflexvm_groups_nexttoken.example.token
}
```

~> Write_only attribute supported on Terraform 1.11.0+. If you want to use resource `fortios_system_license_fortiflex`, make sure you have Terraform version  1.11.0+.

## Argument Reference

The following arguments are supported:

**Either account_id or config_id is required.**

* `account_id` - (Optional/Number) The account ID.
* `config_id` (Optional/Number) The ID of a configuration.
* `folder_path` (Optional/String) Folder path.
* `status` (Optional/List of String) The status of the entitlement.

## Read-Only

The following attributes are exported:

* `token` - (String) Entitlement token. 


