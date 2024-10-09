package http

import (
	tripHTTP "github.com/darksubmarine/booking-fly/domain/entities/trip/inputs/http/gin"
)

// BookingFlyDTO use case DTO to book a fly
type BookingFlyDTO struct {
	tripHTTP.UpdatableDTO
} //@name useCases.BookingFlyDTO
