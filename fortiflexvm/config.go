package fortiflexvm

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
)

var PRODUCT_TYPES = []string{"fgt_vm_bundle", "fmg_vm", "fwb_vm", "fgt_vm_lcs", "fc_ems_op", "faz_vm",
	"fpc_vm", "fad_vm", "fgt_hw", "fwbc_private", "fwbc_public", "fc_ems_cloud"}

func fortiAPIPatch(t interface{}) bool {
	if t == nil {
		return false
	} else if _, ok := t.(string); ok {
		return true
	} else if _, ok := t.(float64); ok {
		return true
	} else if _, ok := t.([]interface{}); ok {
		return true
	}

	return false
}

func convProductTypeName2Id(p_type string) int {
	switch p_type {
	case "FGT_VM_Bundle":
		return 1
	case "FMG_VM":
		return 2
	case "FWB_VM":
		return 3
	case "FGT_VM_LCS":
		return 4
	case "FC_EMS_OP":
		return 5
	case "FAZ_VM":
		return 7
	case "FPC_VM":
		return 8
	case "FAD_VM":
		return 9
	case "FGT_HW":
		return 101
	case "FWBC_PRIVATE":
		return 202
	case "FWBC_PUBLIC":
		return 203
	case "FC_EMS_CLOUD":
		return 204
	default:
		return 0
	}
}

func convProductTypeId2Name(p_id int) string {
	switch p_id {
	case 1:
		return "FGT_VM_Bundle"
	case 2:
		return "FMG_VM"
	case 3:
		return "FWB_VM"
	case 4:
		return "FGT_VM_LCS"
	case 5:
		return "FC_EMS_OP"
	case 7:
		return "FAZ_VM"
	case 8:
		return "FPC_VM"
	case 9:
		return "FAD_VM"
	case 101:
		return "FGT_HW"
	case 202:
		return "FWBC_PRIVATE"
	case 203:
		return "FWBC_PUBLIC"
	case 204:
		return "FC_EMS_CLOUD"
	default:
		return ""
	}
}

func convConfParsId2NameList(p_id int) (string, string, string) {
	switch p_id {
	case 1:
		return "fgt_vm_bundle", "cpu_size", "string"
	case 2:
		return "fgt_vm_bundle", "service_pkg", "string"
	case 3:
		return "fmg_vm", "managed_dev", "int"
	case 4:
		return "fwb_vm", "cpu_size", "string"
	case 5:
		return "fwb_vm", "service_pkg", "string"
	case 6:
		return "fgt_vm_lcs", "cpu_size", "String"
	case 7:
		return "fgt_vm_lcs", "fortiguard_services", "list"
	case 8:
		return "fgt_vm_lcs", "support_service", "string"
	case 9:
		return "fmg_vm", "adom_num", "int"
	case 10:
		return "fgt_vm_bundle", "vdom_num", "int"
	case 11:
		return "fgt_vm_lcs", "vdom_num", "int"
	case 12:
		return "fgt_vm_lcs", "cloud_services", "list"
	case 13:
		return "fc_ems_op", "ztna_num", "int"
	case 14:
		return "fc_ems_op", "epp_ztna_num", "int"
	case 15:
		return "fc_ems_op", "chromebook", "int"
	case 16:
		return "fc_ems_op", "support_service", "string"
	case 21:
		return "faz_vm", "daily_storage", "int"
	case 22:
		return "faz_vm", "adom_num", "int"
	case 23:
		return "faz_vm", "support_service", "string"
	case 24:
		return "fpc_vm", "managed_dev", "int"
	case 25:
		return "fad_vm", "cpu_size", "string"
	case 26:
		return "fad_vm", "service_pkg", "string"
	case 27:
		return "fgt_hw", "device_model", "string"
	case 28:
		return "fgt_hw", "service_pkg", "string"
	case 29:
		return "fgt_hw", "addons", "list"
	case 30:
		return "fmg_vm", "managed_dev", "int"
	case 32:
		return "fwbc_private", "average_throughput", "int"
	case 33:
		return "fwbc_private", "web_applications", "int"
	case 34:
		return "fwbc_public", "average_throughput", "int"
	case 35:
		return "fwbc_public", "web_applications", "int"
	case 36:
		return "fc_ems_op", "addons", "list"
	case 37:
		return "fc_ems_cloud", "ztna_num", "int"
	case 38:
		return "fc_ems_cloud", "ztna_fgf_num", "int"
	case 39:
		return "fc_ems_cloud", "epp_ztna_num", "int"
	case 40:
		return "fc_ems_cloud", "epp_ztna_fgf_num", "int"
	case 41:
		return "fc_ems_cloud", "chromebook", "int"
	case 42:
		return "fc_ems_cloud", "addons", "list"
	case 43:
		return "fgt_vm_bundle", "fortiguard_services", "list"
	case 44:
		return "fgt_vm_bundle", "cloud_services", "list"
	case 45:
		return "fgt_vm_bundle", "support_service", "string"
	default:
		return "", "", ""
	}
}

