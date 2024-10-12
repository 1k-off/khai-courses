package compare

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort_algos_compare/pkg/sort/bubble"
	"sort_algos_compare/pkg/sort/gostd"
	"sort_algos_compare/pkg/sort/merge"
	"sort_algos_compare/pkg/sort/quick"
	"sort_algos_compare/pkg/sort/radixsort"
	"sort_algos_compare/pkg/sort/tim"
	"strings"
	"sync"
	"time"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/components"

	"github.com/go-echarts/go-echarts/v2/opts"
)

var algorithms = []string{"bubble", "quick", "merge", "tim", "radix", "gostd"}
var conditions = []string{
	"small amount of data with the zeros array",
	"small amount of data with the ones array",
	"small amount of data with the ordered asc small values",
	"small amount of data with the ordered asc medium values",
	"small amount of data with the ordered asc big values",
	"small amount of data with the ordered desc small values",
	"small amount of data with the ordered desc medium values",
	"small amount of data with the ordered desc big values",
	"small amount of data with the small random values",
	"small amount of data with the medium random values",
	"small amount of data with the big random values",
	"medium amount of data with the zeros array",
	"medium amount of data with the ones array",
	"medium amount of data with the ordered asc small values",
	"medium amount of data with the ordered asc medium values",
	"medium amount of data with the ordered asc big values",
	"medium amount of data with the ordered desc small values",
	"medium amount of data with the ordered desc medium values",
	"medium amount of data with the ordered desc big values",
	"medium amount of data with the small random values",
	"medium amount of data with the medium random values",
	"medium amount of data with the big random values",
	"big amount of data with the zeros array",
	"big amount of data with the ones array",
	"big amount of data with the ordered asc small values",
	"big amount of data with the ordered asc medium values",
	"big amount of data with the ordered asc big values",
	"big amount of data with the ordered desc small values",
	"big amount of data with the ordered desc medium values",
	"big amount of data with the ordered desc big values",
	"big amount of data with the small random values",
	"big amount of data with the medium random values",
	"big amount of data with the big random values",
}

type SortFunc func([]int) []int

var sortFuncs = map[string]SortFunc{
	"bubble": bubble.Sort,
	"quick":  quick.Sort,
	"merge":  merge.Sort,
	"tim":    tim.Sort,
	"radix":  radixsort.Sort,
	"gostd":  gostd.Sort,
}

func Compare(path string) {
	results := make(map[string]map[string]time.Duration)
	for _, algo := range algorithms {
		results[algo] = make(map[string]time.Duration)
		for _, condition := range conditions {
			data := generateData(condition)
			startTime := time.Now()
			_ = sortFuncs[algo](data)
			elapsed := time.Since(startTime)
			results[algo][condition] = elapsed
		}
	}

	writeMarkdownResults(path+"/comparison_results.md", results)
	generateGraphs(path+"/comparison_graph_linear.html", results, false)
	generateGraphs(path+"/comparison_graph_logarithmic.html", results, true)
}

func CompareMultiThread(path string) {
	results := make(map[string]map[string]time.Duration)
	var mu sync.Mutex
	var wg sync.WaitGroup

	for _, algo := range algorithms {
		results[algo] = make(map[string]time.Duration)
		for _, condition := range conditions {
			wg.Add(1)
			go func(algo, condition string) {
				defer wg.Done()
				data := generateData(condition)
				startTime := time.Now()
				_ = sortFuncs[algo](data)
				elapsed := time.Since(startTime)
				mu.Lock()
				results[algo][condition] = elapsed
				mu.Unlock()
			}(algo, condition)
		}
	}
	wg.Wait()

	writeMarkdownResults(path+"/comparison_results.md", results)
	generateGraphs(path+"/comparison_graph_linear.html", results, false)
	generateGraphs(path+"/comparison_graph_logarithmic.html", results, true)
}

