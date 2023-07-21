---
subcategory: "Configs"
layout: "fortiflexvm"
page_title: "FortiFlexVM: fortiflexvm_config"
description: |-
  Create a new Configuration under a FortiFlex Program.
---

# fortiflexvm_config

Create a new configuration under a FortiFlex Program.

~> By using this resource, you can change the data in the FortiFlex Platform immediately. Yet it may take several hours for the VMs to update their licenses automatically. To update the licenses in the VMs immediately, please reboot your VMs.

## Example Usage

-> You need to specify what type of product you want to create in `product_type` and then fill in the correspond block.

```hcl
resource "fortiflexvm_config" "example1"{
  product_type = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FGT_VM_Bundle_example"
  fgt_vm_bundle {
    cpu_size =  "2"     # "1", "2", "4", "8", "16", "32", "2147483647"
    service_pkg = "ATP" # "FC", "UTM", "ENT", "ATP"
    vdom_num = 11       # 0 ~ 500
  }
}


resource "fortiflexvm_config" "example2"{
  product_type = "FWB_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FWB_VM_example"
  fwb_vm {
    cpu_size =  "2"           # "1", "2", "4", "8", "16"
    service_pkg = "FWBSTD"    # "FWBSTD", "FWBADV"
  }
}


resource "fortiflexvm_config" "example3"{
  product_type = "FMG_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FMG_VM_example"
  fmg_vm {
    managed_dev =  1     # 1 ~ 100000
    adom_num = 1         # 1 ~ 100000
  }
}


resource "fortiflexvm_config" "example4"{
  product_type = "FGT_VM_LCS"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FGT_VM_LCS_example"
  fgt_vm_lcs {
    cpu_size = 3                                # 1 ~ 96
    vdom_num = 3                                # 1 ~ 500
    support_service = "FC247"                   # "FC247", "ASET"
    cloud_services = []                         # "FAMS", "SWNM", "AFAC", "FAZC"
    fortiguard_services = ["IPS", "AVDB"]       # "IPS", "AVDB", "FGSA", "DLDB", "FAIS", "FURLDNS"
  }
}


resource "fortiflexvm_config" "example5"{
  product_type = "FAZ_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FAZ_VM_example"
  faz_vm {
    daily_storage = 11                  # 5 ~ 8300
    support_service = "FAZFC247"        # "FAZFC247"
    adom_num =  0                       # 0 ~ 1200
  }
  status =  "DISABLED"
}


resource "fortiflexvm_config" "example6"{
  product_type = "FPC_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FPC_VM_example"
  fpc_vm {
    managed_dev = 1       # 0 ~ 100000
  }
}


resource "fortiflexvm_config" "example7"{
  product_type = "FAD_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FAD_VM_example"
  fad_vm {
    cpu_size = "1"                # "1", "2", "4", "8", "16", "32"
    service_pkg = "FDVSTD"        # "FDVSTD", "FDVADV", "FDVFC247"
  }
}


resource "fortiflexvm_config" "example8"{
  product_type = "FGT_HW"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FGT_HW_example"
  fgt_hw {
    device_model = "FGT70F"   # "FGT40F", "FGT60F", "FGT70F", "FGT80F", "FG100F", "FGT60E", "FGT61F", "FG100E", "FG101F", "FG200E", "FG200F", "FG201F", "FG4H0F", "FG6H0F"
    service_pkg = "FGHWFC247" # "FGHWFC247", "FGHWFCEL", "FGHWATP", "FGHWUTP", "FGHWENT"
    addons = "NONE"           # only support "NONE" now. Will support "FGHWFCELU" in the future.
  }
}
```

## Argument Reference

The following arguments are supported:

* `product_type` (Required/String) Product type, must be one of the following options:
  * `FAD_VM`: FortiADC Virtual Machine
  * `FAZ_VM`: FortiAnalyzer Virtual Machine
  * `FGT_HW`: FortiGate Hardware
  * `FGT_VM_Bundle`: FortiGate Virtual Machine - Service Bundle
  * `FGT_VM_LCS`: FortiGate Virtual Machine - A La Carte Services
  * `FMG_VM`: FortiManager Virtual Machine
  * `FPC_VM`: FortiPortal Virtual Machine
  * `FWB_VM`: FortiWeb Virtual Machine - Service Bundle
* `program_serial_number` - (Required/String) The serial number of your FortiFlex Program. This serial number should start with `"ELAVMR"`.
* `name` - (Required unless you only update the status/String) The name of your configuration.
* `status` - (Optional/String) Configuration status. If you don't specify, the configuration status keeps unchanged. The default status is `ACTIVE` once you create a configuration. It must be one of the following options:
	* `ACTIVE`: Enable a configuration
	* `DISABLED`: Disable a configuration
