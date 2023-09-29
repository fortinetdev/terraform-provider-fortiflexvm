---
subcategory: "Entitlements"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_entitlements_vm_token"
description: |-
  Regenerate token for a VM.
---

# fortiflexvm_entitlements_vm_token

Regenerate token for a VM.

## Example Usage

```hcl
# if import, use: terraform import fortiflexvm_entitlements_vm_token.example FGVMMLTM23001325.3196
resource "fortiflexvm_entitlements_vm_token" "example"{ 
  config_id = 3196
  serial_number = "FGVMMLTM23001325"
  regenerate_token = true # If set as false, the provider would only provide the token and not regenerate the token.
}
output "entitlement_token"{
    value = fortiflexvm_entitlements_vm_token.example.token
}
```

## Argument Reference

The following arguments are supported:

* `config_id` - (Required/Number) The ID of a FortiFlex Configuration.
* `serial_number` - (Required/String) The ID of the VM entitlement.
* `regenerate_token` - (Required/Boolean) Whether to regenerate a new token. If this argument is `true`, every time you run `terraform apply`, the system will generate a new token for your VM entitlement. Please remember to set it as `false` if you don't want to regenerate the token anymore.


## Attribute Reference

The following attribute is exported:

* `account_id` - (Number) Account ID.
* `id` - (String) The ID of the resource. Its value will be {serial_number}.{config_id}. For example: "FGVMMLTM23001273.3196"
* `token` - (String) The token of the VM entitlement.
* `token_status` - (String) The status of the token. Possible value: "NOTUSED" or "USED"

## Import

```
terraform import fortiflexvm_entitlements_vm_token.labelname {{serial_number}}.{{config_id}}
# For example: terraform import fortiflexvm_entitlements_vm_token.example FGVMMLTM23001325.3196
```
