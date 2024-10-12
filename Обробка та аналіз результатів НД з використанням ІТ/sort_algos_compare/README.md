# Sample data

## Pre-generated data
The `samples` directory contains a set of different pre-generated JSON files with arrays of integers that divided by small (100 elements), 
medium (10k elements) and big (1M elements).  
Samples contains:
- zeros
- ones
- ordered ascending and descending arrays (min value 10, 1000000 , 1000000000)
- random arrays

## Generator
The `gen` app is a command-line tool for generating JSON files containing arrays of integers. It supports various options, 
including filling arrays with random numbers, zeros, ones, or generating sorted arrays in ascending or descending order. 
This app can be useful for sample datasets creation.

The generated array is written to a JSON file named in the format `generated_<timestamp>.json`, 
where <timestamp> represents the date and time when the file was created (e.g., `generated_20241012-0247.json`).

Usage: `gen [options]`
```
  -count int
        Number of integers to generate (default 1000000)
  -help
        Display help message
  -max-value int
        Maximum value for random integers (default 100)
  -min-value int
        Minimum value for random integers (default 0)
  -ones
        Fill all values with one (cannot be used with min-value and max-value)
  -ordered string
        Generate ordered array (asc or desc). Only min-value accepted with this option
  -zeros
        Fill all values with zero (cannot be used with min-value and max-value)

```

### Examples

- Generate an array of 20 random integers between 0 and 100:
```
./gen --count 20 --min-value 0 --max-value 100
```
- Generate an array of 50 integers, all set to zero:
```
./gen --count 50 --zeros
```
- Generate an array of 100 integers, all set to one:
```
./gen --count 100 --ones
```
- Generate an array of 30 integers in ascending order starting from 5:
```
./gen --count 30 --min-value 5 --ordered asc
```
- Generate an array of 15 integers in descending order starting from 10:
```
./gen --count 15 --min-value 10 --ordered desc
```

# Sorting algorithms comparison
This tool is a command-line utility that allows you to sort an array of integers using various sorting algorithms. 
It is designed to help compare the performance of different sorting algorithms, including Bubble Sort, Quick Sort, Merge Sort, Timsort, Radix Sort, 
and the Go standard library's sorting function (pdqsort). 
The utility reads the input array from a JSON file and writes the sorted output to a new JSON file.

## Usage
`sort [options]`
```
  -algo string
        Sorting algorithm to use. Supported values: bubble, quick, merge, tim, radix, gostd. Hint: gostd is the Go standard library sort that used pdqsort (default "bubble")
  -compare
        Compare all sorting algorithms with the pre-defined data samples
  -help
        Display help message
  -input string
        Path to JSON file containing an array of integers to sort (default "input.json")
  -measure-time
        Measure and display the execution time
```

### Examples
-  Sort an array using bubble sort and measure time
```
go run main.go -algo bubble -input input.json -measure-time
```
- Compare All Sorting Algorithms
```
go run main.go -compare
```


## Supported Sorting Algorithms

- Bubble Sort: A simple comparison-based sorting algorithm that repeatedly steps through the list, compares adjacent elements, 
and swaps them if they are in the wrong order. This process repeats until the list is sorted.
- Quick Sort: A divide-and-conquer algorithm that selects a pivot element and partitions the array into two sub-arrays, then recursively sorts them.
- Merge Sort: A stable sorting algorithm that recursively splits the array in half and merges the sorted halves.
- Timsort: A hybrid sorting algorithm derived from Merge Sort and Insertion Sort, designed for efficient performance on real-world data.
- Radix Sort: A non-comparative sorting algorithm that sorts numbers by processing individual digits.
- Go Standard Library Sort (pdqsort): The default sorting implementation in Go, optimized for average-case performance.

# Build

To build the `gen` and `sort` apps, run the following commands:

```
go build -o ./sort.exe ./cmd/sort
go build -o ./gen.exe ./cmd/gen
```

# Sources
- [Radix Sort](https://reintech.io/blog/mastering-radix-sort-algorithm-in-go)
- [Bubble Sort](https://blog.boot.dev/golang/bubble-sort-golang/)
- [Quick Sort](https://blog.boot.dev/golang/quick-sort-golang/)
- [Merge Sort](https://blog.boot.dev/golang/merge-sort-golang/)
- [Tim Sort](https://corte.si/posts/code/timsort/index.html)
- [pdqsort](https://itnext.io/gos-new-sorting-algorithm-pdqsort-822053d7801b)





