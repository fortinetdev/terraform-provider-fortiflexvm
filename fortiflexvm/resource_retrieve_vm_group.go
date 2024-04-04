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

func resourceRetrieveVMGroup() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceRetrieveVMGroupCreate,
		ReadContext:   resourceRetrieveVMGroupRead,
		UpdateContext: resourceRetrieveVMGroupUpdate,
		DeleteContext: resourceRetrieveVMGroupDelete,
		Schema: map[string]*schema.Schema{
			"task_name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"count_num": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Required: true,
			},
			"preempt_interval": &schema.Schema{
				Type:     schema.TypeFloat,
				Optional: true,
				Default:  1.0,
			},
			"refresh_token_when_destroy": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"refresh_token_when_create": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"require_exact_count": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"retrieve_status": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"entitlements": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"account_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"config_id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"description": &schema.Schema{
							Type:     schema.TypeString,
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
						"end_date": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
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
				},
			},
		},
	}
}

func resourceRetrieveVMGroupCreate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	count_number := d.Get("count_num").(int)
	result_entitlements, found_number, diags := retrieveStoppedEntitlements(count_number, d, m)
	if diags.HasError() {
		return diags
	}
	d.Set("entitlements", result_entitlements)
	report_error := d.Get("require_exact_count").(bool)
	if report_error && found_number < count_number {
		diags = append(diags, diag.Diagnostic{
			Severity: diag.Error,
			Summary:  fmt.Sprintf("Unable to retrieve %v entitlement(s).", count_number),
			Detail: fmt.Sprintf("Only retrieve %v entitlement(s), while asking for %v. ", found_number, count_number) +
				"You get this error because you set require_exact_count as true. All retrieved entitlements by this resource are released now.",
		})
		delete_diags := resourceRetrieveVMGroupDelete(ctx, d, m)
		diags = append(diags, delete_diags...)
		return diags
	}
	d.Set("count_num", found_number)
	d.SetId(d.Get("task_name").(string))
	return diags
}

func resourceRetrieveVMGroupRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	count_number := d.Get("count_num").(int)
	result_entitlements := make([]map[string]interface{}, 0, count_number)
	local_entitlements := d.Get("entitlements").([]interface{})
	for _, item := range local_entitlements {
		entitlement := item.(map[string]interface{})
		serial_number := entitlement["serial_number"].(string)
		config_id := entitlement["config_id"].(int)
		// Query again to get the latest information
		resource_id := fmt.Sprintf("%v.%v", serial_number, config_id)
		entitlement, diags = getEntitlementFromId(resource_id, m)
		result_entitlements = appendEntitlement(result_entitlements, entitlement)
	}
	d.Set("entitlements", result_entitlements)
	return diags
}

func resourceRetrieveVMGroupUpdate(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	want_num := d.Get("count_num").(int)
	local_entitlements := d.Get("entitlements").([]interface{})
	local_num := len(local_entitlements)
	if want_num > local_num {
		result_entitlements, found_number, diags := retrieveStoppedEntitlements(want_num-local_num, d, m)
		if diags.HasError() {
			return diags
		}
		for _, entitlement := range result_entitlements {
			local_entitlements = append(local_entitlements, entitlement)
		}
		d.Set("entitlements", local_entitlements)
		d.Set("count_num", local_num+found_number)
	} else if want_num < local_num {
		result_entitlements := make([]map[string]interface{}, 0, want_num)
		for i, item := range local_entitlements {
			entitlement := item.(map[string]interface{})
			if i < want_num {
				result_entitlements = append(result_entitlements, entitlement)
			} else {
				serial_number := entitlement["serial_number"].(string)
				config_id := entitlement["config_id"].(int)
				diags = removeEntitlement(serial_number, config_id, d, m)
				if diags.HasError() {
					return diags
				}
			}
		}
		d.Set("entitlements", result_entitlements)
		d.Set("count_num", len(result_entitlements))
	}
	return diags
}

func resourceRetrieveVMGroupDelete(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	var diags diag.Diagnostics
	entitlements := d.Get("entitlements").([]interface{})
	for _, item := range entitlements {
		entitlement := item.(map[string]interface{})
		serial_number := entitlement["serial_number"].(string)
		config_id := entitlement["config_id"].(int)
		diags = removeEntitlement(serial_number, config_id, d, m)
		if diags.HasError() {
			return diags
		}
	}
	d.SetId("")
	return diags
}

