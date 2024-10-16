// Package gin input
package gin

import (
	"github.com/darksubmarine/booking-fly/domain/entities/user"
	"github.com/darksubmarine/torpedo-lib-go/http/gin_utils"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"
)

// InputGin struct to handle the clean integration with Gin Framework
type InputGin struct {
	*inputGinBase
}

// NewInputGin input handler constructor
func NewInputGin(service user.IService, logger log.ILogger) *InputGin {
	return &InputGin{inputGinBase: newInputGinBase(service, logger)}
}

// Register adds the urls binding it to service methods in the provided Gin Router
func (h *InputGin) Register(g gin.IRouter, middlewares ...*gin_utils.TorpedoMiddleware) {
	h.register(g, middlewares...)
}
