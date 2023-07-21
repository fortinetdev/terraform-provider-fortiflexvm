// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Create and update a VM based on a Configuration.

package fortiflexvm

import (
	"context"
	"fmt"

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
	c.Retries = 1

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
	c.Retries = 1

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

	// check status
	current_status := target_entitlement["status"].(string)
	set_status := ""
	if v, ok := d.GetOk("status"); ok {
		set_status = v.(string)
	}
	if current_status != "ACTIVE" {
		// active entitlements
		if set_status == "ACTIVE" && current_status == "STOPPED" {
			target_entitlement, err = changeVMStatus(target_entitlement["serialNumber"].(string), "reactivate", m)
			if err != nil {
				return diag.FromErr(err)
			}
		} else { // can't update
			switch current_status {
			case "PENDING":
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to Update fortiflexvm_entitlements_hardware",
					Detail:   fmt.Sprintf("Current entitlement status is PENDING. Please use the VM token to activate a virtual machine before using this API."),
				})
			case "STOPPED":
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to Update fortiflexvm_entitlements_hardware",
					Detail:   fmt.Sprintf("Current entitlement status is STOPPED. You can't update a STOPPED entitlement. Please set `status = ACTIVE`."),
				})
			default:
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to Update fortiflexvm_entitlements_hardware",
					Detail:   fmt.Sprintf("Current entitlement status is %v. FlexVM only could update ACTIVE entitlements.", current_status),
				})
			}
			return diags
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
		obj["endDate"] = v
	}
	target_entitlement, err = c.UpdateVmUpdate(&obj)
	if err != nil {
		return diag.FromErr(err)
	}

	// stop entitlement
	if set_status == "STOPPED" {
		target_entitlement, err = changeVMStatus(target_entitlement["serialNumber"].(string), "stop", m)
		if err != nil {
			return diag.FromErr(err)
		}
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
	d.Set("config_id", o["configId"])
	d.Set("description", o["description"])
	d.Set("end_date", o["endDate"])
	d.Set("serial_number", o["serialNumber"])
	d.Set("status", o["status"])
	d.Set("start_date", o["startDate"])
	return diags
}
