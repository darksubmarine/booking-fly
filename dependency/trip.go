package dependency

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	tripHTTP "github.com/darksubmarine/booking-fly/domain/entities/trip/inputs/http/gin"
	tripRepo "github.com/darksubmarine/booking-fly/domain/entities/trip/outputs/memory"
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"
)

type TripProvider struct {
	app.BaseProvider

	// trip service instance to be provided.
	service trip.IService `torpedo.di:"provide"`

	// trip repository instance to be provided.
	repo trip.IRepository `torpedo.di:"provide"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// storageKey is the crypto key to encode encrypted fields at storage level.
	storageKey []byte `torpedo.di:"bind,name=STORAGE_KEY"`

	// apiV1 group to register endpoints
	apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewTripProvider(config conf.Map) *TripProvider {
	return &TripProvider{cfg: config}
}

// Provide provides instances.
func (p *TripProvider) Provide(c app.IContainer) error {

	// -- Repo (output) ---
	p.repo = tripRepo.NewMemoryRepository(p.storageKey)

	// -- Service (business logic)
	p.service = trip.NewService(p.repo, p.logger)

	// -- Controller (input) --
	controller := tripHTTP.NewInputGin(p.service, p.logger)
	controller.Register(p.apiV1)

	return nil
}
