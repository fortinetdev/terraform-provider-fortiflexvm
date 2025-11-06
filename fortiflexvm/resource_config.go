// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Configure a Configuration under a Program.

package fortiflexvm

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceConfig() *schema.Resource {
	return &schema.Resource{
		Create: resourceConfigCreate,
		Read:   resourceConfigRead,
		Update: resourceConfigUpdate,
		Delete: resourceConfigDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"account_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"config_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"program_serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"name": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"product_type": &schema.Schema{
				Type:     schema.TypeString,
				ForceNew: true,
				Required: true,
				Description: `Product Type ID, must be one of the following options:
				FGT_VM_Bundle: FortiGate Virtual Machine - Service Bundle;
				FMG_VM: FortiManager Virtual Machine;
				FWB_VM: FortiWeb Virtual Machine - Service Bundle;
				FGT_VM_LCS: FortiGate Virtual Machine - A La Carte Services;
				FC_EMS_OP: FortiClient EMS On-Prem;
				FAZ_VM: FortiAnalyzer Virtual Machine;
				FPC_VM: FortiPortal Virtual Machine;
				FAD_VM: FortiADC Virtual Machine;
				FORTISOAR_VM: FortiSOAR Virtual Machine;
				FORTIMAIL_VM: FortiMail Virtual Machine;
				FORTINAC_VM: FortiNAC Virtual Machine;
				FGT_HW: FortiGate Hardware;
				FAP_HW: FortiAP Hardware;
				FSW_HW: FortiSwitch Hardware;
				FWBC_PRIVATE: FortiWeb Cloud - Private;
				FWBC_PUBLIC: FortiWeb Cloud - Public;
				FC_EMS_CLOUD: FortiClient EMS Cloud;
				FORTISASE: FortiSASE;
				FORTIEDR: FortiEDR MSSP;
				FORTINDR_CLOUD: FortiNDR Cloud;
				FORTIRECON: FortiRecon;
				SIEM_CLOUD: FortiSIEM Cloud;				
				FORTIAPPSEC: FortiAppSec;
				FORTIDLP: FortiDLP;`,
				ValidateDiagFunc: checkInputValidString("product_type", []string{"FGT_VM_Bundle", "FMG_VM", "FWB_VM", "FGT_VM_LCS",
					"FC_EMS_OP", "FC_EMS_CLOUD", "FAZ_VM", "FPC_VM", "FAD_VM", "FORTISOAR_VM", "FORTIMAIL_VM", "FORTINAC_VM",
					"FGT_HW", "FAP_HW", "FSW_HW", "FWBC_PRIVATE", "FWBC_PUBLIC", "FC_EMS_CLOUD", "FORTISASE",
					"FORTIEDR", "FORTINDR_CLOUD", "FORTIRECON", "SIEM_CLOUD", "FORTIAPPSEC", "FORTIDLP"}),
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				Description: `Configuration status, must be one of the following options:
				ACTIVE: Enable a configuration;
				DISABLED: Disable a configuration.`,
				ValidateDiagFunc: checkInputValidString("status", []string{"ACTIVE", "DISABLED"}),
			},
			"fgt_vm_bundle": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"fortiguard_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"cloud_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fmg_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_dev": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fwb_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fgt_vm_lcs": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fortiguard_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"cloud_services": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fc_ems_op": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"epp_ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"chromebook": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"faz_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"daily_storage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"adom_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fpc_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"managed_dev": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fad_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fortisoar_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"additional_users_license": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fortimail_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"cpu_size": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fortinac_vm": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"endpoints": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"support_service": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fgt_hw": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fap_hw": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fsw_hw": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"device_model": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fwbc_private": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"average_throughput": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"web_applications": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fwbc_public": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"average_throughput": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"web_applications": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fc_ems_cloud": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"ztna_fgf_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"epp_ztna_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"epp_ztna_fgf_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"chromebook": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"fortisase": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"users": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"bandwidth": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"dedicated_ips": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"additional_compute_region": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"locations": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fortiedr": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"endpoints": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"repository_storage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fortindr_cloud": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"metered_usage": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"fortirecon": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"asset_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"network_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"executive_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"vendor_num": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"siem_cloud": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compute_units": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"additional_online_storage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"archive_storage": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"fortiappsec": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_types": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"waf_service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"waf_addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
						"throughput": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"applications": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"qps": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"health_checks": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"fortidlp": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_pkg": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"endpoints": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"addons": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Computed: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func importExistingConfig(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	var err error
	var response_data map[string]interface{}
	config_id := d.Get("config_id").(int)
	request_obj := make(map[string]interface{})
	program_serial_number := d.Get("program_serial_number").(string)
	request_obj["programSerialNumber"] = program_serial_number
	if v, ok := d.GetOk("account_id"); ok {
		request_obj["accountId"] = v
	}
	config_list, err := c.ReadConfigsList(&request_obj)
	if err != nil {
		return fmt.Errorf("can not read configuration list: %v", err)
	}
	response_data, err = findConfigFromList(config_list, config_id)
	if err != nil {
		return err
	}
	if response_data["id"] != nil && response_data["id"] != "" {
		d.SetId(fmt.Sprintf("%v", response_data["id"]))
	} else {
		d.SetId("Config")
	}

	// Update status if needed
	if current_status, ok := response_data["status"].(string); ok {
		if set_status, ok := d.GetOk("status"); ok && current_status != set_status.(string) {
			obj, err := getObjectConfig(d, "id")
			if err != nil {
				return fmt.Errorf("error creating Config resource while getting object: %v", err)
			}

			var op string
			if d.Get("status").(string) == "ACTIVE" {
				op = "enable"
			} else {
				op = "disable"
			}
			response_data, err = c.UpdateConfigStatus(obj, op)
			if err != nil {
				return fmt.Errorf("error update Config status: %v", err)
			}
			if st, ok := response_data["status"].(string); ok {
				if st != d.Get("status").(string) {
					log.Printf("[WARN] Could not update the status of Config %v", d.Id())
				}
			}
		}
	} else {
		log.Printf("[WARN] Could not get status from HTTP response")
	}

	// refresh schema
	err = refreshObjectConfig(d, response_data)
	if err != nil {
		return fmt.Errorf("error refresh Config resource: %v", err)
	}

	return err
}

