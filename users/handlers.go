// This file defines view handlers for the users module

package users

import (
	"fmt"
	"time"
	"strconv"
	"net/http"
	"html/template"
	"database/sql"

	"github.com/gorilla/mux"
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

// UsersRootHandler godoc
// @Summary Serves the /users resource
// @Description Renders the index.html template for the /users resource, showing all users and a form to query a user
// @ID get-all-users
// @Produce html
// @Success 200 {object} string
// @Router /users [get]
func (self *UsersHandler) UsersRootHandler(writer http.ResponseWriter, request *http.Request) {
	user := createTestUser()
	allUsers := GetAllUsers(*self.DbConn)

	data := DataMap {
		"testUser": user,
		"allUsers": allUsers,
	}

	executeTemplate(writer, request, data, "users/index.html")
}

// UserDetailHandler godoc
// @Summary Serves the /users/{id} resource,
// @Description Renders the user.html template for the /users/{id} resource, showing the user by specified path variable {id}
// @ID get-user-by-id
// @Param id path int true "User ID"
// @Produce html
// @Success 200 {object} string
// @Failure 500
// @Router /users/{id} [get]
func (self *UsersHandler) UserDetailHandler(writer http.ResponseWriter, request *http.Request) {
	params := mux.Vars(request)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		self.NotFoundHandler(writer, request)
	} else {
		data := DataMap {
			"user": GetUserById(*self.DbConn, id),
		}
		executeTemplate(writer, request, data, "users/user.html")
	}
}

// NotFoundHandler is the handler for the "404 - Not Found" page
func (self *UsersHandler) NotFoundHandler(writer http.ResponseWriter, request *http.Request) {
	executeTemplate(writer, request, nil, "users/404.html")
}
