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
    service_pkg = "ATP" # "FC", "UTP", "ENT", "ATP"
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
    device_model = "FGT60F"   # "FGT40F", "FGT60F", "FGT70F", "FGT80F", "FG100F", "FGT60E", "FGT61F", "FG100E", "FG101F", "FG200E", 
                              # "FG200F", "FG201F", "FG4H0F", "FG6H0F", "FWF40F", "FWF60F", "FGR60F", "FR70FB", "FGT81F", "FG101E",
                              # "FG4H1F", "FG1K0F", "FG180F", "F2K60F", "FG3K0F", "FG3K1F", "FG3K2F"
    service_pkg = "FGHWFC247" # "FGHWFC247", "FGHWFCEL", "FGHWATP", "FGHWUTP", "FGHWENT"
    addons = []               # List of string, "FGHWFCELU", "FGHWFAMS", "FGHWFAIS", "FGHWSWNM", "FGHWDLDB", "FGHWFAZC", "FGHWSOCA",
                              # "FGHWMGAS", "FGHWSPAL", "FGHWFCSS"
  }
}

resource "fortiflexvm_config" "example9"{
  product_type = "FC_EMS_OP"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FC_EMS_OP_example"
  fc_ems_op {
    ztna_num = 225                # Value should be divisible by 25. Number between 0 and 25000 (inclusive)
    epp_ztna_num = 125            # Value should be divisible by 25. Number between 0 and 25000 (inclusive)
    chromebook = 100              # Value should be divisible by 25. Number between 0 and 25000 (inclusive) 
    support_service = "FCTFC247"  # "FCTFC247"
    addons = []                   # [] or ["BPS"]
  }
  status = "ACTIVE"
}

resource "fortiflexvm_config" "example10"{
  product_type = "FWBC_PUBLIC"
  program_serial_number = "ELAVMS00000XXXXX"
  name = "FWBC_PUBLIC_example"
  fwbc_public {
    average_throughput = 150   # 10, 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600,
                               # 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000,
                               # 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000
    web_applications = 100     # Number between 0 and 2000 (inclusive) 
  }
  status = "ACTIVE"
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional/Number) Account ID. Once the fortiflexvm_config is created, you can't change the account ID of this configuration by changing `account_id`.
* `product_type` (Required/String) Product type, must be one of the following options:
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
* `program_serial_number` - (Required/String) The serial number of your FortiFlex Program. This serial number should start with `"ELAVMR"`.
* `name` - (Required unless you only update the status/String) The name of your configuration.
* `status` - (Optional/String) Configuration status. If you don't specify, the configuration status keeps unchanged. The default status is `ACTIVE` once you create a configuration. It must be one of the following options:
	* `ACTIVE`: Enable a configuration
	* `DISABLED`: Disable a configuration
