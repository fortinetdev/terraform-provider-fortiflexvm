// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Get list of Flex VM Programs.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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

	o, err := c.ReadProgramsList(nil)
	if err != nil {
		return fmt.Errorf("Error describing ProgramsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectProgramsList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing ProgramsList from API: %v", err)
	}

	d.SetId("ProgramsList")

	return nil
}

func dataSourceFlattenProgramsListPrograms(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})

		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "account_id"
		if _, ok := i["accountId"]; ok {
			tmp["account_id"] = dataSourceFlattenProgramsListProgramsAccountId(i["accountId"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "end_date"
		if _, ok := i["endDate"]; ok {
			tmp["end_date"] = dataSourceFlattenProgramsListProgramsEndDate(i["endDate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "has_support_coverage"
		if _, ok := i["hasSupportCoverage"]; ok {
			tmp["has_support_coverage"] = dataSourceFlattenProgramsListProgramsHasSupportCoverage(i["hasSupportCoverage"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial_number"
		if _, ok := i["serialNumber"]; ok {
			tmp["serial_number"] = dataSourceFlattenProgramsListProgramsSerialNumber(i["serialNumber"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "start_date"
		if _, ok := i["startDate"]; ok {
			tmp["start_date"] = dataSourceFlattenProgramsListProgramsStartDate(i["startDate"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenProgramsListProgramsAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenProgramsListProgramsEndDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenProgramsListProgramsHasSupportCoverage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenProgramsListProgramsSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenProgramsListProgramsStartDate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectProgramsList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("programs", dataSourceFlattenProgramsListPrograms(o["programs"], d, "programs")); err != nil {
		if !fortiAPIPatch(o["programs"]) {
			return fmt.Errorf("Error reading programs: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenProgramsListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiFlexVM Ver", " "), e)
}
