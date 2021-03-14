// This file defines view handlers for the users module

package users

import (
	"fmt"
	"time"
	"net/http"
	"html/template"
	"database/sql"
)

// UsersHandler represents the controller for the User model
type UsersHandler struct {
	DbConn *sql.DB
}

// DataMap is a type for passing data to templates
type DataMap map[string]interface{}

const BASE_DATE = "2006-01-02"

// NewUsersHandler creates an instance of the UsersHandler controller that stores
// the database connection instance and returns the pointer to it
func NewUsersHandler(connection *sql.DB) *UsersHandler {
	return &UsersHandler{DbConn: connection}
}

// createTestUser creates a test user to be displayed on the HTML templates
func createTestUser() User {
	dateOfBirth, _ := time.Parse(BASE_DATE, "1992-12-31")
	return User {
		Username: "pepelefrog",
		FirstName: "Pepe",
		LastName: "Perez",
		Birthday: dateOfBirth,
		Sex: "M",
		Biography: "Meme",
	}
}

// executeTemplate parses and executes a template for a request on the specified
// template file, with the entered data struct
func executeTemplate(writer http.ResponseWriter, request *http.Request, data DataMap, file string) {
	parsedTemplate, template_err := template.ParseFiles(file)

	if template_err == nil {
		err := parsedTemplate.Execute(writer, data)
		if err != nil {
			fmt.Println("Error executing template", err)
		}
	} else {
		fmt.Println("Error parsing template", template_err)
	}
}

// UsersRootHandler renders the index.html template for the /users resource
func (self *UsersHandler) UsersRootHandler(writer http.ResponseWriter, request *http.Request) {
	user := createTestUser()
	allUsers := GetAllUsers(*self.DbConn)

	data := DataMap {
		"testUser": user,
		"allUsers": allUsers,
	}

	executeTemplate(writer, request, data, "users/index.html")
}

// UserDetailHandler renders the user.html template for the /users/{id} resource,
// sending it the user struct found for that id
func (self *UsersHandler) UserDetailHandler(writer http.ResponseWriter, request *http.Request) {
	data := DataMap {
		"user": createTestUser(),
	}
	executeTemplate(writer, request, data, "users/user.html")
}
