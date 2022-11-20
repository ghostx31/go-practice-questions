package minwaitingtime

/*
	The CPU scheduling idea where we need to find the minimum waiting time for processes.
	We can arrange the arrange in ascending order and then add the waiting times to find the minimum waiting time for some process.

	We don't need to keep track of the time that the next process needs to wait.
	[3, 5, 6, 9] => 0 + 3 + (5+3) + (3 + 5 + 6) which is equal to 0 + 3*3 + 5*2 + 6*1. Thus the multipliers are equal to the number of queries we're left with.
	We can use this idea to calculate the total waiting time.

	Thus the formula to calculate becomes:
	=> duration of process[0] * (n-1) + (n-2) ... process[n-2]*1 + process[n-1]*0
*/

import "sort"

// Time: I(n - logn); n is number of processes
// Space: O(1)
func minWaitTime(process []int) int {
	sort.Ints(process)

	totalWaitTime := 0
	numOfProcess := len(process)

	for i, time := range process {
		processLeft := numOfProcess - (i + 1)
		totalWaitTime += time * processLeft
	}

	return totalWaitTime
}
