---
subcategory: ""
layout: "fortiflexvm"
page_title: "Usecase"
description: |-
  Common usecases.
---

### EXAMPLE: Create one FortiGate VM Configuration and Entitlement

```hcl
resource "fortiflexvm_config" "example"{
  product_type = "FGT_VM_Bundle"
  program_serial_number = "ELAVMR00000XXXXXX"
  name = "example_configuration"
  fgt_vm_bundle {
    cpu_size            = 8           # 1 ~ 96
    service_pkg         = "FC"        # "FC", "UTP", "ENT", "ATP"
    vdom_num            = 10          # 0 ~ 500
    fortiguard_services = []          # "FGTAVDB", "FGTFAIS", "FGTISSS", "FGTDLDB", "FGTFGSA"
    cloud_services      = []          # "FGTFAMS", "FGTSWNM", "FGTSOCA", "FGTFAZC", "FGTSWOS", "FGTFSPA"
    # support_service = "NONE"        # "NONE", "FGTFCELU"
  }
}

resource "fortiflexvm_entitlements_vm" "example"{ 
  config_id = fortiflexvm_config.example.id
  description = ""
}
```

### EXAMPLE: Import one existing entitlement

To import one existing entitlement, you need to specify its `config_id` and `serial_number`.

```hcl
resource "fortiflexvm_entitlements_vm" "example"{ 
  config_id = 1234
  serial_number = "FGVMXXXX00000000"
}
```

### EXAMPLE: Retrieve STOPPED and PENDING entitlements

With the following configuration, `terraform apply` retrieves `count_num` STOPPED or PENDING entitlements that have the specified `config_id` and an empty description. Terraform changes these entitlements to `ACTIVE` and sets their description to `task_name`.

When you run `terraform destroy`, Terraform refreshes the tokens for the retrieved entitlements, changes their status back to `STOPPED`, and clears their description.


```hcl
resource "fortiflexvm_retrieve_vm_group" "task1" {
  task_name       = "UNIQUE_TASK_NAME"       # Unique task name
  config_id       = 1234                     # Your config ID
  count_num       = 3
  retrieve_status = ["STOPPED", "PENDING"]  # Supported values: "STOPPED", "PENDING"
}
output "task1_tokens" {
  value = { for key, vm in fortiflexvm_retrieve_vm_group.task1.entitlements : vm.serial_number => vm.token }
}
```
