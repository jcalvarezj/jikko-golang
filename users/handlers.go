// This file defines view handlers for the users module

package users

import (
	"fmt"
	"time"
	"net/http"
	"html/template"
)

const BASE_DATE = "2006-01-02"

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

func executeTemplate(writer http.ResponseWriter, request *http.Request, user *User, file string) {
	parsedTemplate, template_err := template.ParseFiles(file)

	if template_err == nil {
		err := parsedTemplate.Execute(writer, *user)
		if err != nil {
			fmt.Println("Error executing template", err)
		}
	} else {
		fmt.Println("Error parsing template", template_err)
	}
}

func UsersRootHandler(writer http.ResponseWriter, request *http.Request) {
	user := createTestUser()
	executeTemplate(writer, request, &user, "users/index.html")
}

func UserDetailHandler(writer http.ResponseWriter, request *http.Request) {
	user := createTestUser()
	executeTemplate(writer, request, &user, "users/user.html")
}
