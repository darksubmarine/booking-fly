package dependency

import (
	"github.com/darksubmarine/booking-fly/domain/entities/user"
	userHTTP "github.com/darksubmarine/booking-fly/domain/entities/user/inputs/http/gin"
	userRepo "github.com/darksubmarine/booking-fly/domain/entities/user/outputs/memory"
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"

	"github.com/darksubmarine/booking-fly/domain/entities/trip"
)

type UserProvider struct {
	app.BaseProvider

	// user service instance to be provided.
	service user.IService `torpedo.di:"provide"`

	// user repository instance to be provided.
	repo user.IRepository `torpedo.di:"provide"`

	// logger instance provided by LoggerProvider.
	logger log.ILogger `torpedo.di:"bind"`

	// storageKey is the crypto key to encode encrypted fields at storage level.
	storageKey []byte `torpedo.di:"bind,name=STORAGE_KEY"`

	// apiV1 group to register endpoints
	apiV1 *gin.RouterGroup `torpedo.di:"bind,name=APIv1"`

	tripSrv trip.IService `torpedo.di:"bind"`

	// private fields initialized by constructor
	cfg conf.Map
}

func NewUserProvider(config conf.Map) *UserProvider {
	return &UserProvider{cfg: config}
}

// Provide provides instances.
func (p *UserProvider) Provide(c app.IContainer) error {

	// -- Repo (output) ---
	p.repo = userRepo.NewMemoryRepository(p.storageKey)

	// -- Service (business logic)
	p.service = user.NewService(p.repo, p.logger, p.tripSrv)

	// -- Controller (input) --
	controller := userHTTP.NewInputGin(p.service, p.logger)
	controller.Register(p.apiV1)

	return nil
}
