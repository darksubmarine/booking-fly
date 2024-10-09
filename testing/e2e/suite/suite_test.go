package suite

import (
	"github.com/darksubmarine/booking-fly/dependency"
	"github.com/darksubmarine/torpedo-lib-go/app"
	"github.com/darksubmarine/torpedo-lib-go/conf"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/suite"
	"net/http/httptest"
	"testing"
	"time"
)

// DomainSuite test suite to validate end to end testing against the exposed domain endpoints.
type DomainSuite struct {
	suite.Suite
	application app.IApp
	router      *gin.Engine
}

// SetupSuite setting up the test suite before to start test executions
func (s *DomainSuite) SetupSuite() {
	// 1. Configuration
	cfg := conf.Load(true, conf.NewYamlLoader("config-test.yaml"))

	// 2. Application container
	s.application = dependency.NewAppContainer(cfg, "test")

	// 3. Application Run!
	go func() { s.application.Run() }()

	// 4. Wait application loaded
	time.Sleep(5 * time.Second)

	// 5. Link dependencies
	s.router = s.application.InvokeByTypeP(&gin.Engine{}).(*gin.Engine)
}

// body helper function to fetch the Response body as string
func (s *DomainSuite) body(recorder *httptest.ResponseRecorder) string {
	return recorder.Body.String()
}

// TestDomainSuite main function that runs the test suite
func TestDomainSuite(t *testing.T) {
	suite.Run(t, new(DomainSuite))
}
