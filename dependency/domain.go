package dependency

import (
	"github.com/darksubmarine/booking-fly/domain"
	"github.com/darksubmarine/booking-fly/domain/inputs/http"

	"github.com/darksubmarine/booking-fly/domain/entities/trip"

	"github.com/darksubmarine/booking-fly/domain/entities/user"

	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"

	"github.com/darksubmarine/booking-fly/domain/use_cases/booking_fly"
)

type DomainProvider struct {
	app.BaseProvider

	// -- Providers --

	// domainCtx provide domain context instance
	domainCtx *domain.Context `torpedo.di:"provide"`

	// domainSrv provide domain service instance
	domainSrv domain.IDomainService `torpedo.di:"provide"`

	// -- Bind services --

	// trip service wired instance.
	tripSrv trip.IService `torpedo.di:"bind"`

	// user service wired instance.
	userSrv user.IService `torpedo.di:"bind"`

	// -- Bind use cases --
	//ucFoo *foo.UseCase `torpedo.di:"bind"`
	//ucBar *bar.UseCase `torpedo.di:"bind"`

	// BookingFly use case wired instance.
	ucBookingFly *booking_fly.UseCase `torpedo.di:"bind"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// apiV1 group to register endpoints
	apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewDomainProvider(config conf.Map) *DomainProvider {
	return &DomainProvider{cfg: config}
}

// Provide provides instances.
func (p *DomainProvider) Provide(c app.IContainer) error {

	p.domainCtx = domain.NewContext(
		p.tripSrv,

		p.userSrv,
	)

	p.domainSrv = domain.NewService(p.domainCtx /*, p.ucFoo, p.ucBar */)

	// API registration
	domainController := http.NewDomainController(p.domainSrv, p.logger)
	domainController.Register(p.apiV1)

	return nil
}
