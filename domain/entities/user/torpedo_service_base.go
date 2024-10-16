// Code generated by Torpedo DO NOT EDIT.

// Package user The user system
package user

import (
	"fmt"
	"github.com/darksubmarine/torpedo-lib-go"
	"github.com/darksubmarine/torpedo-lib-go/context"
	"github.com/darksubmarine/torpedo-lib-go/entity"
	"github.com/darksubmarine/torpedo-lib-go/log"
	"github.com/darksubmarine/torpedo-lib-go/tql"

	"github.com/darksubmarine/booking-fly/domain/entities/trip"

	"reflect"
	"runtime"
)

type hookFn func(ctx context.IDataMap, entity *UserEntity) error
type hookFnById func(ctx context.IDataMap, id string) error

// ServiceBase implements the interface IServiceBase
type ServiceBase struct {
	repo   IRepository
	logger log.ILogger

	// hooks
	hookBuilder *serviceHookBuilder

	// NestedLoading relationships

	tripSrv trip.IService
}

// newServiceBase internal constructor
func newServiceBase(repo IRepository, logger log.ILogger, tripSrv trip.IService) *ServiceBase {
	return &ServiceBase{repo: repo, logger: logger,
		tripSrv: tripSrv}
}

// execHook calls hooks with the entity as param
func (s *ServiceBase) execHook(hFn hookFn, ctx context.IDataMap, entity *UserEntity) (errHook error) {

	defer func() {
		if r := recover(); r != nil {
			hFnName := runtime.FuncForPC(reflect.ValueOf(hFn).Pointer()).Name()
			errHook = fmt.Errorf("recovered from hook %s. Error: %s", hFnName, r)
			s.logger.Error("execHook recovery", "error", errHook, "hook", hFnName)
		}
	}()

	if hFn != nil {
		if err := hFn(ctx, entity); err != nil {
			hFnName := runtime.FuncForPC(reflect.ValueOf(hFn).Pointer()).Name()
			s.logger.Error("execHook error", "error", err, "hook", hFnName)
			return err
		}
	}

	return nil
}

// execHookById calls the hooks with the entity id as param
func (s *ServiceBase) execHookById(hFn hookFnById, ctx context.IDataMap, id string) (errHook error) {

	defer func() {
		if r := recover(); r != nil {
			hFnName := runtime.FuncForPC(reflect.ValueOf(hFn).Pointer()).Name()
			errHook = fmt.Errorf("recovered from hook %s. Error: %s", hFnName, r)
			s.logger.Error("execHookById recovery", "error", errHook, "hook", hFnName)
		}
	}()

	if hFn != nil {
		if err := hFn(ctx, id); err != nil {
			hFnName := runtime.FuncForPC(reflect.ValueOf(hFn).Pointer()).Name()
			s.logger.Error("execHookById error", "error", err, "hook", hFnName)
			return err
		}
	}

	return nil
}

// hookCreate fetchs the new hook instance from the builder
func (s *ServiceBase) hookCreate() iServiceCreateHooks {
	if s.hookBuilder != nil {
		return s.hookBuilder.create()
	}

	return newNoopServiceHooks()
}

// hookRead fetchs the new hook instance from the builder
func (s *ServiceBase) hookRead() iServiceReadHooks {
	if s.hookBuilder != nil {
		return s.hookBuilder.read()
	}

	return newNoopServiceHooks()
}

// hookUpdate fetchs the new hook instance from the builder
func (s *ServiceBase) hookUpdate() iServiceUpdateHooks {
	if s.hookBuilder != nil {
		return s.hookBuilder.update()
	}

	return newNoopServiceHooks()
}

// hookDelete fetchs the new hook instance from the builder
func (s *ServiceBase) hookDelete() iServiceDeleteHooks {
	if s.hookBuilder != nil {
		return s.hookBuilder.delete()
	}

	return newNoopServiceHooks()
}

// Create given a new entity this one is populated with ID and creation timestamp and finally saved into the repository
func (s *ServiceBase) Create(ctx context.IDataMap, ety *UserEntity) (*UserEntity, error) {

	hook := s.hookCreate()
	hookErr := s.execHook(hook.beforeCreate, ctx, ety)
	if hookErr != nil {
		s.logger.Error("before create hook", "error", hookErr)
	}

	now := torpedo_lib.TimeNow()
	toAdd := New()

	entity.Clone(ety, toAdd)

	// Metadata
	toAdd.id = newID()
	toAdd.created = now
	toAdd.updated = now

	// TODO check if field is not in storage... constrain from YAML (unique: true) ... or delegate it to storage engine?
	//      TQL is needed to achieve this goal.

	s.logger.Debug("before save to repo", "ety", toAdd.String())
	if err := s.repo.Save(toAdd); err != nil {
		s.logger.Error("saving to repo", "error", err, "ety", toAdd.String())
		return nil, err
	}

	hookErr = s.execHook(hook.afterCreate, ctx, toAdd)
	if hookErr != nil {
		s.logger.Error("after create hook", "error", hookErr)
	}

	return toAdd, nil
}

