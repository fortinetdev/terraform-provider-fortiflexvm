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

Create or import configuration
```hcl
// Create a new configuration
resource "fortiflexvm_config" "create_example" {
  product_type          = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FGT_VM_Bundle_example"
  fgt_vm_bundle {
    cpu_size            = 2           # 1 ~ 96
    service_pkg         = "ATP"       # "FC", "UTP", "ENT", "ATP"
    vdom_num            = 10          # 0 ~ 500
    fortiguard_services = ["FGTAVDB"] # "FGTAVDB", "FGTFAIS", "FGTISSS", "FGTDLDB", "FGTFGSA", "FGTFCSS"
    cloud_services      = []          # "FGTFAMS", "FGTSWNM", "FGTSOCA", "FGTFAZC", "FGTSWOS", "FGTFSPA"
    # support_service = "FGTFCELU"    # "NONE", "FGTFCELU"
  }
}

// Import an existing configuration
resource "fortiflexvm_config" "import_example" {
  config_id             = 12345
  product_type          = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS00000XXXXX"
}

// Import an existing configuration and update it
// You need to run "terraform apply" twice.
// The first one is for import and the second one is for update
resource "fortiflexvm_config" "import_and_update" {
  status                = "ACTIVE"
  config_id             = 12345
  product_type          = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FGT_VM_Bundle_example"
  fgt_vm_bundle {
    cpu_size            = 2           # 1 ~ 96
    service_pkg         = "ATP"       # "FC", "UTP", "ENT", "ATP"
    vdom_num            = 10          # 0 ~ 500
    fortiguard_services = ["FGTAVDB"] # "FGTAVDB", "FGTFAIS", "FGTISSS", "FGTDLDB", "FGTFGSA", "FGTFCSS"
    cloud_services      = []          # "FGTFAMS", "FGTSWNM", "FGTSOCA", "FGTFAZC", "FGTSWOS", "FGTFSPA"
    # support_service = "FGTFCELU"    # "NONE", "FGTFCELU"
  }
}
```

