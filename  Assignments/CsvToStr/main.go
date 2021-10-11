package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
)

func main() {
	csvFile, err := os.Open("Strings.csv")
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(csvFile)
	for {
		line, err := reader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
		var val []string
		val = append(val, line[0])
		val = append(val, line[1])

		fmt.Println(val)

	}

}
