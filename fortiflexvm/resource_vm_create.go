// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Create one or more VMs based on a Configuration.

package fortiflexvm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVmsCreate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmsCreateCreate,
		Read:   resourceVmsCreateRead,
		Update: resourceVmsCreateRead,
		Delete: resourceVmsCreateRead,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"vm_count": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"end_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}

func resourceVmsCreateCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectVmsCreate(d)
	if err != nil {
		return fmt.Errorf("Error creating Vms resource while getting object: %v", err)
	}

	_, err = c.CreateVms(obj)

	if err != nil {
		return fmt.Errorf("Error creating Vms resource: %v", err)
	}

	d.SetId("VmsCreate")

	return nil
}

func resourceVmsCreateRead(d *schema.ResourceData, m interface{}) error {
	return nil
}

func expandVmsCreateConfigId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmsCreateCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmsCreateDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmsCreateEndDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmsCreateFolderPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVmsCreate(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("config_id"); ok {
		t, err := expandVmsCreateConfigId(d, v, "config_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configId"] = t
		}
	}

	if v, ok := d.GetOk("vm_count"); ok {
		t, err := expandVmsCreateCount(d, v, "vm_count")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["count"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandVmsCreateDescription(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("end_date"); ok {
		t, err := expandVmsCreateEndDate(d, v, "end_date")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["endDate"] = t
		}
	}

	if v, ok := d.GetOk("folder_path"); ok {
		t, err := expandVmsCreateFolderPath(d, v, "folder_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["folderPath"] = t
		}
	}

	return &obj, nil
}
