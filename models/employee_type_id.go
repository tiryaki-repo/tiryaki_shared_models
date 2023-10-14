package models

type EmployeeTypeId int

const (
	Admin   EmployeeTypeId = EmployeeTypeId(1000)
	Manager EmployeeTypeId = EmployeeTypeId(2000)
	Staff   EmployeeTypeId = EmployeeTypeId(3000)
)
