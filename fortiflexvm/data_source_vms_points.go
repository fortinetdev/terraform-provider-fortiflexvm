// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Get point usage for Virtual Machines.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceVmsPoints() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVmsPointRead,
		Schema: map[string]*schema.Schema{
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
			"vms": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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

func dataSourceVmsPointRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVmsPoint(d)
	if err != nil {
		return fmt.Errorf("Error reading VmsPoint data source while getting required parameters: %v", err)
	}

	o, err := c.ReadVmsPoint(obj)
	if err != nil {
		return fmt.Errorf("Error describing VmsPoint: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectVmsPoint(d, o)
	if err != nil {
		return fmt.Errorf("Error describing VmsPoint from API: %v", err)
	}

	d.SetId("VmsPoint")

	return nil
}

func dataSourceFlattenVmsPointVms(v interface{}, d *schema.ResourceData) []map[string]interface{} {
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

		if _, ok := i["points"]; ok {
			tmp["points"] = dataSourceFlattenVmsPointVmsPoints(i["points"], d)
		}

		if _, ok := i["serialNumber"]; ok {
			tmp["serial_number"] = dataSourceFlattenVmsPointVmsSerialNumber(i["serialNumber"], d)
		}

		result = append(result, tmp)
	}

	return result
}

func dataSourceFlattenVmsPointVmsSerialNumber(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsPointVmsPoints(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceRefreshObjectVmsPoint(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vms", dataSourceFlattenVmsPointVms(o["vms"], d)); err != nil {
		if !fortiAPIPatch(o["vms"]) {
			return fmt.Errorf("Error reading vms: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenVmsPointFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiFlexVM Ver", " "), e)
}

func expandVmsPointConfigId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmsPointStartDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmsPointEndDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVmsPoint(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("config_id"); ok {
		t, err := expandVmsPointConfigId(d, v, "config_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configId"] = t
		}
	}

	if v, ok := d.GetOk("start_date"); ok {
		t, err := expandVmsPointStartDate(d, v, "start_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["startDate"] = t
		}
	}

	if v, ok := d.GetOk("end_date"); ok {
		t, err := expandVmsPointEndDate(d, v, "end_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endDate"] = t
		}
	}

	return &obj, nil
}
