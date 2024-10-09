// Package trip The user fly trip reservations
package trip

import (
	"github.com/darksubmarine/torpedo-lib-go/log"
)

// IService interface that aggregates IServiceBase. All of your use cases should be added here.
type IService interface {
	IServiceBase // DO NOT REMOVE this line. IServiceBase defines all CRUD operations for the entity

	/* Add your use cases here */

}

// Service defines your use cases. Extends from ServiceBase to get the CRUD operations
type Service struct {
	*ServiceBase // DO NOT REMOVE this line. ServiceBase implements IServiceBase interface
}

// NewService Service constructor function
func NewService(repo IRepository, logger log.ILogger) *Service {
	srv := &Service{ServiceBase: newServiceBase(repo, logger)}
	srv.initHooks()
	return srv
}

// initHooks this method allow you to set the different hooks that are called on each step of the entity life cycle.
func (s *Service) initHooks() {
	/*
		MODIFY THIS CODE TO SET YOUR HOOKS HERE
	*/

	s.hookBuilder = newServiceHooksBuilder(

		// builder hooks function for create operation
		func() iServiceCreateHooks {
			return newNoopServiceHooks()
		},

		// builder hooks function for read operation
		func() iServiceReadHooks {
			return newNoopServiceHooks()
		},

		// builder hooks function for update operation
		func() iServiceUpdateHooks {
			return newNoopServiceHooks()
		},

		// builder hooks function for delete operation
		func() iServiceDeleteHooks {
			return newNoopServiceHooks()
		},
	)
}
