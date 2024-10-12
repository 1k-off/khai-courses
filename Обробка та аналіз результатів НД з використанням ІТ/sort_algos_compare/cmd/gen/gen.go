package main

import (
	"flag"
	"log"
	"sort_algos_compare/pkg/gen"
	"sort_algos_compare/pkg/util"
)

func main() {
	minValue := flag.Int("min-value", 0, "Minimum value for random integers")
	maxValue := flag.Int("max-value", 100, "Maximum value for random integers")
	count := flag.Int("count", 1000000, "Number of integers to generate")
	zeros := flag.Bool("zeros", false, "Fill all values with zero (cannot be used with min-value and max-value)")
	ones := flag.Bool("ones", false, "Fill all values with one (cannot be used with min-value and max-value)")
	ordered := flag.String("ordered", "", "Generate ordered array (asc or desc). Only min-value accepted with this option")
	sampleData := flag.Bool("sample-data", false, "Generate sample data")
	help := flag.Bool("help", false, "Display help message")

	flag.Parse()

	if *help || (flag.NFlag() == 0) {
		flag.Usage()
		return
	}

	if *sampleData {
		gen.SampleData()
		return
	}
	if *zeros && (flag.Lookup("min-value").Value.String() != "0" || flag.Lookup("max-value").Value.String() != "100") {
		log.Fatalf("--zeros cannot be used with --min-value and --max-value")
	}
	if *ones && (flag.Lookup("min-value").Value.String() != "0" || flag.Lookup("max-value").Value.String() != "100") {
		log.Fatalf("--ones cannot be used with --min-value and --max-value")
	}
	if *zeros && *ones {
		log.Fatalf("--zeros and --ones cannot be used together")
	}
	if *ordered != "" && *zeros {
		log.Fatalf("--ordered cannot be used with --zeros")
	}
	if *ordered != "" && *ones {
		log.Fatalf("--ordered cannot be used with --ones")
	}
	if *ordered != "" && *ordered != "asc" && *ordered != "desc" {
		log.Fatalf("--ordered must be either 'asc' or 'desc'")
	}

	var data []int

	if *zeros {
		data = gen.Zeros(*count)
		util.ToJSON(data)
		return
	}

	if *ones {
		data = gen.Ones(*count)
		util.ToJSON(data)
		return
	}

	if *ordered == "asc" || *ordered == "desc" {
		data = gen.Ordered(*count, *minValue, *ordered)
		util.ToJSON(data)
		return
	}

	if *minValue >= *maxValue {
		log.Fatalf("--min-value must be less than --max-value")
	}

	data = gen.Random(*count, *minValue, *maxValue)
	util.ToJSON(data)
}
