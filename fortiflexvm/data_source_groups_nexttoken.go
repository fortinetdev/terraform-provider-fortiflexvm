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
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"entitlements": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
	c.Retries = 1

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

	if len(request_obj) == 0 {
		return fmt.Errorf("Either config_id or folder_path is required.")
	}

	// Send request
	o, err := c.ReadGroupsNexttoken(&request_obj)
	if err != nil {
		return fmt.Errorf("Error describing GroupsNexttoken: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	// Update status
	err = dataSourceRefreshObjectGroupsNexttoken(d, o)
	if err != nil {
		return fmt.Errorf("Error describing GroupsNexttoken from API: %v", err)
	}

	resource_id := fmt.Sprintf("%v.%v", config_id, folder_path)
	d.SetId(resource_id)

	return nil
}

func dataSourceRefreshObjectGroupsNexttoken(d *schema.ResourceData, o map[string]interface{}) error {
	var err error
	if err = d.Set("entitlements", dataSourceFlattenGroupsNexttokenEntitlementsElement(o["entitlements"])); err != nil {
		if !fortiAPIPatch(o["entitlements"]) {
			return fmt.Errorf("Error reading entitlements: %v", err)
		}
	}
	return nil
}

func dataSourceFlattenGroupsNexttokenEntitlementsElement(v interface{}) []map[string]interface{} {
	if v == nil {
		return nil
	}
	result := make([]map[string]interface{}, 0, 1)
	tmp := make(map[string]interface{})
	i := v.(map[string]interface{})
	if _, ok := i["configId"]; ok {
		tmp["config_id"] = i["configId"]
	}
	if _, ok := i["description"]; ok {
		tmp["description"] = i["description"]
	}
	if _, ok := i["serialNumber"]; ok {
		tmp["serial_number"] = i["serialNumber"]
	}
	if _, ok := i["startDate"]; ok {
		tmp["start_date"] = i["startDate"]
	}
	if _, ok := i["endDate"]; ok {
		tmp["end_date"] = i["endDate"]
	}
	if _, ok := i["status"]; ok {
		tmp["status"] = i["status"]
	}
	if _, ok := i["token"]; ok {
		tmp["token"] = i["token"]
	}
	if _, ok := i["tokenStatus"]; ok {
		tmp["token_status"] = i["tokenStatus"]
	}
	result = append(result, tmp)
	return result
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
		if _, ok := i["configId"]; ok {
			tmp["config_id"] = i["configId"]
		}
		if _, ok := i["description"]; ok {
			tmp["description"] = i["description"]
		}
		if _, ok := i["serialNumber"]; ok {
			tmp["serial_number"] = i["serialNumber"]
		}
		if _, ok := i["startDate"]; ok {
			tmp["start_date"] = i["startDate"]
		}
		if _, ok := i["endDate"]; ok {
			tmp["end_date"] = i["endDate"]
		}
		if _, ok := i["status"]; ok {
			tmp["status"] = i["status"]
		}
		if _, ok := i["token"]; ok {
			tmp["token"] = i["token"]
		}
		if _, ok := i["tokenStatus"]; ok {
			tmp["token_status"] = i["tokenStatus"]
		}
		result = append(result, tmp)
	}
	return result
}
