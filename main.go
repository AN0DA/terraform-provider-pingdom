package main

import (
	"github.com/AN0DA/terraform-provider-pingdom/pingdom"
	"github.com/hashicorp/terraform/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: pingdom.Provider,
	})
}
