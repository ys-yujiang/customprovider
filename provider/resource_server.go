package provider

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func resourceServer() *schema.Resource {
	fmt.Print()
	return &schema.Resource{
		Create: resourceServerCreate,
		Read:   resourceServerRead,
		Update: resourceServerUpdate,
		Delete: resourceServerDelete,
		Schema: map[string]*schema.Schema{
			"address": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "A description of an item",
			},
			// "tags": {
			// 	Type:        schema.TypeSet,
			// 	Optional:    true,
			// 	Description: "An optional list of tags, represented as a key, value pair",
			// 	Elem:        &schema.Schema{Type: schema.TypeString},
			// },
		},
	}
}

func resourceServerCreate(d *schema.ResourceData, m interface{}) error {
	address := d.Get("address").(string)
	description := d.Get("description").(string)
	// tfTags := d.Get("tags").(*schema.Set).List()
	// tags := make([]string, len(tfTags))
	// for i, tfTag := range tfTags {
	// 	tags[i] = tfTag.(string)
	// }
	d.SetId(address)
	d.SetId(description)
	// d.SetId(tags)
	return resourceServerRead(d, m)
}

func resourceServerRead(d *schema.ResourceData, m interface{}) error {
	// client := m.(*MyClient)
	// attempt to read from an upstream API
	// obj, ok := client.Get(d.Id())
	// if the resource does not exist, inform terraform.
	// We want to immediately prevent further processing
	// if !ok {
	// 	d.SetId("")
	// 	return nil
	// }
	// d.Set("address", obj.Addres)
	return nil
}

func resourceServerUpdate(d *schema.ResourceData, m interface{}) error {
	// Enable partial state mode
	// d.Partial(true)
	// if d.HasChange("address") {
	// 	// tye update the address
	// 	if err := updateAddress(d, m); err != nil {
	// 		return err
	// 	}
	// 	d.SetPartial("address")
	// }
	//if we were to returh here, before disabling partial mode below,
	// then only the "address" field would be saved

	// we succeeded,d isable partial mode. This causes terraform to save
	// all fields again
	//d.Partial(false)
	return resourceServerRead(d, m)
}

func resourceServerDelete(d *schema.ResourceData, m interface{}) error {
	return nil
}
