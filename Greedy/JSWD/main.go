package main

import (
	"fmt"
	"sort"
)

// Job sequencing with deadline

//Sequence all the so that max profit will be produced.

func main() {

	// jobs with index 0 profit and index 1 Deadline.
	// jobsWithProfitandDeadline := [][]int{
	// 	{150, 2},
	// 	{70, 1},
	// 	{200, 2},
	// 	{80, 1},
	// }

	jobsWithProfitandDeadline := [][]int{{80, 6}, {20, 3}, {90, 4}, {50, 4}, {30, 3}, {60, 1}, {85, 2}}

	profit := sequenceTheJob(jobsWithProfitandDeadline)

	fmt.Println(profit)

}

// sequenceTheJob funtion with sequence the job and provide max profit out of given input.
func sequenceTheJob(inputJobs [][]int) int {

	jobs := inputJobs
	profit := 0
	//sort the provided input in decreasing order of profit.
	sort.Slice(jobs, func(i, j int) bool {
		return jobs[i][0] > jobs[j][0]
	})
	// get the max deadline present in given jobs
	max := 0
	for _, job := range jobs {
		if job[1] > max {
			max = job[1]
		}
	}
	for i := max; i > 0; i-- {
		fmt.Println("for deadline  : ", i)
		for _, job := range jobs {
			fmt.Println("Entered here")
			if job[1] >= i {
				profit += job[0]
				job[1] = -1
				break
			}
		}
	}
	return profit

}
