// Package mongodb is an output adapter to store entities in MongoDB
package mongodb

import (
	"github.com/darksubmarine/booking-fly/domain/entities/trip"
)

type DMO struct {
	*trip.EntityDMO `bson:"-"` // Do not remove it. This will let you add custom encrypted fields and more.

	/* your custom fields goes here */
}
