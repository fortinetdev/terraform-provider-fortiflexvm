---
subcategory: "Configs"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_configs_list"
description: |-
  Get list of configurations for a FortiFlex Program.
---

# Data Source: fortiflexvm_configs_list
Get list of configurations for a FortiFlex Program.

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

* `account_id` - (Optional/Number) Account ID.
* `program_serial_number` - (Required/String) The unique serial number of the Program.

## Attribute Reference

The following attributes are exported:

* `configs` - (List of Object) The list of Configurations for a Program. The structure of [`configs` block](#nestedatt--configs) is documented below.
* `id` - (String) An ID for the resource. Its value is variable `program_serial_number`.


<a id="nestedatt--configs"></a>
The `configs` block contains:

* `fad_vm` - (List of Object) FortiADC Virtual Machine. The structure of [`configs.fad_vm` block](#nestedobjatt--configs--fad_vm) is documented below.
* `faz_vm` - (List of Object) FortiAnalyzer Virtual Machine. The structure of [`configs.faz_vm` block](#nestedobjatt--configs--faz_vm) is documented below.
* `fc_ems_op` - (List of Object) FortiClient EMS On-Prem. The structure of [`configs.fc_ems_op` block](#nestedobjatt--configs--fc_ems_op) is documented below.
* `fgt_hw` - (List of Object) FortiGate Hardware. The structure of [`configs.fgt_hw` block](#nestedobjatt--configs--fgt_hw) is documented below.
* `fgt_vm_bundle` - (List of Object) FortiGate Virtual Machine. The structure of [`configs.fgt_vm_bundle` block](#nestedobjatt--configs--fgt_vm_bundle) is documented below.
* `fgt_vm_lcs` - (List of Object) FortiGate Virtual Machine (A La Carte Services). The structure of [`configs.fgt_vm_lcs` block](#nestedobjatt--configs--fgt_vm_lcs) is documented below.
* `fmg_vm` - (List of Object) FortiManager Virtual Machine. The structure of [`configs.fmg_vm` block](#nestedobjatt--configs--fmg_vm) is documented below.
* `fpc_vm` - (List of Object) FortiPortal Virtual Machine. The structure of [`configs.fpc_vm` block](#nestedobjatt--configs--fpc_vm) is documented below.
* `fwb_vm` - (List of Object) FortiWeb Virtual Machine. The structure of [`configs.fwb_vm` block](#nestedobjatt--configs--fwb_vm) is documented below.
* `fwbc_private` - (List of Object) FortiWeb Cloud - Private. The structure of [`configs.fwbc_private` block](#nestedobjatt--configs--fwbc_private) is documented below.
* `fwbc_public` - (List of Object) FortiWeb Cloud - Public. The structure of [`configs.fwbc_public` block](#nestedobjatt--configs--fwbc_public) is documented below.
* `account_id` - (Optional/Number) Account ID.
* `id` - (Number) The unqiue number of the configuration.
* `name` - (String) Configuration name.
* `product_type` - (String) Configuration type. Possible values: 
  * `FAD_VM`: FortiADC Virtual Machine
  * `FAZ_VM`: FortiAnalyzer Virtual Machine
  * `FC_EMS_OP`: FortiClient EMS On-Prem
  * `FGT_HW`: FortiGate Hardware
  * `FGT_VM_Bundle`: FortiGate Virtual Machine - Service Bundle
  * `FGT_VM_LCS`: FortiGate Virtual Machine - A La Carte Services
  * `FMG_VM`: FortiManager Virtual Machine
  * `FPC_VM`: FortiPortal Virtual Machine
  * `FWB_VM`: FortiWeb Virtual Machine - Service Bundle
  * `FWBC_PRIVATE`: FortiWeb Cloud - Private
  * `FWBC_PUBLIC`: FortiWeb Cloud - Public
* `program_serial_number` - (String) The unique serial number of the FortiFlex Program this configuration belongs to.
* `status` - (String) The status of this configuration. `ACTIVATE` or `DISABLED`.


<a id="nestedobjatt--configs--fad_vm"></a>
The `configs.fad_vm` block contains:

* `cpu_size` - (String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`, `"32"`.
* `service_pkg` - (String) The value of this attribute is one of `"FDVSTD"` (Standard), `"FDVADV"` (Advanced) or `"FDVFC247"` (FortiCare Premium).


<a id="nestedobjatt--configs--faz_vm"></a>
The `configs.faz_vm` block contains:

* `adom_num` - (Number) Number of ADOMs. A number between 0 and 1200 (inclusive).
* `daily_storage` - (Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).
* `support_service` - (String) Support Service. Possible value: `"FAZFC247"` (FortiCare Premium).


<a id="nestedobjatt--configs--fc_ems_op"></a>
The `configs.fc_ems_op` block contains:

* `ztna_num` - (Number) ZTNA/VPN (number of endpoints). Value should be divisible by 25. Number between 0 and 25000 (inclusive).
* `epp_ztna_num` - (Number) EPP/ATP + ZTNA/VPN (number of endpoints). Value should be divisible by 25. Number between 0 and 25000 (inclusive).
* `chromebook` - (Number) Chromebook (number of endpoints). Value should be divisible by 25. Number between 0 and 25000 (inclusive).
* `support_service` - (String) Possible value: `"FCTFC247"` (FortiCare Premium).
* `addons` - (List of String) Possible value: `"BPS"` (FortiCare Best Practice).


<a id="nestedobjatt--configs--fgt_hw"></a>
The `configs.fgt_hw` block contains:

* `device_model` - (String) Device Model. Possible values: 
  * `"FGT40F"`: FortiGate-40F
  * `"FGT60F"`: FortiGate-60F
  * `"FGT70F"`: FortiGate-70F
  * `"FGT80F"`: FortiGate-80F
  * `"FG100F"`: FortiGate-100F
  * `"FGT60E"`: FortiGate-60E
  * `"FGT61F"`: FortiGate-61F
  * `"FG100E"`: FortiGate-100E
  * `"FG101F"`: FortiGate-101F
  * `"FG200E"`: FortiGate-200E
  * `"FG200F"`: FortiGate-200F
  * `"FG201F"`: FortiGate-201F
  * `"FG4H0F"`: FortiGate-400F
  * `"FG6H0F"`: FortiGate-600F
  * `"FWF40F"`: FortiWifi-40F
  * `"FWF60F"`: FortiWifi-60F
  * `"FGR60F"`: FortiGateRugged-60F
  * `"FR70FB"`: FortiGateRugged-70F
  * `"FGT81F"`: FortiGate-81F
  * `"FG101E"`: FortiGate-101E
  * `"FG4H1F"`: FortiGate-401F
  * `"FG1K0F"`: FortiGate-1000F
  * `"FG180F"`: FortiGate-1800F
  * `"F2K60F"`: FortiGate-2600F
  * `"FG3K0F"`: FortiGate-3000F
  * `"FG3K1F"`: FortiGate-3001F
  * `"FG3K2F"`: FortiGate-3200F
* `service_pkg` - (String) Possible values: `"FGHWFC247"` (FortiCare Premium), `"FGHWFCEL"` (FortiCare Elite), `"FGHWATP"` (ATP), `"FGHWUTP"` (UTP) or `"FGHWENT"` (Enterprise).
* `addons` - (List of String) Possible values:
  * `"FGHWFCELU"`: FortiCare Elite Upgrade
  * `"FGHWFAMS"`: FortiGate Cloud Management
  * `"FGHWFAIS"`: AI-Based In-line Sandbox
  * `"FGHWSWNM"`: SD-WAN Underlay
  * `"FGHWDLDB"`: FortiGuard DLP
  * `"FGHWFAZC"`: FortiAnalyzer Cloud
  * `"FGHWSOCA"`: SOCaaS
  * `"FGHWMGAS"`: Managed FortiGate
  * `"FGHWSPAL"`: SD-WAN Connector for FortiSASE
  * `"FGHWFCSS"`: FortiConverter Service


<a id="nestedobjatt--configs--fgt_vm_bundle"></a>
The `configs.fgt_vm_bundle` block contains:

* `cpu_size` - (String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`, `"32"` or `"2147483647"` (unlimited).
* `service_pkg` - (String) The value of this attribute is one of `"FC"` (FortiCare), `"UTP"` (UTP), `"ENT"` (Enterprise) or `"ATP"` (ATP).
* `vdom_num` - (Number) Number of VDOMs. A number between 0 and 500 (inclusive).


<a id="nestedobjatt--configs--fgt_vm_lcs"></a>
The `configs.fgt_vm_lcs` block contains:

* `cloud_services` - (List of String) The cloud services this FortiGate Virtual Machine supports. The combination of:
  * `"FAMS"`: FortiGate Cloud
  * `"SWNM"`: SD-WAN Underlay
  * `"AFAC"`: FortiAnalyzer Cloud with SOCaaS
  * `"FAZC"`: FortiAnalyzer Cloud
* `cpu_size` - (String) The number of CPUs. A number between 1 and 96 (inclusive).
* `fortiguard_services` - (List of String) The FortiGuard services this FortiGate Virtual Machine supports. The combination of:
  * `"IPS"`: Intrusion Prevention
  * `"AVDB"`: Advanced Malware
  * `"FURLDNS"`: Web, DNS & Video Filtering
  * `"FGSA"`: Security Rating
  * `"DLDB"`: DLP
  * `"FAIS"`: AI-Based InLine Sandbox
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


<a id="nestedobjatt--configs--fwbc_private"></a>
The `configs.fwbc_private` block contains:

* `average_throughput` - (Number) Average Throughput (Mbps). Possible values: 10, 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600, 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000.
* `web_applications` - (Number) Number between 0 and 2000 (inclusive).


<a id="nestedobjatt--configs--fwbc_public"></a>
The `configs.fwbc_public` block contains:

* `average_throughput` - (Number) Average Throughput (Mbps). Possible values: 10, 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600, 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000.
* `web_applications` - (Number) Number between 0 and 2000 (inclusive).
