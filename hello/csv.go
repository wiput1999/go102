package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func ReadCSV() {
	f, _ := os.Open("oscar_age_male.csv")
	r := csv.NewReader(f)

	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	result := make(map[string]int)

	for i, record := range records {
		if i == 0 {
			continue
		}

		result[record[3]]++
	}

	for key, value := range result {
		if value > 1 {
			fmt.Println(key)
		}
	}
}
