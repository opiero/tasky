package database

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type CsvDatabase struct {
	Columns
	csvPath string
	data    [][]string
}

func (c *CsvDatabase) loadDatabase() {
	filename := c.csvPath

	file, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Error creating file pointer to database: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Error loading database into memory: %v", err)
	}

	c.data = records
}

func NewCsvDatabase(csvPath string) *CsvDatabase {
	database := new(CsvDatabase)
	database.csvPath = csvPath
	database.loadDatabase()

	return database
}

func (c *CsvDatabase) PrintDatabase() error {
	for i := 0; i < len(c.data); i++ {
		for j := 0; j < len(c.data[i]); j++ {
			fmt.Printf("%v ", c.data[i][j])
		}
		fmt.Println()
	}

	return nil
}
