package models
type Employee struct{
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Age   int    `db:"age" json:"age"`
}

type EmployeeRequest struct{
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Age   int    `db:"age" json:"age"`
}