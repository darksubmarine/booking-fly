// Code generated by Torpedo DO NOT EDIT.

// Package user The user system
package user

import "github.com/darksubmarine/torpedo-lib-go/context"

type iServiceCreateHooks interface {
	beforeCreate(ctx context.IDataMap, entity *UserEntity) error
	afterCreate(ctx context.IDataMap, entity *UserEntity) error
}

type iServiceReadHooks interface {
	beforeRead(ctx context.IDataMap, id string) error
	afterRead(ctx context.IDataMap, entity *UserEntity) error
}

type iServiceUpdateHooks interface {
	beforeUpdate(ctx context.IDataMap, entity *UserEntity) error
	afterUpdate(ctx context.IDataMap, entity *UserEntity) error
}

type iServiceDeleteHooks interface {
	beforeDelete(ctx context.IDataMap, id string) error
	afterDelete(ctx context.IDataMap, id string) error
}

// serviceHooks object
type serviceHookBuilder struct {
	noop   *noopServiceHooks
	create func() iServiceCreateHooks
	read   func() iServiceReadHooks
	update func() iServiceUpdateHooks
	delete func() iServiceDeleteHooks
}

// newServiceHooksBuilder Service Hooks (CRUD) builder constructor
func newServiceHooksBuilder(createFn func() iServiceCreateHooks, readFn func() iServiceReadHooks,
	updateFn func() iServiceUpdateHooks, deleteFn func() iServiceDeleteHooks) *serviceHookBuilder {

	noop := newNoopServiceHooks()

	return &serviceHookBuilder{
		noop: newNoopServiceHooks(),
		create: func() iServiceCreateHooks {
			if createFn == nil {
				return noop
			}
			return createFn()
		},
		read: func() iServiceReadHooks {
			if readFn == nil {
				return noop
			}
			return readFn()
		},
		update: func() iServiceUpdateHooks {
			if updateFn == nil {
				return noop
			}
			return updateFn()
		},
		delete: func() iServiceDeleteHooks {
			if deleteFn == nil {
				return noop
			}
			return deleteFn()
		},
	}
}

// noopServiceHooks no operational service hooks
type noopServiceHooks struct{}

func newNoopServiceHooks() *noopServiceHooks { return &noopServiceHooks{} }

func (n *noopServiceHooks) beforeCreate(ctx context.IDataMap, entity *UserEntity) error { return nil }
func (n *noopServiceHooks) afterCreate(ctx context.IDataMap, entity *UserEntity) error  { return nil }

func (n *noopServiceHooks) beforeUpdate(ctx context.IDataMap, entity *UserEntity) error { return nil }
func (n *noopServiceHooks) afterUpdate(ctx context.IDataMap, entity *UserEntity) error  { return nil }

func (n *noopServiceHooks) beforeDelete(ctx context.IDataMap, id string) error { return nil }
func (n *noopServiceHooks) afterDelete(ctx context.IDataMap, id string) error  { return nil }

func (n *noopServiceHooks) beforeRead(ctx context.IDataMap, id string) error         { return nil }
func (n *noopServiceHooks) afterRead(ctx context.IDataMap, entity *UserEntity) error { return nil }