* `fad_vm` - (Block List) You must fill in this block if your `product_type` is `"FAD_VM"`. The structure of [`fad_vm` block](#nestedobjatt--fad_vm) is documented below.
* `faz_vm` - (Block List) You must fill in this block if your `product_type` is `"FAZ_VM"`. The structure of [`faz_vm` block](#nestedobjatt--faz_vm) is documented below.
* `fgt_hw` - (Block List) You must fill in this block if your `product_type` is `"FGT_HW"`. The structure of [`fgt_hw` block](#nestedobjatt--fgt_hw) is documented below.
* `fgt_vm_bundle` - (Block List) You must fill in this block if your `product_type` is `"FGT_VM_Bundle"`. The structure of [`fgt_vm_bundle` block](#nestedatt--fgt_vm_bundle) is documented below.
* `fgt_vm_lcs` - (Block List) You must fill in this block if your `product_type` is `"FGT_VM_LCS"`. The structure of [`fgt_vm_lcs` block](#nestedatt--fgt_vm_lcs) is documented below.
* `fmg_vm` - (Block List) You must fill in this block if your `product_type` is `"FMG_VM"`. The structure of [`fmg_vm` block](#nestedatt--fmg_vm) is documented below.
* `fpc_vm` - (Block List) You must fill in this block if your `product_type` is `"FPC_VM"`. The structure of [`fpc_vm` block](#nestedobjatt--fpc_vm) is documented below.
* `fwb_vm` - (Block List) You must fill in this block if your `product_type` is `"FWB_VM"`. The structure of [`fwb_vm` block](#nestedatt--fwb_vm) is documented below.


<a id="nestedobjatt--fad_vm"></a>
The `fad_vm` block contains:

* `cpu_size` - (Required if `product_type = "FAD_VM"`/String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`, `"32"`.
* `service_pkg` - (Required if `product_type = "FAD_VM"`/String) The value of this attribute is one of `"FDVSTD"` (Standard), `"FDVADV"` (Advanced) or `"FDVFC247"` (FortiCare Premium).


<a id="nestedblock--faz_vm"></a>
The `faz_vm` block contains:

* `adom_num` - (Required if `product_type = "FAZ_VM"`/Number) Number of ADOMs. A number between 0 and 1200 (inclusive).
* `daily_storage` - (Required if `product_type = "FAZ_VM"`/Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).
* `support_service` - (Required if `product_type = "FAZ_VM"`/String) Support Service. Currently, the only available option is `"FAZFC247"` (FortiCare Premium). The default value is `"FAZFC247"`.


<a id="nestedobjatt--fgt_hw"></a>
The `fgt_hw` block contains:

* `device_model` - (Required if `product_type = "FGT_HW"`/String) Device Model. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`, `"32"` or `"2147483647"` (unlimited).
* `service_pkg` - (Required if `product_type = "FGT_HW"`/String) The value of this attribute is one of `"FGHWFC247"` (FortiCare Premium), `"FGHWFCEL"` (FortiCare Elite), `"FGHWATP"` (ATP), `"FGHWUTP"` (UTP) or `"FGHWENT"` (Enterprise).
* `addons` - (Required if `product_type = "FGT_HW"`/String) Only support `"NONE"` now. Will support `"FGHWFCELU"` (FortiCare Elite Upgrade) in the future.


<a id="nestedblock--fgt_vm_bundle"></a>
The `fgt_vm_bundle` block contains:

* `cpu_size` - (Required if `product_type = "FGT_VM_Bundle"`/String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`,  `"32"` or `"2147483647"` (unlimited). 
* `service_pkg` - (Required if `product_type = "FGT_VM_Bundle"`/String) The value of this attribute is one of `"FC"` (FortiCare), `"UTM"`, `"ENT"` (Enterprise) or `"ATP"`.
* `vdom_num` - (Optional/Number) Number of VDOMs. A number between 0 and 500 (inclusive). The default number is 0.


<a id="nestedblock--fgt_vm_lcs"></a>
The `fgt_vm_lcs` block contains:

* `cloud_services` - (Optional/List of String) The cloud services this FortiGate Virtual Machine supports. The default value is an empty list. It should be a combination of:
  * `"FAMS"`: FortiGate Cloud
  * `"SWNM"`: SD-WAN Underlay
  * `"AFAC"`: FortiAnalyzer Cloud with SOCaaS
  * `"FAZC"`: FortiAnalyzer Cloud
* `cpu_size` - (Required if `product_type = "FGT_VM_LCS"`/String) The number of CPUs. A number between 1 and 96 (inclusive).
* `fortiguard_services` - (Optional/List of String) The fortiguard services this FortiGate Virtual Machine supports. The default value is an empty list. It should be a combination of:
  * `"IPS"`: Intrusion Prevention
  * `"AVDB"`: Advanced Malware
  * `"FGSA"`: Security Rating
  * `"DLDB"`: DLP
  * `"FAIS"`: AI-Based InLine Sandbox
  * `"FURLDNS"`: Web, DNS & Video Filtering
* `support_service` - (Required if `product_type = "FGT_VM_LCS"`/String) `"FC247"` (FortiCare 24x7) or `"ASET"` (FortiCare Elite).
* `vdom_num` - (Optional/Number) Number of VDOMs. A number between 1 and 500 (inclusive). The default number is 1.


<a id="nestedblock--fmg_vm"></a>
The `fmg_vm` block contains:

* `adom_num` - (Optional/Number) Number of ADOMs. A number between 1 and 100000 (inclusive). The default value is 1.
* `managed_dev` - (Optional/Number) Number of managed devices. A number between 1 and 100000 (inclusive). The default value is 1.


<a id="nestedblock--fpc_vm"></a>
The `fpc_vm` block contains:

* `managed_dev` - (Required if `product_type = "FPC_VM"`/Number) Number of managed devices. A number between 0 and 100000 (inclusive).


<a id="nestedblock--fwb_vm"></a>
The `fwb_vm` block contains:

* `cpu_size` - (Required if `product_type = "FWB_VM"`/String) Number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"` or `"16"`.
* `service_pkg` - (Required if `product_type = "FWB_VM"`/String) Service Package. Valid values: `"FWBSTD"` (Standard) or `"FWBADV"` (Advanced).


## Attribute Reference

The following attribute is exported:

* `id` - (String) An ID for the resource.

## Import

FortiFlex Configuration can be imported by using the following steps:

First, specify the `program_serial_number` when you configure the provider.
```
provider "fortiflexvm" {
  username = "ABCDEFG"
  password = "HIJKLMN"
  import_options= toset(["program_serial_number=ELAVMS000000XXXX"])
}
```

Then, use the following command to import the configuration.
```
terraform import fortiflexvm_config.labelname {{id}}
```