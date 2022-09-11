package main

import (
	"fmt"
	"sort"
)

func main() {

	// capacity := 25
	// profit := []int{100, 180, 140, 150}
	// weight := []int{5, 12, 10, 8}

	capacity := 34
	profit := []int{80, 50, 100, 70, 40}
	weight := []int{10, 11, 12, 4, 3}

	maxProfit := knapSack(capacity, weight, profit)
	fmt.Println(maxProfit)

}

type knapSackStruct struct {
	weight     int
	profit     int
	unitProfit float64
}

func knapSack(capacity int, weight []int, profit []int) float64 {

	p := 0.0

	containers := make([]knapSackStruct, len(weight))

	for i := 0; i < len(weight); i++ {

		containers[i] = knapSackStruct{weight[i], profit[i], float64(profit[i]) / float64(weight[i])}
	}

	fmt.Println("====Before Sorting")

	fmt.Println(weight)
	fmt.Println(profit)

	sort.Slice(containers, func(i, j int) bool {
		return containers[i].unitProfit > containers[j].unitProfit
	})

	fmt.Println("====After Sorting")

	for i := 0; i < len(containers) && capacity != 0; i++ {
		if capacity >= containers[i].weight {
			capacity -= containers[i].weight
			p += float64(containers[i].profit)
		} else {
			p += float64(containers[i].profit * capacity / containers[i].weight)
			capacity = 0
		}
	}

	return p

}
