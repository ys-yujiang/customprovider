package provider

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

// Data Provider
func myIPAddressServer() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIPAddress,
		Schema: map[string]*schema.Schema{
			"myip": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func dataSourceIPAddress(d *schema.ResourceData, m interface{}) error {
	response, err := http.Get("http://ifconfig.co")
	if err != nil {
		return fmt.Errorf("error getting ip address: %s ,", err)
	}
	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	//fmt.Println("myipaddress", contents)
	d.Set("myip", string(contents))
	d.SetId(string(contents))
	return nil
}
