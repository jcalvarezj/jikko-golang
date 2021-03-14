// This file defines models for the users module

package users

import (
	"time"
	"fmt"
)

type User struct {
	Username string
	FirstName string
	LastName string
	Birthday time.Time
	Sex string
	Biography string
}

func (user User) GetFullName() string {
	return fmt.Sprintf("%s %s", user.FirstName, user.LastName)
}
