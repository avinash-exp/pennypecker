package main

import (
    "github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
    "github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// version is set during build via ldflags
var version = "0.1.0-alpha.1"

func main() {
    plugin.Serve(&plugin.ServeOpts{
        ProviderFunc: func() *schema.Provider {
            return &schema.Provider{
                Schema:         map[string]*schema.Schema{},
                ResourcesMap:   map[string]*schema.Resource{},
                DataSourcesMap: map[string]*schema.Resource{},
            }
        },
    })
}