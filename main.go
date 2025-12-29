package main

import (
	"fmt"
	"os"

	"github.com/avinash-exp/pennypecker/provider"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

// Version is set during build via ldflags
var Version = "0.1.0-alpha.1"

func main() {
	// Handle --version flag
	if len(os.Args) > 1 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
		fmt.Println("pennypecker version", Version)
		return
	}

	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: provider.New,
	})
}