func createNewConfig(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	var err error
	var response_data map[string]interface{}
	obj, err := getObjectConfig(d, "create")
	if err != nil {
		return fmt.Errorf("error creating Config resource while getting object: %v", err)
	}

	response_data, err = c.CreateConfig(obj)
	if err != nil {
		return fmt.Errorf("error creating Config resource: %v", err)
	}

	if response_data["id"] != nil && response_data["id"] != "" {
		d.SetId(fmt.Sprintf("%v", response_data["id"]))
	} else {
		d.SetId("Config")
	}

	// Update status if needed
	if current_status, ok := response_data["status"].(string); ok {
		if set_status, ok := d.GetOk("status"); ok && current_status != set_status.(string) {
			obj, err := getObjectConfig(d, "id")
			if err != nil {
				return fmt.Errorf("error creating Config resource while getting object: %v", err)
			}

			var op string
			if d.Get("status").(string) == "ACTIVE" {
				op = "enable"
			} else {
				op = "disable"
			}
			response_data, err = c.UpdateConfigStatus(obj, op)
			if err != nil {
				return fmt.Errorf("error update Config status: %v", err)
			}
			if st, ok := response_data["status"].(string); ok {
				if st != d.Get("status").(string) {
					log.Printf("[WARN] Could not update the status of Config %v", d.Id())
				}
			}
		}
	} else {
		log.Printf("[WARN] Could not get status from HTTP response")
	}

	// refresh schema
	err = refreshObjectConfig(d, response_data)
	if err != nil {
		return fmt.Errorf("error refresh Config resource: %v", err)
	}

	return nil
}

func resourceConfigCreate(d *schema.ResourceData, m interface{}) error {
	config_id := d.Get("config_id").(int)
	if config_id != 0 {
		// Import existing one
		return importExistingConfig(d, m)
	} else {
		// Create a new configuration
		return createNewConfig(d, m)
	}
}

func resourceConfigRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if d.Get("program_serial_number") == "" {
		psn := importOptionChecking(m.(*FortiClient).ImportOptions, "program_serial_number")
		if err := d.Set("program_serial_number", psn); err != nil {
			return fmt.Errorf("error set params program_serial_number: %v", err)
		}
	}
	obj, err := getObjectConfig(d, "read")
	if err != nil {
		return fmt.Errorf("error reading Config while getting required parameters: %v", err)
	}

	o, err := c.ReadConfigsList(obj)
	if err != nil {
		return fmt.Errorf("error reading Config resource: %v", err)
	}

	co, err := getConfigReadResponse(o, d.Id())
	if err != nil {
		d.SetId("")
		return err
	}

	err = refreshObjectConfig(d, co)
	if err != nil {
		return fmt.Errorf("error reading Config resource from API: %v", err)
	}
	return nil
}

