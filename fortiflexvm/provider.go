// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Xing Li (@lix-fortinet), Hongbin Lu (@fgtdev-hblu),

// Description: Provider for FlexVM

package fortiflexvm

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Provider creates and returns the FlexVM *schema.Provider.
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
			"fortiflexvm_programs_list":    dataSourceProgramsList(),
			"fortiflexvm_configs_list":     dataSourceConfigsList(),
			"fortiflexvm_vms_list":         dataSourceVmsList(),
			"fortiflexvm_vms_points":       dataSourceVmsPoints(),
			"fortiflexvm_groups_list":      dataSourceGroupsList(),
			"fortiflexvm_groups_nexttoken": dataSourceGroupsNexttoken(),
		},

		ResourcesMap: map[string]*schema.Resource{
			"fortiflexvm_config":     resourceConfig(),
			"fortiflexvm_vms_create": resourceVmsCreate(),
			"fortiflexvm_vms_update": resourceVmUpdate(),
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