Examples of creating configurations.
```hcl
# FortiGate Virtual Machine - Service Bundle
resource "fortiflexvm_config" "FGT_VM_Bundle" {
  product_type          = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FGT_VM_Bundle_example"
  fgt_vm_bundle {
    cpu_size            = "2"         # 1 ~ 96
    service_pkg         = "FC"        # "FC", "UTP", "ENT", "ATP"
    vdom_num            = 10          # 0 ~ 500
    fortiguard_services = ["FGTAVDB"] # "FGTAVDB", "FGTFAIS", "FGTISSS", "FGTDLDB", "FGTFGSA"
    cloud_services      = []          # "FGTFAMS", "FGTSWNM", "FGTSOCA", "FGTFAZC", "FGTSWOS", "FGTFSPA"
    support_service     = "NONE"      # "FGTFCELU", "NONE"
  }
}

# FortiManager Virtual Machine
resource "fortiflexvm_config" "FMG_VM" {
  product_type          = "FMG_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FMG_VM_example"
  fmg_vm {
    managed_dev = 1 # 1 ~ 100000
    adom_num    = 0 # 0 ~ 100000
  }
}

# FortiWeb Virtual Machine - Service Bundle
resource "fortiflexvm_config" "FWB_VM" {
  product_type          = "FWB_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FWB_VM_example"
  fwb_vm {
    cpu_size    = "2"      # "1", "2", "4", "8", "16"
    service_pkg = "FWBSTD" # "FWBSTD", "FWBADV", "FWBENT"
  }
}

# FortiGate Virtual Machine - A La Carte Services
resource "fortiflexvm_config" "FGT_VM_LCS" {
  product_type          = "FGT_VM_LCS"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FGT_VM_LCS_example"
  fgt_vm_lcs {
    cpu_size            = 4               # 1 ~ 96
    vdom_num            = 2               # 0 ~ 500
    support_service     = "FC247"         # "FC247", "ASET"
    cloud_services      = []              # "FAMS", "SWNM", "AFAC", "FAZC", "FSPA", "SWOS"
    fortiguard_services = ["IPS", "AVDB"] # "IPS", "AVDB", "FURLDNS", "FGSA", "ISSS", "DLDB", "FAIS"
  }
}

# FortiClient EMS On-Prem
resource "fortiflexvm_config" "FC_EMS_OP" {
  product_type          = "FC_EMS_OP"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FC_EMS_OP_example"
  fc_ems_op {
    ztna_num        = 225        # Number between 0 and 25,000 (inclusive)
    epp_ztna_num    = 125        # Number between 0 and 25,000 (inclusive)
    chromebook      = 100        # Number between 0 and 25,000 (inclusive) 
    support_service = "FCTFC247" # "FCTFC247"
    addons          = ["BPS"]    # [] or ["BPS"]
  }
}

# FortiAnalyzer Virtual Machine
resource "fortiflexvm_config" "FAZ_VM" {
  product_type          = "FAZ_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FAZ_VM_example"
  faz_vm {
    daily_storage   = 20         # 5 ~ 8300
    adom_num        = 5          # 0 ~ 1200
    support_service = "FAZFC247" # "FAZFC247"
    addons          = []         # "FAZISSS", "FAZFGSA", "FAZAISN"
  }
}

# FortiPortal Virtual Machine
resource "fortiflexvm_config" "FPC_VM" {
  product_type          = "FPC_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FPC_VM_example"
  fpc_vm {
    managed_dev = 1 # 0 ~ 100000
  }
}

# FortiADC Virtual Machine
resource "fortiflexvm_config" "FAD_VM" {
  product_type          = "FAD_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FAD_VM_example"
  fad_vm {
    cpu_size    = "1"        # "1", "2", "4", "8", "16", "32"
    service_pkg = "FDVFC247" # "FDVFC247", "FDVNET", "FDVAPP", "FDVAI"
  }
}

# FortiSOAR Virtual Machine
resource "fortiflexvm_config" "FORTISOAR_VM" {
  product_type          = "FORTISOAR_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTISOAR_VM_example"
  fortisoar_vm {
    service_pkg              = "FSRE" # "FSRE", "FSRM", "FSRD", "FSRR"
    additional_users_license = 0      # 0 ~ 1000
    addons                   = []     # "FSRTIMS"
  }
}

# FortiMail Virtual Machine
resource "fortiflexvm_config" "FORTIMAIL_VM" {
  product_type          = "FORTIMAIL_VM"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTIMAIL_VM_example"
  fortimail_vm {
    cpu_size    = "2"       # "1", "2", "4", "8", "16", "32"
    service_pkg = "FMLBASE" # "FMLBASE", "FMLATP"
    addons      = []        # "FMLFEMS", "FMLFCAS", "FMLFEOP", "FMLFEEC"
  }
}

# FortiGate Hardware
resource "fortiflexvm_config" "FGT_HW" {
  product_type          = "FGT_HW"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FGT_HW_example"
  fgt_hw {
    device_model = "FGT60F"    # For all possible values, please check https://fndn.fortinet.net/index.php?/fortiapi/954-fortiflex/5009/
                               # "FGT40F", "FGT60F", "FGT70F", "FGT80F", "FG100F", "FGT60E", "FGT61F", "FG100E", "FG101F", "FG200E", 
                               # "FG200F", "FG201F", "FG4H0F", "FG6H0F", "FWF40F", "FWF60F", "FGR60F", "FR70FB", "FGT81F", "FG101E",
                               # "FG4H1F", "FG1K0F", "FG180F", "F2K60F", "FG3K0F", "FG3K1F", "FG3K2F" ...
    service_pkg = "FGHWFC247"  # "FGHWFC247", "FGHWFCEL", "FGHWATP", "FGHWUTP", "FGHWENT", "FGHWFCESN"
    addons      = ["FGHWDLDB"] # List of string, "FGHWFCELU", "FGHWFAMS", "FGHWFAIS", "FGHWSWNM", "FGHWDLDB", "FGHWFAZC", "FGHWSOCA",
                               # "FGHWMGAS", "FGHWSPAL", "FGHWISSS", "FGHWSWOS", "FGHWAVDB", "FGHWNIDS", "FGHWFGSA", "FGHWFURL", "FGHWFSFG"
  }
}

# FortiAP Hardware
resource "fortiflexvm_config" "FAP_HW" {
  product_type          = "FAP_HW"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FAP_HW_example"
  fap_hw {
    device_model = "FP23JF"     # For all possible values, please check https://fndn.fortinet.net/index.php?/fortiapi/954-fortiflex/5010/
                                # "FP23JF", "FP221E", "FP223E", "FP231F", "FP231G", "FP233G", "FP234F"
                                # "FP234G", "FP431F", "FP431G", "FP432F", "F432FR", "FP432G", "FP433F"
                                # "FP433G", "FP441K", "FP443K", "FP831F", "PU231F", "PU234F", "PU422E"
                                # "PU431F", "PU432F", "PU433F", "FP222E", "FP224E", "FP231E"
    service_pkg = "FAPHWFC247"  # "FAPHWFC247", "FAPHWFCEL"
    addons      = []            # List of string, "FAPHWFSFG"
  }
}

# FortiSwitch Hardware
resource "fortiflexvm_config" "FSW_HW" {
  product_type          = "FSW_HW"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FSW_HW_example"
  fsw_hw {
    device_model = "S108EN"     # For all possible values, please check https://fndn.fortinet.net/index.php?/fortiapi/954-fortiflex/5011/
                                # "S108EN", "S108EF", "S108EP", "S108FN", "S108FF", "S108FP", "S124EN", "S124EF", 
                                # "S124EP", "S124FN", "S124FF", "S124FP", "S148EN", "S148EP", "S148FN", "S148FF",
                                # "S148FP", "S224DF", "S224EN", "S224EP", "S248DN", "S248EF", "S248EP", "S424DN",
                                # "S424DF", "S424DP", "S424EN", "S424EF", "S424EI", "S424EP", "S448DN", "S448DP",
                                # "S448EN", "S448EF", "S448EP", "S524DN", "S524DF", "S548DN", "S548DF", "S624FN",
                                # "S624FF", "S648FN", "S648FF", "FS1D24", "FS1E24", "FS1D48", "FS1E48", "FS2F48",
                                # "FS3D32", "FS3E32", "S426EF", "ST1E24", "SR12DP", "SR24DN", "SM10GF", "SR16FP", "SR24FP"
    service_pkg = "FSWHWFC247"  # "FSWHWFC247", "FSWHWFCEL"
  }
}

# FortiWeb Cloud - Public
resource "fortiflexvm_config" "FWBC_PUBLIC" {
  product_type          = "FWBC_PUBLIC"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FWBC_PUBLIC_example"
  fwbc_public {
    average_throughput = 150 # 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600
                             # 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 
                             # 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000
    web_applications = 50    # Number between 1 and 5000 (inclusive) 
  }
}

# FortiClient EMS Cloud
resource "fortiflexvm_config" "FC_EMS_CLOUD" {
  product_type          = "FC_EMS_CLOUD"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FC_EMS_CLOUD_example"
  fc_ems_cloud {
    ztna_num         = 225     # Value should be 0 or between 25 and 25000 (inclusive)
    ztna_fgf_num     = 225     # Value should be 0 or between 25 and 25000 (inclusive)
    epp_ztna_num     = 125     # Value should be 0 or between 25 and 25000 (inclusive)
    epp_ztna_fgf_num = 125     # Value should be 0 or between 25 and 25000 (inclusive)
    chromebook       = 100     # Value should be 0 or between 25 and 25000 (inclusive) 
    addons           = ["BPS"] # [] or ["BPS"]
  }
}

# FortiSASE
resource "fortiflexvm_config" "FORTISASE" {
  product_type          = "FORTISASE"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTISASE_example"
  fortisase {
    users                     = 50         # Number between 50 and 50,000 (inclusive)
    service_pkg               = "FSASESTD" # "FSASESTD" (Standard), "FSASEADV" (Advanced) or "FSASECOM" (Comprehensive)
    bandwidth                 = 1000       # Number between 25 and 10,000 (inclusive)
    dedicated_ips             = 4          # Number between 4 and 65,534 (inclusive)
    additional_compute_region = 0          # Number between 0 and 16 (inclusive)
    locations                 = 0          # Number between 0 and 8 (inclusive)
  }
}

# FortiEDR MSSP
resource "fortiflexvm_config" "FORTIEDR" {
  product_type          = "FORTIEDR"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTIEDR_example"
  fortiedr {
    service_pkg        = "FEDRPDR"   # Only support "FEDRPDR" (Discover/Protect/Respond) now
    addons             = ["FEDRXDR"] # Empty list or ["FEDRXDR"]
    repository_storage = 0           # Number between 0 and 30720 (inclusive)
  }
}

# FortiNDR Cloud
resource "fortiflexvm_config" "FORTINDR_CLOUD" {
  product_type          = "FORTINDR_CLOUD"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTINDR_example"
  # fortindr_cloud doesn't have any required parameters
  fortindr_cloud {
  }
}

# FortiRecon
resource "fortiflexvm_config" "FORTIRECON" {
  product_type          = "FORTIRECON"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTIRECON_example"
  fortirecon {
    service_pkg   = "FRNEASM" # "FRNEASM", "FRNEASMBP" or "FRNEASMBPACI"
    asset_num     = 200       # Number of Monitored Assets. Number between 200 and 1,000,000 (inclusive). Value should be divisible by 50.
    network_num   = 0         # Internal Attack Surface Monitoring. Number between 0 and 100 (inclusive)
    executive_num = 0         # Executive Monitoring. Number between 0 and 1,000 (inclusive). This value can only be set to 0 if `service_pkg` is `"FRNEASM"` or `"FRNEASMBP"`.
    vendor_num    = 0         # Vendor Monitoring. Number between 0 and 1,000 (inclusive) This value can only be set to 0 if `service_pkg` is `"FRNEASM"` or `"FRNEASMBP"`.
  }
}

# FortiSIEM Cloud
resource "fortiflexvm_config" "SIEM_CLOUD" {
  product_type          = "SIEM_CLOUD"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "SIEM_CLOUD_example"
  siem_cloud {
    compute_units = 10               # Number between 10 and 600 (inclusive). Value should be divisible by 10.
    additional_online_storage =  500 # Number between 500 and 60,000 (inclusive). Value should be divisible by 500.
                                     # It can be scaled up in an increment of 500 but scaling down is NOT allowed.
    archive_storage = 0              # Number between 0 and 60,000 (inclusive). Value should be divisible by 500.
                                     # It can be scaled up in an increment of 500 but scaling down is NOT allowed.
  }
}

# FortiAppSec
resource "fortiflexvm_config" "FORTIAPPSEC" {
  product_type          = "FORTIAPPSEC"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTIAPPSEC_example"
  fortiappsec {
    service_types   = ["UCWAF", "UCGSLB"] # Possible values: "UCWAF", "UCGSLB"
    waf_service_pkg = "UCWAFSTD"          # "UCWAFSTD", "UCWAFADV", "UCWAFENT". The 'Cloud WAF Service Package' can be upgraded but downgrading is NOT allowed.
    waf_addons      = []                  # Possible values:"UCSOCA"
  }
}

# FortiDLP
resource "fortiflexvm_config" "FORTIDLP" {
  product_type          = "FORTIDLP"
  program_serial_number = "ELAVMS00000XXXXX"
  name                  = "FORTIDLP_example"
  fortidlp {
    service_pkg = "DLPSTD" # "DLPSTD", "DLPENT", "DLPENTP"
    endpoints   = 25       # Number between 25 and 100000 (inclusive).
    addons      = []       # Possible values:"BPS", BPS must be active for at least 90 straight days once enabled.
  }
}
```

