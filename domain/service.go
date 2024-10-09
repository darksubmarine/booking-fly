// Package domain domain entry point
package domain

type IDomainService interface {
	iDomainServiceBase
}

type Service struct {
	*serviceBase
}

func NewService(ctx *Context) *Service {
	return &Service{serviceBase: &serviceBase{ctx: ctx}}
}