func resourceConfigUpdate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	obj, err := getObjectConfig(d, "update")
	if err != nil {
		return fmt.Errorf("error updating Config resource while getting object: %v", err)
	}

	o, err := c.UpdateConfig(obj)
	if err != nil {
		return fmt.Errorf("error updating Config resource: %v", err)
	}

	if o["id"] != nil && o["id"] != "" {
		d.SetId(fmt.Sprintf("%v", o["id"]))
	} else {
		d.SetId("Config")
	}

	if st, ok := o["status"].(string); ok {
		if statusV, ok := d.GetOk("status"); ok && st != statusV.(string) {
			obj, err = getObjectConfig(d, "id")
			if err != nil {
				return fmt.Errorf("error creating Config resource while getting object: %v", err)
			}

			var op string
			if d.Get("status").(string) == "ACTIVE" {
				op = "enable"
			} else {
				op = "disable"
			}

			o, err = c.UpdateConfigStatus(obj, op)
			if err != nil {
				return fmt.Errorf("error update Config status: %v", err)
			}
			if st, ok := o["status"].(string); ok {
				if st != d.Get("status").(string) {
					log.Printf("[WARN] Could not update the status of Config %v", d.Id())
				}
			}
		}
	} else {
		log.Printf("[WARN] Could not get status from HTTP response")
	}

	err = refreshObjectConfig(d, o)
	if err != nil {
		return fmt.Errorf("error refresh Config resource: %v", err)
	}

	return nil
}

func resourceConfigDelete(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client

	if d.Get("status").(string) != "DISABLED" {
		obj, err := getObjectConfig(d, "id")
		if err != nil {
			return fmt.Errorf("error creating Config resource while getting object: %v", err)
		}

		o, err := c.UpdateConfigStatus(obj, "disable")
		if err != nil {
			return fmt.Errorf("error update Config status: %v", err)
		}
		if st, ok := o["status"].(string); ok {
			if st != d.Get("status").(string) {
				log.Printf("[WARN] Could not update the status of Config %v", d.Id())
			}
		}

		err = refreshObjectConfig(d, o)
		if err != nil {
			return fmt.Errorf("error refresh Config resource: %v", err)
		}
	}

	d.SetId("")
	return nil
}

func getConfigReadResponse(o map[string]interface{}, mkey string) (map[string]interface{}, error) {
	var err error
	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", mkey)
		err = fmt.Errorf("response is nil")
		return nil, err
	}

	if configs, ok := o["configs"]; ok {
		if configsList, ok := configs.([]interface{}); ok {
			for _, conf := range configsList {
				if confMap, ok := conf.(map[string]interface{}); ok {
					cId := fmt.Sprintf("%v", confMap["id"])
					if cId == mkey {
						return confMap, nil
					}
				}
			}
		}
	}
	err = fmt.Errorf("Config not been created")
	return nil, err
}

func flattenConfigProductType(v interface{}) interface{} {
	var rst interface{}
	rst = ""
	if pt, ok := v.(map[string]interface{}); ok {
		if p_id, ok := pt["id"]; ok {
			rst = convProductTypeId2Name(int(p_id.(float64)))
			if rst == "" {
				log.Printf("[ERROR] Can not recognise Product Type ID: %v", p_id)
			}
		}
	}
	return rst
}

func flattenConfigParameters(v interface{}) interface{} {
	if v == nil {
		return nil
	}

	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil
	}

	result := make([]map[string]interface{}, 0, 1)
	tmp := make(map[string]interface{})
	for _, r := range l {
		param := r.(map[string]interface{})
		_, cName, dataType := convConfParsId2NameList(int(param["id"].(float64)))
		if cName == "" {
			log.Printf("NEW PARAM: %v", param["id"])
			continue
		}
		if cValue, ok := param["value"]; ok {
			switch dataType {
			case "int":
				tmp[cName], _ = strconv.Atoi((cValue.(string)))
			case "string":
				tmp[cName] = cValue.(string)
			case "list":
				if _, ok := tmp[cName]; !ok {
					tmp[cName] = []interface{}{}
				}
				if cValue != "NONE" {
					tmp[cName] = append(tmp[cName].([]interface{}), cValue)
				}
			default:
				tmp[cName] = cValue.(string)
			}
		}
	}
	result = append(result, tmp)

	return result
}

