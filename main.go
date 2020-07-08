package main

import (
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
	"github.com/toddlers/terraform-provider-ipaddress/provider"
)

// func main() {
// 	plugin.Serve(&plugin.ServeOpts{
// 		ProviderFunc: func() terraform.ResourceProvider {
// 			return provider.Provider()
// 		},
// 	})
// }

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.Provider,
	})
}
