// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Get list of Flex VM Configurations for a Program.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceConfigsList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceConfigsListRead,
		Schema: map[string]*schema.Schema{
			"program_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"configs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"program_serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"product_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"fgt_vm_bundle": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_size": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_pkg": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"vdom_num": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"fmg_vm": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"managed_dev": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"adom_num": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
						"fwb_vm": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_size": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_pkg": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"fgt_vm_lcs": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cpu_size": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"fortiguard_services": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
									"support_service": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"vdom_num": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"cloud_services": &schema.Schema{
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"faz_vm": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"daily_storage": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"adom_num": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
									"support_service": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
						"fpc_vm": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"managed_dev": &schema.Schema{
										Type:     schema.TypeInt,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceConfigsListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectConfigsList(d)
	if err != nil {
		return fmt.Errorf("Error reading ConfigsList data source while getting required parameters: %v", err)
	}

	o, err := c.ReadConfigsList(obj)
	if err != nil {
		return fmt.Errorf("Error describing ConfigsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectConfigsList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing ConfigsList from API: %v", err)
	}

	d.SetId("ConfigsList")

	return nil
}

func dataSourceFlattenConfigsListConfigs(v interface{}, d *schema.ResourceData) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		if _, ok := i["id"]; ok {
			tmp["id"] = dataSourceFlattenConfigsListConfigsId(i["id"], d)
		}

		if _, ok := i["programSerialNumber"]; ok {
			tmp["program_serial_number"] = dataSourceFlattenConfigsListConfigsProgramSerialNumber(i["programSerialNumber"], d)
		}

		if _, ok := i["name"]; ok {
			tmp["name"] = dataSourceFlattenConfigsListConfigsName(i["name"], d)
		}

		if _, ok := i["productType"]; ok {
			tmp["product_type"] = dataSourceFlattenConfigsListConfigsProductType(i["productType"], d)
		}

		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenConfigsListConfigsStatus(i["status"], d)
		}

		if tmp["product_type"] == "FGT_VM_Bundle" {
			if _, ok := i["parameters"]; ok {
				tmp["fgt_vm_bundle"] = dataSourceFlattenConfigsListConfigsParameters(i["parameters"], d, "fgt_vm_bundle")
			}
		} else if tmp["product_type"] == "FMG_VM" {
			if _, ok := i["parameters"]; ok {
				tmp["fmg_vm"] = dataSourceFlattenConfigsListConfigsParameters(i["parameters"], d, "fmg_vm")
			}
		} else if tmp["product_type"] == "FWB_VM" {
			if _, ok := i["parameters"]; ok {
				tmp["fwb_vm"] = dataSourceFlattenConfigsListConfigsParameters(i["parameters"], d, "fwb_vm")
			}
		} else if tmp["product_type"] == "FGT_VM_LCS" {
			if _, ok := i["parameters"]; ok {
				tmp["fgt_vm_lcs"] = dataSourceFlattenConfigsListConfigsParameters(i["parameters"], d, "fgt_vm_lcs")
			}
		} else if tmp["product_type"] == "FAZ_VM" {
			if _, ok := i["parameters"]; ok {
				tmp["faz_vm"] = dataSourceFlattenConfigsListConfigsParameters(i["parameters"], d, "faz_vm")
			}
		} else if tmp["product_type"] == "FPC_VM" {
			if _, ok := i["parameters"]; ok {
				tmp["fpc_vm"] = dataSourceFlattenConfigsListConfigsParameters(i["parameters"], d, "fpc_vm")
			}
		}

		result = append(result, tmp)
	}

	return result
}

func dataSourceFlattenConfigsListConfigsId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenConfigsListConfigsProgramSerialNumber(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenConfigsListConfigsName(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenConfigsListConfigsProductType(v interface{}, d *schema.ResourceData) interface{} {
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

func dataSourceFlattenConfigsListConfigsStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenConfigsListConfigsParameters(v interface{}, d *schema.ResourceData, pt string) interface{} {
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

		pType, cName, dataType := convConfParsId2NameList(int(i["id"].(float64)))
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
				switch dataType {
				case "int":
					tmp[cName], _ = strconv.Atoi((cValue.(string)))
				case "string":
					tmp[cName] = cValue.(string)
				default:
					tmp[cName] = cValue.(string)
				}

			}
		}

	}
	result = append(result, tmp)

	return result
}

func dataSourceRefreshObjectConfigsList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("configs", dataSourceFlattenConfigsListConfigs(o["configs"], d)); err != nil {
		if !fortiAPIPatch(o["configs"]) {
			return fmt.Errorf("Error reading configs: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenConfigsListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiFlexVM Ver", " "), e)
}

func expandConfigsListProgramSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectConfigsList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("program_serial_number"); ok {
		t, err := expandConfigsListProgramSerialNumber(d, v, "program_serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["programSerialNumber"] = t
		}
	}

	return &obj, nil
}
