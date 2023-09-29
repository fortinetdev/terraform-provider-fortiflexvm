## 2.1.1 (Unreleased)

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