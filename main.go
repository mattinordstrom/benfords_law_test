package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

/*
1	30.1%
2	17.6%
3	12.5%
4	9.7%
5	7.9%
6	6.7%
7	5.8%
8	5.1%
9	4.6%
*/
var percentage = []float64{30.1, 17.6, 12.5, 9.7, 7.9, 6.7, 5.8, 5.1, 4.6}

var quantityColumnIndex = 0
var unitPriceColumnIndex = 1

func main() {
	fmt.Println(":::: Benfords law tester ::::")
	fmt.Println(percentage)
	fmt.Println(" ")

	file, err := os.Open("Online_Retail_mod.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Number of rows:", len(rows))

	numberOfBadValues := 0
	numbers := map[string]float64{
		"1": 0, "2": 0, "3": 0, "4": 0, "5": 0, "6": 0, "7": 0, "8": 0, "9": 0,
	}

	for _, row := range rows {
		quantity, _ := strconv.ParseFloat(row[quantityColumnIndex], 64)
		unitPrice, _ := strconv.ParseFloat(row[unitPriceColumnIndex], 64)

		rowAmount := quantity * unitPrice

		firstDigit := fmt.Sprintf("%f", rowAmount)[:1]
		firstDigitAsInt, _ := strconv.ParseInt(firstDigit, 10, 0)
		if firstDigit == "" || firstDigit == "-" || firstDigitAsInt < 1 {
			numberOfBadValues++
			continue
		}

		numbers[firstDigit] = numbers[firstDigit] + 1
	}

	fmt.Println("Number of bad values:", numberOfBadValues)
	fmt.Println(" ")
	fmt.Println(numbers)

	percentResult := map[string]string{
		"1": fmt.Sprintf("%.2f", (numbers["1"]/float64(len(rows))*100)) + "%",
		"2": fmt.Sprintf("%.2f", (numbers["2"]/float64(len(rows))*100)) + "%",
		"3": fmt.Sprintf("%.2f", (numbers["3"]/float64(len(rows))*100)) + "%",
		"4": fmt.Sprintf("%.2f", (numbers["4"]/float64(len(rows))*100)) + "%",
		"5": fmt.Sprintf("%.2f", (numbers["5"]/float64(len(rows))*100)) + "%",
		"6": fmt.Sprintf("%.2f", (numbers["6"]/float64(len(rows))*100)) + "%",
		"7": fmt.Sprintf("%.2f", (numbers["7"]/float64(len(rows))*100)) + "%",
		"8": fmt.Sprintf("%.2f", (numbers["8"]/float64(len(rows))*100)) + "%",
		"9": fmt.Sprintf("%.2f", (numbers["9"]/float64(len(rows))*100)) + "%",
	}

	fmt.Println(" ")
	fmt.Println(percentResult)

}