* `fad_vm` - (Block List) You must fill in this block if your `product_type` is `"FAD_VM"`. The structure of [`fad_vm` block](#nestedblock--fad_vm) is documented below.
* `faz_vm` - (Block List) You must fill in this block if your `product_type` is `"FAZ_VM"`. The structure of [`faz_vm` block](#nestedblock--faz_vm) is documented below.
* `fc_ems_op` - (Block List) You must fill in this block if your `product_type` is `"FC_EMS_OP"`. The structure of [`fc_ems_op` block](#nestedblock--fc_ems_op) is documented below.
* `fgt_hw` - (Block List) You must fill in this block if your `product_type` is `"FGT_HW"`. The structure of [`fgt_hw` block](#nestedblock--fgt_hw) is documented below.
* `fgt_vm_bundle` - (Block List) You must fill in this block if your `product_type` is `"FGT_VM_Bundle"`. The structure of [`fgt_vm_bundle` block](#nestedblock--fgt_vm_bundle) is documented below.
* `fgt_vm_lcs` - (Block List) You must fill in this block if your `product_type` is `"FGT_VM_LCS"`. The structure of [`fgt_vm_lcs` block](#nestedblock--fgt_vm_lcs) is documented below.
* `fmg_vm` - (Block List) You must fill in this block if your `product_type` is `"FMG_VM"`. The structure of [`fmg_vm` block](#nestedblock--fmg_vm) is documented below.
* `fpc_vm` - (Block List) You must fill in this block if your `product_type` is `"FPC_VM"`. The structure of [`fpc_vm` block](#nestedblock--fpc_vm) is documented below.
* `fwb_vm` - (Block List) You must fill in this block if your `product_type` is `"FWB_VM"`. The structure of [`fwb_vm` block](#nestedblock--fwb_vm) is documented below.
* `fwbc_private` - (Block List) You must fill in this block if your `product_type` is `"FWBC_PRIVATE"`. The structure of [`fwbc_private` block](#nestedblock--fwbc_private) is documented below.
* `fwbc_public` - (Block List) You must fill in this block if your `product_type` is `"FWBC_PUBLIC"`. The structure of [`fwbc_public` block](#nestedblock--fwbc_public) is documented below.


<a id="nestedblock--fad_vm"></a>
The `fad_vm` block contains:

* `cpu_size` - (Required if `product_type = "FAD_VM"`/String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`, `"32"`.
* `service_pkg` - (Required if `product_type = "FAD_VM"`/String) Options: `"FDVSTD"` (Standard), `"FDVADV"` (Advanced) or `"FDVFC247"` (FortiCare Premium).


<a id="nestedblock--faz_vm"></a>
The `faz_vm` block contains:

* `adom_num` - (Required if `product_type = "FAZ_VM"`/Number) Number of ADOMs. A number between 0 and 1200 (inclusive).
* `daily_storage` - (Required if `product_type = "FAZ_VM"`/Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).
* `support_service` - (Required if `product_type = "FAZ_VM"`/String) Support Service. Option: `"FAZFC247"` (FortiCare Premium).


<a id="nestedblock--fc_ems_op"></a>
The `fc_ems_op` block contains:

* `ztna_num` - (Required if `product_type = "FC_EMS_OP"`/Number) ZTNA/VPN (number of endpoints). Value should be divisible by 25. Number between 0 and 25000 (inclusive).
* `epp_ztna_num` - (Required if `product_type = "FC_EMS_OP"`/Number) EPP/ATP + ZTNA/VPN (number of endpoints). Value should be divisible by 25. Number between 0 and 25000 (inclusive).
* `chromebook` - (Required if `product_type = "FC_EMS_OP"`/Number) Chromebook (number of endpoints). Value should be divisible by 25. Number between 0 and 25000 (inclusive).
* `support_service` - (Required if `product_type = "FC_EMS_OP"`/String) Option: `"FCTFC247"` (FortiCare Premium).
* `addons` - (Optional/List of String) The default value is an empty list. Options: `"BPS"` (FortiCare Best Practice).


<a id="nestedblock--fgt_hw"></a>
The `fgt_hw` block contains:

* `device_model` - (Required if `product_type = "FGT_HW"`/String) Device Model. Options: 
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
* `service_pkg` - (Required if `product_type = "FGT_HW"`/String) Options: `"FGHWFC247"` (FortiCare Premium), `"FGHWFCEL"` (FortiCare Elite), `"FGHWATP"` (ATP), `"FGHWUTP"` (UTP) or `"FGHWENT"` (Enterprise).
* `addons` - (Optional/List of String) The default value is an empty list. Options: 
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

<a id="nestedblock--fgt_vm_bundle"></a>
The `fgt_vm_bundle` block contains:

* `cpu_size` - (Required if `product_type = "FGT_VM_Bundle"`/String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`,  `"32"` or `"2147483647"` (unlimited). 
* `service_pkg` - (Required if `product_type = "FGT_VM_Bundle"`/String) The value of this attribute is one of `"FC"` (FortiCare), `"UTP"` (UTP), `"ENT"` (Enterprise) or `"ATP"` (ATP).
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
  * `"FURLDNS"`: Web, DNS & Video Filtering
  * `"FGSA"`: Security Rating
  * `"DLDB"`: DLP
  * `"FAIS"`: AI-Based InLine Sandbox
* `support_service` - (Required if `product_type = "FGT_VM_LCS"`/String) Options: `"FC247"` (FortiCare 24x7) or `"ASET"` (FortiCare Elite).
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
* `service_pkg` - (Required if `product_type = "FWB_VM"`/String) Service Package. Options: `"FWBSTD"` (Standard) or `"FWBADV"` (Advanced).


<a id="nestedblock--fwbc_private"></a>
The `fwbc_private` block contains:

* `average_throughput` - (Required if `product_type = "FWBC_PRIVATE"`/Number) Average Throughput (Mbps). Options: 10, 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600, 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000.
* `web_applications` - (Required if `product_type = "FWBC_PRIVATE"`/Number) Number between 0 and 2000 (inclusive).


<a id="nestedblock--fwbc_public"></a>
The `fwbc_public` block contains:

* `average_throughput` - (Required if `product_type = "FWBC_PUBLIC"`/Number) Average Throughput (Mbps). Options: 10, 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600, 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000.
* `web_applications` - (Required if `product_type = "FWBC_PUBLIC"`/Number) Number between 0 and 2000 (inclusive) 


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