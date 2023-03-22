---
layout: "fortiflexvm"
page_title: "Provider: FortiFlexVM"
sidebar_current: "docs-fortiflexvm-index"
description: |-
  The FortiFlexVM provider interacts with FortiFlexVM.
---

# FortiFlexVM Provider

The FortiFlexVM provider is used to interact with the resources supported by FortiFlexVM. We need to configure the provider with the proper credentials before it can be used. Please use the navigation on the left to read more details about the available resources.


## Example Usage

```hcl
# Configure the Provider for FortiFlexVM
provider "fortiflexvm" {
  username = "ABCDEFG"
  password = "HIJKLMN"
}

# Create one VM
resource "fortiflexvm_vms_create" "example"{
  config_id = 42
  description = "Create through Terraform"
  end_date = "2023-11-11 00:00:00"
  folder_path = "My Assets"
  vm_count = 1
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

You can provide your credentials via the `FLEXVM_ACCESS_USERNAME` and `FLEXVM_ACCESS_PASSWORD` environment variables. Note that setting your FortiFlexVM credentials using static credentials variables will override the environment variables.

Usage:

```shell
$ export "FLEXVM_ACCESS_USERNAME"="ABCDEFG"
$ export "FLEXVM_ACCESS_PASSWORD"="HIJKLMN"
```

-> If you provide your credentials through environment variables, and the variables contain the character "!", please put double quotes around the exclamation mark "!" to avoid the autoescaping problem.
For example, if your password is "123!456", please use the command  `$export "FLEXVM_ACCESS_PASSWORD"="123""!""456"`.

Then configure the FortiFlexVM Provider as follows:

```hcl
provider "fortiflexvm" {}
```



## Argument Reference

The following arguments are supported:

- `username` - (Optional/String) Your username. It must be provided, but it can also be sourced from the `FLEXVM_ACCESS_USERNAME` environment variable.
- `password` - (Optional/String) Your password. It must be provided, but it can also be sourced from the `FLEXVM_ACCESS_PASSWORD` environment variable.
- `import_options` - (Optional/List of Object)  This parameter is only used for import in some special cases. When the resource to be imported includes pkg parameter, you need to assign a value to the parameter here, for example:

    ```hcl
    provider "fortiflexvm" {
      username = "ABCDEFG"
      password = "HIJKLMN"

      import_options = ["pkg=default"]
    }
    ```
