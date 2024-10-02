package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"log/slog"
	"os"
)

func NewAppContainer(cfg conf.Map, environment string) app.IApp {

	// Options to initialize the internal application container logger.
	opts := app.ContainerOpts{Log: app.ContainerLogsOpts{W: os.Stdout, L: slog.LevelInfo}}

	// returns the app.IApp instance.
	return app.NewContainer(opts).
		WithProvider(NewLoggerProvider(cfg.FetchSubMapP("logger"))).
		WithProvider(NewStorageKeyProvider(cfg.FetchSubMapP("domain", "storage"))).
		WithProvider(NewHttpServerProvider(cfg.FetchSubMapP("http"))).

		// entity providers

		WithProvider(NewTripProvider(cfg.FetchSubMapP("domain"))).
		WithProvider(NewUserProvider(cfg.FetchSubMapP("domain"))).

		// Use Case providers

		WithProvider(NewUseCaseBookingFlyProvider()).

		// domain provider
		WithProvider(NewDomainProvider(cfg.FetchSubMapP("domain")))

	// Use Case providers
	//WithProvider(NewUseCaseFooProvider())
	//WithProvider(NewUseCaseBarProvider())

	/* here you can add all needed providers */
}
