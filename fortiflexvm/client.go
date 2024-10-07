package fortiflexvm

import (
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fortisdk "github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/sdkcore"
)

// FortiClient contains the basic FortiFlex SDK connection information to FortiFlex
// It can be used to as a client of FortiFlex for the plugin
type FortiClient struct {
	Client        *fortisdk.FortiSDKClient
	ImportOptions *schema.Set // Only used in terraform import
}

// providerConfigure creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))
	username := d.Get("username").(string)
	password := d.Get("password").(string)
	client, err := fortisdk.NewClient(username, password)
	if err != nil {
		return nil, err
	}
	return &FortiClient{
		Client:        client,
		ImportOptions: d.Get("import_options").(*schema.Set),
	}, nil
}

func importOptionChecking(ImportOptions *schema.Set, para string) string {
	v := ImportOptions.List()
	if len(v) == 0 {
		return ""
	}

	for _, v1 := range v {
		if v2, ok := v1.(string); ok {
			v3 := strings.Split(v2, "=")

			if len(v3) == 2 { // Example "program_serial_number=******"
				if v3[0] == para {
					return v3[1]
				}
			}
		}
	}

	return ""
}
