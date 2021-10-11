package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	csvFile, err := os.Open("Data.csv")
	if err != nil {
		panic(err)
	}
	d_count := 0
	w_count := map[string]int{}
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		w_count[line[0]]++
		if w_count[line[0]] > 1 {
			d_count = w_count[line[0]]

		}
	}
	if d_count == 0 {

		fmt.Println("No duplicates")

	} else {
		fmt.Printf("No of duplicates: %d", d_count)
	}

}
