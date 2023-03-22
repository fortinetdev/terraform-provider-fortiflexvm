package fortiflexvm

import (
	"os"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

var testAccProviders map[string]terraform.ResourceProvider
var testAccProvider *schema.Provider

func init() {
	testAccProvider = Provider().(*schema.Provider)
	testAccProviders = map[string]terraform.ResourceProvider{
		"fortiflexvm": testAccProvider,
	}
}

func TestProvider(t *testing.T) {
	if err := Provider().(*schema.Provider).InternalValidate(); err != nil {
		t.Fatalf("err: %s", err)
	}
}

func testAccPreCheck(t *testing.T) {
	if v := os.Getenv("FLEXVM_ACCESS_USERNAME"); v == "" {
		t.Fatal("FLEXVM_ACCESS_USERNAME must be set for acceptance tests")
	}
	if v := os.Getenv("FLEXVM_ACCESS_PASSWORD"); v == "" {
		t.Fatal("FLEXVM_ACCESS_PASSWORD must be set for acceptance tests")
	}
}
