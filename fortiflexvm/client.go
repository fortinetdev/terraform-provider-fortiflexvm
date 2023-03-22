package fortiflexvm

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/auth"
	forticlient "github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/sdkcore"
)

// Config gets the authentication information from the given metadata
type Config struct {
	Username      string
	Password      string
	ImportOptions *schema.Set
}

// FortiClient contains the basic FlexVM SDK connection information to FlexVM
// It can be used to as a client of FlexVM for the plugin
type FortiClient struct {
	Client *forticlient.FortiSDKClient
	Cfg    *Config
}

// CreateClient creates a FortiClient Object with the authentication information.
// It returns the FortiClient Object for the use when the plugin is initialized.
func (c *Config) CreateClient() (interface{}, error) {
	var fClient FortiClient

	err := createFlexVMClient(&fClient, c)
	if err != nil {
		return nil, fmt.Errorf("Error create flexvm client: %v", err)
	}
	return &fClient, nil
}

func createFlexVMClient(fClient *FortiClient, c *Config) error {
	config := &tls.Config{}

	auth := auth.NewAuth(c.Username, c.Password)

	if auth.Username == "" {
		_, err := auth.GetEnvUsername()
		if err != nil {
			return fmt.Errorf("Error reading Username")
		}
	}

	if auth.Password == "" {
		_, err := auth.GetEnvPassword()
		if err != nil {
			return fmt.Errorf("Error reading Password")
		}
	}

	tr := &http.Transport{
		TLSClientConfig: config,
	}

	client := &http.Client{
		Transport: tr,
		Timeout:   time.Second * 250,
	}

	fc, err := forticlient.NewClient(auth, client)

	if err != nil {
		return fmt.Errorf("connection error: %v", err)
	}

	err = fc.GenerateToken()
	if err != nil {
		return fmt.Errorf("Fail to generate Token: %v", err)
	}

	fClient.Cfg = c
	fClient.Client = fc

	return nil
}
