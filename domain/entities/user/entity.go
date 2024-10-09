// Package user The user system
package user

// UserEntity The user system
type UserEntity struct {
	*entityBase // DO NOT REMOVE IT
}

// New is a UserEntity constructor function
func New() *UserEntity {
	return &UserEntity{entityBase: newEntityBase()}
}
