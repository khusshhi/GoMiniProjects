package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	file, err := os.Open("SampleText.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	cnt := wordCount(file)
	fmt.Println("Number of count is", cnt["lorem"])
}

func wordCount(rdr io.Reader) map[string]int {
	cnt := map[string]int{}
	sc := bufio.NewScanner(rdr)
	sc.Split(bufio.ScanWords)
	for sc.Scan() {
		word := sc.Text()

		cnt[word]++
	}
	return cnt
}
