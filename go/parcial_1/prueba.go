package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	var filePath = "datos.csv"

	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	fmt.Println(records)

	var new = make([][]string, len(records[0]))
	fmt.Println(len(records))
	// var new [len(records[0])][len(records)]string
	fmt.Println(new)

	// for i := 0; i < len(records); i++ {
	// 	for j := 0; j < len(records[0]); j++ {
	// 		// fmt.Println(records[i][j])
	// 		new[j] = append(new[j], records[i][j])
	// 		// new[j] = records[i][j]
	// 	}
	// }

}
