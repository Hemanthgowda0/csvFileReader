package main

import (
	csv "github.com/Hemanthgowda0/csvFileReader/CSV"
	json "github.com/Hemanthgowda0/csvFileReader/JSON"
	"log"
)

func main() {
	students, err := csv.ReadCsvFile("students.csv")
	if err != nil {
		log.Fatal("error reading csv data:", err)
	}

	err = json.WriteJsonFile("output.json", students)
	if err != nil {
		log.Fatal("error writing JSON data", err)
	}

	log.Fatal("data has been written to output.json")
}
