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
                                Description: "",
                        },
                },
                ResourcesMap: map[string]*schema.Resource{},
        }
}


// func Provider() terraform.ResourceProvider {
// return &schema.Provider{
//         Schema: map[string]*schema.Schema{
//                 "address": {
//                         Type:        schema.TypeString,
//                         Required:    true,
//                         DefaultFunc: schema.EnvDefaultFunc("VAULT_ADDR", nil),
//                         Description: "URL of the root of the target Vault server    .",
//                 },