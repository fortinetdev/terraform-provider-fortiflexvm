package config

import (
	"net/http"

	"github.com/terraform-providers/terraform-provider-fortiflexvm/sdk/auth"
)

// Config provides configuration to a FortiFlex client instance
// It saves authentication information and a http connection
// for FortiFlex Client instance to create New connction to FortiFlex
// and Send data to FortiFlex,  etc. (needs to be extended later.)
type Config struct {
	Auth    *auth.Auth
	HTTPCon *http.Client
}
