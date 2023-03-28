package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/bufbuild/plugins/internal/plugin"
)

func main() {
	flag.Parse()
	if len(flag.Args()) != 1 {
		flag.Usage()
		os.Exit(2)
	}
	basedir := flag.Args()[0]

	plugins := make([]*plugin.Plugin, 0)
	if err := plugin.Walk(basedir, func(plugin *plugin.Plugin) error {
		plugins = append(plugins, plugin)
		return nil
	}); err != nil {
		log.Fatalf("failed to walk directory: %v", err)
	}

	includedPlugins, err := plugin.FilterByPluginsEnv(plugins, os.Getenv("PLUGINS"))
	if err != nil {
		log.Fatalf("failed to filter plugins by PLUGINS env var: %v", err)
	}

	for _, includedPlugin := range includedPlugins {
		if _, err := fmt.Fprintln(os.Stdout, includedPlugin.Path); err != nil {
			log.Fatalf("failed to print plugin: %v", err)
		}
	}
}
