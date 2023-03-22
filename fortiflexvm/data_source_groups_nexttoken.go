// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Get next available (unused) token.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vms": &schema.Schema{
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

	obj, err := getObjectGroupsNexttoken(d)
	if err != nil {
		return fmt.Errorf("Error reading GroupsNexttoken data source while getting required parameters: %v", err)
	}

	o, err := c.ReadGroupsNexttoken(obj)
	if err != nil {
		return fmt.Errorf("Error describing GroupsNexttoken: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectGroupsNexttoken(d, o)
	if err != nil {
		return fmt.Errorf("Error describing GroupsNexttoken from API: %v", err)
	}

	d.SetId("GroupsNexttoken")

	return nil
}

func dataSourceFlattenGroupsNexttokenVms(v interface{}, d *schema.ResourceData) []map[string]interface{} {
	if v == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, 1)

	tmp := make(map[string]interface{})
	i := v.(map[string]interface{})

	if _, ok := i["configId"]; ok {
		tmp["config_id"] = dataSourceFlattenGroupsNexttokenVmsConfigId(i["configId"], d)
	}

	if _, ok := i["description"]; ok {
		tmp["description"] = dataSourceFlattenGroupsNexttokenVmsDescription(i["description"], d)
	}

	if _, ok := i["serialNumber"]; ok {
		tmp["serial_number"] = dataSourceFlattenGroupsNexttokenVmsSerialNumber(i["serialNumber"], d)
	}

	if _, ok := i["startDate"]; ok {
		tmp["start_date"] = dataSourceFlattenGroupsNexttokenVmsStartDate(i["startDate"], d)
	}

	if _, ok := i["endDate"]; ok {
		tmp["end_date"] = dataSourceFlattenGroupsNexttokenVmsEndDate(i["endDate"], d)
	}

	if _, ok := i["status"]; ok {
		tmp["status"] = dataSourceFlattenGroupsNexttokenVmsStatus(i["status"], d)
	}

	if _, ok := i["token"]; ok {
		tmp["token"] = dataSourceFlattenGroupsNexttokenVmsToken(i["token"], d)
	}

	if _, ok := i["tokenStatus"]; ok {
		tmp["token_status"] = dataSourceFlattenGroupsNexttokenVmsTokenStatus(i["tokenStatus"], d)
	}

	result = append(result, tmp)

	return result
}

func dataSourceFlattenGroupsNexttokenVmsConfigId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenGroupsNexttokenVmsDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenGroupsNexttokenVmsSerialNumber(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenGroupsNexttokenVmsStartDate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenGroupsNexttokenVmsEndDate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenGroupsNexttokenVmsStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenGroupsNexttokenVmsToken(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenGroupsNexttokenVmsTokenStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceRefreshObjectGroupsNexttoken(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vms", dataSourceFlattenGroupsNexttokenVms(o["vms"], d)); err != nil {
		if !fortiAPIPatch(o["vms"]) {
			return fmt.Errorf("Error reading vms: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenGroupsNexttokenFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiFlexVM Ver", " "), e)
}

func expandGroupsNexttokenConfigId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandGroupsNexttokenFolderPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectGroupsNexttoken(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("config_id"); ok {
		t, err := expandGroupsNexttokenConfigId(d, v, "config_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configId"] = t
		}
	}

	if v, ok := d.GetOk("folder_path"); ok {
		t, err := expandGroupsNexttokenFolderPath(d, v, "folder_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folderPath"] = t
		}
	}

	return &obj, nil
}