func convConfParsNameList2Id(p_type, c_name string) int {
	switch p_type {
	case "fgt_vm_bundle":
		switch c_name {
		case "cpu_size":
			return 1
		case "service_pkg":
			return 2
		case "vdom_num":
			return 10
		case "fortiguard_services":
			return 43
		case "cloud_services":
			return 44
		case "support_service":
			return 45
		default:
			return 0
		}
	case "fmg_vm":
		switch c_name {
		case "managed_dev":
			return 30
		case "adom_num":
			return 9
		default:
			return 0
		}
	case "fwb_vm":
		switch c_name {
		case "cpu_size":
			return 4
		case "service_pkg":
			return 5
		default:
			return 0
		}
	case "fgt_vm_lcs":
		switch c_name {
		case "cpu_size":
			return 6
		case "fortiguard_services":
			return 7
		case "support_service":
			return 8
		case "vdom_num":
			return 11
		case "cloud_services":
			return 12
		default:
			return 0
		}
	case "fc_ems_op":
		switch c_name {
		case "ztna_num":
			return 13
		case "epp_ztna_num":
			return 14
		case "chromebook":
			return 15
		case "support_service":
			return 16
		case "addons":
			return 36
		default:
			return 0
		}
	case "faz_vm":
		switch c_name {
		case "daily_storage":
			return 21
		case "adom_num":
			return 22
		case "support_service":
			return 23
		default:
			return 0
		}
	case "fpc_vm":
		switch c_name {
		case "managed_dev":
			return 24
		default:
			return 0
		}
	case "fad_vm":
		switch c_name {
		case "cpu_size":
			return 25
		case "service_pkg":
			return 26
		default:
			return 0
		}
	case "fgt_hw":
		switch c_name {
		case "device_model":
			return 27
		case "service_pkg":
			return 28
		case "addons":
			return 29
		default:
			return 0
		}
	case "fwbc_private":
		switch c_name {
		case "average_throughput":
			return 32
		case "web_applications":
			return 33
		default:
			return 0
		}
	case "fwbc_public":
		switch c_name {
		case "average_throughput":
			return 34
		case "web_applications":
			return 35
		default:
			return 0
		}
	case "fc_ems_cloud":
		switch c_name {
		case "ztna_num":
			return 37
		case "ztna_fgf_num":
			return 38
		case "epp_ztna_num":
			return 39
		case "epp_ztna_fgf_num":
			return 40
		case "chromebook":
			return 41
		case "addons":
			return 42
		default:
			return 0
		}
	default:
		return 0
	}
}

func isInterfaceEmpty(i interface{}) bool {
	if i == nil {
		return true
	}

	switch i.(type) {
	case string:
		return i.(string) == ""
	case int:
		return i.(int) == 0
	case []interface{}:
		return len(i.([]interface{})) == 0
	case map[string]interface{}:
		return len(i.(map[string]interface{})) == 0
	default:
		return false
	}
}

func importOptionChecking(c *Config, para string) string {
	v := c.ImportOptions.List()
	if len(v) == 0 {
		return ""
	}

	for _, v1 := range v {
		if v2, ok := v1.(string); ok {
			v3 := strings.Split(v2, "=")

			if len(v3) == 2 { // Example "program_serial_number=******"
				if v3[0] == para {
					return v3[1]
				}
			}
		}
	}

	return ""
}

func contains(str_list []string, target string) bool {
	for _, str := range str_list {
		if target == str {
			return true
		}
	}
	return false
}

