package time_tracker

import (
	"github.com/SerafimKuzmin/sd/src/cmd/time_tracker/flags"
	echo "github.com/labstack/echo/v4"
)

type base struct {
	Logger   flags.LoggerFlags `toml:"logger"`
	services *baseServices
}

type baseServices struct {
	Logger echo.Logger
	// Tracer          *otel.Tracer
	// MetricsRegistry *metrics.Registry
}

func (b *base) Init(e *echo.Echo) (*baseServices, error) {
	services := &baseServices{}
	logger := b.Logger.Init(e)
	services.Logger = logger
	b.services = services

	return services, nil
}
