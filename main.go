package main

import (
	"github.com/hashicorp/terraform/plugin"
	"github.com/protolabs-oss/terraform-provider-windns/windns"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: windns.Provider,
	})
}
