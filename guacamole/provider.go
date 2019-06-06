package guacamole

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	guacapi "github.com/mdanidl/guac-api"
)

func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_URL", nil),
				Description: "",
			},
			"user": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_USER", nil),
				Description: "",
			},
			"password": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_PASSWORD", nil),
				Description: "",
			},
		},
		ResourcesMap: map[string]*schema.Resource{
			"guacamole_user": resourceUser(),
		},
		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	client := &guacapi.Guac{
		URI:      d.Get("url").(string),
		Username: d.Get("user").(string),
		Password: d.Get("password").(string),
	}
	err := client.Connect()
	if err != nil {
		return nil, fmt.Errorf("Failed to connect to Guacamole server: %s", err)
	}
	return client, nil
}
