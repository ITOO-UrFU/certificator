package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
)

func main() {

	jsonStr := `[{
         "Name": "Adam",
         "Age": 36,
         "Job": "CEO"
 }, {
         "Name": "Eve",
         "Age": 34,
         "Job": "CFO"
 }, {
         "Name": "Mike",
         "Age": 38,
         "Job": "COO"
 }]`

	type Person struct {
		Name string
		Age  int
		Job  string
	}
	var people []Person

	var personMap []map[string]interface{}

	err := json.Unmarshal([]byte(jsonStr), &personMap)

	if err != nil {
		panic(err)
	}

	for _, personData := range personMap {

		// convert map to array of Person struct
		var p Person
		p.Name = fmt.Sprintf("%s", personData["Name"])
		p.Age, _ = strconv.Atoi(fmt.Sprintf("%v", personData["Age"]))
		p.Job = fmt.Sprintf("%s", personData["Job"])
		people = append(people, p)

	}
	log.Println(personMap[0]["Name"])
	fmt.Println(people)

}
