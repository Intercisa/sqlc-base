-- name: CreateEmployees :one
INSERT INTO employees (first_name, last_name, email, phone_number, hire_date, job_title, salary, department_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: CreateDepartmant :one
INSERT INTO departments (name) VALUES ($1)
RETURNING *;

-- name: SelectEmployee :one
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
    employees.id = $1;