func generateData(condition string) []int {
	var filePath string
	switch condition {
	case "small amount of data with the zeros array":
		filePath = "samples/small/zeros.json"
	case "small amount of data with the ones array":
		filePath = "samples/small/ones.json"
	case "small amount of data with the ordered asc small values":
		filePath = "samples/small/ordered_asc_small.json"
	case "small amount of data with the ordered asc medium values":
		filePath = "samples/small/ordered_asc_medium.json"
	case "small amount of data with the ordered asc big values":
		filePath = "samples/small/ordered_asc_big.json"
	case "small amount of data with the ordered desc small values":
		filePath = "samples/small/ordered_desc_small.json"
	case "small amount of data with the ordered desc medium values":
		filePath = "samples/small/ordered_desc_medium.json"
	case "small amount of data with the ordered desc big values":
		filePath = "samples/small/ordered_desc_big.json"
	case "small amount of data with the small random values":
		filePath = "samples/small/random_small.json"
	case "small amount of data with the medium random values":
		filePath = "samples/small/random_medium.json"
	case "small amount of data with the big random values":
		filePath = "samples/small/random_big.json"
	case "medium amount of data with the zeros array":
		filePath = "samples/medium/zeros.json"
	case "medium amount of data with the ones array":
		filePath = "samples/medium/ones.json"
	case "medium amount of data with the ordered asc small values":
		filePath = "samples/medium/ordered_asc_small.json"
	case "medium amount of data with the ordered asc medium values":
		filePath = "samples/medium/ordered_asc_medium.json"
	case "medium amount of data with the ordered asc big values":
		filePath = "samples/medium/ordered_asc_big.json"
	case "medium amount of data with the ordered desc small values":
		filePath = "samples/medium/ordered_desc_small.json"
	case "medium amount of data with the ordered desc medium values":
		filePath = "samples/medium/ordered_desc_medium.json"
	case "medium amount of data with the ordered desc big values":
		filePath = "samples/medium/ordered_desc_big.json"
	case "medium amount of data with the small random values":
		filePath = "samples/medium/random_small.json"
	case "medium amount of data with the medium random values":
		filePath = "samples/medium/random_medium.json"
	case "medium amount of data with the big random values":
		filePath = "samples/medium/random_big.json"
	case "big amount of data with the zeros array":
		filePath = "samples/big/zeros.json"
	case "big amount of data with the ones array":
		filePath = "samples/big/ones.json"
	case "big amount of data with the ordered asc small values":
		filePath = "samples/big/ordered_asc_small.json"
	case "big amount of data with the ordered asc medium values":
		filePath = "samples/big/ordered_asc_medium.json"
	case "big amount of data with the ordered asc big values":
		filePath = "samples/big/ordered_asc_big.json"
	case "big amount of data with the ordered desc small values":
		filePath = "samples/big/ordered_desc_small.json"
	case "big amount of data with the ordered desc medium values":
		filePath = "samples/big/ordered_desc_medium.json"
	case "big amount of data with the ordered desc big values":
		filePath = "samples/big/ordered_desc_big.json"
	case "big amount of data with the small random values":
		filePath = "samples/big/random_small.json"
	case "big amount of data with the medium random values":
		filePath = "samples/big/random_medium.json"
	case "big amount of data with the big random values":
		filePath = "samples/big/random_big.json"
	default:
		log.Fatalf("Unknown condition: %s", condition)
	}

	fileData, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	var data []int
	err = json.Unmarshal(fileData, &data)
	if err != nil {
		log.Fatalf("Invalid input in file %s: %v", filePath, err)
	}

	return data
}

func writeMarkdownResults(outputFilePath string, results map[string]map[string]time.Duration) {
	var builder strings.Builder
	builder.WriteString("# Sorting Algorithm Comparison Results\n\n")
	builder.WriteString("| Condition | " + strings.Join(algorithms, " | ") + " |\n")
	builder.WriteString("|" + strings.Repeat(" --- |", len(algorithms)+1) + "\n")

	for _, condition := range conditions {
		builder.WriteString("| " + condition + " ")
		for _, algo := range algorithms {
			builder.WriteString(fmt.Sprintf("| %v ", results[algo][condition]))
		}
		builder.WriteString("|\n")
	}

	f, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Failed to create output markdown file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(builder.String())
	if err != nil {
		log.Fatalf("Failed to write to output markdown file: %v", err)
	}

	fmt.Printf("Markdown results written to %s\n", outputFilePath)
}

