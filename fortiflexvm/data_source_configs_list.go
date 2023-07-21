// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Get list of configurations for a program.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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
						"fad_vm": &schema.Schema{
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
						"fgt_hw": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"device_model": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"service_pkg": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"addons": &schema.Schema{
										Type:     schema.TypeString,
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

	// Prepare data
	request_obj := make(map[string]interface{})
	program_serial_number := d.Get("program_serial_number").(string)
	request_obj["programSerialNumber"] = program_serial_number

	// Send request
	o, err := c.ReadConfigsList(&request_obj)
	if err != nil {
		return fmt.Errorf("Error describing ConfigsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	// Update status
	err = dataSourceRefreshObjectConfigsList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing ConfigsList from API: %v", err)
	}

	d.SetId(program_serial_number)

	return nil
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
			tmp["id"] = i["id"]
		}
		if _, ok := i["programSerialNumber"]; ok {
			tmp["program_serial_number"] = i["programSerialNumber"]
		}
		if _, ok := i["name"]; ok {
			tmp["name"] = i["name"]
		}
		if _, ok := i["status"]; ok {
			tmp["status"] = i["status"]
		}
		if _, ok := i["productType"]; ok {
			tmp["product_type"] = dataSourceFlattenConfigsListConfigsProductType(i["productType"])
			if _, ok := i["parameters"]; ok {
				product_type := tmp["product_type"].(string)
				product_type_lower := strings.ToLower(product_type)
				tmp[product_type_lower] = dataSourceFlattenConfigsListConfigsParameters(i["parameters"])
			}
		}
		result = append(result, tmp)
	}

	return result
}

func dataSourceFlattenConfigsListConfigsProductType(v interface{}) interface{} {
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

func dataSourceFlattenConfigsListConfigsParameters(v interface{}) interface{} {
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
