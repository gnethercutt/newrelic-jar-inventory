package main

import (
	sdkArgs "github.com/newrelic/infra-integrations-sdk/args"
	"github.com/newrelic/infra-integrations-sdk/log"
	"github.com/newrelic/infra-integrations-sdk/sdk"
	"os/exec"
	"strings"
)
import "regexp"
import "fmt"

type argumentList struct {
	sdkArgs.DefaultArgumentList
	Jrepath string `default:"/usr/bin" help:"Path to the directory containing the jar binary from the JRE/JDK."`
	Jarfile string `help:"Path to the Jar file to inspect for version information."`
}

const (
	integrationName    = "org.nethercutt.jarinventory"
	integrationVersion = "0.1.0"
)

var args argumentList

func populateInventory(inventory sdk.Inventory, jarfile string, jrepath string) error {
	cmd := exec.Command("/bin/sh", "-c", fmt.Sprintf("%s/jar tvf %s", jrepath, jarfile))
	output, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}
	r, _ := regexp.Compile(`.*/(.*)\-(\d[0-9_\.\-\w]*)\.jar`)
	for _, line := range strings.Split(string(output), "\n") {
		if strings.Contains(line, ".jar") {
			f := strings.Fields(line)
			if len(f) > 0 {
				result := r.FindStringSubmatch(f[7])
				if len(result) > 2 {
					inventory.SetItem(result[1], "version", result[2])
				}
			}
		}
	}
	return nil
}

func main() {
	integration, err := sdk.NewIntegration(integrationName, integrationVersion, &args)
	fatalIfErr(err)

	if args.All || args.Inventory {
		fatalIfErr(populateInventory(integration.Inventory, args.Jarfile, args.Jrepath))
	}

	fatalIfErr(integration.Publish())
}

func fatalIfErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