func refreshObjectConfig(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if value, ok := o["accountId"]; ok {
		d.Set("account_id", value)
	}
	if value, ok := o["id"]; ok {
		d.Set("config_id", value)
	}
	if value, ok := o["programSerialNumber"]; ok {
		d.Set("program_serial_number", value)
	}
	if value, ok := o["name"]; ok {
		d.Set("name", value)
	}
	if err = d.Set("product_type", flattenConfigProductType(o["productType"])); err != nil {
		if !fortiAPIPatch(o["productType"]) {
			return fmt.Errorf("error reading product_type: %v", err)
		}
	}
	if value, ok := o["status"]; ok {
		d.Set("status", value)
	}

	// Initialize product variables. This can fix the problem of inconsistent output.
	for _, type_name := range PRODUCT_TYPES {
		empty_interface := make([]map[string]interface{}, 0)
		d.Set(type_name, empty_interface)
	}

	// Set param
	pType := strings.ToLower(d.Get("product_type").(string))
	if err = d.Set(pType, flattenConfigParameters(o["parameters"])); err != nil {
		if !fortiAPIPatch(o["parameters"]) {
			return fmt.Errorf("error reading %v: %v", pType, err)
		}
	}

	return nil
}

func expandConfigProductType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	typeId := convProductTypeName2Id(v.(string))
	if typeId == 0 {
		err := fmt.Errorf("product_type invalid: %v, should be one of [%v]", v.(string),
			"FGT_VM_Bundle, FMG_VM, FWB_VM, FGT_VM_LCS, FC_EMS_OP, FAZ_VM, FPC_VM, FAD_VM, FORTINAC_VM, "+
				"FORTISOAR_VM, FORTIMAIL_VM, FGT_HW, FAP_HW, FSW_HW, FWBC_PRIVATE, FWBC_PUBLIC, "+
				"FC_EMS_CLOUD, FORTISASE, FORTIEDR, FORTINDR_CLOUD, FORTIRECON, SIEM_CLOUD, "+
				"FORTIAPPSEC, FORTIDLP")
		return typeId, err
	}
	return typeId, nil
}

func expandConfigParameters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	for ck, cv := range l[0].(map[string]interface{}) {
		// if isInterfaceEmpty(cv) {
		// 	continue
		// }
		ckId := convConfParsNameList2Id(pre, ck)
		if ckId == 0 {
			err := fmt.Errorf("could not get target argument ID, this is a plugin error")
			log.Printf("[ERROR] %v", err)
			return result, err
		}
		if ckId == 47 || ckId == 60 || ckId == 85 || ckId == 86 || ckId == 87 || ckId == 88 { // This argument is read only
			continue
		}
		if cvList, ok := cv.([]interface{}); ok {
			for _, csv := range cvList {
				tmp := make(map[string]interface{})
				tmp["id"] = ckId
				tmp["value"] = csv
				result = append(result, tmp)
			}
			if len(cvList) == 0 { // if this list is empty, send "NONE" to the fortiflex server
				tmp := make(map[string]interface{})
				tmp["id"] = ckId
				tmp["value"] = "NONE"
				result = append(result, tmp)
			}
		} else {
			if pre == "fgt_vm_bundle" { // version 2.2.0, allow fgt_vm_bundle->support_service empty
				if ck == "support_service" && cv == "" {
					cv = "NONE"
				}
			}
			tmp := make(map[string]interface{})
			tmp["id"] = ckId
			tmp["value"] = cv
			result = append(result, tmp)
		}
	}

	return result, nil
}

func getObjectConfig(d *schema.ResourceData, rType string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if rType == "update" || rType == "id" {
		obj["id"] = d.Id()
	}

	if rType == "create" || rType == "read" {
		if value, ok := d.GetOk("program_serial_number"); ok {
			obj["programSerialNumber"] = value
		}
		if value, ok := d.GetOk("account_id"); ok {
			obj["accountId"] = value
		}
	}

	if rType == "create" || rType == "update" {
		if value, ok := d.GetOk("name"); ok {
			obj["name"] = value
		}

		var pType string
		if v, ok := d.GetOk("product_type"); ok {
			pType = v.(string)
			if rType == "create" {
				t, err := expandConfigProductType(d, v, "product_type")
				if err != nil {
					return &obj, err
				} else if t != nil {
					obj["productTypeId"] = t
				}
			}
		}

		pTypeLower := strings.ToLower(pType)
		if v, ok := d.GetOk(pTypeLower); ok {
			t, err := expandConfigParameters(d, v, pTypeLower)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["parameters"] = t
			}
		}
	}

	return &obj, nil
}
