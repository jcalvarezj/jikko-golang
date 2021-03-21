// This file defines the DAO functions for the models

package users

import (
	"fmt"
	"database/sql"
)

// GetAllUsers returns a slice containing instances of User related to all the 
// records in the "user" table. If there are no records, an empty slice is returned
func GetAllUsers(dbConn *sql.DB) []User {
	var allUsers []User
	results, err := (*dbConn).Query("SELECT username, first_name, last_name, birthday, sex, biography FROM user")
	if err != nil {
		fmt.Println("An error has occurred when performing the query. ", err)
		return []User {}
	} else {
		for results.Next() {
			var user User
			scanErr := results.Scan(&user.Username, &user.FirstName, &user.LastName,
						&user.Birthday, &user.Sex, &user.Biography)
			if scanErr != nil {
				fmt.Println("An error has occurred when retrieving data. ", scanErr)
			}
			allUsers = append(allUsers, user)
		}
		return allUsers
	}
}

// GetUserById returns an instance of User, found by id in the "user" table
func GetUserById(dbConn *sql.DB, id int) *User {
	results, err := (*dbConn).Query("SELECT username, first_name, last_name, birthday, sex, biography FROM user WHERE id = ?", id)
	if err != nil {
		fmt.Println("An error has occurred when performing the query. ", err)
		return nil
	} else {
		var user User
		success := results.Next()
		if !success {
			return nil
		} else {
			err := results.Scan(&user.Username, &user.FirstName, &user.LastName,
						&user.Birthday, &user.Sex, &user.Biography)
			switch err {
				case sql.ErrNoRows:
					fmt.Println("No rows were found")
					return nil
				case nil:
					return &user
				default:
					fmt.Println("An error has occurred when retrieving data. ", err)
					return nil
			}
		}
	}
}