func generateGraphs(outputFilePath string, results map[string]map[string]time.Duration, isLogarithmic bool) {
	page := components.NewPage()
	graphCategories := []string{
		"small amount of data",
		"medium amount of data",
		"big amount of data",
	}
	valueTypes := []string{"small", "medium", "big"}

	scaleType := "Linear"
	if isLogarithmic {
		scaleType = "Logarithmic"
	}
	page.PageTitle = fmt.Sprintf("Sorting Algorithm Comparison - %s Scale", scaleType)

	for _, category := range graphCategories {
		for _, valueType := range valueTypes {
			graph := generateSingleGraph(results, category, valueType, isLogarithmic)
			if graph != nil {
				page.AddCharts(graph)
			}
		}
	}

	f, err := os.Create(outputFilePath)
	if err != nil {
		log.Fatalf("Failed to create output graph file: %v", err)
	}
	defer f.Close()

	err = page.Render(f)
	if err != nil {
		log.Fatalf("Failed to render graph: %v", err)
	}

	fmt.Printf("%s scale graph results written to %s\n", scaleType, outputFilePath)
}

func generateSingleGraph(results map[string]map[string]time.Duration, category, valueType string, isLogarithmic bool) *charts.Line {
	line := charts.NewLine()
	scaleType := "Linear"
	if isLogarithmic {
		scaleType = "Logarithmic"
	}

	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title: fmt.Sprintf("%s - %s values - %s Scale", category, valueType, scaleType),
			Left:  "center",
			TitleStyle: &opts.TextStyle{
				FontSize: 20,
			},
			SubtitleStyle: &opts.TextStyle{
				FontSize: 14,
			},
		}),
		charts.WithLegendOpts(opts.Legend{
			Top:    "bottom",
			Orient: "horizontal",
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Type: func() string {
				if isLogarithmic {
					return "log"
				}
				return "value"
			}(),
			Name: "Time (milliseconds)",
		}),
		charts.WithXAxisOpts(opts.XAxis{
			AxisLabel: &opts.AxisLabel{
				Show:     opts.Bool(true),
				Rotate:   45,
				Interval: "0",
			},
		}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show:  opts.Bool(true),
			Right: "10%",
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  opts.Bool(true),
					Type:  "png",
					Title: "Save as PNG",
				},
				DataZoom: &opts.ToolBoxFeatureDataZoom{
					Show: opts.Bool(true),
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show: opts.Bool(true),
				},
			},
		}),
		charts.WithTooltipOpts(opts.Tooltip{
			Show:    opts.Bool(true),
			Trigger: "axis",
			AxisPointer: &opts.AxisPointer{
				Type: "cross",
			},
		}),
	)

	relevantConditions := filterConditions(conditions, category, valueType)

	// Set x-axis once
	line.SetXAxis(relevantConditions)

	// Add series for each algorithm
	for _, algo := range algorithms {
		var yValues []opts.LineData
		for _, condition := range relevantConditions {
			yValues = append(yValues, opts.LineData{
				Value: results[algo][condition].Milliseconds(),
				Name:  condition,
			})
		}
		line.AddSeries(algo, yValues).
			SetSeriesOptions(
				charts.WithLineChartOpts(opts.LineChart{Smooth: opts.Bool(true)}),
				charts.WithLabelOpts(opts.Label{Show: opts.Bool(true)}),
			)
	}

	return line
}

func filterConditions(conditions []string, category, valueType string) []string {
	var filtered []string
	for _, condition := range conditions {
		if strings.Contains(condition, category) && strings.Contains(condition, valueType) {
			filtered = append(filtered, condition)
		}
	}
	return filtered
}
