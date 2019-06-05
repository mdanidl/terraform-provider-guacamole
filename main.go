package main

import (
        "github.com/hashicorp/terraform/plugin"
		"github.com/hashicorp/terraform/terraform"
		"github.com/mdanidl/terraform-provider-guacamole/guacamole"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return guacamole.Provider()
                },
        })
}