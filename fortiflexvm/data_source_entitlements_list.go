// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Get list of entitlements for a Configuration.

package fortiflexvm

import (
	"fmt"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceEntitlementsList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceEntitlementsListRead,
		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"program_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"token_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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

func dataSourceEntitlementsListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	// Prepare data
	request_obj := make(map[string]interface{})
	config_id := d.Get("config_id").(int)
	account_id := d.Get("account_id").(int)
	program_serial_number := d.Get("program_serial_number").(string)
	recource_id := ""
	if config_id == 0 && (account_id == 0 || program_serial_number == "") {
		return fmt.Errorf("either config_id or (account_id + program_serial_number) should be provided in request payload")
	}
	if config_id != 0 {
		recource_id = strconv.Itoa(config_id)
	} else {
		recource_id = fmt.Sprintf("%v.%v", account_id, program_serial_number)
	}

	if v, ok := d.GetOk("account_id"); ok {
		request_obj["accountId"] = v
	}
	if v, ok := d.GetOk("config_id"); ok {
		request_obj["configId"] = v
	}
	if v, ok := d.GetOk("description"); ok {
		request_obj["description"] = v
	}
	if v, ok := d.GetOk("program_serial_number"); ok {
		request_obj["programSerialNumber"] = v
	}
	if v, ok := d.GetOk("serial_number"); ok {
		request_obj["serialNumber"] = v
	}
	if v, ok := d.GetOk("status"); ok {
		request_obj["status"] = v
	}
	if v, ok := d.GetOk("token_status"); ok {
		request_obj["tokenStatus"] = v
	}

	// Send request
	o, err := c.ReadEntitlementsList(&request_obj)
	if err != nil {
		return fmt.Errorf("error describing EntitlementsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	// Update status
	err = dataSourceRefreshObjectEntitlementsList(d, o)
	if err != nil {
		return fmt.Errorf("error describing EntitlementsList from API: %v", err)
	}

	d.SetId(recource_id)

	return nil
}

func dataSourceRefreshObjectEntitlementsList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("entitlements", dataSourceFlattenEntitlementsListEntitlements(o["entitlements"])); err != nil {
		if !fortiAPIPatch(o["entitlements"]) {
			return fmt.Errorf("error reading entitlements: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenEntitlementsListEntitlements(v interface{}) []map[string]interface{} {
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
