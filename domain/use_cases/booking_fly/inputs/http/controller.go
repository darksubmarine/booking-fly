package http

import (
	"errors"
	"fmt"
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	tripHTTP "github.com/darksubmarine/booking-fly/domain/entities/trip/inputs/http/gin"
	"github.com/darksubmarine/booking-fly/domain/use_cases/booking_fly"
	"github.com/darksubmarine/torpedo-lib-go/api"
	"github.com/darksubmarine/torpedo-lib-go/entity"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Controller struct that handles HTTP Requests for the use case
type Controller struct {
	logger       log.ILogger
	ucBookingFly *booking_fly.UseCase
}

// NewController controller constructor function
func NewController(useCase *booking_fly.UseCase, logger log.ILogger) *Controller {
	return &Controller{ucBookingFly: useCase, logger: logger}
}

// BookingFlyEndpoint HTTP handler function that calls the use case DoBooking method.
// @Summary Books a fly
// @Schemes http https
// @Description Books a fly given a user and the trip information
// @Tags UseCases
// @Accept json
// @Produce json
// @Param trip body BookingFlyDTO true "The user fly trip reservation"
// @Success 200 {object} trip.FullDTO
// @Failure 400 {object} api.Error
// @Failure 500 {object} api.Error
// @Router /booking [post]
func (c *Controller) BookingFlyEndpoint(ctx *gin.Context) {

	var dto BookingFlyDTO
	if err := ctx.ShouldBindJSON(&dto); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorBindingJSON(err))
		return
	}

	tripModel := trip.New()
	if err := entity.From(&dto, tripModel); err != nil {
		ctx.JSON(http.StatusBadRequest, api.ErrorBuildingEntityFromDTO(err))
		return
	}

	tripCreated, err := c.ucBookingFly.DoBooking(tripModel)
	if err != nil {

		if errors.Is(err, booking_fly.ErrUserNotFound) {
			ctx.JSON(http.StatusBadRequest, api.NewError("4007", err))
		} else {
			ctx.JSON(http.StatusInternalServerError, api.ErrorEntityCreation(err))
		}
		return
	}

	var tripDTO = tripHTTP.NewFullDTO()
	if err := entity.To(tripCreated, tripDTO); err != nil {
		ctx.JSON(http.StatusBadRequest, api.NewError("4006", fmt.Errorf("error copying model to DTO: %w", err)))
		return
	}

	ctx.JSON(http.StatusOK, tripDTO)
}
