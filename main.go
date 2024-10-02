package main

import (
	"fmt"
	"github.com/darksubmarine/booking-fly/dependency"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"os"
)

const (
	appEnvironment = "ENVIRONMENT"
	verboseInit    = false
)

func main() {

	// 1. Configuration
	env := os.Getenv(appEnvironment)
	if env == "" {
		panic("missing ENVIRONMENT!")
	}
	cfg := fetchConfig(env)

	// 2. Application container
	app := dependency.NewAppContainer(cfg, env)

	// 3. Application Run!
	app.Run()

}

func fetchConfig(env string) conf.Map {
	return conf.Load(verboseInit, conf.NewYamlLoader(fmt.Sprintf("./config-%s.yaml", env)))
}
