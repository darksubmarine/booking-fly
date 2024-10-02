// Package domain domain entry point
package domain

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	"github.com/darksubmarine/booking-fly/domain/entities/user"
)

type Context struct {
	*contextBase
}

func NewContext(trip_ trip.IService, user_ user.IService) *Context {
	return &Context{&contextBase{
		TripSrv: trip_,
		UserSrv: user_,
	}}
}
