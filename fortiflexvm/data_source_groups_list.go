// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Get list of Flex VM Groups.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceGroupsList() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceGroupsListRead,
		Schema: map[string]*schema.Schema{
			"groups": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"folder_path": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"available_tokens": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"used_tokens": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceGroupsListRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadGroupsList(nil)
	if err != nil {
		return fmt.Errorf("Error describing GroupsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectGroupsList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing GroupsList from API: %v", err)
	}

	d.SetId("GroupsList")

	return nil
}

func dataSourceFlattenGroupsListGroups(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "folder_path"
		if _, ok := i["folderPath"]; ok {
			tmp["folder_path"] = dataSourceFlattenGroupsListGroupsFolderPath(i["folderPath"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "available_tokens"
		if _, ok := i["availableTokens"]; ok {
			tmp["available_tokens"] = dataSourceFlattenGroupsListGroupsAvailableTokens(i["availableTokens"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "used_tokens"
		if _, ok := i["usedTokens"]; ok {
			tmp["used_tokens"] = dataSourceFlattenGroupsListGroupsUsedTokens(i["usedTokens"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenGroupsListGroupsFolderPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenGroupsListGroupsAvailableTokens(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenGroupsListGroupsUsedTokens(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectGroupsList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("groups", dataSourceFlattenGroupsListGroups(o["groups"], d, "groups")); err != nil {
		if !fortiAPIPatch(o["groups"]) {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenGroupsListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiFlexVM Ver", " "), e)
}
