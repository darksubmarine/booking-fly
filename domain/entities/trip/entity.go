// Package trip The user fly trip reservations
package trip

// TripEntity The user fly trip reservations
type TripEntity struct {
	*entityBase // DO NOT REMOVE IT
}

// New is a TripEntity constructor function
func New() *TripEntity {
	return &TripEntity{entityBase: newEntityBase()}
}