## Argument Reference

The following arguments are supported:

* `account_id` - (Optional/Number) Account ID. Once the fortiflexvm_config is created, you can't change the account ID of this configuration by changing `account_id`.
* `config_id` - (Optional/Number) Configuration ID. If you specify this argument, this resource will import this configuration rather than create a new one.
* `product_type` - (Required/String) Product type, must be one of the following options:
  * `FAD_VM`: FortiADC Virtual Machine
  * `FAP_HW`: FortiAP Hardware
  * `FAZ_VM`: FortiAnalyzer Virtual Machine
  * `FC_EMS_CLOUD`: FortiClient EMS Cloud
  * `FC_EMS_OP`: FortiClient EMS On-Prem
  * `FGT_HW`: FortiGate Hardware
  * `FGT_VM_Bundle`: FortiGate Virtual Machine - Service Bundle
  * `FGT_VM_LCS`: FortiGate Virtual Machine - A La Carte Services
  * `FMG_VM`: FortiManager Virtual Machine
  * `FPC_VM`: FortiPortal Virtual Machine
  * `FSW_HW`: FortiSwitch Hardware
  * `FWB_VM`: FortiWeb Virtual Machine - Service Bundle
  * `FWBC_PRIVATE`: FortiWeb Cloud - Private
  * `FWBC_PUBLIC`: FortiWeb Cloud - Public
  * `FORTIAPPSEC`: FortiAppSec
  * `FORTIDLP`: FortiDLP
  * `FORTIEDR`: FortiEDR MSSP
  * `FORTIMAIL_VM`: FortiMail Virtual Machine
  * `FORTINAC_VM`: FortiNAC Virtual Machine
  * `FORTINDR_CLOUD`: FortiNDR Cloud
  * `FORTIRECON`: FortiRecon
  * `FORTISASE`: FortiSASE
  * `FORTISOAR_VM`: FortiSOAR Virtual Machine
  * `SIEM_CLOUD`: FortiSIEM Cloud
* `program_serial_number` - (Required/String) The serial number of your FortiFlex Program. This serial number should start with `"ELAVMR"`.
* `name` - (Required unless you only update the status/String) The name of your configuration.
* `status` - (Optional/String) Configuration status. If you don't specify, the configuration status keeps unchanged. The default status is `ACTIVE` once you create a configuration. It must be one of the following options:
	* `ACTIVE`: Enable a configuration
	* `DISABLED`: Disable a configuration
