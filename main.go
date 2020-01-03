package main

import (
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm"
	"github.com/IBM-Cloud/terraform-provider-ibm/version"
	"github.com/hashicorp/terraform-plugin-sdk/plugin"
)

func main() {
	log.Println("IBM Cloud Provider version", version.Version, version.VersionPrerelease, version.GitCommit)
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: ibm.Provider,
	})
}
