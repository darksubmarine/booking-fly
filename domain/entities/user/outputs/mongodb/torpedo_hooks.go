// Code generated by TORPEDO DO NOT EDIT.

// Package mongodb is an output adapter to store entities in MongoDB
package mongodb

import "github.com/darksubmarine/torpedo-lib-go/tql"

// ISaveHooks interface that defines the before and after methods to be call at db save operation
type ISaveHooks interface {
	BeforeSave(dmo *EntityDMOMongoDB) error
	AfterSave(dmo *EntityDMOMongoDB, err error) error
}

// IFetchByIdHooks defines the before and after methods to be call at db fetch by id operation
type IFetchByIdHooks interface {
	BeforeFetchById(id string) error
	AfterFetchById(dmo *EntityDMOMongoDB, err error) error
}

// IUpdateHooks interface that defines the before and after methods to be call at db update operation
type IUpdateHooks interface {
	BeforeUpdate(dmo *EntityDMOMongoDB) error
	AfterUpdate(dmo *EntityDMOMongoDB, err error) error
}

// IDeleteByIdHooks defines the before and after methods to be call at db delete by id operation
type IDeleteByIdHooks interface {
	BeforeDeleteById(id string) error
	AfterDeleteById(err error) error
}

// IDeleteByHooks defines the before and after methods to be call at db delete by field operation
type IDeleteByHooks interface {
	BeforeDeleteBy(field, val string) error
	AfterDeleteBy(count int64, err error) error
}

// IQueryHooks defines the before and after methods to be call at db query operation
type IQueryHooks interface {
	BeforeQuery(q *tql.Query) error
	AfterQuery(err error) error
}

// NoopHooks implement the hook interfaces to call when a dev doesn't set it.
type NoopHooks struct{}

func (o *NoopHooks) BeforeSave(dmo *EntityDMOMongoDB) error                { return nil }
func (o *NoopHooks) AfterSave(dmo *EntityDMOMongoDB, err error) error      { return nil }
func (o *NoopHooks) BeforeFetchById(id string) error                       { return nil }
func (o *NoopHooks) AfterFetchById(dmo *EntityDMOMongoDB, err error) error { return nil }
func (o *NoopHooks) BeforeUpdate(dmo *EntityDMOMongoDB) error              { return nil }
func (o *NoopHooks) AfterUpdate(dmo *EntityDMOMongoDB, err error) error    { return nil }
func (o *NoopHooks) BeforeDeleteById(id string) error                      { return nil }
func (o *NoopHooks) AfterDeleteById(err error) error                       { return nil }
func (o *NoopHooks) BeforeDeleteBy(field, val string) error                { return nil }
func (o *NoopHooks) AfterDeleteBy(count int64, err error) error            { return nil }
func (o *NoopHooks) BeforeQuery(q *tql.Query) error                        { return nil }
func (o *NoopHooks) AfterQuery(err error) error                            { return nil }

// Hooks struct used as parameter to set the user builder hooks in the builder instance
type Hooks struct {
	Save       func() ISaveHooks
	FetchById  func() IFetchByIdHooks
	Update     func() IUpdateHooks
	DeleteById func() IDeleteByIdHooks
	DeleteBy   func() IDeleteByHooks
	Query      func() IQueryHooks
}

// HookBuilder struct called from the mongodb repo object. If the user has no set the desired builder function, a noop is set by default
type HookBuilder struct {
	noop *NoopHooks

	save       func() ISaveHooks
	fetchById  func() IFetchByIdHooks
	update     func() IUpdateHooks
	deleteById func() IDeleteByIdHooks
	deleteBy   func() IDeleteByHooks
	query      func() IQueryHooks
}

// NewHookBuilder creates a hooks builder with a noop hook by default which can be overwritten by developers via hooks parameter
func NewHookBuilder(hooks *Hooks) *HookBuilder {
	hb := new(HookBuilder)
	hb.noop = new(NoopHooks)

	if hooks == nil {
		return hb
	}

	if hooks.Save != nil {
		hb.save = hooks.Save
	}

	if hooks.FetchById != nil {
		hb.fetchById = hooks.FetchById
	}

	if hooks.Update != nil {
		hb.update = hooks.Update
	}

	if hooks.DeleteById != nil {
		hb.deleteById = hooks.DeleteById
	}

	if hooks.DeleteBy != nil {
		hb.deleteBy = hooks.DeleteBy
	}

	if hooks.Query != nil {
		hb.query = hooks.Query
	}

	return hb
}

func (h *HookBuilder) Save() ISaveHooks {
	if h.save == nil {
		return h.noop
	}
	return h.save()
}

func (h *HookBuilder) FetchById() IFetchByIdHooks {
	if h.fetchById == nil {
		return h.noop
	}
	return h.fetchById()
}

func (h *HookBuilder) Update() IUpdateHooks {
	if h.update == nil {
		return h.noop
	}
	return h.update()
}

func (h *HookBuilder) DeleteById() IDeleteByIdHooks {
	if h.deleteById == nil {
		return h.noop
	}
	return h.deleteById()
}

func (h *HookBuilder) DeleteBy() IDeleteByHooks {
	if h.deleteBy == nil {
		return h.noop
	}
	return h.deleteBy()
}

func (h *HookBuilder) Query() IQueryHooks {
	if h.query == nil {
		return h.noop
	}
	return h.query()
}
