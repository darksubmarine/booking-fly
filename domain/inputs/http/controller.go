package http

import (
	"github.com/darksubmarine/booking-fly/domain"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"
)

type DomainController struct {
	logger    log.ILogger
	domainSrv domain.IDomainService
}

func NewDomainController(domainSrv domain.IDomainService, logger log.ILogger) *DomainController {
	return &DomainController{domainSrv: domainSrv, logger: logger}
}

func (d *DomainController) Register(router gin.IRouter) {
	/* Register here your use cases endpoints

	   For instance:
	      Do not forget to import your use case:
	            import onboardingHTTP "github.com/darksubmarine/blog-app/domain/use_cases/onboarding/inputs/http"

	      Add here the use case endpoint:
	            router.POST("/onboarding", onboardingHTTP.NewController(d.domainSrv, d.logger).RegisterNewUser)
	*/
}
