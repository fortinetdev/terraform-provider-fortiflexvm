// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Create and update a VM based on a Configuration.

package fortiflexvm

import (
	"context"
	"fmt"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceEntitlementsHW() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEntitlementsHWCreate,
		ReadContext:   resourceEntitlementsHWRead,
		UpdateContext: resourceEntitlementsHWUpdate,
		DeleteContext: resourceEntitlementsHWDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
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
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"start_date": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceEntitlementsHWCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*FortiClient).Client

	// Send request
	obj := make(map[string]interface{})
	obj["configId"] = d.Get("config_id")
	serial_number := d.Get("serial_number").(string)
	serial_number_list := []string{serial_number}
	obj["serialNumbers"] = serial_number_list
	if v, ok := d.GetOk("end_date"); ok {
		obj["endDate"] = v
	}
	target_entitlement, err := c.CreateEntitlementsHW(&obj)
	if err != nil {
		return diag.FromErr(err)
	}

	_, set_description := d.GetOk("description")
	status, set_status := d.GetOk("status")

	// Update status
	diags = refreshObjectEntitlementsHW(d, target_entitlement)
	if diags.HasError() {
		return diags
	}

	resource_id := fmt.Sprintf("%v.%v", target_entitlement["serialNumber"], target_entitlement["configId"])
	d.SetId(resource_id)

	if set_description || (set_status && status.(string) == "STOPPED") {
		diags = resourceEntitlementsHWUpdate(ctx, d, m)
	}

	return diags
}

func resourceEntitlementsHWRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Send request
	target_entitlement, diags := getEntitlementFromId(d.Id(), m)
	if diags.HasError() {
		return diags
	}

	// Update status
	diags = refreshObjectEntitlementsHW(d, target_entitlement)
	return diags
}

func resourceEntitlementsHWUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	var err error
	c := m.(*FortiClient).Client

	// Get ID
	serial_number, previous_config_id, diags := splitID(d.Id())
	if diags.HasError() {
		return diags
	}

	// Check status first
	target_entitlement, diags := getEntitlementFromId(d.Id(), m)
	if diags.HasError() {
		return diags
	}

	// Check status
	current_status := target_entitlement["status"].(string)
	set_status := d.Get("status")
	if set_status != "" && current_status != set_status {
		if set_status == "ACTIVE" {
			target_entitlement, err = changeVMStatus(target_entitlement["serialNumber"].(string), "reactivate", m)
		} else if set_status == "STOPPED" {
			target_entitlement, err = changeVMStatus(target_entitlement["serialNumber"].(string), "stop", m)
		}
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Send update request
	obj := make(map[string]interface{})
	obj["serialNumber"] = serial_number
	obj["configId"] = previous_config_id
	if v, ok := d.GetOk("config_id"); ok { // if specify new config ID, use it
		obj["configId"] = v
	}
	// TODO check it later
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v
	}
	if v, ok := d.GetOk("end_date"); ok {
		err_flag := false
		current_end_date, err := time.Parse(time.RFC3339, target_entitlement["endDate"].(string))
		if err != nil {
			err_flag = true
		}
		user_end_date, err := time.Parse(time.RFC3339, v.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Unable to parsing end_date, ignoring update end_date",
				Detail:   fmt.Sprintf("Unable to parsing %v, please check the format.", v),
			})
			return diags
		}
		if !err_flag && current_end_date.Before(user_end_date) {
			obj["endDate"] = v
		}
	}
	target_entitlement, err = c.UpdateVmUpdate(&obj)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update status
	diags = refreshObjectEntitlementsHW(d, target_entitlement)
	if diags.HasError() {
		return diags
	}
	resource_id := fmt.Sprintf("%v.%v", target_entitlement["serialNumber"], target_entitlement["configId"])
	d.SetId(resource_id)

	return diags
}

func resourceEntitlementsHWDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	// If entitlement is ACTIVE, stop it.
	if d.Get("status").(string) == "ACTIVE" {
		// Get ID
		serial_number, _, diags := splitID(d.Id())
		if diags.HasError() {
			return diags
		}
		// Send delete request
		_, err := changeVMStatus(serial_number, "stop", m)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// Update status
	d.SetId("")
	return diags
}

func refreshObjectEntitlementsHW(d *schema.ResourceData, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// can't set folder_path
	if value, ok := o["accountId"]; ok {
		d.Set("account_id", value)
	}
	if value, ok := o["configId"]; ok {
		d.Set("config_id", value)
	}
	if value, ok := o["description"]; ok {
		d.Set("description", value)
	}
	if value, ok := o["endDate"]; ok {
		d.Set("end_date", value)
	}
	if value, ok := o["serialNumber"]; ok {
		d.Set("serial_number", value)
	}
	if value, ok := o["status"]; ok {
		d.Set("status", value)
	}
	if value, ok := o["startDate"]; ok {
		d.Set("start_date", value)
	}
	return diags
}
