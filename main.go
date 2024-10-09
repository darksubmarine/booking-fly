package main

import (
	"fmt"
	"github.com/darksubmarine/booking-fly/dependency"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"os"

	_ "github.com/darksubmarine/booking-fly/oas"
)

const (
	appEnvironment = "ENVIRONMENT"
	verboseInit    = false
)

// @title           Booking Fly Example API
// @version         1.0
// @description     This is a sample server
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
