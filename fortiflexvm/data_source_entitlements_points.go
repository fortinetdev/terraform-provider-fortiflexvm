// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Get point usage for Virtual Machines.

package fortiflexvm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEntitlementsPoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEntitlementsPointRead,
		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"start_date": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"end_date": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
						"points": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"serial_number": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceEntitlementsPointRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	// Prepare data
	request_obj := make(map[string]interface{})
	config_id := d.Get("config_id").(int)
	start_date := d.Get("start_date").(string)
	end_date := d.Get("end_date").(string)

	request_obj["configId"] = config_id
	request_obj["startDate"] = start_date
	request_obj["endDate"] = end_date
	if v, ok := d.GetOk("account_id"); ok {
		request_obj["accountId"] = v
	}

	// Send request
	o, err := c.ReadEntitlementsPoint(&request_obj)
	if err != nil {
		return fmt.Errorf("error describing EntitlementsPoint: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	// Update status
	err = dataSourceRefreshObjectEntitlementsPoint(d, o)
	if err != nil {
		return fmt.Errorf("error describing EntitlementsPoint from API: %v", err)
	}

	resource_id := fmt.Sprintf("%v.%v.%v", config_id, start_date, end_date)
	d.SetId(resource_id)

	return nil
}

func dataSourceRefreshObjectEntitlementsPoint(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("entitlements", dataSourceFlattenEntitlementsPointEntitlements(o["entitlements"], d)); err != nil {
		if !fortiAPIPatch(o["entitlements"]) {
			return fmt.Errorf("error reading entitlements: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenEntitlementsPointEntitlements(v interface{}, d *schema.ResourceData) []map[string]interface{} {
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
		if value, ok := i["points"]; ok {
			tmp["points"] = value
		}
		if value, ok := i["serialNumber"]; ok {
			tmp["serial_number"] = value
		}
		result = append(result, tmp)
	}

	return result
}
