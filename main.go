/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"tasky/database"
)

func main() {
	filename := "./iris.csv"
	newDatabase := database.NewCsvDatabase(filename)
	newDatabase.PrintDatabase()
	// cmd.Execute()
}
