// Package memory is an output adapter to store entities in memory
package memory

import (
	"github.com/darksubmarine/booking-fly/domain/entities/user"
)

type DMO struct {
	*user.EntityDMO // Do not remove it. This will let you add custom encrypted fields and more.

	/* your custom fields goes here */
}
