// This file helps create the Database connection

package repository

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

const RDBMS = "mysql"
const DATABASE = "jikkodb"
const USER = "jikkouser"
const PASS = "jikkopass"
const HOST = "127.0.0.1:3306"

// ConnectToDatabase attempts connection with the database and returns both
// the connection instance and the error
func ConnectToDatabase() (*sql.DB, error){
	connStr := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", USER, PASS, HOST, DATABASE)
	return sql.Open(RDBMS, connStr)
}
