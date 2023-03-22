// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Update a VM's setting.

package fortiflexvm

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceVmUpdate() *schema.Resource {
	return &schema.Resource{
		Create: resourceVmUpdateUpdate,
		Read:   resourceVmUpdateRead,
		Update: resourceVmUpdateUpdate,
		Delete: resourceVmUpdateDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"end_date": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ValidateDiagFunc: func(v interface{}, p cty.Path) diag.Diagnostics {
					value := v.(string)
					var diags diag.Diagnostics
					if value != "ACTIVE" && value != "STOPPED" {
						diag := diag.Diagnostic{
							Severity: diag.Error,
							Summary:  "Wrong value",
							Detail:   fmt.Sprintf("Invalid status: %v. Valid values: ACTIVE or STOPPED", value),
						}
						diags = append(diags, diag)
					}
					return diags
				},
			},
			"regenerate_token": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func getVMInformation(d *schema.ResourceData, m interface{}, serial_number string) (map[string]interface{}, error) {
	c := m.(*FortiClient).Client
	target := make(map[string]interface{})
	obj, _, err := getObjectVmUpdate(d, "read")
	if err != nil {
		return target, fmt.Errorf("Error reading VmUpdate while getting required parameters: %v", err)
	}
	o, err := c.ReadVmsList(obj)
	if err != nil {
		return target, fmt.Errorf("Error reading VmUpdate resource: %v", err)
	}
	target, err = getVmUpdateReadResponse(o, d.Get("serial_number").(string))
	return target, nil
}

func updateVMStatus(d *schema.ResourceData, m interface{}, op string) error {
	c := m.(*FortiClient).Client
	obj, _, err := getObjectVmUpdate(d, "")
	if err != nil {
		return fmt.Errorf("Error creating VmUpdate resource while getting object: %v", err)
	}

	o, err := c.UpdateVmUpdateStatus(obj, op)
	if err != nil {
		return fmt.Errorf("Error update VmUpdate status: %v", err)
	}
	if st, ok := o["status"].(string); ok {
		if st != d.Get("status").(string) {
			log.Printf("[WARN] Could not update the status of VmUpdate %v", d.Id())
		}
	}
	return nil
}

func resourceVmUpdateUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, needUpdate, err := getObjectVmUpdate(d, "update")
	if err != nil {
		return fmt.Errorf("Error updating VmUpdate resource while getting object: %v", err)
	}
	var o map[string]interface{}
	var statusRV string
	if needUpdate {
		// We can't update if the VM status is PENDING. Need to consider the scenario when VM status is STOPPED or EXPIRED
		target, err := getVMInformation(d, m, d.Get("serial_number").(string))
		log.Printf("target: %v", target)
		if err != nil {
			return err
		}
		current_status := target["status"]

		switch current_status {
		case "PENDING":
			return fmt.Errorf(`VM's current status is PENDING. 
You can't update a VM if its current status is "PENDING".
Please use the VM token to activate a virtual machine before using this API.`)
		case "EXPIRED":
			if statusCV, ok := d.GetOk("status"); ok && statusCV == "ACTIVE" {
				err = updateVMStatus(d, m, "reactivate")
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("Can't update the VM. Its current status is EXPIRED. Please set `status = ACTIVE` and specify `end_date`.")
			}
		case "STOPPED":
			if statusCV, ok := d.GetOk("status"); ok && statusCV == "ACTIVE" {
				err = updateVMStatus(d, m, "reactivate")
				if err != nil {
					return err
				}
			} else {
				return fmt.Errorf("Can't update the VM. Its current status is STOPPED. Please set `status = ACTIVE`.")
			}
		}

		o, err = c.UpdateVmUpdate(obj)
		if err != nil {
			return fmt.Errorf("Error updating VmUpdate resource: %v", err)
		}

		log.Printf(strconv.Itoa(c.Retries))
		if st, ok := o["status"].(string); ok {
			statusRV = st
		} else {
			log.Printf("[WARN] Could not get status from HTTP response")
		}
	}

	if statusCV, ok := d.GetOk("status"); ok && d.HasChange("status") && statusRV != statusCV.(string) {
		obj, _, err = getObjectVmUpdate(d, "")
		if err != nil {
			return fmt.Errorf("Error creating VmUpdate resource while getting object: %v", err)
		}

		var op string
		if statusCV.(string) == "ACTIVE" {
			op = "reactivate"
		} else {
			op = "stop"
		}

		o, err = c.UpdateVmUpdateStatus(obj, op)
		if err != nil {
			var error_result map[string]interface{}
			json.Unmarshal([]byte(err.Error()), &error_result)
			log.Printf("Error: %v", error_result)
			if error_result["message"] == "Unable to update VM from current status." {
				// This error happens when you try to set ACTIVE -> ACTIVE, or STOPPED -> STOPPED and your `config_id` is not specified.
				// By specifying `config_id`, this code would know the VM's current status in advance and avoid setting ACTIVE -> ACTIVE or STOPPED -> STOPPED.
				return fmt.Errorf("Please specify `config_id`")
			} else {
				return fmt.Errorf("Error update VmUpdate status: %v", err)
			}
		}
		if st, ok := o["status"].(string); ok {
			if st != d.Get("status").(string) {
				log.Printf("[WARN] Could not update the status of VmUpdate %v", d.Id())
			}
		}
	}

	if v, ok := d.GetOk("regenerate_token"); ok && v.(bool) {
		obj, _, err = getObjectVmUpdate(d, "")
		o, err = c.UpdateVmUpdateRegenerateToken(obj)
		if err != nil {
			return fmt.Errorf("Error update VmUpdate to regenerate token: %v", err)
		}
	}

	if o["serialNumber"] != nil && o["serialNumber"] != "" {
		d.SetId(fmt.Sprintf("%v", o["serialNumber"]))
	} else {
		d.SetId("VmUpdate")
	}

	err = refreshObjectVmUpdate(d, o)
	if err != nil {
		return fmt.Errorf("Error refresh VmUpdate resource: %v", err)
	}

	return nil
}

func resourceVmUpdateDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if d.Get("status").(string) == "ACTIVE" {
		obj, _, err := getObjectVmUpdate(d, "id")
		if err != nil {
			return fmt.Errorf("Error creating VmUpdate resource while getting object: %v", err)
		}

		o, err := c.UpdateVmUpdateStatus(obj, "stop")
		if err != nil {
			return fmt.Errorf("Error update VmUpdate status: %v", err)
		}
		if st, ok := o["status"].(string); ok {
			if st != d.Get("status").(string) {
				log.Printf("[WARN] Could not update the status of VmUpdate %v", d.Id())
			}
		}

		err = refreshObjectVmUpdate(d, o)
		if err != nil {
			return fmt.Errorf("Error refresh VmUpdate resource: %v", err)
		}
	}

	return nil
}

func resourceVmUpdateRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	if d.Get("config_id") == "" {
		psn := importOptionChecking(m.(*FortiClient).Cfg, "config_id")
		if err := d.Set("config_id", psn); err != nil {
			return fmt.Errorf("Error set params config_id: %v", err)
		}
	}
	obj, _, err := getObjectVmUpdate(d, "read")
	if err != nil {
		return fmt.Errorf("Error reading VmUpdate while getting required parameters: %v", err)
	}

	o, err := c.ReadVmsList(obj)
	if err != nil {
		return fmt.Errorf("Error reading VmUpdate resource: %v", err)
	}

	co, err := getVmUpdateReadResponse(o, d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	err = refreshObjectVmUpdate(d, co)
	if err != nil {
		return fmt.Errorf("Error reading VmUpdate resource from API: %v", err)
	}
	return nil
}

func getVmUpdateReadResponse(o map[string]interface{}, mkey string) (map[string]interface{}, error) {
	var err error
	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", mkey)
		err = fmt.Errorf("Response is nil")
		return nil, err
	}

	if vms, ok := o["vms"]; ok {
		if vmsList, ok := vms.([]interface{}); ok {
			for _, vm := range vmsList {
				if vmMap, ok := vm.(map[string]interface{}); ok {
					cId := fmt.Sprintf("%v", vmMap["serialNumber"])
					if cId == mkey {
						return vmMap, nil
					}
				}
			}
		}
	}
	err = fmt.Errorf("VmUpdate not been created")
	return nil, err
}

func flattenVmUpdateConfigId(v interface{}, d *schema.ResourceData) interface{} {
	var rst interface{}
	switch v.(type) {
	case float64:
		rst = int(v.(float64))
	default:
		rst = v
	}
	return rst
}

func flattenVmUpdateDescription(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenVmUpdateEndDate(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func flattenVmUpdateStatus(v interface{}, d *schema.ResourceData) interface{} {
	return v
}

func refreshObjectVmUpdate(d *schema.ResourceData, o map[string]interface{}) error {
	var err error
	if o == nil {
		return nil
	}

	if err = d.Set("config_id", flattenVmUpdateConfigId(o["configId"], d)); err != nil {
		if !fortiAPIPatch(o["configId"]) {
			return fmt.Errorf("Error reading config_id: %v", err)
		}
	}

	if err = d.Set("description", flattenVmUpdateDescription(o["description"], d)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("end_date", flattenVmUpdateEndDate(o["endDate"], d)); err != nil {
		if !fortiAPIPatch(o["endDate"]) {
			return fmt.Errorf("Error reading end_date: %v", err)
		}
	}

	if err = d.Set("status", flattenVmUpdateStatus(o["status"], d)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	d.Set("regenerate_token", false)

	return nil
}

func expandVmUpdateSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmUpdatecConfigId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmUpdatecDescription(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandVmUpdatecEndDate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectVmUpdate(d *schema.ResourceData, rType string) (*map[string]interface{}, bool, error) {
	obj := make(map[string]interface{})

	var needUpdate bool

	if rType != "read" {
		if v, ok := d.GetOk("serial_number"); ok {
			t, err := expandVmUpdateSerialNumber(d, v, "serial_number")
			if err != nil {
				return &obj, needUpdate, err
			} else if t != nil {
				obj["serialNumber"] = t
			}
		}
	}

	if rType == "update" || rType == "read" {
		if v, ok := d.GetOk("config_id"); ok {
			t, err := expandVmUpdatecConfigId(d, v, "config_id")
			if err != nil {
				return &obj, needUpdate, err
			} else if t != nil {
				obj["configId"] = t
				if d.HasChanges("config_id") {
					needUpdate = true
				}
			}
		}
	}

	if rType == "update" {
		if v, ok := d.GetOk("description"); ok && d.HasChanges("description") {
			t, err := expandVmUpdatecDescription(d, v, "description")
			if err != nil {
				return &obj, needUpdate, err
			} else if t != nil {
				obj["description"] = t
				needUpdate = true
			}
		}

		if v, ok := d.GetOk("end_date"); ok && d.HasChanges("end_date") {
			t, err := expandVmUpdatecEndDate(d, v, "end_date")
			if err != nil {
				return &obj, needUpdate, err
			} else if t != nil {
				obj["endDate"] = t
				needUpdate = true
			}
		}
	}

	return &obj, needUpdate, nil
}
