package fortiflexvm

import (
	"log"
	"strings"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	fortisdk "github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/sdkcore"
)

// FortiClient contains the basic FortiFlex SDK connection information to FortiFlex
// It can be used to as a client of FortiFlex for the plugin
type FortiClient struct {
	Client              *fortisdk.FortiSDKClient
	AccountID           int
	ProgramSerialNumber string
	ImportOptions       *schema.Set // Only used in terraform import
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
		Client:              client,
		AccountID:           d.Get("account_id").(int),
		ProgramSerialNumber: d.Get("program_serial_number").(string),
		ImportOptions:       d.Get("import_options").(*schema.Set),
	}, nil
}

func getAccountID(d *schema.ResourceData, m interface{}) (interface{}, bool) {
	if value, ok := d.GetOk("account_id"); ok {
		return value, true
	}

	if client, ok := m.(*FortiClient); ok && client.AccountID != 0 {
		return client.AccountID, true
	}

	return nil, false
}

func getProgramSerialNumber(d *schema.ResourceData, m interface{}) (string, bool) {
	if value, ok := d.GetOk("program_serial_number"); ok {
		return value.(string), true
	}

	if client, ok := m.(*FortiClient); ok && client.ProgramSerialNumber != "" {
		return client.ProgramSerialNumber, true
	}

	return "", false
}

type rawConfigGetter interface {
	GetRawConfigAt(cty.Path) (cty.Value, diag.Diagnostics)
}

func isConfigured(d rawConfigGetter, name string) bool {
	rawValue, rawDiags := d.GetRawConfigAt(cty.GetAttrPath(name))
	return !rawDiags.HasError() && !rawValue.IsNull() && rawValue.IsWhollyKnown()
}

func getConfiguredString(d *schema.ResourceData, name string) (string, bool) {
	if !isConfigured(d, name) {
		return "", false
	}
	value, ok := d.Get(name).(string)
	if !ok || value == "" {
		return "", false
	}
	return value, true
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
