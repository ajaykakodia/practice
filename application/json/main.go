package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Age   int    `json:"age"`
}

func main() {
	p1 := person{
		First: "Ajay",
		Last:  "Yadav",
		Age:   35,
	}

	p2 := person{
		First: "Rekha",
		Last:  "Yadav",
		Age:   35,
	}

	persons := []person{p1, p2}
	pjson, err := json.Marshal(persons)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Printf("%+v\n", persons)
	fmt.Println(string(pjson))

	pdata := `[{"first":"Ajay","last":"Yadav","age":35},{"first":"Rekha","last":"Yadav","age":35}]`
	pbytes := []byte(pdata)

	fmt.Printf("%T\n", pdata)
	fmt.Printf("%T\n", pbytes)

	var prsns []person
	err = json.Unmarshal(pbytes, &prsns)
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Printf("%+v\n", prsns)
}
