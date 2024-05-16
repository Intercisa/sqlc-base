// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Department struct {
	ID   int32
	Name string
}

type Employee struct {
	ID           int32
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  pgtype.Text
	HireDate     pgtype.Date
	JobTitle     string
	Salary       pgtype.Numeric
	DepartmentID pgtype.Int4
}
