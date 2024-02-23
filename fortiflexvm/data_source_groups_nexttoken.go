// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Get next available (unused) token.

package fortiflexvm

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceGroupsNexttoken() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGroupsNexttokenRead,
		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"entitlements": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"config_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"start_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"end_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"token": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"token_status": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGroupsNexttokenRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	// Prepare data
	request_obj := make(map[string]interface{})
	config_id := "none"
	folder_path := "none"

	if v, ok := d.GetOk("config_id"); ok {
		config_id = strconv.Itoa(v.(int))
		request_obj["configId"] = v
	}
	if v, ok := d.GetOk("folder_path"); ok {
		folder_path = v.(string)
		request_obj["folderPath"] = v
	}
	if v, ok := d.GetOk("account_id"); ok {
		request_obj["accountId"] = v
	}
	if v, ok := d.GetOk("status"); ok {
		request_obj["status"] = v
	}
	if len(request_obj) == 0 {
		return fmt.Errorf("either config_id or folder_path is required")
	}

	// Send request
	o, err := c.ReadGroupsNexttoken(&request_obj)
	if err != nil {
		return fmt.Errorf("error describing GroupsNexttoken: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	// Update status
	err = dataSourceRefreshObjectGroupsNexttoken(d, o)
	if err != nil {
		return fmt.Errorf("error describing GroupsNexttoken from API: %v", err)
	}

	resource_id := fmt.Sprintf("%v.%v", config_id, folder_path)
	d.SetId(resource_id)

	return nil
}

func dataSourceRefreshObjectGroupsNexttoken(d *schema.ResourceData, o map[string]interface{}) error {
	var err error
	switch v := o["entitlements"].(type) {
	case map[string]interface{}:
		entitlements_list := make([]interface{}, 0, 1)
		entitlements_list = append(entitlements_list, o["entitlements"].(map[string]interface{}))
		if err = d.Set("entitlements", dataSourceFlattenGroupsNexttokenEntitlements(entitlements_list)); err != nil {
			if !fortiAPIPatch(o["entitlements"]) {
				return fmt.Errorf("error reading entitlements: %v", err)
			}
		}
	case []interface{}:
		if err = d.Set("entitlements", dataSourceFlattenGroupsNexttokenEntitlements(o["entitlements"])); err != nil {
			if !fortiAPIPatch(o["entitlements"]) {
				return fmt.Errorf("error reading entitlements: %v", err)
			}
		}
	default:
		fmt.Printf("Unsupported type: %T\n", v)
	}
	return nil
}

func dataSourceFlattenGroupsNexttokenEntitlements(v interface{}) []map[string]interface{} {
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
		if value, ok := i["accountId"]; ok {
			tmp["account_id"] = value
		}
		if value, ok := i["configId"]; ok {
			tmp["config_id"] = value
		}
		if value, ok := i["description"]; ok {
			tmp["description"] = value
		}
		if value, ok := i["serialNumber"]; ok {
			tmp["serial_number"] = value
		}
		if value, ok := i["startDate"]; ok {
			tmp["start_date"] = value
		}
		if value, ok := i["endDate"]; ok {
			tmp["end_date"] = value
		}
		if value, ok := i["status"]; ok {
			tmp["status"] = value
		}
		if value, ok := i["token"]; ok {
			tmp["token"] = value
		}
		if value, ok := i["tokenStatus"]; ok {
			tmp["token_status"] = value
		}
		result = append(result, tmp)
	}
	return result
}
