// Code generated by Torpedo DO NOT EDIT.

// Package domain domain entry point
package domain

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	"github.com/darksubmarine/booking-fly/domain/entities/user"
)

type contextBase struct {
	TripSrv trip.IService
	UserSrv user.IService
}
