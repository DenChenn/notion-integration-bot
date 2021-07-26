package model

type Field struct {
	Key string
	Value string
}

type DepartmentDetail struct {
	Action string
	AssigneeEmail string
	Title string
	FieldSet []Field
}