// Code generated by Torpedo DO NOT EDIT.

// Package trip The user fly trip reservations
package trip

import "github.com/darksubmarine/torpedo-lib-go"

// newID returns an ID string representation based on the yaml configuration: UUID or ULID
func newID() string {
	return torpedo_lib.Ulid()
}
