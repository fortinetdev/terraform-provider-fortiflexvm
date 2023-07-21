// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Get list of Groups.

package fortiflexvm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
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

	// Send request
	o, err := c.ReadGroupsList(nil)
	if err != nil {
		return fmt.Errorf("Error describing GroupsList: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	// Update status
	err = dataSourceRefreshObjectGroupsList(d, o)
	if err != nil {
		return fmt.Errorf("Error describing GroupsList from API: %v", err)
	}

	d.SetId("GroupsList")

	return nil
}

func dataSourceRefreshObjectGroupsList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("groups", dataSourceFlattenGroupsListGroups(o["groups"])); err != nil {
		if !fortiAPIPatch(o["groups"]) {
			return fmt.Errorf("Error reading groups: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenGroupsListGroups(v interface{}) []map[string]interface{} {
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
		if _, ok := i["folderPath"]; ok {
			tmp["folder_path"] = i["folderPath"]
		}
		if _, ok := i["availableTokens"]; ok {
			tmp["available_tokens"] = i["availableTokens"]
		}
		if _, ok := i["usedTokens"]; ok {
			tmp["used_tokens"] = i["usedTokens"]
		}
		result = append(result, tmp)
	}

	return result
}
