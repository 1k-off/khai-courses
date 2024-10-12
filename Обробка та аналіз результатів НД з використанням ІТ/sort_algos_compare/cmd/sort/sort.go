package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	sortcompare "sort_algos_compare/pkg/compare"
	"sort_algos_compare/pkg/sort/bubble"
	"sort_algos_compare/pkg/sort/gostd"
	"sort_algos_compare/pkg/sort/merge"
	"sort_algos_compare/pkg/sort/quick"
	"sort_algos_compare/pkg/sort/radixsort"
	"sort_algos_compare/pkg/sort/tim"
	"sort_algos_compare/pkg/util"
	"time"
)

func main() {
	algo := flag.String("algo", "bubble", "Sorting algorithm to use. Supported values: bubble, quick, merge, tim, radix, gostd. Hint: gostd is the Go standard library sort that used pdqsort")
	measureTime := flag.Bool("measure-time", false, "Measure and display the execution time")
	inputFilePath := flag.String("input", "input.json", "Path to JSON file containing an array of integers to sort")
	compare := flag.Bool("compare", false, "Compare all sorting algorithms with the pre-defined data samples")
	// wanna kill your laptop? uncomment this and flag check
	//compareMultiThread := flag.Bool("compare-multithread", false, "Compare all sorting algorithms with the pre-defined data samples using multithreading")
	help := flag.Bool("help", false, "Display help message")

	flag.Parse()

	if *help || (flag.NFlag() == 0) {
		flag.Usage()
		return
	}

	if *compare {
		outputPath := "compare-sample"
		err := util.CreateDirIfNotExist(outputPath)
		if err != nil {
			log.Println(err)
		}
		sortcompare.Compare(outputPath)
		return
	}

	//if *compareMultiThread {
	//	outputPath := "compare-sample-multithread"
	//	err := util.CreateDirIfNotExist(outputPath)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	sortcompare.CompareMultiThread(outputPath)
	//	return
	//}

	fileData, err := os.ReadFile(*inputFilePath)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	var inputInts []int
	err = json.Unmarshal(fileData, &inputInts)
	if err != nil {
		log.Fatalf("Invalid input: %v", err)
	}

	var sorted []int
	startTime := time.Now()

	switch *algo {
	case "radix":
		sorted = radixsort.Sort(inputInts)
	case "bubble":
		sorted = bubble.Sort(inputInts)
	case "quick":
		sorted = quick.Sort(inputInts)
	case "merge":
		sorted = merge.Sort(inputInts)
	case "tim":
		sorted = tim.Sort(inputInts)
	case "gostd":
		sorted = gostd.Sort(inputInts)
	default:
		log.Fatalf("unknown algorithm: %s", *algo)
	}

	outputFilePath := fmt.Sprintf("sorted_%s_%s.json", *algo, time.Now().Format("20060102-1504"))

	if *measureTime {
		elapsed := time.Since(startTime)
		log.Printf("Execution time: %s", elapsed)
	}

	sortedData, err := json.Marshal(sorted)
	if err != nil {
		log.Fatalf("Failed to marshal sorted data: %v", err)
	}

	err = os.WriteFile(outputFilePath, sortedData, 0644)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("Sorted array written to %s\n", outputFilePath)
}
