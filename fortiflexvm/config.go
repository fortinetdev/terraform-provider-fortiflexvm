package fortiflexvm

import (
	"net"
	"os"
	"strconv"
	"strings"
)

func validateConvIPMask2CIDR(oNewIP, oOldIP string) string {
	if oNewIP != oOldIP && strings.Contains(oNewIP, "/") && strings.Contains(oOldIP, " ") {
		line := strings.Split(oOldIP, " ")
		if len(line) >= 2 {
			ip := line[0]
			mask := line[1]
			prefixSize, _ := net.IPMask(net.ParseIP(mask).To4()).Size()
			return ip + "/" + strconv.Itoa(prefixSize)
		}
	}
	return oOldIP
}

func fortiStringValue(t interface{}) string {
	if v, ok := t.(string); ok {
		return v
	} else {
		return ""
	}
}

func fortiIntValue(t interface{}) int {
	if v, ok := t.(float64); ok {
		return int(v)
	} else {
		return 0
	}
}

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

func isImportTable() bool {
	itable := os.Getenv("FLEXVM_IMPORT_TABLE")
	if itable == "false" {
		return false
	}
	return true
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
	case "FAZ_VM":
		return 7
	case "FPC_VM":
		return 8
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
	case 7:
		return "FAZ_VM"
	case 8:
		return "FPC_VM"
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
	case 21:
		return "faz_vm", "daily_storage", "int"
	case 22:
		return "faz_vm", "adom_num", "int"
	case 23:
		return "faz_vm", "support_service", "string"
	case 24:
		return "fpc_vm", "managed_dev", "int"
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
		default:
			return 0
		}
	case "fmg_vm":
		switch c_name {
		case "managed_dev":
			return 3
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
