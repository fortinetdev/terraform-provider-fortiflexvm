// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Configure a Configuration under a Program.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
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
			"program_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"product_type": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				// ExactlyOneOf: []string{"FGT_VM_Bundle", "FMG_VM", "FWB_VM", "FGT_VM_LCS", "FAZ_VM", "FPC_VM"},
				Description: `Product Type ID, must be one of the following options:
				FGT_VM_Bundle: FortiGate Virtual Machine - Service Bundle;
				FMG_VM: FortiManager Virtual Machine;
				FWB_VM: FortiWeb Virtual Machine - Service Bundle;
				FGT_VM_LCS: FortiGate Virtual Machine - A La Carte Services;
				FAZ_VM: FortiAnalyzer Virtual Machine;
				FPC_VM: FortiPortal Virtual Machine.`,
				ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					value := v.(string)
					valid_values := []string{"FGT_VM_Bundle", "FMG_VM", "FWB_VM", "FGT_VM_LCS", "FAZ_VM", "FPC_VM"}
					var diags diag.Diagnostics
					flag := false
					for _, valid_value := range valid_values {
						if valid_value == value {
							flag = true
							break
						}
					}
					if !flag {
						diag := diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Wrong value",
							Detail:   fmt.Sprintf("Invalid status: %v. Valid values: %v", value, valid_values),
						}
						diags = append(diags, diag)
					}
					return diags
				},
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				// ExactlyOneOf: []string{"ACTIVE", "DISABLED"},
				Description: `Configuration status, must be one of the following options:
				ACTIVE: Enable a configuration;
				DISABLED: Disable a configuration.`,
				ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					value := v.(string)
					var diags diag.Diagnostics
					if value != "ACTIVE" && value != "DISABLED" {
						diag := diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Wrong value",
							Detail:   fmt.Sprintf("Invalid status: %v. Valid values: ACTIVE or DISABLED", value),
						}
						diags = append(diags, diag)
					}
					return diags
				},
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
		},
	}
}

func resourceConfigCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectConfig(d, "create")
	if err != nil {
		return fmt.Errorf("Error creating Config resource while getting object: %v", err)
	}

	o, err := c.CreateConfig(obj)

	if err != nil {
		return fmt.Errorf("Error creating Config resource: %v", err)
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
				return fmt.Errorf("Error creating Config resource while getting object: %v", err)
			}

			var op string
			if d.Get("status").(string) == "ACTIVE" {
				op = "enable"
			} else {
				op = "disable"
			}

			o, err = c.UpdateConfigStatus(obj, op)
			if err != nil {
				return fmt.Errorf("Error update Config status: %v", err)
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
		return fmt.Errorf("Error refresh Config resource: %v", err)
	}

	return nil
}

func resourceConfigUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectConfig(d, "update")
	if err != nil {
		return fmt.Errorf("Error updating Config resource while getting object: %v", err)
	}

	o, err := c.UpdateConfig(obj)
	if err != nil {
		return fmt.Errorf("Error updating Config resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["id"] != nil && o["id"] != "" {
		d.SetId(fmt.Sprintf("%v", o["id"]))
	} else {
		d.SetId("Config")
	}

	if st, ok := o["status"].(string); ok {
		if statusV, ok := d.GetOk("status"); ok && st != statusV.(string) {
			obj, err = getObjectConfig(d, "id")
			if err != nil {
				return fmt.Errorf("Error creating Config resource while getting object: %v", err)
			}

			var op string
			if d.Get("status").(string) == "ACTIVE" {
				op = "enable"
			} else {
				op = "disable"
			}

			o, err = c.UpdateConfigStatus(obj, op)
			if err != nil {
				return fmt.Errorf("Error update Config status: %v", err)
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
		return fmt.Errorf("Error refresh Config resource: %v", err)
	}

	return nil
}

func resourceConfigDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if d.Get("status").(string) != "DISABLED" {
		obj, err := getObjectConfig(d, "id")
		if err != nil {
			return fmt.Errorf("Error creating Config resource while getting object: %v", err)
		}

		o, err := c.UpdateConfigStatus(obj, "disable")
		if err != nil {
			return fmt.Errorf("Error update Config status: %v", err)
		}
		if st, ok := o["status"].(string); ok {
			if st != d.Get("status").(string) {
				log.Printf("[WARN] Could not update the status of Config %v", d.Id())
			}
		}

		err = refreshObjectConfig(d, o)
		if err != nil {
			return fmt.Errorf("Error refresh Config resource: %v", err)
		}
	}

	return nil
}

func resourceConfigRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if d.Get("program_serial_number") == "" {
		psn := importOptionChecking(m.(*FortiClient).Cfg, "program_serial_number")
		if err := d.Set("program_serial_number", psn); err != nil {
			return fmt.Errorf("Error set params program_serial_number: %v", err)
		}
	}
	obj, err := getObjectConfig(d, "read")
	if err != nil {
		return fmt.Errorf("Error reading Config while getting required parameters: %v", err)
	}

	o, err := c.ReadConfigsList(obj)
	if err != nil {
		return fmt.Errorf("Error reading Config resource: %v", err)
	}

	co, err := getConfigReadResponse(o, d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	err = refreshObjectConfig(d, co)
	if err != nil {
		return fmt.Errorf("Error reading Config resource from API: %v", err)
	}
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

func flattenConfigProgramSerialNumber(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenConfigName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenConfigProductType(v interface{}, d *schema.ResourceData) interface{} {
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

func flattenConfigStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenConfigParameters(v interface{}, d *schema.ResourceData, pt string) interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	tmp := make(map[string]interface{})
	for _, r := range l {
		i := r.(map[string]interface{})

		pType, cName, _ := convConfParsId2NameList(int(i["id"].(float64)))
		if pType != pt {
			log.Printf("[ERROR] Got incorrect parameter ID of Product Type %v, should be type %v", pType, pt)
			return nil
		}
		if cValue, ok := i["value"]; ok {
			if cName == "fortiguard_services" || cName == "cloud_services" {
				if argList, ok := tmp[cName]; ok {
					tmp[cName] = append(argList.([]interface{}), cValue)
				} else {
					tmp[cName] = []interface{}{cValue}
				}
			} else {
				tmp[cName] = cValue.(string)
			}
		}

	}
	result = append(result, tmp)

	return result
}

func refreshObjectConfig(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("program_serial_number", flattenConfigProgramSerialNumber(o["programSerialNumber"], d)); err != nil {
		if !fortiAPIPatch(o["programSerialNumber"]) {
			return fmt.Errorf("Error reading program_serial_number: %v", err)
		}
	}

	if err = d.Set("name", flattenConfigName(o["name"], d)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("product_type", flattenConfigProductType(o["productType"], d)); err != nil {
		if !fortiAPIPatch(o["productType"]) {
			return fmt.Errorf("Error reading product_type: %v", err)
		}
	}

	if err = d.Set("status", flattenConfigStatus(o["status"], d)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	// initialize product variables. This can fix the problem of inconsistent output.
	pTypeList := []string{"fgt_vm_bundle", "fmg_vm", "fwb_vm", "fgt_vm_lcs", "faz_vm", "fpc_vm"}
	for _, type_name := range pTypeList {
		empty_interface := make([]map[string]interface{}, 0)
		d.Set(type_name, empty_interface)
	}

	pType := strings.ToLower(d.Get("product_type").(string))
	if err = d.Set(pType, flattenConfigParameters(o["parameters"], d, pType)); err != nil {
		if !fortiAPIPatch(o["parameters"]) {
			return fmt.Errorf("Error reading %v: %v", pType, err)
		}
	}

	return nil
}

func expandConfigProgramSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandConfigName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandConfigProductType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	var typeId interface{}
	typeId = convProductTypeName2Id(v.(string))
	if typeId == 0 {
		err := fmt.Errorf("product_type invalid: %v, should be one of [FGT_VM_Bundle, FMG_VM, FWB_VM, FGT_VM_LCS]", v.(string))
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
				tmp := make(map[string]interface{})
				tmp["id"] = ckId
				tmp["value"] = csv
				result = append(result, tmp)
			}
		} else {
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
		if v, ok := d.GetOk("program_serial_number"); ok {
			t, err := expandConfigProgramSerialNumber(d, v, "program_serial_number")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["programSerialNumber"] = t
			}
		}
	}

	if rType == "create" || rType == "update" {
		if v, ok := d.GetOk("name"); ok {
			t, err := expandConfigName(d, v, "name")
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["name"] = t
			}
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