* `fad_vm` - (Block List) You must fill in this block if your `product_type` is `"FAD_VM"`. The structure of [`fad_vm` block](#nestedblock--fad_vm) is documented below.
* `fap_hw` - (Block List) You must fill in this block if your `product_type` is `"FAP_HW"`. The structure of [`fap_hw` block](#nestedblock--fap_hw) is documented below.
* `faz_vm` - (Block List) You must fill in this block if your `product_type` is `"FAZ_VM"`. The structure of [`faz_vm` block](#nestedblock--faz_vm) is documented below.
* `fc_ems_cloud` - (Block List) You must fill in this block if your `product_type` is `"FC_EMS_CLOUD"`. The structure of [`fc_ems_cloud` block](#nestedblock--fc_ems_cloud) is documented below.
* `fc_ems_op` - (Block List) You must fill in this block if your `product_type` is `"FC_EMS_OP"`. The structure of [`fc_ems_op` block](#nestedblock--fc_ems_op) is documented below.
* `fgt_hw` - (Block List) You must fill in this block if your `product_type` is `"FGT_HW"`. The structure of [`fgt_hw` block](#nestedblock--fgt_hw) is documented below.
* `fgt_vm_bundle` - (Block List) You must fill in this block if your `product_type` is `"FGT_VM_Bundle"`. The structure of [`fgt_vm_bundle` block](#nestedblock--fgt_vm_bundle) is documented below.
* `fgt_vm_lcs` - (Block List) You must fill in this block if your `product_type` is `"FGT_VM_LCS"`. The structure of [`fgt_vm_lcs` block](#nestedblock--fgt_vm_lcs) is documented below.
* `fmg_vm` - (Block List) You must fill in this block if your `product_type` is `"FMG_VM"`. The structure of [`fmg_vm` block](#nestedblock--fmg_vm) is documented below.
* `fpc_vm` - (Block List) You must fill in this block if your `product_type` is `"FPC_VM"`. The structure of [`fpc_vm` block](#nestedblock--fpc_vm) is documented below.
* `fsw_hw` - (Block List) You must fill in this block if your `product_type` is `"FSW_HW"`. The structure of [`fsw_hw` block](#nestedblock--fsw_hw) is documented below.
* `fwb_vm` - (Block List) You must fill in this block if your `product_type` is `"FWB_VM"`. The structure of [`fwb_vm` block](#nestedblock--fwb_vm) is documented below.
* `fortinac_vm` - (Block List) You must fill in this block if your `product_type` is `"FORTINAC_VM"`. The structure of [`fortinac_vm` block](#nestedblock--fortinac_vm) is documented below.
* `fwbc_private` - (Block List) You must fill in this block if your `product_type` is `"FWBC_PRIVATE"`. The structure of [`fwbc_private` block](#nestedblock--fwbc_private) is documented below.
* `fwbc_public` - (Block List) You must fill in this block if your `product_type` is `"FWBC_PUBLIC"`. The structure of [`fwbc_public` block](#nestedblock--fwbc_public) is documented below.
* `fortisase` - (Block List) You must fill in this block if your `product_type` is `"FORTISASE"`. The structure of [`fortisase` block](#nestedblock--fortisase) is documented below.
* `fortiedr` - (Block List) You must fill in this block if your `product_type` is `"FORTIEDR"`. The structure of [`fortiedr` block](#nestedblock--fortiedr) is documented below.
* `fortimail_vm` - (Block List) You must fill in this block if your `product_type` is `"FORTIMAIL_VM"`. The structure of [`fortimail_vm` block](#nestedblock--fortimail_vm) is documented below.
* `fortindr_cloud` - (Block List) You must fill in this block if your `product_type` is `"FORTINDR_CLOUD"`. The structure of [`fortindr_cloud` block](#nestedblock--fortindr_cloud) is documented below.
* `fortirecon` - (Block List) You must fill in this block if your `product_type` is `"FORTIRECON"`. The structure of [`fortirecon` block](#nestedblock--fortirecon) is documented below.
* `fortisoar_vm` - (Block List) You must fill in this block if your `product_type` is `"FORTISOAR_VM"`. The structure of [`fortisoar_vm` block](#nestedblock--fortisoar_vm) is documented below.
* `siem_cloud` - (Block List) You must fill in this block if your `product_type` is `"SIEM_CLOUD"`. The structure of [`siem_cloud` block](#nestedblock--siem_cloud) is documented below.
* `fortiappsec` - (Block List) You must fill in this block if your `product_type` is `"FORTIAPPSEC"`. The structure of [`fortiappsec` block](#nestedblock--fortiappsec) is documented below.
* `fortidlp` - (Block List) You must fill in this block if your `product_type` is `"FORTIDLP"`. The structure of [`fortidlp` block](#nestedblock--fortidlp) is documented below.

<a id="nestedblock--fad_vm"></a>
The `fad_vm` block contains:

* `cpu_size` - (Required if `product_type = "FAD_VM"`/String) The number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"`, `"16"`, `"32"`.
* `service_pkg` - (Required if `product_type = "FAD_VM"`/String) Options: `"FDVFC247"` (FortiCare Premium), `"FDVNET"` (Network Security), `"FDVAPP"` (Application Security),  `"FDVAI"` (AI Security).


<a id="nestedblock--fap_hw"></a>
The `fap_hw` block contains:

* `device_model` - (Required if `product_type = "FAP_HW"`/String) Device Model. For all possible values, please check https://fndn.fortinet.net/index.php?/fortiapi/954-fortiflex/5010/. Options:
  * `"FP23JF"`: FortiAP-23JF
  * `"FP221E"`: FortiAP-221E
  * `"FP223E"`: FortiAP-223E
  * `"FP231F"`: FortiAP-231F
  * `"FP231G"`: FortiAP-231G
  * `"FP233G"`: FortiAP-233G
  * `"FP234F"`: FortiAP-234F
  * `"FP234G"`: FortiAP-234G
  * `"FP431F"`: FortiAP-431F
  * `"FP431G"`: FortiAP-431G
  * `"FP432F"`: FortiAP-432F
  * `"F432FR"`: FortiAP-432FR
  * `"FP432G"`: FortiAP-432G
  * `"FP433F"`: FortiAP-433F
  * `"FP433G"`: FortiAP-433G
  * `"FP441K"`: FortiAP-441K
  * `"FP443K"`: FortiAP-443K
  * `"FP831F"`: FortiAP-831F
  * `"PU231F"`: FortiAP-U231F
  * `"PU234F"`: FortiAP-U234F
  * `"PU422E"`: FortiAP-U422EV
  * `"PU431F"`: FortiAP-U431F
  * `"PU432F"`: FortiAP-U432F
  * `"PU433F"`: FortiAP-U433F
  * `"FP222E"`: FortiAP-222E
  * `"FP224E"`: FortiAP-224E
  * `"FP231E"`: FortiAP-231E
* `service_pkg` - (Required if `product_type = "FAP_HW"`/String) Possible values: `"FAPHWFC247"` (FortiCare Premium), `"FAPHWFCEL"` (FortiCare Elite).
* `addons` - (Optional/List of String) The default value is an empty list. Possible values:
  * `"FAPHWFSFG"`: FortiSASE Cloud Managed AP

<a id="nestedblock--fortinac_vm"></a>
The `fortinac_vm` block contains:

* `service_pkg` - (Required if `product_type = "FORTINAC_VM"`/String) Options: `"FNCPLUS"` (Plus), `"FNCPRO"` (Pro).
* `endpoints` - (Required if `product_type = "FORTINAC_VM"`/Number) Number of endpoints. Number between 25 and 100000 (inclusive).
* `support_service` - (Required if `product_type = "FORTINAC_VM"`/String) Support Service. Option: `"FCTFC247"` (FortiCare Premium).

<a id="nestedblock--faz_vm"></a>
The `faz_vm` block contains:

* `addons` - (Optional) The default value is an empty list. Options: `"FAZISSS"` (OT Security Service), `"FAZFGSA"` (Attack Surface Security Service), `"FAZAISN"` (FortiAI Service).
* `adom_num` - (Required if `product_type = "FAZ_VM"`/Number) Number of ADOMs. A number between 0 and 1200 (inclusive).
* `daily_storage` - (Required if `product_type = "FAZ_VM"`/Number) Daily Storage (GB). A number between 5 and 8300 (inclusive).
* `support_service` - (Required if `product_type = "FAZ_VM"`/String) Support Service. Option: `"FAZFC247"` (FortiCare Premium).

<a id="nestedblock--fc_ems_cloud"></a>
The `fc_ems_cloud` block contains:

* `ztna_num` - (Required if `product_type = "FC_EMS_CLOUD"`/Number) ZTNA/VPN (number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `ztna_fgf_num` - (Required if `product_type = "FC_EMS_CLOUD"`/Number) ZTNA/VPN + FortiGuard Forensics(number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `epp_ztna_num` - (Required if `product_type = "FC_EMS_CLOUD"`/Number) EPP/ATP + ZTNA/VPN (number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `epp_ztna_fgf_num` - (Required if `product_type = "FC_EMS_CLOUD"`/Number) EPP/ATP + ZTNA/VPN + FortiGuard Forensics (number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `chromebook` - (Required if `product_type = "FC_EMS_CLOUD"`/Number) Chromebook (number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `addons` - (Optional/List of String) The default value is an empty list. Options: `"BPS"` (FortiCare Best Practice).


<a id="nestedblock--fc_ems_op"></a>
The `fc_ems_op` block contains:

* `ztna_num` - (Required if `product_type = "FC_EMS_OP"`/Number) ZTNA/VPN (number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `epp_ztna_num` - (Required if `product_type = "FC_EMS_OP"`/Number) EPP/ATP + ZTNA/VPN (number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `chromebook` - (Required if `product_type = "FC_EMS_OP"`/Number) Chromebook (number of endpoints). Value should be 0 or between 25 and 25000 (inclusive).
* `support_service` - (Required if `product_type = "FC_EMS_OP"`/String) Option: `"FCTFC247"` (FortiCare Premium).
* `addons` - (Optional/List of String) The default value is an empty list. Options: `"BPS"` (FortiCare Best Practice).

<a id="nestedblock--fgt_hw"></a>
The `fgt_hw` block contains:

* `device_model` - (Required if `product_type = "FGT_HW"`/String) Device Model. For all possible values, please check https://fndn.fortinet.net/index.php?/fortiapi/954-fortiflex/5009/. Options: 
  * `"FGT40F"`: FortiGate 40F
  * `"FWF40F"`: FortiWifi 40F
  * `"FGT60E"`: FortiGate 60E
  * `"FGT60F"`: FortiGate 60F
  * `"FWF60F"`: FortiWifi 60F
  * `"FGR60F"`: FortiGateRugged 60F
  * `"FGT61F"`: FortiGate 61F
  * `"FGT70F"`: FortiGate 70F
  * `"FR70FB"`: FortiGateRugged 70F
  * `"FGT80F"`: FortiGate 80F
  * `"FGT81F"`: FortiGate 81F
  * `"FG100E"`: FortiGate 100E
  * `"FG100F"`: FortiGate 100F
  * `"FG101E"`: FortiGate 101E
  * `"FG101F"`: FortiGate 101F
  * `"FG200E"`: FortiGate 200E
  * `"FG200F"`: FortiGate 200F
  * `"FG201F"`: FortiGate 201F
  * `"FG4H0F"`: FortiGate 400F
  * `"FG4H1F"`: FortiGate 401F
  * `"FG6H0F"`: FortiGate 600F
  * `"FG1K0F"`: FortiGate 1000F
  * `"FG180F"`: FortiGate 1800F
  * `"F2K60F"`: FortiGate 2600F
  * `"FG3K0F"`: FortiGate 3000F
  * `"FG3K1F"`: FortiGate 3001F
  * `"FG3K2F"`: FortiGate 3200F
  * `"FG40FI"`: FortiGate 40F-3G4G
  * `"FW40FI"`: FortiWifi 40F-3G4G
  * `"FWF61F"`: FortiWifi 61F
  * `"FR60FI"`: FortiGateRugged 60F 3G4G
  * `"FGT71F"`: FortiGate 71F
  * `"FG80FP"`: FortiGate 80F-PoE
  * `"FG80FB"`: FortiGate 80F-Bypass
  * `"FG80FD"`: FortiGate 80F DSL
  * `"FWF80F"`: FortiWiFi 80F-2R
  * `"FW80FS"`: FortiWiFi 80F-2R-3G4G-DSL
  * `"FWF81F"`: FortiWiFi 81F 2R
  * `"FW81FS"`: FortiWiFi 81F-2R-3G4G-DSL
  * `"FW81FD"`: FortiWiFi 81F-2R-3G4G-PoE
  * `"FW81FP"`: FortiWiFi 81F 2R POE
  * `"FG81FP"`: FortiGate 81F-PoE
  * `"FGT90G"`: FortiGate 90G
  * `"FGT91G"`: FortiGate 91G
  * `"FG201E"`: FortiGate 201E
  * `"FG4H0E"`: FortiGate 400E
  * `"FG4HBE"`: FortiGate 400E BYPASS
  * `"FG4H1E"`: FortiGate 401E
  * `"FD4H1E"`: FortiGate 401E DC
  * `"FG6H0E"`: FortiGate 600E
  * `"FG6H1E"`: FortiGate 601E
  * `"FG6H1F"`: FortiGate 601F
  * `"FG9H0G"`: FortiGate 900G
  * `"FG9H1G"`: FortiGate 901G
  * `"FG1K1F"`: FortiGate 1001F
  * `"FG181F"`: FortiGate 1801F
  * `"FG3K7F"`: FortiGate 3700F
  * `"FG39E6"`: FortiGate 3960E
  * `"FG441F"`: FortiGate 4401F
* `service_pkg` - (Required if `product_type = "FGT_HW"`/String) Options:
  * `"FGHWFC247"`: FortiCare Premium
  * `"FGHWFCEL"`: FortiCare Elite
  * `"FGHWATP"`: ATP
  * `"FGHWUTP"`: UTP
  * `"FGHWENT"`: Enterprise
  * `"FGHWFCESN"`: FortiCare Essential
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
  * `"FGHWISSS"`: FortiGuard OT Security Service
  * `"FGHWSWOS"`: SD-WAN Overlay-as-a-Service
  * `"FGHWAVDB"`: Advanced Malware Protection
  * `"FGHWNIDS"`: Intrusion Prevention
  * `"FGHWFGSA"`: Attack Surface Security Service
  * `"FGHWFURL"`: Web, DNS & Video Filtering
  * `"FGHWFSFG"`: FortiSASE Subscription

<a id="nestedblock--fgt_vm_bundle"></a>
The `fgt_vm_bundle` block contains:

* `cloud_services` - (Optional/List of String) Cloud Services. The default value is an empty list. It should be a combination of:
  * `"FGTFAMS"`: FortiGate Cloud Management
  * `"FGTSWNM"`: SD-WAN Underlay
  * `"FGTSOCA"`: SOCaaS
  * `"FGTFAZC"`: FortiAnalyzer Cloud
  * `"FGTSWOS"`: Cloud-based Overlay-as-a-Service
  * `"FGTFSPA"`: SD-WAN Connector for FortiSASE
* `cpu_size` - (Required if `product_type = "FGT_VM_Bundle"`/String) The number of CPUs. Number between 1 and 96 (inclusive). 
* `fortiguard_services` - (Optional/List of String) FortiGuard Services. The default value is an empty list. It should be a combination of:
  * `"FGTAVDB"`: Advanced Malware Protection
  * `"FGTFAIS"`: AI-Based In-line Sandbox
  * `"FGTISSS"`: FortiGuard OT Security Service
  * `"FGTDLDB"`: FortiGuard DLP
  * `"FGTFGSA"`: FortiGuard Attack Surface Security Service
* `service_pkg` - (Required if `product_type = "FGT_VM_Bundle"`/String) The value of this attribute is one of `"FC"` (FortiCare), `"UTP"` (UTP), `"ENT"` (Enterprise) or `"ATP"` (ATP).
* `support_service` - (Optional/List of String) Support service. The default value is "NONE". Support values:
  * `"FGTFCELU"`: FC Elite Upgrade
* `vdom_num` - (Optional/Number) Number of VDOMs. A number between 0 and 500 (inclusive). The default number is 0.

<a id="nestedblock--fgt_vm_lcs"></a>
The `fgt_vm_lcs` block contains:

* `cloud_services` - (Optional/List of String) The cloud services this FortiGate Virtual Machine supports. The default value is an empty list. It should be a combination of:
  * `"FAMS"`: FortiGate Cloud
  * `"SWNM"`: SD-WAN Underlay
  * `"AFAC"`: FortiAnalyzer Cloud with SOCaaS
  * `"FAZC"`: FortiAnalyzer Cloud
  * `"FSPA"`: SD-WAN Connector for FortiSASE SPA
  * `"SWOS"`: Cloud-based Overlay-as-a-Service
* `cpu_size` - (Required if `product_type = "FGT_VM_LCS"`/String) The number of CPUs. A number between 1 and 96 (inclusive).
* `fortiguard_services` - (Optional/List of String) The fortiguard services this FortiGate Virtual Machine supports. The default value is an empty list. It should be a combination of:
  * `"IPS"`: Intrusion Prevention
  * `"AVDB"`: Advanced Malware
  * `"FURLDNS"`: Web, DNS & Video Filtering
  * `"FGSA"`: Security Rating
  * `"ISSS"`: OT Security Service
  * `"DLDB"`: DLP
  * `"FAIS"`: AI-Based InLine Sandbox
* `support_service` - (Required if `product_type = "FGT_VM_LCS"`/String) Options: `"FC247"` (FortiCare 24x7) or `"ASET"` (FortiCare Elite).
* `vdom_num` - (Optional/Number) Number of VDOMs. A number between 0 and 500 (inclusive). The default number is 0.


<a id="nestedblock--fmg_vm"></a>
The `fmg_vm` block contains:

* `adom_num` - (Optional/Number) Number of ADOMs. A number between 0 and 100000 (inclusive). The default value is 0.
* `managed_dev` - (Optional/Number) Number of managed devices. A number between 1 and 100000 (inclusive). The default value is 1.


<a id="nestedblock--fpc_vm"></a>
The `fpc_vm` block contains:

* `managed_dev` - (Required if `product_type = "FPC_VM"`/Number) Number of managed devices. A number between 0 and 100000 (inclusive).

<a id="nestedblock--fsw_hw"></a>
The `fsw_hw` block contains:

* `device_model` - (Required if `product_type = "FSW_HW"`/String) Device Model. For all possible values, please check https://fndn.fortinet.net/index.php?/fortiapi/954-fortiflex/5011/. Possible values: 
	* `"S108EN"`: FortiSwitch-108E
  * `"S108EF"`: FortiSwitch-108E-FPOE
  * `"S108EP"`: FortiSwitch-108E-POE
  * `"S108FN"`: FortiSwitch-108F
  * `"S108FF"`: FortiSwitch-108F-FPOE
  * `"S108FP"`: FortiSwitch-108F-POE
  * `"S124EN"`: FortiSwitch-124E
  * `"S124EF"`: FortiSwitch-124E-FPOE
  * `"S124EP"`: FortiSwitch-124E-POE
  * `"S124FN"`: FortiSwitch-124F
  * `"S124FF"`: FortiSwitch-124F-FPOE
  * `"S124FP"`: FortiSwitch-124F-POE
  * `"S148EN"`: FortiSwitch-148E
  * `"S148EP"`: FortiSwitch-148E-POE
  * `"S148FN"`: FortiSwitch-148F
  * `"S148FF"`: FortiSwitch-148F-FPOE
  * `"S148FP"`: FortiSwitch-148F-POE
  * `"S224DF"`: FortiSwitch-224D-FPOE
  * `"S224EN"`: FortiSwitch-224E
  * `"S224EP"`: FortiSwitch-224E-POE
  * `"S248DN"`: FortiSwitch-248D
  * `"S248EF"`: FortiSwitch-248E-FPOE
  * `"S248EP"`: FortiSwitch-248E-POE
  * `"S424DN"`: FortiSwitch-424D
  * `"S424DF"`: FortiSwitch-424D-FPOE
  * `"S424DP"`: FortiSwitch-424D-POE
  * `"S424EN"`: FortiSwitch-424E
  * `"S424EF"`: FortiSwitch-424E-FPOE
  * `"S424EI"`: FortiSwitch-424E-Fiber
  * `"S424EP"`: FortiSwitch-424E-POE
  * `"S448DN"`: FortiSwitch-448D
  * `"S448DP"`: FortiSwitch-448D-POE
  * `"S448EN"`: FortiSwitch-448E
  * `"S448EF"`: FortiSwitch-448E-FPOE
  * `"S448EP"`: FortiSwitch-448E-POE
  * `"S524DN"`: FortiSwitch-524D
  * `"S524DF"`: FortiSwitch-524D-FPOE
  * `"S548DN"`: FortiSwitch-548D
  * `"S548DF"`: FortiSwitch-548D-FPOE
  * `"S624FN"`: FortiSwitch-624F
  * `"S624FF"`: FortiSwitch-624F-FPOE
  * `"S648FN"`: FortiSwitch-648F
  * `"S648FF"`: FortiSwitch-648F-FPOE
  * `"FS1D24"`: FortiSwitch-1024D
  * `"FS1E24"`: FortiSwitch-1024E
  * `"FS1D48"`: FortiSwitch-1048D
  * `"FS1E48"`: FortiSwitch-1048E
  * `"FS2F48"`: FortiSwitch-2048F
  * `"FS3D32"`: FortiSwitch-3032D
  * `"FS3E32"`: FortiSwitch-3032E
  * `"S426EF"`: FortiSwitch-M426E-FPOE
  * `"ST1E24"`: FortiSwitch-T1024E
  * `"SR12DP"`: FortiSwitchRugged-112D-POE
  * `"SR24DN"`: FortiSwitchRugged-124D
  * `"SM10GF"`: FortiSwitch-110G-FPOE
  * `"SR16FP"`: FortiSwitchRugged-216F-POE
  * `"SR24FP"`: FortiSwitchRugged 424F-POE
* `service_pkg` - (Required if `product_type = "FSW_HW"`/String) Possible values: `"FSWHWFC247"` (FortiCare Premium), `"FSWHWFCEL"` (FortiCare Elite).


<a id="nestedblock--fwb_vm"></a>
The `fwb_vm` block contains:

* `cpu_size` - (Required if `product_type = "FWB_VM"`/String) Number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"` or `"16"`.
* `service_pkg` - (Required if `product_type = "FWB_VM"`/String) Service Package. Options: `"FWBSTD"` (Standard), `"FWBADV"` (Advanced) or `"FWBENT"` (Advanced).


<a id="nestedblock--fwbc_private"></a>
The `fwbc_private` block contains:

* `average_throughput` - (Required if `product_type = "FWBC_PRIVATE"`/Number) Average Throughput (Mbps). Options: 10, 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600, 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000.
* `web_applications` - (Required if `product_type = "FWBC_PRIVATE"`/Number) Number between 1 and 5000 (inclusive).


<a id="nestedblock--fwbc_public"></a>
The `fwbc_public` block contains:

* `average_throughput` - (Required if `product_type = "FWBC_PUBLIC"`/Number) Average Throughput (Mbps). Options: 25, 50, 75, 100, 150, 200, 250, 300, 350, 400, 450, 500, 600, 700, 800, 900, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 5500, 6000, 6500, 7000, 7500, 8000, 8500, 9000, 9500, 10000.
* `web_applications` - (Required if `product_type = "FWBC_PUBLIC"`/Number) Number between 1 and 5000 (inclusive) 

<a id="nestedblock--fortisase"></a>
The `fortisase` block contains:

* `users` - (Required if `product_type = "FORTISASE"`/Number) Number between 50 and 50,000 (inclusive).
* `service_pkg` - (Required if `product_type = "FORTISASE"`/String) `"FSASESTD"` (Standard), `"FSASEADV"` (Advanced) or `"FSASECOM"` (Comprehensive)
* `bandwidth` - (Required if `product_type = "FORTISASE"`/Number) Mbps. Number between 25 and 10,000 (inclusive).
* `dedicated_ips` - (Required if `product_type = "FORTISASE"`/Number) Number between 4 and 65,534 (inclusive).
* `additional_compute_region` - (Optional/Number) Additional Compute Region. Number between 0 and 16 (inclusive). The 'Additional Compute Region' can be scaled up in an increment of 1 but scaling down is NOT allowed.
* `locations` - (Optional/Number) SD-WAN On-Ramp Locations. Number between 0 and 8 (inclusive). The 'SD-WAN On-Ramp Locations' can be scaled up in an increment of 1 but scaling down is NOT allowed.

<a id="nestedblock--fortiedr"></a>
The `fortiedr` block contains:

* `service_pkg` - (Required if `product_type = "FORTIEDR"`/String) `"FEDRPDR"` (Discover/Protect/Respond).
* `endpoints` - (Read only/Number) Number of endpoints. Read only.
* `addons` - (Optional/List of String) The default value is an empty list. Options: `"FEDRXDR"` (XDR).
* `repository_storage` - (Optional/Number) Number between 0 and 30720 (inclusive). The default value is 0.

<a id="nestedblock--fortimail_vm"></a>
The `fortimail_vm` block contains:

* `cpu_size` - (Required if `product_type = "FORTIMAIL_VM"`/String) Number of CPUs. The value of this attribute is one of `"1"`, `"2"`, `"4"`, `"8"` or `"16"`.
* `service_pkg` - (Required if `product_type = "FORTIMAIL_VM"`/String) `"FMLBASE"` (Base Bundle) or `"FMLATP"` (ATP Bundle).
* `addons` - (Optional/List of String) The default value is an empty list. Options: 
  * `"FMLFEMS"`: Advanced Management
  * `"FMLFCAS"`: Dynamic Content Analysis
  * `"FMLFEOP"`: Cloud Email API Integration
  * `"FMLFEEC"`: Email Continuity

<a id="nestedblock--fortindr_cloud"></a>
The `fortindr_cloud` block contains:

* `metered_usage` - (Read only/Number) Metered Usage. Read only. Can't be set.

<a id="nestedblock--fortirecon"></a>
The `fortirecon` block contains:

* `service_pkg` - (Required if `product_type = "FORTITRECON"`/String) Possible values are:
  * `"FRNEASM"`: External Attack Surface Monitoring
  * `"FRNEASMBP"`: External Attack Surface Monitoring & Brand Protect
  * `"FRNEASMBPACI"`: External Attack Surface Monitoring & Brand Protect & Adversary Centric Intelligence
* `asset_num` - (Required if `product_type = "FORTITRECON"`/Number) Number of Monitored Assets. Number between 200 and 1,000,000 (inclusive). Value should be divisible by 50.
* `network_num` - (Optional/Number) Internal Attack Surface Monitoring. Number between 0 and 100 (inclusive)
* `executive_num` - (Optional/Number) Executive Monitoring. Number between 0 and 1,000 (inclusive). This value can only be set to 0 if `service_pkg` is `"FRNEASM"` or `"FRNEASMBP"`.
* `vendor_num` - (Optional/Number) Vendor Monitoring. Number between 0 and 1,000 (inclusive) This value can only be set to 0 if `service_pkg` is `"FRNEASM"` or `"FRNEASMBP"`.

<a id="nestedblock--fortisoar_vm"></a>
The `fortisoar_vm` block contains:

* `service_pkg` - (Required if `product_type = "FORTISOAR_VM"`/String) Service Package. Possible values are:
  * `"FSRE"`: Enterprise Edition
  * `"FSRM"`: Multi Tenant Edition - Manager
  * `"FSRD"`: Multi Tenant Edition - Tenant Node - Single User
  * `"FSRR"`: Multi Tenant Edition - Tenant Node - Multi User
* `additional_users_license` - (Optional/Number) Additional Users License. Number between 0 and 1000 (inclusive)
* `addons` - (Optional/List of String) The default value is an empty list. Options: 
  * `"FSRTIMS"`: Threat Intelligence Management

<a id="nestedblock--siem_cloud"></a>
The `siem_cloud` block contains:

* `compute_units` - (Required if `product_type = "SIEM_CLOUD"`/Number) Number of Compute Units. Number between 10 and 600 (inclusive). Value should be divisible by 10.
* `additional_online_storage` - (Required if `product_type = "SIEM_CLOUD"`/Number) Additional Online Storage. Number between 500 and 60,000 (inclusive). Value should be divisible by 500. The 'Additional Online Storage' can be scaled up in an increment of 500 but scaling down is NOT allowed.
* `archive_storage` - (Optional/Number) Archive storage. Number between 0 and 60,000 (inclusive). Value should be divisible by 500. The 'Archive Storage' can be scaled up in an increment of 500 but scaling down is NOT allowed.

<a id="nestedblock--fortiappsec"></a>
The `fortiappsec` block contains:

* `service_types` - (Optional/List of String) Service Types. The default value is an empty list. Possible values: `"UCWAF"` (Cloud WAF), `"UCGSLB"` (Cloud GSLB).
* `waf_service_pkg` - (Required if `service_types` contains `"UCWAF"`/String) Cloud WAF Service Package. `"UCWAFSTD"` (Standard), `"UCWAFADV"` (Advanced) or `"UCWAFENT"` (Enterprise). The 'Cloud WAF Service Package' can be upgraded but downgrading is NOT allowed.
* `waf_addons` - (Optional/List of String) The default value is an empty list. Options: `"UCSOCA"` (SOCaaS).
* `throughput` - (Read only/Number) Throughput. Read only. Can't be set.
* `applications` - (Read only/Number) Number of applications. Read only. Can't be set.
* `qps` - (Read only/Number) QPS. Read only. Can't be set.
* `health_checks` - (Read only/Number) Health checks. Read only. Can't be set.

<a id="nestedblock--fortidlp"></a>
The `fortidlp` block contains:

* `service_pkg` - (Required if `product_type = "FORTIDLP"`/String) `"DLPSTD"` (Standard), `"DLPENT"` (Enterprise), `"DLPENTP"` (Enterprise Premium).
* `endpoints` - (Required if `product_type = "FORTIDLP"`/Number) Number of endpoints. Number between 25 and 100000 (inclusive).
* `addons` - (Optional/List of String) The default value is an empty list. Options: `"BPS"` (Best Practice Service). Best Practice Service must be active for at least 90 straight days once enabled.

## Attribute Reference

The following attribute is exported:

* `id` - (String) An ID for the resource.

## Import

FortiFlex Configuration can be imported by using the following steps:

Method 1: Specify `config_id`
```
resource "fortiflexvm_config" "import_example" {
  config_id             = 12345
  product_type          = "FGT_VM_Bundle"
  program_serial_number = "ELAVMS00000XXXXX"
}
```


Method 2: Use `terraform import`
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