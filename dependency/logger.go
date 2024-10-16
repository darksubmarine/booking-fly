package dependency

import (
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"log/slog"
	"os"
)

type LoggerProvider struct {
	app.BaseProvider

	// logger instance to be provided
	logger log.ILogger `torpedo.di:"provide"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewLoggerProvider(config conf.Map) *LoggerProvider {
	return &LoggerProvider{cfg: config}
}

// Provide provides the logger instance.
func (p *LoggerProvider) Provide(c app.IContainer) error {
	// read log level from configuration. Info will be by default if not found.
	sLvl := conf.SlogLevel(p.cfg.FetchStringOrElse("info", "level"))

	// creating log instance as singleton.
	p.logger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: sLvl}))

	return nil
}
