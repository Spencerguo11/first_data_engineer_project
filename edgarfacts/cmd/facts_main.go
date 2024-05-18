package main

import (
	"fmt"

	facts "github.com/spencerguo11/first_data_engineer_project/edgarfacts/internal/facts"
)

func main() {
	cik := "0000886982"
	organization := "individual"
	name := "Spencer Guo"
	email := "spencerguo12@gmail.com"
	companyFacts, err := facts.LoadFacts(cik, organization, name, email)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(companyFacts))
}
