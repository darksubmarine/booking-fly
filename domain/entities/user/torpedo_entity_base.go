// Code generated by Torpedo DO NOT EDIT.

// Package user The user system
package user

import (
	"bytes"
	"fmt"
	"github.com/darksubmarine/torpedo-lib-go/entity"
	"github.com/darksubmarine/torpedo-lib-go/trn"
	"github.com/darksubmarine/torpedo-lib-go/validator"

	"github.com/darksubmarine/booking-fly/domain/entities/trip"
)

const (
	fieldId       = "id"
	fieldCreated  = "created"
	fieldUpdated  = "updated"
	fieldName     = "name"
	fieldEmail    = "email"
	fieldPassword = "password"
	fieldPlan     = "plan"
	fieldMiles    = "miles"
)

var _fieldsMap entity.FieldMap
var _fieldsMetadata map[string]*entity.FieldMetadata

func init() {
	instance := New()
	_fieldsMap = entity.ToFieldMap(instance)
	_fieldsMetadata = entity.FieldsMetadata(instance)
}

// FieldsMetadata returns the user Entity fields metadata.
func FieldsMetadata() map[string]*entity.FieldMetadata { return _fieldsMetadata }

type entityBase struct {
	id      string
	created int64
	updated int64

	// name The user full name
	name string

	// email The user contact email
	email string

	// password The user system password
	password string

	// plan The user membership plan
	plan string

	// miles The accumulated flyer miles
	miles int64

	// NestedLoading relationships

	trips []*trip.TripEntity

	validators map[string]validator.IValidator
}

func newEntityBase() *entityBase {
	return new(entityBase).init()
}

func (e *entityBase) init() *entityBase {
	e.validators = map[string]validator.IValidator{}

	e.validators[fieldPlan] = validator.NewList([]string{"GOLD", "SILVER", "BRONZE"})

	return e
}

func (e *entityBase) FieldsMetadata() map[string]*entity.FieldMetadata { return _fieldsMetadata }

func (e *entityBase) TRN() *trn.TRN { return entity.TRN(Name, e.id) }

func (e *entityBase) Id() string     { return e.id }
func (e *entityBase) Created() int64 { return e.created }
func (e *entityBase) Updated() int64 { return e.updated }

func (e *entityBase) SetId(id string)          { e.id = id }
func (e *entityBase) SetCreated(created int64) { e.created = created }
func (e *entityBase) SetUpdated(updated int64) { e.updated = updated }

// Name The user full name
func (e *entityBase) Name() string { return e.name }

// SetName The user full name
func (e *entityBase) SetName(name string) error {

	e.name = name
	return nil
}

// Email The user contact email
func (e *entityBase) Email() string { return e.email }

// SetEmail The user contact email
func (e *entityBase) SetEmail(email string) error {

	e.email = email
	return nil
}

// Password The user system password
func (e *entityBase) Password() string { return e.password }

// SetPassword The user system password
func (e *entityBase) SetPassword(password string) error {

	e.password = password
	return nil
}

// Plan The user membership plan
func (e *entityBase) Plan() string { return e.plan }

// SetPlan The user membership plan
func (e *entityBase) SetPlan(plan string) error {

	if !e.validators[fieldPlan].Value(plan).IsValid() {
		return ErrInvalidPlan
	}

	e.plan = plan
	return nil
}

// Miles The accumulated flyer miles
func (e *entityBase) Miles() int64 { return e.miles }

// SetMiles The accumulated flyer miles
func (e *entityBase) SetMiles(miles int64) error {

	e.miles = miles
	return nil
}

// NestedLoading relationships

// Trips returns a list of trips
func (e *entityBase) Trips() []*trip.TripEntity {
	return e.trips
}

// SetTrips sets a list of trips
func (e *entityBase) SetTrips(trips []*trip.TripEntity) {
	e.trips = trips
}

// String returns the string representation of the entityBase
func (e *entityBase) String() string {
	buf := bytes.NewBufferString("")
	buf.WriteString("user.entityBase{ ")
	buf.WriteString(fmt.Sprintf("%s=%v ", "id", e.id))
	buf.WriteString(fmt.Sprintf("%s=%v ", "created", e.created))
	buf.WriteString(fmt.Sprintf("%s=%v ", "updated", e.updated))
	buf.WriteString(fmt.Sprintf("%s=%v ", "name", e.name))
	buf.WriteString(fmt.Sprintf("%s=%v ", "email", e.email))
	buf.WriteString(fmt.Sprintf("%s=%v ", "password", e.password))
	buf.WriteString(fmt.Sprintf("%s=%v ", "plan", e.plan))
	buf.WriteString(fmt.Sprintf("%s=%v ", "miles", e.miles))
	buf.WriteString("}")
	return buf.String()
}
