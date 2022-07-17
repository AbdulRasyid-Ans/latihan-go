package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Status   string
	Payload  []Employee
	Query    Query
	Optional Optional
}

type Employee struct {
	ID    string
	Name  string
	Email string
}

type Query struct {
	Limit      uint
	Page       uint
	Total_Page uint
}

type Optional struct {
	Config string
}

func main() {
	jsonData := `
	{
		"status": "success",
		"payload": [
			{
				"id": "1",
				"name": "Haccktiv",
				"email": "email@email.com"
			},{
				"id": "2",
				"name": "Haccktiv2",
				"email": "email2@email.com"
			}
		],
		"query": {
			"limit": 10,
			"page": 1,
			"total_page": 2
		},
		"optional": {
			"config": "nanananana"
		}
	}
	`

	var emp Data

	err := json.Unmarshal([]byte(jsonData), &emp)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("%+v\n", emp)
	fmt.Println(emp.Optional)
}
