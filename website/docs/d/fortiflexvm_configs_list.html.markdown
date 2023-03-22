---
subcategory: "Configs"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_configs_list"
description: |-
  Get list of Flex VM Configurations for a Program.
---

# Data Source: fortiflexvm_configs_list
Get list of Flex VM Configurations for a Program.

## Example Usage

```hcl
data "fortiflexvm_configs_list" "example" {
    program_serial_number = "ELAVMS000000XXXX"
}

output "my_configs_list"{
    value = data.fortiflexvm_configs_list.example
}
```

## Argument Reference

The following argument is required:

* `program_serial_number` - (Required/String) The unique serial number of the Flex VM Program.

## Attribute Reference

The following attributes are exported:

* `configs` - (List of Object) The list of Flex VM Configurations for a Program. The structure of [`configs` block](#nestedatt--configs) is documented below.
* `id` - (String) An ID for the resource.
* `program_serial_number` - (String) The unique serial number of this Flex VM Program.


<a id="nestedatt--configs"></a>
The `configs` block contains:

* `faz_vm` - (List of Object) List of FortiAnalyzer Virtual Machines this Flex VM Project contains. The structure of [`configs.faz_vm` block](#nestedobjatt--configs--faz_vm) is documented below.
* `fgt_vm_bundle` - (List of Object) List of FortiGate Virtual Machines (Service Bundle) this Flex VM Project contains. The structure of [`configs.fgt_vm_bundle` block](#nestedobjatt--configs--fgt_vm_bundle) is documented below.
* `fgt_vm_lcs` - (List of Object) List of FortiGate Virtual Machines (A La Carte Services) this Flex VM Project contains. The structure of [`configs.fgt_vm_lcs` block](#nestedobjatt--configs--fgt_vm_lcs) is documented below.
* `fmg_vm` - (List of Object) List of FortiManager Virtual Machines this Flex VM Project contains. The structure of [`configs.fmg_vm` block](#nestedobjatt--configs--fmg_vm) is documented below.
* `fpc_vm` - (List of Object) List of FortiPortal Virtual Machines this Flex VM Project contains. The structure of [`configs.fpc_vm` block](#nestedobjatt--configs--fpc_vm) is documented below.
* `fwb_vm` - (List of Object) List of FortiWeb Virtual Machines this Flex VM Project contains. The structure of [`configs.fwb_vm` block](#nestedobjatt--configs--fwb_vm) is documented below.
* `id` = (Number) The unqiue number of the Flex VM Configuration.
* `name` - (String) Flex VM Configuration name.
* `product_type` - (String) Flex VM Configuration type. Possible values: 
  * `FAZ_VM`: FortiAnalyzer Virtual Machine
  * `FGT_VM_Bundle`: FortiGate Virtual Machine - Service Bundle
  * `FGT_VM_LCS`: FortiGate Virtual Machine - A La Carte Services
  * `FMG_VM`: FortiManager Virtual Machine
  * `FPC_VM`: FortiPortal Virtual Machine
  * `FWB_VM`: FortiWeb Virtual Machine - Service Bundle
* `program_serial_number` - (String) The unique serial number of the Flex VM Program this Flex VM Configuration belongs to.
* `status` - (String) The status of this Flex VM Configuration. `ACTIVATE` or `DISABLED`.


<a id="nestedobjatt--configs--faz_vm"></a>
The `configs.faz_vm` block contains:

* `adom_num` - (Number) Number of ADOMs. A number between 0 and 1200 (inclusive).
* `daily_storage` - (Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).
* `support_service` - (String) Support Service. `"FAZFC247"` (FortiCare Premium).

<a id="nestedobjatt--configs--fgt_vm_bundle"></a>
The `configs.fgt_vm_bundle` block contains:

* `cpu_size` - (String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`,  `"32"` or `"2147483647"` (unlimited).
* `service_pkg` - (String) The value of this attribute is one of `"FC"` (FortiCare), `"UTM"`, `"ENT"` (Enterprise) or `"ATP"`.
* `vdom_num` - (Number) Number of VDOMs. A number between 0 and 500 (inclusive).


<a id="nestedobjatt--configs--fgt_vm_lcs"></a>
The `configs.fgt_vm_lcs` block contains:

* `cloud_services` - (List of String) The cloud services this FortiGate Virtual Machine supports. The combination of:
  * `"FAMS"` (FortiGate Cloud)
  * `"SWNM"` (SD-WAN Cloud)
  * `"FMGC"` (FortiManager Cloud)
  * `"AFAC"` (FortiAnalyzer Cloud with SOCaaS)
* `cpu_size` - (String) The number of CPUs. A number between 1 and 96 (inclusive).
* `fortiguard_services` - (List of String) The FortiGuard services this FortiGate Virtual Machine supports. The combination of:
  * `"IPS"` (Intrusion Prevention)
  * `"AVDB"` (Advanced Malware)
  * `"FURL"` (Web & Video Filtering)
  * `"IOTH"` (IOT Detection)
  * `"FGSA"` (Security Rating)
  * `"ISSS"` (Industrial Security)
* `support_service` - (String) `"FC247"` (FortiCare 24x7) or `"ASET"` (FortiCare Elite).
* `vdom_num` - (Number) Number of VDOMs. A number between 0 and 500 (inclusive).


<a id="nestedobjatt--configs--fmg_vm"></a>
The `configs.fmg_vm` block contains:

* `adom_num` - (Number) Number of ADOMs. A number between 1 and 100000 (inclusive).
* `managed_dev` - (Number) Number of managed devices. A number between 1 and 100000 (inclusive).

<a id="nestedobjatt--configs--fpc_vm"></a>
The `configs.fpc_vm` block contains:

* `managed_dev` - (Number) Number of managed devices. A number between 0 and 100000 (inclusive).


<a id="nestedobjatt--configs--fwb_vm"></a>
The `configs.fwb_vm` block contains:

* `cpu_size` - (String) Number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"` or `"16"`.
* `service_pkg` - (String) Service Package. `"FWBSTD"` (Standard) or `"FWBADV"` (Advanced).


