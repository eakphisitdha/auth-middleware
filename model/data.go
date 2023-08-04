package model

import "database/sql"

type Data struct {
	ID    int
	Name  string
	Field sql.NullString
	//
	// add your data
	//
}

type DataResponse struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Field string `json:"field"`
	//
	// add your data
	//
}