func removeEntitlement(serial_number string, config_id int, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	c := m.(*FortiClient).Client
	var diags diag.Diagnostics
	var err error
	// Remove description
	request_obj := make(map[string]interface{})
	request_obj["serialNumber"] = serial_number
	request_obj["configId"] = config_id
	request_obj["description"] = ""
	_, err = c.UpdateVmUpdate(&request_obj)
	if err != nil {
		return diag.FromErr(err)
	}
	// Stop
	_, err = changeVMStatus(serial_number, "stop", m)
	if err != nil {
		return diag.FromErr(err)
	}
	// Refresh token
	if d.Get("refresh_token_when_destroy").(bool) {
		request_obj = make(map[string]interface{})
		request_obj["serialNumber"] = serial_number
		_, err = c.UpdateVmUpdateRegenerateToken(&request_obj)
		if err != nil {
			return diag.FromErr(err)
		}
	}
	return diags
}

func retrieveStoppedEntitlements(want_num int, d *schema.ResourceData, m interface{}) ([]map[string]interface{}, int, diag.Diagnostics) {
	var diags diag.Diagnostics
	task_name := d.Get("task_name").(string)
	found_number := 0
	result_entitlements := make([]map[string]interface{}, 0, want_num)
	preempt_interval := d.Get("preempt_interval").(float64)
	rawAllowStatus := d.Get("retrieve_status").([]interface{})
	allow_status := []string{}
	if len(rawAllowStatus) == 0 {
		allow_status = append(allow_status, "STOPPED")
	}
	for _, v := range rawAllowStatus {
		allow_status = append(allow_status, v.(string))
	}
	c := m.(*FortiClient).Client
	// Retrieve all entitlements with the required config_id
	request_obj := make(map[string]interface{})
	request_obj["configId"] = d.Get("config_id").(int)
	return_obj, err := c.ReadEntitlementsList(&request_obj)
	if err != nil {
		return nil, 0, diag.FromErr(err)
	}
	// Find the target entitlements
	all_entitlements := return_obj["entitlements"].([]interface{})
	for _, item := range all_entitlements {
		if found_number >= want_num {
			break
		}
		entitlement := item.(map[string]interface{})
		serial_number := entitlement["serialNumber"].(string)
		config_id := entitlement["configId"]
		// Query again to get the latest information
		resource_id := fmt.Sprintf("%v.%v", serial_number, config_id)
		entitlement, _ = getEntitlementFromId(resource_id, m)
		current_status := entitlement["status"]
		found := false
		for _, s := range allow_status {
			if current_status == s {
				found = true
				break
			}
		}
		if found && (entitlement["description"] == "" || entitlement["description"] == nil) {
			// Preempt this entitlement
			request_obj = make(map[string]interface{})
			request_obj["serialNumber"] = serial_number
			request_obj["configId"] = config_id
			request_obj["description"] = task_name
			_, err = c.UpdateVmUpdate(&request_obj)
			if err != nil {
				continue
			}
			// Sleep preempt_interval second and check description again
			time.Sleep(time.Duration(int64(preempt_interval * 1e9)))
			entitlement, _ = getEntitlementFromId(resource_id, m)
			if entitlement["description"] != task_name {
				// If this entitlement has been used by other tasks, skip.
				continue
			}
			// Use this entitlement
			entitlement, err = changeVMStatus(serial_number, "reactivate", m)
			if err != nil {
				continue
			}
			// Refresh token
			if d.Get("refresh_token_when_create").(bool) {
				request_obj = make(map[string]interface{})
				request_obj["serialNumber"] = serial_number
				entitlement, err = c.UpdateVmUpdateRegenerateToken(&request_obj)
				if err != nil {
					continue
				}
			}
			found_number += 1
			result_entitlements = appendEntitlement(result_entitlements, entitlement)
		}
	}
	return result_entitlements, found_number, diags
}

func appendEntitlement(result_entitlements []map[string]interface{}, entitlement map[string]interface{}) []map[string]interface{} {
	tmp := make(map[string]interface{})
	if value, ok := entitlement["accountId"]; ok {
		tmp["account_id"] = value
	}
	if value, ok := entitlement["configId"]; ok {
		tmp["config_id"] = value
	}
	if value, ok := entitlement["description"]; ok {
		tmp["description"] = value
	}
	if value, ok := entitlement["serialNumber"]; ok {
		tmp["serial_number"] = value
	}
	if value, ok := entitlement["startDate"]; ok {
		tmp["start_date"] = value
	}
	if value, ok := entitlement["endDate"]; ok {
		tmp["end_date"] = value
	}
	if value, ok := entitlement["status"]; ok {
		tmp["status"] = value
	}
	if value, ok := entitlement["token"]; ok {
		tmp["token"] = value
	}
	if value, ok := entitlement["tokenStatus"]; ok {
		tmp["token_status"] = value
	}
	result_entitlements = append(result_entitlements, tmp)
	return result_entitlements
}
