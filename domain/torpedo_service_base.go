// Code generated by Torpedo DO NOT EDIT.

// Package domain domain entry point
package domain

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
	"github.com/darksubmarine/booking-fly/domain/entities/user"
)

type iDomainServiceBase interface {
	Context() *Context
	Trip() trip.IService
	User() user.IService
}

type serviceBase struct {
	ctx *Context
}

func (s *serviceBase) Context() *Context {
	return s.ctx
}

func (s *serviceBase) Trip() trip.IService {
	return s.ctx.TripSrv
}
func (s *serviceBase) User() user.IService {
	return s.ctx.UserSrv
}
