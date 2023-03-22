// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Get list of Flex VM Virtual Machines for a Configuration.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceVmsList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceVmsListRead,
		Schema: map[string]*schema.Schema{
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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

func dataSourceVmsListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVmsList(d)
	if err != nil {
		return fmt.Errorf("Error reading VmsList data source while getting required parameters: %v", err)
	}

	o, err := c.ReadVmsList(obj)
	if err != nil {
		return fmt.Errorf("Error describing VmsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectVmsList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing VmsList from API: %v", err)
	}

	d.SetId("VmsList")

	return nil
}

func dataSourceFlattenVmsListVms(v interface{}, d *schema.ResourceData) []map[string]interface{} {
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
			tmp["config_id"] = dataSourceFlattenVmsListVmsConfigId(i["configId"], d)
		}

		if _, ok := i["description"]; ok {
			tmp["description"] = dataSourceFlattenVmsListVmsDescription(i["description"], d)
		}

		if _, ok := i["serialNumber"]; ok {
			tmp["serial_number"] = dataSourceFlattenVmsListVmsSerialNumber(i["serialNumber"], d)
		}

		if _, ok := i["startDate"]; ok {
			tmp["start_date"] = dataSourceFlattenVmsListVmsStartDate(i["startDate"], d)
		}

		if _, ok := i["endDate"]; ok {
			tmp["end_date"] = dataSourceFlattenVmsListVmsEndDate(i["endDate"], d)
		}

		if _, ok := i["status"]; ok {
			tmp["status"] = dataSourceFlattenVmsListVmsStatus(i["status"], d)
		}

		if _, ok := i["token"]; ok {
			tmp["token"] = dataSourceFlattenVmsListVmsToken(i["token"], d)
		}

		if _, ok := i["tokenStatus"]; ok {
			tmp["token_status"] = dataSourceFlattenVmsListVmsTokenStatus(i["tokenStatus"], d)
		}

		result = append(result, tmp)
	}

	return result
}

func dataSourceFlattenVmsListVmsConfigId(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsListVmsDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsListVmsSerialNumber(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsListVmsStartDate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsListVmsEndDate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsListVmsStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsListVmsToken(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceFlattenVmsListVmsTokenStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func dataSourceRefreshObjectVmsList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("vms", dataSourceFlattenVmsListVms(o["vms"], d)); err != nil {
		if !fortiAPIPatch(o["vms"]) {
			return fmt.Errorf("Error reading vms: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenVmsListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiFlexVM Ver", " "), e)
}

func expandVmsListConfigId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVmsList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("config_id"); ok {
		t, err := expandVmsListConfigId(d, v, "config_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configId"] = t
		}
	}

	return &obj, nil
}
