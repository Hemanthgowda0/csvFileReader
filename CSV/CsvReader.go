package csv

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	model "github.com/Hemanthgowda0/csvFileReader/Model"
)

func ReadCsvFile(filePath string) ([]model.Student, error) {
	csvFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)
	csvData, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var students []model.Student
	for i, row := range csvData {
		if i == 0 {
			continue
		}

		id, _ := strconv.Atoi(row[0])
		age, _ := strconv.Atoi(row[3])
		marks, _ := strconv.Atoi(row[5])
		joining_date, _ := time.Parse("2023-01-02", row[6])

		student := model.Student{
			ID:          id,
			Firstname:   row[1],
			LastName:    row[2],
			Age:         age,
			Email:       row[4],
			Marks:       marks,
			JoiningDate: joining_date,
			Result:      row[7],
		}
		students = append(students, student)
	}

	return students, nil

}
