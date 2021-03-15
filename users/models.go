// This file defines models for the users module

package users

import (
	"time"
	"fmt"
)

// User represents the user entity
type User struct {
	Username string
	FirstName string
	LastName string
	Birthday time.Time
	Sex string
	Biography string
}

// GetFullName returns the Full Name as "FirstName LastName"
func (user User) GetFullName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
