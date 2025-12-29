package provider

import (
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// New returns the Terraform provider schema
func New() *schema.Provider {
	return &schema.Provider{
		Schema:         map[string]*schema.Schema{},
		ResourcesMap:   map[string]*schema.Resource{},
		DataSourcesMap: map[string]*schema.Resource{},
	}
}
