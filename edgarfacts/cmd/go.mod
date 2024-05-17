module github.com/spencerguo11/first_data_engineer_project/edgarfacts/cmd

go 1.22.3

replace "github.com/spencerguo11/first_data_engineer_project/edgarfacts/internal/facts" => "../facts"

require (
	github.com/spencerguo11/first_data_engineer_project/edgarfacts/internal/facts v0.0.0-00010101000000-000000000000

)