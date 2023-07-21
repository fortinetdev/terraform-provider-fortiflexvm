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

func resourceEntitlementsVMToken() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceEntitlementsVMTokenUpdate,
		ReadContext:   resourceEntitlementsVMTokenRead,
		UpdateContext: resourceEntitlementsVMTokenUpdate,
		DeleteContext: resourceEntitlementsVMTokenDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},
		Schema: map[string]*schema.Schema{
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"regenerate_token": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
				ForceNew: true,
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

func resourceEntitlementsVMTokenRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// Send request
	target_entitlement, diags := getEntitlementFromId(d.Id(), m)
	if diags.HasError() {
		return diags
	}
	d.Set("config_id", target_entitlement["configId"])
	d.Set("serial_number", target_entitlement["serialNumber"])
	d.Set("regenerate_token", false)
	d.Set("token", target_entitlement["token"])
	d.Set("token_status", target_entitlement["tokenStatus"])
	diags = refreshObjectEntitlementsVMToken(d, target_entitlement)
	return diags
}

func resourceEntitlementsVMTokenUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*FortiClient).Client
	c.Retries = 1

	config_id := d.Get("config_id")
	serial_number := d.Get("serial_number")
	regenerate_token := d.Get("regenerate_token").(bool)

	resource_id := fmt.Sprintf("%v.%v", serial_number, config_id)

	if regenerate_token {
		request_obj := make(map[string]interface{})
		request_obj["serialNumber"] = serial_number
		_, err := c.UpdateVmUpdateRegenerateToken(&request_obj)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	d.SetId(resource_id)

	diags = resourceEntitlementsVMTokenRead(ctx, d, m)
	return diags
}

func resourceEntitlementsVMTokenDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	d.SetId("")
	return diags
}

func refreshObjectEntitlementsVMToken(d *schema.ResourceData, o map[string]interface{}) diag.Diagnostics {
	var diags diag.Diagnostics

	result := make([]map[string]interface{}, 0, 1)
	tmp := make(map[string]interface{})
	if _, ok := o["configId"]; ok {
		tmp["config_id"] = o["configId"]
	}
	if _, ok := o["description"]; ok {
		tmp["description"] = o["description"]
	}
	if _, ok := o["serialNumber"]; ok {
		tmp["serial_number"] = o["serialNumber"]
	}
	if _, ok := o["startDate"]; ok {
		tmp["start_date"] = o["startDate"]
	}
	if _, ok := o["endDate"]; ok {
		tmp["end_date"] = o["endDate"]
	}
	if _, ok := o["status"]; ok {
		tmp["status"] = o["status"]
	}
	if _, ok := o["token"]; ok {
		tmp["token"] = o["token"]
	}
	if _, ok := o["tokenStatus"]; ok {
		tmp["token_status"] = o["tokenStatus"]
	}
	result = append(result, tmp)
	d.Set("entitlements", result)

	return diags
}
