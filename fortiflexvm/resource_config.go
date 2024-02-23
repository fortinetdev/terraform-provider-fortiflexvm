// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Configure a Configuration under a Program.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigCreate,
		Read:   resourceConfigRead,
		Update: resourceConfigUpdate,
		Delete: resourceConfigDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"program_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_type": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				Description: `Product Type ID, must be one of the following options:
				FGT_VM_Bundle: FortiGate Virtual Machine - Service Bundle;
				FMG_VM: FortiManager Virtual Machine;
				FWB_VM: FortiWeb Virtual Machine - Service Bundle;
				FGT_VM_LCS: FortiGate Virtual Machine - A La Carte Services;
				FC_EMS_OP: FortiClient EMS On-Prem;
				FAZ_VM: FortiAnalyzer Virtual Machine;
				FPC_VM: FortiPortal Virtual Machine;
				FAD_VM: FortiADC Virtual Machine;
				FGT_HW: FortiGate Hardware;
				FWBC_PRIVATE: FortiWeb Cloud - Private;
				FWBC_PUBLIC: FortiWeb Cloud - Public;
				FC_EMS_CLOUD: FortiClient EMS Cloud.`,
				ValidateDiagFunc: checkInputValidString("product_type", []string{"FGT_VM_Bundle", "FMG_VM", "FWB_VM", "FGT_VM_LCS",
					"FC_EMS_OP", "FC_EMS_CLOUD", "FAZ_VM", "FPC_VM", "FAD_VM", "FGT_HW", "FWBC_PRIVATE", "FWBC_PUBLIC"}),
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: `Configuration status, must be one of the following options:
				ACTIVE: Enable a configuration;
				DISABLED: Disable a configuration.`,
				ValidateDiagFunc: checkInputValidString("status", []string{"ACTIVE", "DISABLED"}),
			},
			"fgt_vm_bundle": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortiguard_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"cloud_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fmg_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_dev": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fwb_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fgt_vm_lcs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiguard_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cloud_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fc_ems_op": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"epp_ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"chromebook": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"faz_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"daily_storage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fpc_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_dev": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fad_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fgt_hw": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fwbc_private": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"average_throughput": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"web_applications": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fwbc_public": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"average_throughput": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"web_applications": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fc_ems_cloud": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ztna_fgf_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"epp_ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"epp_ztna_fgf_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"chromebook": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceConfigCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	obj, err := getObjectConfig(d, "create")
	if err != nil {
		return fmt.Errorf("error creating Config resource while getting object: %v", err)
	}

	response_data, err := c.CreateConfig(obj)

	if err != nil {
		return fmt.Errorf("error creating Config resource: %v", err)
	}

	if response_data["id"] != nil && response_data["id"] != "" {
		d.SetId(fmt.Sprintf("%v", response_data["id"]))
	} else {
		d.SetId("Config")
	}

	// set status
	if current_status, ok := response_data["status"].(string); ok {
		if set_status, ok := d.GetOk("status"); ok && current_status != set_status.(string) {
			obj, err = getObjectConfig(d, "id")
			if err != nil {
				return fmt.Errorf("error creating Config resource while getting object: %v", err)
			}

			var op string
			if d.Get("status").(string) == "ACTIVE" {
				op = "enable"
			} else {
				op = "disable"
			}
			response_data, err = c.UpdateConfigStatus(obj, op)
			if err != nil {
				return fmt.Errorf("error update Config status: %v", err)
			}
			if st, ok := response_data["status"].(string); ok {
				if st != d.Get("status").(string) {
					log.Printf("[WARN] Could not update the status of Config %v", d.Id())
				}
			}
		}
	} else {
		log.Printf("[WARN] Could not get status from HTTP response")
	}

	// refresh schema
	err = refreshObjectConfig(d, response_data)
	if err != nil {
		return fmt.Errorf("error refresh Config resource: %v", err)
	}

	return nil
}

func resourceConfigRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if d.Get("program_serial_number") == "" {
		psn := importOptionChecking(m.(*FortiClient).Cfg, "program_serial_number")
		if err := d.Set("program_serial_number", psn); err != nil {
			return fmt.Errorf("error set params program_serial_number: %v", err)
		}
	}
	obj, err := getObjectConfig(d, "read")
	if err != nil {
		return fmt.Errorf("error reading Config while getting required parameters: %v", err)
	}

	o, err := c.ReadConfigsList(obj)
	if err != nil {
		return fmt.Errorf("error reading Config resource: %v", err)
	}

	co, err := getConfigReadResponse(o, d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	err = refreshObjectConfig(d, co)
	if err != nil {
		return fmt.Errorf("error reading Config resource from API: %v", err)
	}
	return nil
}

func resourceConfigUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	obj, err := getObjectConfig(d, "update")
	if err != nil {
		return fmt.Errorf("error updating Config resource while getting object: %v", err)
	}

	o, err := c.UpdateConfig(obj)
	if err != nil {
		return fmt.Errorf("error updating Config resource: %v", err)
	}

	if o["id"] != nil && o["id"] != "" {
		d.SetId(fmt.Sprintf("%v", o["id"]))
	} else {
		d.SetId("Config")
	}

	if st, ok := o["status"].(string); ok {
		if statusV, ok := d.GetOk("status"); ok && st != statusV.(string) {
			obj, err = getObjectConfig(d, "id")
			if err != nil {
				return fmt.Errorf("error creating Config resource while getting object: %v", err)
			}

			var op string
			if d.Get("status").(string) == "ACTIVE" {
				op = "enable"
			} else {
				op = "disable"
			}

			o, err = c.UpdateConfigStatus(obj, op)
			if err != nil {
				return fmt.Errorf("error update Config status: %v", err)
			}
			if st, ok := o["status"].(string); ok {
				if st != d.Get("status").(string) {
					log.Printf("[WARN] Could not update the status of Config %v", d.Id())
				}
			}
		}
	} else {
		log.Printf("[WARN] Could not get status from HTTP response")
	}

	err = refreshObjectConfig(d, o)
	if err != nil {
		return fmt.Errorf("error refresh Config resource: %v", err)
	}

	return nil
}

func resourceConfigDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if d.Get("status").(string) != "DISABLED" {
		obj, err := getObjectConfig(d, "id")
		if err != nil {
			return fmt.Errorf("error creating Config resource while getting object: %v", err)
		}

		o, err := c.UpdateConfigStatus(obj, "disable")
		if err != nil {
			return fmt.Errorf("error update Config status: %v", err)
		}
		if st, ok := o["status"].(string); ok {
			if st != d.Get("status").(string) {
				log.Printf("[WARN] Could not update the status of Config %v", d.Id())
			}
		}

		err = refreshObjectConfig(d, o)
		if err != nil {
			return fmt.Errorf("error refresh Config resource: %v", err)
		}
	}

	d.SetId("")
	return nil
}

func getConfigReadResponse(o map[string]interface{}, mkey string) (map[string]interface{}, error) {
	var err error
	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", mkey)
		err = fmt.Errorf("Response is nil")
		return nil, err
	}

	if configs, ok := o["configs"]; ok {
		if configsList, ok := configs.([]interface{}); ok {
			for _, conf := range configsList {
				if confMap, ok := conf.(map[string]interface{}); ok {
					cId := fmt.Sprintf("%v", confMap["id"])
					if cId == mkey {
						return confMap, nil
					}
				}
			}
		}
	}
	err = fmt.Errorf("Config not been created")
	return nil, err
}

func flattenConfigProductType(v interface{}) interface{} {
	var rst interface{}
	rst = ""
	if pt, ok := v.(map[string]interface{}); ok {
		if p_id, ok := pt["id"]; ok {
			rst = convProductTypeId2Name(int(p_id.(float64)))
			if rst == "" {
				log.Printf("[ERROR] Can not recognise Product Type ID: %v", p_id)
			}
		}
	}
	return rst
}

func flattenConfigParameters(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, 1)
	tmp := make(map[string]interface{})
	for _, r := range l {
		param := r.(map[string]interface{})
		_, cName, dataType := convConfParsId2NameList(int(param["id"].(float64)))
		if cName == "" {
			log.Printf("NEW PARAM: %v", param["id"])
			continue
		}
		if cValue, ok := param["value"]; ok {
			switch dataType {
			case "int":
				tmp[cName], _ = strconv.Atoi((cValue.(string)))
			case "string":
				tmp[cName] = cValue.(string)
			case "list":
				if _, ok := tmp[cName]; ok == false {
					tmp[cName] = []interface{}{}
				}
				if cValue != "NONE" {
					tmp[cName] = append(tmp[cName].([]interface{}), cValue)
				}
			default:
				tmp[cName] = cValue.(string)
			}
		}
	}
	result = append(result, tmp)

	return result
}

