package services

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLDatastore struct {
	*sql.DB
}

var connection MySQLDatastore

func NewMySQLDatastore() (*MySQLDatastore, error) {

	connection, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/Hit_man")

	if err != nil {
		return nil, err
	}

	return &MySQLDatastore{
		DB: connection}, nil
}

// func GetdataStoreConnection() (*MySQLDatastore, error) {
// 	return nil, connection
// }

// func init() {
// 	connection = NewMySQLDatastore()
// }
