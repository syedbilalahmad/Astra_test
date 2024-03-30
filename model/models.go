package model

type Data struct {
	Emp []Employee `json:"data"`
}

type Employee struct {
	Email  string
	Name   string
	Mobile string
}
