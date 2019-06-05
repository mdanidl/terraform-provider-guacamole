package guacamole

import (
	"github.com/hashicorp/terraform/helper/schema"
)

func Provider() *schema.Provider {
        return &schema.Provider{
                Schema: map[string]*schema.Schema{
                        "url": {
                                Type: schema.TypeString,
                                Required: true,
                                DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_URL", nil),
                                Description: "",
                        },
                        "user": {
                                Type: schema.TypeString,
                                Required: true,
                                DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_USER", nil),
                                Description: "",
                        },
                        "password": {
                                Type: schema.TypeString,
                                Required: true,
                                DefaultFunc: schema.EnvDefaultFunc("GUACAMOLE_PASSWORD", nil),
                                Description: "",
                        },
                },
                ResourcesMap: map[string]*schema.Resource{
                        "guacamole_user" : resourceUser(),
                },
        }
}