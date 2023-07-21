// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Get list of programs.

package fortiflexvm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceProgramsList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceProgramsListRead,
		Schema: map[string]*schema.Schema{
			"programs": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"end_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"has_support_coverage": &schema.Schema{
							Type:     schema.TypeBool,
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
					},
				},
			},
		},
	}
}

func dataSourceProgramsListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	// Send request
	o, err := c.ReadProgramsList(nil)
	if err != nil {
		return fmt.Errorf("Error describing ProgramsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	// Update status
	err = dataSourceRefreshObjectProgramsList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing ProgramsList from API: %v", err)
	}

	d.SetId("ProgramsList")

	return nil
}

func dataSourceRefreshObjectProgramsList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error
	if err = d.Set("programs", dataSourceFlattenProgramsListPrograms(o["programs"])); err != nil {
		if !fortiAPIPatch(o["programs"]) {
			return fmt.Errorf("Error reading programs: %v", err)
		}
	}
	return nil
}

func dataSourceFlattenProgramsListPrograms(v interface{}) []map[string]interface{} {
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
		if _, ok := i["accountId"]; ok {
			tmp["account_id"] = i["accountId"]
		}
		if _, ok := i["endDate"]; ok {
			tmp["end_date"] = i["endDate"]
		}
		if _, ok := i["hasSupportCoverage"]; ok {
			tmp["has_support_coverage"] = i["hasSupportCoverage"]
		}
		if _, ok := i["serialNumber"]; ok {
			tmp["serial_number"] = i["serialNumber"]
		}
		if _, ok := i["startDate"]; ok {
			tmp["start_date"] = i["startDate"]
		}
		result = append(result, tmp)
	}

	return result
}