func refreshObjectConfig(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if value, ok := o["accountId"]; ok {
		d.Set("account_id", value)
	}
	if value, ok := o["programSerialNumber"]; ok {
		d.Set("program_serial_number", value)
	}
	if value, ok := o["name"]; ok {
		d.Set("name", value)
	}
	if err = d.Set("product_type", flattenConfigProductType(o["productType"])); err != nil {
		if !fortiAPIPatch(o["productType"]) {
			return fmt.Errorf("error reading product_type: %v", err)
		}
	}
	if value, ok := o["status"]; ok {
		d.Set("status", value)
	}

	// initialize product variables. This can fix the problem of inconsistent output.
	for _, type_name := range PRODUCT_TYPES {
		empty_interface := make([]map[string]interface{}, 0)
		d.Set(type_name, empty_interface)
	}

	// Set param
	pType := strings.ToLower(d.Get("product_type").(string))
	if err = d.Set(pType, flattenConfigParameters(o["parameters"])); err != nil {
		if !fortiAPIPatch(o["parameters"]) {
			return fmt.Errorf("error reading %v: %v", pType, err)
		}
	}

	return nil
}

func expandConfigProductType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	var typeId interface{}
	typeId = convProductTypeName2Id(v.(string))
	if typeId == 0 {
		err := fmt.Errorf("product_type invalid: %v, should be one of [%v]", v.(string),
			"FGT_VM_Bundle, FMG_VM, FWB_VM, FGT_VM_LCS, FC_EMS_OP, FAZ_VM, FPC_VM, FAD_VM, FGT_HW, FWBC_PRIVATE, FWBC_PUBLIC, FC_EMS_CLOUD")
		return typeId, err
	}
	return typeId, nil
}

func expandConfigParameters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	for ck, cv := range l[0].(map[string]interface{}) {
		// if isInterfaceEmpty(cv) {
		// 	continue
		// }
		ckId := convConfParsNameList2Id(pre, ck)
		if ckId == 0 {
			err := fmt.Errorf("Could not get target argument ID, this is a plugin error.")
			log.Printf("[ERROR] %v", err)
			return result, err
		}
		if cvList, ok := cv.([]interface{}); ok {
			for _, csv := range cvList {
				// 				if pre == "fgt_vm_lcs" { // input check
				// 					fortiguard_services_valid_values := []string{"IPS", "AVDB", "FURL", "IOTH", "FGSA", "ISSS"}
				// 					cloud_services_valid_values := []string{"FAMS", "SWNM", "FMGC", "AFAC"}
				// 					if ck == "fortiguard_services" && !contains(fortiguard_services_valid_values, csv.(string)) {
				// 						return result, fmt.Errorf(`Invalid fgt_vm_lcs.fortiguard_services input %v
				// Valid values (you can select multiple values): %v`, csv.(string), fortiguard_services_valid_values)
				// 					} else if ck == "cloud_services" && !contains(cloud_services_valid_values, csv.(string)) {
				// 						return result, fmt.Errorf(`Invalid fgt_vm_lcs.cloud_services input %v
				// Valid values (you can select multiple values): %v`, csv.(string), cloud_services_valid_values)
				// 					}
				// 				}
				tmp := make(map[string]interface{})
				tmp["id"] = ckId
				tmp["value"] = csv
				result = append(result, tmp)
			}
			if len(cvList) == 0 { // if this list is empty, send "NONE" to the fortiflex server
				tmp := make(map[string]interface{})
				tmp["id"] = ckId
				tmp["value"] = "NONE"
				result = append(result, tmp)
			}
		} else {
			if pre == "fgt_vm_bundle" { // version 2.2.0, allow fgt_vm_bundle->support_service empty
				if ck == "support_service" && cv == "" {
					cv = "NONE"
				}
			}
			tmp := make(map[string]interface{})
			tmp["id"] = ckId
			tmp["value"] = cv
			result = append(result, tmp)
		}
	}

	return result, nil
}

func getObjectConfig(d *schema.ResourceData, rType string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if rType == "update" || rType == "id" {
		obj["id"] = d.Id()
	}

	if rType == "create" || rType == "read" {
		if value, ok := d.GetOk("program_serial_number"); ok {
			obj["programSerialNumber"] = value
		}
		if value, ok := d.GetOk("account_id"); ok {
			obj["accountId"] = value
		}
	}

	if rType == "create" || rType == "update" {
		if value, ok := d.GetOk("name"); ok {
			obj["name"] = value
		}

		var pType string
		if v, ok := d.GetOk("product_type"); ok {
			pType = v.(string)
			if rType == "create" {
				t, err := expandConfigProductType(d, v, "product_type")
				if err != nil {
					return &obj, err
				} else if t != nil {
					obj["productTypeId"] = t
				}
			}
		}

		pTypeLower := strings.ToLower(pType)
		if v, ok := d.GetOk(pTypeLower); ok {
			t, err := expandConfigParameters(d, v, pTypeLower)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["parameters"] = t
			}
		}
	}

	return &obj, nil
}
