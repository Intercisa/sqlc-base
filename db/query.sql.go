// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: query.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createDepartmant = `-- name: CreateDepartmant :one
INSERT INTO departments (name) VALUES ($1)
RETURNING id, name
`

func (q *Queries) CreateDepartmant(ctx context.Context, name string) (Department, error) {
	row := q.db.QueryRow(ctx, createDepartmant, name)
	var i Department
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const createEmployees = `-- name: CreateEmployees :one
INSERT INTO employees (first_name, last_name, email, phone_number, hire_date, job_title, salary, department_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id, first_name, last_name, email, phone_number, hire_date, job_title, salary, department_id
`

type CreateEmployeesParams struct {
	FirstName    string
	LastName     string
	Email        string
	PhoneNumber  pgtype.Text
	HireDate     pgtype.Date
	JobTitle     string
	Salary       pgtype.Numeric
	DepartmentID pgtype.Int4
}

func (q *Queries) CreateEmployees(ctx context.Context, arg CreateEmployeesParams) (Employee, error) {
	row := q.db.QueryRow(ctx, createEmployees,
		arg.FirstName,
		arg.LastName,
		arg.Email,
		arg.PhoneNumber,
		arg.HireDate,
		arg.JobTitle,
		arg.Salary,
		arg.DepartmentID,
	)
	var i Employee
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.HireDate,
		&i.JobTitle,
		&i.Salary,
		&i.DepartmentID,
	)
	return i, err
}

const selectEmployee = `-- name: SelectEmployee :one
SELECT
    employees.id AS employee_id,
    employees.first_name,
    employees.last_name,
    employees.email,
    employees.phone_number,
    employees.hire_date,
    employees.job_title,
    employees.salary,
    departments.id AS department_id,
    departments.name AS department_name
FROM
    employees
JOIN
    departments
ON
    employees.department_id = departments.id
WHERE
    employees.id = $1
`

type SelectEmployeeRow struct {
	EmployeeID     int32
	FirstName      string
	LastName       string
	Email          string
	PhoneNumber    pgtype.Text
	HireDate       pgtype.Date
	JobTitle       string
	Salary         pgtype.Numeric
	DepartmentID   int32
	DepartmentName string
}

func (q *Queries) SelectEmployee(ctx context.Context, id int32) (SelectEmployeeRow, error) {
	row := q.db.QueryRow(ctx, selectEmployee, id)
	var i SelectEmployeeRow
	err := row.Scan(
		&i.EmployeeID,
		&i.FirstName,
		&i.LastName,
		&i.Email,
		&i.PhoneNumber,
		&i.HireDate,
		&i.JobTitle,
		&i.Salary,
		&i.DepartmentID,
		&i.DepartmentName,
	)
	return i, err
}