func checkInputValidString(parameter_name string, valid_values []string) func(interface{}, cty.Path) diag.Diagnostics {
	return func(v interface{}, p cty.Path) diag.Diagnostics {
		value := v.(string)
		var diags diag.Diagnostics
		if flag := contains(valid_values, value); !flag {
			diag := diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Invalid value of parameter: %v", parameter_name),
				Detail:   fmt.Sprintf("Invalid %v value: %v\nValid values: %v", parameter_name, value, valid_values),
			}
			diags = append(diags, diag)
		}
		return diags
	}
}

func checkInputValidStringList(parameter_name string, valid_values []string) func(interface{}, cty.Path) diag.Diagnostics {
	return func(v interface{}, p cty.Path) diag.Diagnostics {
		values := v.([]string)
		var diags diag.Diagnostics
		flag := true
		for _, value := range values {
			flag = flag && contains(valid_values, value)
			if !flag {
				diag := diag.Diagnostic{
					Severity: diag.Error,
					Summary:  fmt.Sprintf("Invalid value of parameter: %v", parameter_name),
					Detail:   fmt.Sprintf("Invalid %v value: %v\nValid values (you can select multiple values): %v", parameter_name, value, valid_values),
				}
				diags = append(diags, diag)
				break
			}
		}
		return diags
	}
}

func checkInputValidInt(parameter_name string, lower_bound int, upper_bound int) func(interface{}, cty.Path) diag.Diagnostics {
	return func(v interface{}, p cty.Path) diag.Diagnostics {
		value := v.(int)
		var diags diag.Diagnostics
		if value < lower_bound || value > upper_bound {
			diag := diag.Diagnostic{
				Severity: diag.Error,
				Summary:  fmt.Sprintf("Invalid value of parameter: %v", parameter_name),
				Detail:   fmt.Sprintf("Invalid %v value: %v\nValid values: number between %v and %v (inclusive)", parameter_name, value, lower_bound, upper_bound),
			}
			diags = append(diags, diag)
		}
		return diags
	}
}

func splitID(resource_id string) (string, string, diag.Diagnostics) {
	var diags diag.Diagnostics
	split_parts := strings.Split(resource_id, ".")
	if len(split_parts) != 2 {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  "Unable to handle id in fortiflexvm_entitlement",
			Detail:   "Incorrect id format. Please use 'serial_number' + '.' + 'config_id', example: 'FGVMMLTM12345678.123' ",
		})
		return "", "", diags
	}
	serial_number := split_parts[0]
	config_id := split_parts[1]
	_, err := strconv.Atoi(config_id)
	if err != nil {
		return "", "", diag.FromErr(fmt.Errorf("The id you import in fortiflexvm_entitlement is incorrect."+
			"Please use 'serial_number' + '.' + 'config_id', example: 'FGVMMLTM12345678.123'."+
			"Your serial_number: %s, your config_id: %s (should be an integer).", serial_number, config_id))
	}
	return serial_number, config_id, diags
}

func getEntitlementFromId(resource_id string, m interface{}) (map[string]interface{}, diag.Diagnostics) {
	c := m.(*FortiClient).Client
	c.Retries = 1

	serial_number, config_id, diags := splitID(resource_id)
	if diags.HasError() {
		return nil, diags
	}

	// Get entitlements list
	obj := make(map[string]interface{})
	obj["configId"] = config_id
	return_data, err := c.ReadEntitlementsList(&obj)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	// Find target entitlement
	target_entitlement, err := findEntitlementFromList(return_data, serial_number)
	if err != nil {
		return nil, diag.FromErr(err)
	}

	return target_entitlement, diags
}

func findEntitlementFromList(entitlement_list map[string]interface{}, serial_number string) (map[string]interface{}, error) {
	if entitlement_list == nil {
		return nil, fmt.Errorf("Response from FlexVM API is nil")
	}

	if ent_list, ok := entitlement_list["entitlements"].([]interface{}); ok {
		for _, data_ent := range ent_list {
			if ent, ok := data_ent.(map[string]interface{}); ok {
				cId := fmt.Sprintf("%v", ent["serialNumber"])
				if cId == serial_number {
					return ent, nil
				}
			}
		}
	}
	return nil, fmt.Errorf("Target entitlement %v not exist", serial_number)
}

func changeVMStatus(serial_number string, action string, m interface{}) (map[string]interface{}, error) {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj := make(map[string]interface{})
	obj["serialNumber"] = serial_number
	target_entitlement, err := c.UpdateVmUpdateStatus(&obj, action)
	return target_entitlement, err
}
