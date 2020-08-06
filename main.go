package main

import (
	"encoding/csv"
	"github.com/gocarina/gocsv"
	"os"
)

// 一些测试

func main () {
	file, err := os.Open("./readFile/data/c.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var clients interface{}

	// gocsv.UnmarshalFile(file, clients)

	csvReader := csv.NewReader(file)
	csvReader.Comma = '\t'
	gocsv.UnmarshalCSVWithoutHeaders(csvReader, clients)
	// gocsv.UnmarshalWithoutHeaders(csvReader, clients)



}
