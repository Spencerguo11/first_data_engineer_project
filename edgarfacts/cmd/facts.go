// Declare Package
package main

// Make Imports
import (
	"fmt"

	facts "github.com/spencerguo11/first_data_engineer_project/edgarfacts/internal/facts"
)

// Main Function
func main() {
	// Load Company Facts for Goldman Sachs
	cik := "0000886982"
	organization := "individual"
	name := "Spencer Guo"
	email := "spencerguo12@gmail.com"
	facts, err := facts.LoadFacts(cik, organization, name, email)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(facts))
}
