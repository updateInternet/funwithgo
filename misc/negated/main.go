package main

import (
	"fmt"
	"math"
)

type interval struct {
	// Inclusive
	begin, end int32
}

func negateIntervals(intervals []interval) []interval {
	if len(intervals) == 0 {
		return []interval{
			interval{
				begin: math.MinInt32,
				end:   math.MaxInt32,
			},
		}
	}
	var negated []interval
	first := intervals[0]
	if first.begin != math.MinInt32 {
		// Special case where the first output is "before" the input.
		negated = append(negated, interval{
			begin: math.MinInt32,
			end:   first.begin - 1,
		})
	}
	// In the loop we always add an interval "after" the current input.
	for i, current := range intervals {
		var begin, end int32
		if current.end == math.MaxInt32 {
			break
		}
		begin = current.end + 1
		if i == len(intervals)-1 {
			end = math.MaxInt32
		} else {
			end = intervals[i+1].begin - 1
		}
		negated = append(negated, interval{begin: begin, end: end})
	}
	return negated
}

func main() {
	input := []interval{{10, 20}, {30, 200}, {1000, math.MaxInt32}}
	fmt.Printf("%v => %v\n", input, negateIntervals(input))
}
