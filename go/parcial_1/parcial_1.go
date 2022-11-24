package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func min_arr(arr []string) int {
	var min int
	for i, tmp := range arr {
		tmp2, _ := strconv.Atoi(tmp)
		if i == 0 || tmp2 < min {
			min = tmp2
		}
	}
	return min
}

func max_arr(arr []string) int {
	var max = 0
	for i, tmp := range arr {
		tmp2, _ := strconv.Atoi(tmp)
		if i == 0 || tmp2 > max {
			max = tmp2
		}
	}
	return max
}

func avg_arr(arr []string) int {
	var sum = 0
	for _, tmp := range arr {
		tmp2, _ := strconv.Atoi(tmp)
		sum = sum + tmp2
	}
	return sum / len(arr)
}

func varianza_arr(arr []string, avg int) float64 {
	var sum float64
	var n = len(arr)

	for _, item := range arr {
		item2, _ := strconv.Atoi(item)
		tmp := item2 - avg
		tmp2 := math.Pow(float64(tmp), 2)
		sum = sum + tmp2
	}
	return (float64(1) / float64(n)) * sum
}

func desviacion_arr(varianza float64) float64 {
	return math.Sqrt(varianza)
}

func readCsvFile(filePath string) [][]string {
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

	return records
}

func moda_str(arr []string) map[string]int {
	freq := make(map[string]int)
	for _, num := range arr {
		freq[num] = freq[num] + 1
	}
	// fmt.Println(freq)
	return freq
}

// len(arr) = 4 // numero de filas
// len(arr[0]) = 4 // numero de columnas
// fil = i
// col = j
// [[1,1 1,2 1,3]
// [2,1 2,2 2,3]
// [3,1 3,2 3,3]
// [4,1 4,2 4,3]]

// len(arr) = 4
// [[1,1 2,1 3,1 4,1]
// [1,2 2,2 3,2 4,2]
// [1,3 2,3 3,3 4,3]]

func me_transpose(records [][]string) [][]string {
	var new = make([][]string, len(records[0]))
	fmt.Println(len(records))
	// var new [len(records[0])] [len(records)] string

	for i := 0; i < len(records); i++ {
		for j := 0; j < len(records[0]); j++ {
			// fmt.Println(records[i][j])
			new[j] = append(new[j], records[i][j])
			// new[j] = records[i][j]
		}
	}
	return new
}

func main() {
	records := readCsvFile("datos.csv") // se obtiene la matriz

	fmt.Println(records)

	// transpuesta mia
	var new = me_transpose(records)
	fmt.Println(new)

	for i := 0; i < len(new); i++ {
		// fmt.Println(new[i][1:len(new[0])])
		tmp := new[i][1]
		if _, err := strconv.Atoi(tmp); err == nil {

			fmt.Printf("\n Columnas numericas enteras \n")
			// numero maximo
			maximo := max_arr(new[i][1:len(new[0])])
			fmt.Printf("%s \t maximo: %d \n", new[i][0], maximo)

			// numero minimo
			minimo := min_arr(new[i][1:len(new[0])])
			fmt.Printf("%s \t minimo: %d \n", new[i][0], minimo)

			avg := avg_arr(new[i][1:len(new[0])])
			fmt.Printf("%s \t avg: %d \n", new[i][0], avg)

			varianza := varianza_arr(new[i][1:len(new[0])], avg)
			fmt.Printf("%s \t varianza: %f \n", new[i][0], varianza)

			desviacion := desviacion_arr(varianza)
			fmt.Printf("%s \t desviacion: %f \n\n", new[i][0], desviacion)
		} else {
			// calcular media
			fmt.Printf("\n Columnas string \n")
			moda := moda_str(new[i][1:len(new[0])])
			fmt.Println(moda)
		}
	}
}
