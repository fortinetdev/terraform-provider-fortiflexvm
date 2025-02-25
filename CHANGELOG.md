## 2.4.1 (Unreleased)

## 2.4.0 (February 25, 2025)

FEATURES:

* **New Ephemeral Resource:** `fortiflexvm_groups_nexttoken`

## 2.3.4 (December 11, 2024)

IMPROVEMENTS:

* Supported 1 new configuration: `FortiRecon`.


## 2.3.3 (October 7, 2024)

IMPROVEMENTS:

* Supported 1 new configuration: `FortiSIEM Cloud` (`siem_cloud`)
* Added `additional_compute_region` to `fortisase` configuration.
* Improved SDK logic.

## 2.3.2 (June 17, 2024)

IMPROVEMENTS:

* Supported 2 new configurations: `FortiAP Hardware` (`fap_hw`) and `FortiSwitch Hardware` (`fsw_hw`).
* Configuration `faz_vm` added one new argument `addons`.
* `fortiflexvm_entitlements_vm` supported new input argument `skip_pending`.


## 2.3.1 (April 3, 2024)

IMPROVEMENTS:

* Supported 2 new configurations: `FortiSASE` (`fortisase`) and `FortiEDR` (`fortiedr`).
* `fortiflexvm_config` supported input argument `config_id`. You can import existing configurations by specifying this argument.
* `fortiflexvm_retrieve_vm_group` supported input argument `retrieve_status`. It can retrieve both PENDING and STOPPED entitlements if you set `retrieve_status = ["STOPPED", "PENDING"]`.
* `fortiflexvm_retrieve_vm_group` supported input argument `require_exact_count`. The default value is false, if set as true, the resource will release retrieved entitlements and report an error if the resource doesn't get enough `count_num` entitlements.
* Reported a warning rather than an error if `end_date` in `fortiflexvm_entitlements_vm` is set incorrectly.

## 2.3.0 (Feburary 23, 2024)

FEATURES:

* **New Resource:** `fortiflexvm_retrieve_vm_group`

IMPROVEMENTS:

* Data resource `fortiflexvm_entitlements_list` supports new input arguments `description`,  `serial_number`, `status`, `token_status`.
* Data resource `fortiflexvm_groups_nexttoken` supports new input argument `status`.
* Improve the logic of the resource `fortiflexvm_entitlements_vm` and `fortiflexvm_entitlements_cloud`. If you specify the argument `serial_number` and `config_id` at the same time, it will import an existing resource rather than create a new one.
* Resource `fortiflexvm_entitlements_vm` supports new input argument `refresh_token_when_destroy`. If set it as true, the token of this entitlement will be refreshed when you use `terraform destroy`.

## 2.2.1 (Feburary 9, 2024)

IMPROVEMENTS:

* Add update entitlements scenario use cases in the document.
* Improve examples format.

## 2.2.0 (December 5, 2023)

FEATURES:

* **New Resource:** `fortiflexvm_entitlements_cloud`

IMPROVEMENTS:

* New product: `FortiClient EMS Cloud` (`fc_ems_clous`)
* Configuration `fgt_vm_bundle` supports `fortiguard_services`, `cloud_services` and `support_service`.

## 2.1.0 (September 29, 2023)

IMPROVEMENTS:

* Support 3 new products: `FortiClient EMS On-Prem`, `FortiWeb Cloud - Private` and `FortiWeb Cloud - Public`.
* Support return value `accound_id` for all resources and data sources.
* `data_source_entitlements_list` supports `account_id` and `program_serial_number`.

## 2.0.0 (July 17, 2023)

BREAKING CHANGES:

* Resource `data_source_vms_points` has been renamed to `data_source_entitlements_points`. Rename attribute `vms` to `entitlements`.
* Resource `data_source_vms_list` has been renamed to `data_source_entitlements_list`. Rename attribute `vms` to `entitlements`.
* Rename `vms` in `fortiflexvm_groups_nexttoken` to `entitlements`.
* Resource `fortiflexvm_vms_create` and `fortiflexvm_vms_update` are deleted, use `fortiflexvm_entitlements_vm` instead.


FEATURES:

* **New Resource:** `fortiflexvm_entitlements_vm`
* **New Resource:** `fortiflexvm_entitlements_hardware`
* **New Resource:** `fortiflexvm_entitlements_vm_token`

IMPROVEMENTS:

* Resource `resource_config` support `fad_vm` (FortiADC Virtual Machine) and `fgt_hw` (FortiGate Hardware)

BUG FIXES:

* Fix an issue where users could receive "Invalid security token".


## 1.0.0 (March 22, 2023)

FEATURES:

* **New Resource:** `fortiflexvm_config`
* **New Data Source:** `fortiflexvm_configs_list`
* **New Data Source:** `fortiflexvm_groups_list`
* **New Data Source:** `fortiflexvm_groups_nexttoken`
* **New Data Source:** `fortiflexvm_programs_list`
* **New Resource:** `fortiflexvm_vms_create`
* **New Data Source:** `fortiflexvm_vms_list`
* **New Data Source:** `fortiflexvm_vms_points`
* **New Resource:** `fortiflexvm_vms_update`