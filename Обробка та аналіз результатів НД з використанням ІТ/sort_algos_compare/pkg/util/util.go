package util

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateDirIfNotExist(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return os.Mkdir(path, os.ModePerm)
	}
	return nil
}

func ToJSON(data []int) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	// Generate output file name
	timestamp := time.Now().Format("20060102-1504")
	outputFileName := fmt.Sprintf("generated_%s.json", timestamp)

	// Write JSON data to file
	err = os.WriteFile(outputFileName, jsonData, 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("Generated JSON written to %s\n", outputFileName)
}

func ToNamedJson(data []int, name string) {
	jsonData, err := json.Marshal(data)
	if err != nil {
		log.Fatalf("Failed to marshal data: %v", err)
	}

	// Generate output file name
	outputFileName := fmt.Sprintf("%s.json", name)

	// Write JSON data to file
	err = os.WriteFile(outputFileName, jsonData, 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("Generated JSON written to %s\n", outputFileName)
}
