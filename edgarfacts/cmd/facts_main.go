package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	facts "github.com/spencerguo11/first_data_engineer_project/edgarfacts/internal/facts"
	storage "github.com/spencerguo11/first_data_engineer_project/edgarfacts/internal/storage"
)

func main() {
	/*
		cik := "0000886982"
		organization := "individual"
		name := "Spencer Guo"
		email := "spencerguo12@gmail.com"
		facts, err := facts.LoadFacts(cik, organization, name, email)
		if err != nil {
			panic(err)
		}

		// fmt.Println(string(companyFacts))

		// Upload to Storage
		bucketName := "gofordataengineersstorage"
		filepath := fmt.Sprintf("sec/edgar/facts/stage/%s.json", cik)
		err = storage.UploadBytes(facts, bucketName, filepath)
		if err != nil {
			panic(err)
		}

		// Log Upload
		fmt.Printf("Upload: %s\n", cik)

	*/

	// Parse Command Line Arguements
	var cik string
	var organization string
	var name string
	var email string

	flag.StringVar(&cik, "cik", "", "CIK Number")
	flag.StringVar(&organization, "organization", "", "Your Organization")
	flag.StringVar(&name, "name", "", "Your Name")
	flag.StringVar(&email, "email", "", "Your Email")

	flag.Parse()

	// Validate Command Line Arguments
	if len(cik) != 10 { // Checking the cik number
		panic("CIK must be of length 10")
	}

	if organization == "" {
		panic("Please provide the name of your organization")
	}

	if name == "" {
		panic("Please provide your name")
	}

	if email == "" {
		panic("Please provide your email address")
	}

	// Load the Environment Variables
	bucketName := os.Getenv("BUCKET")
	folderPath := os.Getenv("STAGE")
	if bucketName == "" || folderPath == "" {
		panic("Error reading ENV")
	}

	// Configure Logger
	logger := log.New(os.Stdout, "", log.LstdFlags)

	// Load Data
	logger.Printf("Loading Facts for %s\n", cik)
	facts, err := facts.LoadFacts(cik, organization, name, email)
	if err != nil {
		panic(err)
	}

	// Upload to Google Storage
	fileName := fmt.Sprintf("%s.json", cik)
	filePath := filepath.Join(folderPath, fileName)

	logger.Printf("Uploading Facts to %s on bucket %s\n", fileName, bucketName)

	err = storage.UploadBytes(facts, bucketName, filePath)
	if err != nil {
		panic(err)
	}

}
