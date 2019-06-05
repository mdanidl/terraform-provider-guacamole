package main

import (
        "github.com/hashicorp/terraform/plugin"
		"github.com/hashicorp/terraform/terraform"
		"github.com/terraform-providers/terraform-provider-guacamole/guacamole"
)

func main() {
        plugin.Serve(&plugin.ServeOpts{
                ProviderFunc: func() terraform.ResourceProvider {
                        return Provider()
                },
        })
}