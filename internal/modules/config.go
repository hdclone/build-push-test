package modules

import (
	"broadcaster/internal/config"
	"flag"
	stdLog "log"
	"os"
)

func Config() *config.Config {
	return Register("config", func(s string) (Module, error) {
		var configFile string

		flagSet := flag.NewFlagSet("service", flag.ExitOnError)
		flagSet.StringVar(&configFile, "c", "config.yml", "config file")
		if err := flagSet.Parse(os.Args[1:]); err != nil {
			stdLog.Fatal(err)
		}
		cfg, err := config.Load(configFile)
		if err != nil {
			stdLog.Fatal(err)
		}
		return cfg, nil
	}).(*config.Config)
}
