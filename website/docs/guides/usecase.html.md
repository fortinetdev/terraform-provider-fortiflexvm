---
subcategory: ""
layout: "fortiflexvm"
page_title: "Usecase"
description: |-
  Common usecase.
---

### EXAMPLE: Create one FortiGate VM Configuration and Entitlement

```hcl
resource "fortiflexvm_config" "example"{
  product_type = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS0000000601"
  name = "example_configuration"
  fgt_vm_bundle {
    cpu_size            = 8           # 1 ~ 96
    service_pkg         = "FC"        # "FC", "UTP", "ENT", "ATP"
    vdom_num            = 10          # 0 ~ 500
    fortiguard_services = []          # "FGTAVDB", "FGTFAIS", "FGTISSS", "FGTDLDB", "FGTFGSA", "FGTFCSS"
    cloud_services      = []          # "FGTFAMS", "FGTSWNM", "FGTSOCA", "FGTFAZC", "FGTSWOS", "FGTFSPA"
    # support_service = "FGTFCELU" # "NONE", "FGTFCELU"
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

### EXAMPLE: Retrieve STOPPED entitlements

By using the following code, when you use `terraform apply`, terraform will retrieve `count_num` STOPPED entitlements whose config id is `config_id` and description is empty. Terraform will change the status of those entitlements from `STOPPED` to `ACTIVE` and change their description to `task_name`.

When you use `terraform destroy`, the retrieved entitlements will refresh their token, change their status back to `STOPPED` and change their description to empty.


```hcl
resource "fortiflexvm_retrieve_vm_group" "task1" {
  task_name = "UNIQUE_TASK_NAME" # Unique task name
  config_id = 1234               # Your config ID
  count_num = 3
}
output "task1_tokens" {
  value = { for key, vm in fortiflexvm_retrieve_vm_group.task1.entitlements : vm.serial_number => vm.token }
}
```