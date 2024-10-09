package dependency

import (
	"github.com/darksubmarine/booking-fly/domain/use_cases/booking_fly"
	BookingFlyHTTP "github.com/darksubmarine/booking-fly/domain/use_cases/booking_fly/inputs/http"
	"github.com/gin-gonic/gin"

	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	"github.com/darksubmarine/booking-fly/domain/entities/user"

	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/log"
)

type UseCaseBookingFlyProvider struct {
	app.BaseProvider

	// useCaseBookingFly provides an booking_fly.UseCase instance.
	useCaseBookingFly *booking_fly.UseCase `torpedo.di:"provide"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// userSrv instance of user service.
	userSrv user.IService `torpedo.di:"bind"`

	// tripSrv instance of trip service.
	tripSrv trip.IService `torpedo.di:"bind"`

	// Uncomment following lines if your use case contains http input.
	// api router group to add endpoints under /api prefix
	apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"`
}

func NewUseCaseBookingFlyProvider() *UseCaseBookingFlyProvider {
	return &UseCaseBookingFlyProvider{}
}

// Provide provides the use case instance.
func (p *UseCaseBookingFlyProvider) Provide(c app.IContainer) error {
	p.useCaseBookingFly = booking_fly.NewUseCase(p.logger, p.userSrv, p.tripSrv)

	p.apiV1.POST("/booking",
		BookingFlyHTTP.NewController(p.useCaseBookingFly, p.logger).BookingFlyEndpoint)

	return nil
}
