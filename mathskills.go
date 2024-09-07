package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	Data := Read()
	Sorted := QuickSort(Data)
	resfile, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Error creating file : ", err)
	}
	defer resfile.Close()
	var s string
	for _, v := range Sorted {
		s += strconv.FormatFloat(v, 'f', 0, 64) + " "
	}
	_, err = io.WriteString(resfile, s)
	if err != nil {
		log.Fatal("Error writing to file : ", err)
	}

	average := Average(Sorted)
	median := Mediane(Sorted)
	variance, deviation := Variance(Sorted)
	fmt.Println("File processed successfully :) ")
	fmt.Println("Average is :", Round(average))
	fmt.Println("Median is :", Round(median))
	fmt.Println("Variance is :", Round(variance))
	fmt.Println("Deviation is :", Round(deviation))
}

func Round(x float64) int {
	var rounded int
	if x >= 0 {
		rounded = int(x + 0.5)
	} else {
		rounded = int(x - 0.5)
	}
	return rounded
}

func Read() []float64 {
	var Population []float64

	content, err := os.ReadFile("popdata.txt")
	if err != nil {
		log.Fatal("couldn't read file")
	}

	Split := strings.Split(strings.TrimSpace(string(content)), "\n")

	for _, v := range Split {
		s, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Print("Parse failed", string(v), "in this value")
		} else {
			Population = append(Population, s)
		}
	}
	return Population
}

func QuickSort(Population []float64) []float64 {
	var before, after, pivotlist []float64

	if len(Population) < 2 {
		return Population
	}
	first := Population[0]
	middle := Population[(len(Population)-1)/2]
	last := Population[len(Population)-1]
	var Pivot float64

	if (first > middle) && (first < last) {
		Pivot = first
	} else if (middle > first) && (middle < last) {
		Pivot = middle
	} else {
		Pivot = last
	}

	for _, value := range Population {
		if value < Pivot {
			before = append(before, value)
		} else if value > Pivot {
			after = append(after, value)
		} else {
			pivotlist = append(pivotlist, value)
		}
	}
	var Sorted []float64
	Sorted = append(Sorted, QuickSort(before)...)
	Sorted = append(Sorted, pivotlist...)
	Sorted = append(Sorted, QuickSort(after)...)

	return Sorted
}

func Average(Sorted []float64) float64 {
	var sum float64
	var average float64
	for _, value := range Sorted {
		sum += value
	}
	average = sum / float64(len(Sorted))
	return average
}

func Mediane(Sorted []float64) float64 {
	var mediane float64

	if len(Sorted)%2 == 0 {
		mediane = (Sorted[(len(Sorted)/2-1)] + Sorted[(len(Sorted)/2)]) / 2
	} else {
		mediane = (Sorted[len(Sorted)/2])
	}
	return mediane
}

func Variance(Sorted []float64) (float64, float64) {
	var res float64
	var variance float64
	var deviation float64
	average := Average(Sorted)
	for _, value := range Sorted {
		res = (value - average) * (value - average)
		variance += res
	}
	variance = variance / float64(len(Sorted))
	deviation = math.Sqrt(variance)

	return variance, deviation
}
