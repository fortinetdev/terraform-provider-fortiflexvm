package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
	"github.com/terraform-providers/terraform-provider-fortiflexvm/fortiflexvm"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: fortiflexvm.Provider})
}
