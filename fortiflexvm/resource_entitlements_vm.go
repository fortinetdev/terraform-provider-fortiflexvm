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
			"folder_path": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
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
			"refresh_token_when_destroy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
		},
	}
}

func resourceEntitlementsVMCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*FortiClient).Client

	// If the user does not specify serial_number, create a new one, else, retrieve the old one.
	config_id := d.Get("config_id")
	serial_number := d.Get("serial_number")
	var target_entitlement map[string]interface{}
	if serial_number != "" {
		// Query existing entitlement
		resource_id := fmt.Sprintf("%v.%v", serial_number, config_id)
		target_entitlement, diags = getEntitlementFromId(resource_id, m)
		if diags.HasError() {
			return diags
		}
	} else {
		// Send create request
		obj := make(map[string]interface{})
		obj["configId"] = config_id
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
		var err error
		target_entitlement, err = c.CreateEntitlementsVM(&obj)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	resource_id := fmt.Sprintf("%v.%v", target_entitlement["serialNumber"], target_entitlement["configId"])
	d.SetId(resource_id)

	if serial_number != "" {
		// Only send update request if the user specifies serial_number
		diags = resourceEntitlementsVMUpdate(ctx, d, m)
	} else {
		var err error
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
		diags = refreshObjectEntitlementsVM(d, target_entitlement)
	}
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

func tryParseISO8601(timeStr string) (time.Time, error) {
	var layouts = []string{
		time.RFC3339,
		"2006-01-02T15:04:05",
		"2006-01-02T15:04",
		"2006-01-02T15",
		"2006-01-02",
		"2006-01-02T15:04:05.999999999",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, timeStr)
		if err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("unsupport time format: %s", timeStr)
}

func resourceEntitlementsVMUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
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
			if current_status == "PENDING" {
				diags = append(diags, diag.Diagnostic{
					Severity: diag.Warning,
					Summary:  "Unable to change status from PENDING to ACTIVE",
					Detail: "The current entitlement status is PENDING. You can't manually change PENDING to ACTIVE. " +
						"Once you use the token, the entitlement becomes ACTIVE.",
				})
			}
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
	if v, ok := d.GetOk("description"); ok {
		obj["description"] = v
	}
	if v, ok := d.GetOk("end_date"); ok {
		now := time.Now()
		current_end_date := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
		user_end_date, err := tryParseISO8601(v.(string))
		if err != nil {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "Unable to parsing end_date, ignoring update end_date",
				Detail:   fmt.Sprintf("Unable to parsing %v, please check the format.", v),
			})
		}
		if current_end_date.Before(user_end_date) {
			obj["endDate"] = v
		} else {
			diags = append(diags, diag.Diagnostic{
				Severity: diag.Warning,
				Summary:  "end_date can not be before today's date, ignoring update end_date",
				Detail:   fmt.Sprintf("today's date: %v, end_date: %v.", current_end_date, user_end_date),
			})
		}
	}
	target_entitlement, err = c.UpdateVmUpdate(&obj)
	if err != nil {
		return diag.FromErr(err)
	}

	// Update status
	update_diags := refreshObjectEntitlementsVM(d, target_entitlement)
	if update_diags.HasError() {
		return update_diags
	}
	resource_id := fmt.Sprintf("%v.%v", target_entitlement["serialNumber"], target_entitlement["configId"])
	d.SetId(resource_id)

	return diags
}

func resourceEntitlementsVMDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	c := m.(*FortiClient).Client

	// Get ID
	serial_number, _, diags := splitID(d.Id())
	if diags.HasError() {
		return diags
	}

	// Send delete request
	target_entitlement, diags := getEntitlementFromId(d.Id(), m)
	if target_entitlement["status"] != "STOPPED" {
		_, err := changeVMStatus(serial_number, "stop", m)
		if err != nil {
			return diag.FromErr(err)
		}
	}

	// If refresh_token_when_destroy, refresh token
	if d.Get("refresh_token_when_destroy").(bool) {
		request_obj := make(map[string]interface{})
		request_obj["serialNumber"] = serial_number
		_, err := c.UpdateVmUpdateRegenerateToken(&request_obj)
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
	if value, ok := o["token"]; ok {
		d.Set("token", value)
	}
	if value, ok := o["tokenStatus"]; ok {
		d.Set("token_status", value)
	}
	return diags
}
