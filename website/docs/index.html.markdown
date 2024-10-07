---
layout: "fortiflexvm"
page_title: "Provider: FortiFlexVM"
sidebar_current: "docs-fortiflexvm-index"
description: |-
  The FortiFlexVM provider interacts with FortiFlex.
---

# FortiFlexVM Provider

The FortiFlexVM provider is used to interact with the resources supported by FortiFlex. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources. (The terms `FortiFlexVM`, `FortiFlex` and `FlexVM` refer to the same product. Due to historical reasons, the provider name is defined as `FortiFlexVM`.)


## Example Usage

```hcl
terraform {
  required_providers {
    fortiflexvm = {
      version = "2.3.3"
      source  = "fortinetdev/fortiflexvm"
    }
  }
}

# Configure the Provider for FortiFlexVM
provider "fortiflexvm" {
  username = "ABCDEFG"
  password = "HIJKLMN"

  # If you want to import resource_config, please specify your program_serial_number here
  import_options = ["program_serial_number=ELAVMS00000XXXXX"]
}

# Create one congifuration
# If import, please add `import_options = ["program_serial_number=ELAVMS00000XXXXX"]` in `provider "fortiflexvm"`
# Then use: terraform import fortiflexvm_config.labelname <your config_id>
resource "fortiflexvm_config" "example"{
  product_type = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "example_name"
  fgt_vm_bundle {
    cpu_size =  "2"     # "1", "2", "4", "8", "16", "32", "2147483647"
    service_pkg = "ATP" # "FC", "UTM", "ENT", "ATP"
    vdom_num = 10       # 0 ~ 500
  }
}


# Create one VM entitlement
# If import, use: terraform import fortiflexvm_entitlements_vm.labelname <serial_number>.<config_id>
# For example: terraform import fortiflexvm_entitlements_vm.example FGVMMLTM23001273.3196
resource "fortiflexvm_entitlements_vm" "example"{ 
  config_id = fortiflexvm_config.example.id
  description = "Your description" # Optional.
  end_date = "2023-11-12T00:00:00" # Optional. If not set or empty "", it will use the program's end date automatically.
  # folder_path = "My Assets" # Optional. If not set, new VM will be in "My Assets"
  # status = "ACTIVE" # "ACTIVE" or "STOPPED". Optional. It has many restrictions. Not recommended to set it manually.
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
- `import_options` - (Optional/List of Object)  This parameter is only used for import in some special cases. When the resource to be imported includes pkg parameter, you need to assign a value to the parameter here, for example:

    ```hcl
    provider "fortiflexvm" {
      username = "ABCDEFG"
      password = "HIJKLMN"

      import_options = ["pkg=default"]
    }
    ```
