package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

// terraform-<TYPE>-<NAME>
//In the case above, the plugin is of type "provider" and of name "example".
// func Provider() *schema.Provider {
// 	return &schema.Provider{
// 		ResourcesMap: map[string]*schema.Resource{
// 			"example_server": resourceServer(),
// 		},
// 	}
// }

// func Provider() *schema.Provider {
// 	return &schema.Provider{
// 		ResourcesMap: map[string]*schema.Resource{
// 			"example_ipaddress": resourceServer(),
// 		},
// 	}
// }

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		ResourcesMap: map[string]*schema.Resource{
			"sp_ipaddress": resourceServer(),
		},
		DataSourcesMap: map[string]*schema.Resource{
			"sp_myipaddress": myIPAddressServer(),
		},
	}
}
