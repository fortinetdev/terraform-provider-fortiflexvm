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

func resourceEntitlementsVM() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEntitlementsVMCreate,
		ReadContext:   resourceEntitlementsVMRead,
		UpdateContext: resourceEntitlementsVMUpdate,
		DeleteContext: resourceEntitlementsVMDelete,
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
			"folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
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
			"token": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"token_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func resourceEntitlementsVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*FortiClient).Client
	c.Retries = 1

	// Send request
	obj := make(map[string]interface{})
	obj["configId"] = d.Get("config_id")
	obj["count"] = 1
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v
	}
	if v, ok := d.GetOk("folder_path"); ok {
		obj["folderPath"] = v
	}
	if v, ok := d.GetOk("end_date"); ok {
		obj["endDate"] = v
	}
	target_entitlement, err := c.CreateEntitlementsVM(&obj)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update status
	diags = refreshObjectEntitlementsVM(d, target_entitlement)
	if diags.HasError() {
		return diags
	}

	resource_id := fmt.Sprintf("%v.%v", target_entitlement["serialNumber"], target_entitlement["configId"])
	d.SetId(resource_id)
	return diags
}

func resourceEntitlementsVMRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Send request
	target_entitlement, diags := getEntitlementFromId(d.Id(), m)
	if diags.HasError() {
		return diags
	}

	// Update status
	diags = refreshObjectEntitlementsVM(d, target_entitlement)
	return diags
}

func resourceEntitlementsVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
		if set_status == "ACTIVE" && (current_status == "STOPPED" || current_status == "EXPIRED") {
			target_entitlement, err = changeVMStatus(target_entitlement["serialNumber"].(string), "reactivate", m)
			if err != nil {
				return diag.FromErr(err)
			}
		} else { // can't update
			switch current_status {
			case "PENDING":
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to Update fortiflexvm_entitlements_vm",
					Detail:   fmt.Sprintf("Current entitlement status is PENDING. Please use the VM token to activate a virtual machine before using this API."),
				})
			case "STOPPED":
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to Update fortiflexvm_entitlements_vm",
					Detail:   fmt.Sprintf("Current entitlement status is STOPPED. You can't update a STOPPED entitlement. Please set `status = ACTIVE`."),
				})
			default:
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Error,
					Summary:  "Unable to Update fortiflexvm_entitlements_vm",
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
	diags = refreshObjectEntitlementsVM(d, target_entitlement)
	if diags.HasError() {
		return diags
	}
	resource_id := fmt.Sprintf("%v.%v", target_entitlement["serialNumber"], target_entitlement["configId"])
	d.SetId(resource_id)

	return diags
}

func resourceEntitlementsVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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

func refreshObjectEntitlementsVM(d *schema.ResourceData, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	// can't set folder_path
	d.Set("config_id", o["configId"])
	d.Set("description", o["description"])
	d.Set("end_date", o["endDate"])
	d.Set("serial_number", o["serialNumber"])
	d.Set("status", o["status"])
	d.Set("start_date", o["startDate"])
	d.Set("token", o["token"])
	d.Set("token_status", o["tokenStatus"])
	return diags
}
