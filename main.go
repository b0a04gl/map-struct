package main

import (
	"fmt"
	"github.com/map-struct/mapper"
)

func main() {
	data := map[string]interface{}{
		"Name":    "Tommy Shelby",
		"Age":     34,
		"Address": "441, Baker's street, Buckingham",
	}

	var person mapper.Person
	err := mapper.MapToStruct(data, &person)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Mapped Person:", person)
	}
}
