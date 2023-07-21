// Copyright 2023 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation: Xing Li (@lix-fortinet), Xinwei Du (@dux-fortinet), Hongbin Lu (@fgtdev-hblu)

// Description: Provider for FortiFlex

package fortiflexvm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider creates and returns the FortiFlex *schema.Provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"username": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "The API username.",
			},

			"password": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "",
			},

			"import_options": &schema.Schema{
				Type:     schema.TypeSet,
				Elem:     &schema.Schema{Type: schema.TypeString},
				Optional: true,
				Computed: true,
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"fortiflexvm_programs_list":       dataSourceProgramsList(),
			"fortiflexvm_configs_list":        dataSourceConfigsList(),
			"fortiflexvm_entitlements_list":   dataSourceEntitlementsList(),
			"fortiflexvm_entitlements_points": dataSourceEntitlementsPoints(),
			"fortiflexvm_groups_list":         dataSourceGroupsList(),
			"fortiflexvm_groups_nexttoken":    dataSourceGroupsNexttoken(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"fortiflexvm_config":                resourceConfig(),
			"fortiflexvm_entitlements_vm":       resourceEntitlementsVM(),
			"fortiflexvm_entitlements_hardware": resourceEntitlementsHW(),
			"fortiflexvm_entitlements_vm_token": resourceEntitlementsVMToken(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {

	// Init client config with the values from TF files
	config := Config{
		Username:      d.Get("username").(string),
		Password:      d.Get("password").(string),
		ImportOptions: d.Get("import_options").(*schema.Set),
	}

	// Create Client for later connections
	return config.CreateClient()
}