// Read returns a pointer to UserEntity given its id
func (s *ServiceBase) Read(ctx context.IDataMap, id string) (*UserEntity, error) {
	hook := s.hookRead()
	hookErr := s.execHookById(hook.beforeRead, ctx, id)
	if hookErr != nil {
		s.logger.Error("before read hook", "error", hookErr)
	}

	s.logger.Debug("reading from repo", "id", id)
	entity, err := s.repo.FetchByID(id)
	if err != nil {
		s.logger.Error("reading from repo", "id", id, "error", err)
		return nil, err
	}

	s.nestedLoadingTrip(entity)

	hookErr = s.execHook(hook.afterRead, ctx, entity)
	if hookErr != nil {
		s.logger.Error("after read hook", "error", hookErr)
	}
	return entity, nil
}

// Update returns a pointer to UserEntity after update it
func (s *ServiceBase) Update(ctx context.IDataMap, ety *UserEntity) (*UserEntity, error) {
	hook := s.hookUpdate()
	hookErr := s.execHook(hook.beforeUpdate, ctx, ety)
	if hookErr != nil {
		s.logger.Error("before update hook", "error", hookErr)
	}

	current, err := s.repo.FetchByID(ety.Id())
	if err != nil {
		return nil, err
	}

	var toUpdate = New()
	entity.Clone(ety, toUpdate)

	// Metadata
	toUpdate.created = current.created
	toUpdate.updated = torpedo_lib.TimeNow()

	s.logger.Debug("before update to repo", "ety", toUpdate.String())
	if err := s.repo.Update(toUpdate); err != nil {
		s.logger.Error("updating to repo", "error", err, "ety", toUpdate.String())
		return nil, err
	}

	hookErr = s.execHook(hook.afterUpdate, ctx, toUpdate)
	if hookErr != nil {
		s.logger.Error("after update hook", "error", hookErr)
	}

	return toUpdate, nil
}

// Delete removes the entity given its id
func (s *ServiceBase) Delete(ctx context.IDataMap, id string) error {
	hook := s.hookDelete()
	hookErr := s.execHookById(hook.beforeDelete, ctx, id)
	if hookErr != nil {
		s.logger.Error("before delete hook", "error", hookErr)
	}

	s.logger.Debug("deleting from repo", "id", id)
	err := s.repo.DeleteByID(id)
	if err != nil {
		s.logger.Error("deleting from repo", "error", err, "id", id)
		return err // prevents call the after delete hook due to an error
	}

	hookErr = s.execHookById(hook.afterDelete, ctx, id)
	if hookErr != nil {
		s.logger.Error("before delete hook", "error", hookErr)
	}

	return err
}

// Query executes the given query (TQL) and returns the query result
func (s *ServiceBase) Query(q *tql.Query) (*tql.Result, error) {

	if err := q.Validate(_fieldsMap); err != nil {
		return nil, err
	}

	if q.IsCursorPagination() && len(q.Projection) > 0 && q.CursorPaginationSort() != nil {
		sortField := q.CursorPaginationSort().Field
		isOnProjection := false
		for _, prjField := range q.Projection {
			if prjField == sortField {
				isOnProjection = true
				break
			}
		}
		if !isOnProjection {
			return nil, tql.ErrInvalidSortFieldNotProjectionMember
		}
	}

	results, err := s.repo.Query(q, _fieldsMap)
	if err != nil {
		return nil, err
	}

	qros := make([]interface{}, len(results))
	for i, ue := range results {
		qro := &EntityQRO{}

		// Projection field names to QRO attributes
		normalizedFields := make([]string, len(q.Projection))
		for j, f := range q.Projection {
			normalizedFields[j] = _fieldsMetadata[f].QroName()
		}

		if err := qro.HydrateFromEntity(ue, normalizedFields...); err != nil {
			return nil, tql.ErrQueryResultObjectBuild
		}
		qros[i] = qro
	}

	var count = int64(len(qros))
	var prev interface{}
	var next interface{}

	if q.HasPagination() {
		if q.IsOffsetPagination() {
			if q.PaginationOffsetPage() > 1 {
				prev = q.PaginationOffsetPage() - 1
			}

			if count > q.PaginationRAWItems() {
				qros = qros[:count-1] // removing last item used to check if we have next page
				next = q.PaginationOffsetPage() + 1
			}
		} else if q.IsCursorPagination() {
			if len(qros) > 0 {
				var currentCursorSort = tql.PaginationCursorOrder(q.PaginationCursorOrder())
				if id := qros[len(qros)-1].(*EntityQRO).Id_; id != nil {
					var cursor *tql.Cursor
					if q.HasSort() {
						sortVal := qros[len(qros)-1].(*EntityQRO).FieldValue(q.CursorPaginationSort().Field)
						sortField := q.CursorPaginationSort().Field
						sortType := q.CursorPaginationSort().Kind
						cursor = tql.NewPaginationCursorFrom(*id, currentCursorSort, sortField, sortVal, sortType)
					} else {
						cursor = tql.NewPaginationCursorFromPivot(*id, currentCursorSort)
					}
					next = cursor.Next()
				}
			}
		}
	}

	tqlResult := tql.NewResult(qros, len(qros), prev, next)
	return tqlResult, nil
}

func (s *ServiceBase) nestedLoadingTrip(entity *UserEntity) error {

	trips, err := s.tripSrv.BelongsToUser(entity.id, 100, 1)
	if err != nil {
		return err
	}

	entity.trips = trips
	return nil
}
