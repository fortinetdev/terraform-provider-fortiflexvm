package config

import (
	"net/http"

	"github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/auth"
)

// Config provides configuration to a FlexVM client instance
// It saves authentication information and a http connection
// for FlexVM Client instance to create New connction to FlexVM
// and Send data to FlexVM,  etc. (needs to be extended later.)
type Config struct {
	Auth    *auth.Auth
	HTTPCon *http.Client
}
