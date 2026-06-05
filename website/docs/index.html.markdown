---
layout: "fortiflexvm"
page_title: "Provider: FortiFlexVM"
sidebar_current: "docs-fortiflexvm-index"
description: |-
  The FortiFlexVM provider interacts with FortiFlex.
---

# FortiFlexVM Provider

The FortiFlexVM provider manages resources supported by FortiFlex.
Before using the provider, please configure it with valid credentials. Use the navigation on the left to learn more about the available resources. The terms `FortiFlexVM`, `FortiFlex`, and `FlexVM` refer to the same product. For historical reasons, the provider name remains `FortiFlexVM`.

## How to use this provider
To use this provider, [generate an API token for FortiFlexVM](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/guides/fortiflexvm_token) and review the [use cases](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/guides/usecase).

### Key concepts
**Configuration**: A configuration defines the parameters used to create entitlements and calculate charges.
- [fortiflexvm_config](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/resources/fortiflexvm_config): Create a new configuration or update an existing configuration.

**Entitlements**: Entitlements are licenses created from a configuration. Charges apply when an entitlement is `ACTIVE`; no charges apply when its status is `PENDING` or `STOPPED`.
- [fortiflexvm_entitlements_vm](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/resources/fortiflexvm_entitlements_vm): Create or update a VM entitlement based on a configuration.
- [fortiflexvm_entitlements_hardware](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/resources/fortiflexvm_entitlements_hardware): Create or update a hardware entitlement based on a configuration.
- [fortiflexvm_entitlements_cloud](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/resources/fortiflexvm_entitlements_cloud): Create or update a cloud entitlement based on a configuration.
- [fortiflexvm_entitlements_vm_token](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/resources/fortiflexvm_entitlements_vm_token): Regenerate the token for an entitlement.
- [fortiflexvm_retrieve_vm_group](https://registry.terraform.io/providers/fortinetdev/fortiflexvm/latest/docs/resources/fortiflexvm_retrieve_vm_group): Retrieve existing `STOPPED` VM entitlements, or both `STOPPED` and `PENDING` VM entitlements, with empty descriptions and change them to `ACTIVE`.

~> Due to the design of FortiFlex, configuration and entitlement records cannot be fully deleted. After you run `terraform destroy`, you will no longer be charged for the resources you created, but the configuration created by `fortiflexvm_config` and the entitlement created by `fortiflexvm_entitlements_vm` remain visible in the GUI. Their status changes to `STOPPED`. We recommend creating a `fortiflexvm_config` once and reusing it instead of creating a separate configuration for each entitlement.


## Example Usage：Create one configuration and one entitlement

```hcl
terraform {
  required_providers {
    fortiflexvm = {
      version = "~> 2.0"
      source  = "fortinetdev/fortiflexvm"
    }
  }
}

# Configure the Provider for FortiFlexVM
provider "fortiflexvm" {
  username = "ABCDEFG"
  password = "HIJKLMN"

  # Optional.
  # account_id = 12345
  # program_serial_number = "ELAVMR00000XXXXX"
}

# Create one congifuration
# If import exisiting one, please specify `config_id`
resource "fortiflexvm_config" "example"{
  product_type = "FGT_VM_Bundle"
  program_serial_number = "ELAVMR00000XXXXX"
  name = "example_name"
  fgt_vm_bundle {
    cpu_size            = 4      # 1 ~ 96
    service_pkg         = "FC"   # "FC", "UTP", "ENT", "ATP"
    vdom_num            = 10     # 0 ~ 500
    fortiguard_services = []     # "FGTAVDB", "FGTFAIS", "FGTISSS", "FGTDLDB", "FGTFGSA"
    cloud_services      = []     # "FGTFAMS", "FGTSWNM", "FGTSOCA", "FGTFAZC", "FGTSWOS", "FGTFSPA"
    support_service     = "NONE" # "NONE", "FGTFCELU"
  }
}


# Create one VM entitlement
# If import existing one, please specify `serial_number`
resource "fortiflexvm_entitlements_vm" "example"{ 
  config_id = fortiflexvm_config.example.id
  description = "Your description" # Optional.
  # end_date = "2023-11-12T00:00:00" # Optional. If not set or empty "", it will use the program's end date automatically.
  # folder_path = "My Assets" # Optional. If not set, new VM will be in "My Assets"
}
output "new_entitlement"{
    value = fortiflexvm_entitlements_vm.example
}
output "new_entitlement_token"{
    value = fortiflexvm_entitlements_vm.example.token
}

```


## Authentication

The FortiFlexVM provider offers a means of providing credentials for authentication. The following methods are supported:

- Static credentials
- Environment variables


### Static credentials

Static credentials can be provided by `username` and `password` parameters in the FortiFlexVM provider block.

Usage:

```hcl
provider "fortiflexvm" {
  username = "ABCDEFG"
  password = "HIJKLMN"
}
```

### Environment variables

You can provide your credentials via the `FORTIFLEX_ACCESS_USERNAME` and `FORTIFLEX_ACCESS_PASSWORD` environment variables. Note that setting your FortiFlexVM credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FORTIFLEX_ACCESS_USERNAME"="ABCDEFG"
$ export "FORTIFLEX_ACCESS_PASSWORD"="HIJKLMN"
```

-> If you provide your credentials through environment variables, and the variables contain the character "!", please put double quotes around the exclamation mark "!" to avoid the autoescaping problem.
For example, if your password is "123!456", please use the command  `$export "FORTIFLEX_ACCESS_PASSWORD"="123""!""456"`.

Then configure the FortiFlexVM Provider as follows:

```hcl
provider "fortiflexvm" {}
```



## Argument Reference

The following arguments are supported:

- `username` - (Optional/String) Your username. It must be provided, but it can also be sourced from the `FORTIFLEX_ACCESS_USERNAME` environment variable.
- `password` - (Optional/String) Your password. It must be provided, but it can also be sourced from the `FORTIFLEX_ACCESS_PASSWORD` environment variable.
- `account_id` - (Optional/Number) The default account ID. Resource or data source arguments take precedence when specified.
- `program_serial_number` - (Optional/String) The default FortiFlex Program serial number. Resource or data source arguments take precedence when specified.
- `import_options` - (Deprecated/Optional/List of Object) Deprecated. Specify `program_serial_number` directly instead.
