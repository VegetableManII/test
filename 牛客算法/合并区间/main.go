package main

import (
	"fmt"
	"sort"
)

type Interval struct {
	Start int
	End   int
}
type IntervalSlice []*Interval

func main() {
	res := merge([]*Interval{
		{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6},
	})
	fmt.Printf("%v\n", res)
}

/**
 *
 * @param intervals Interval类一维数组
 * @return Interval类一维数组
 */
func merge(intervals []*Interval) []*Interval {
	// write code here
	if len(intervals) == 0 {
		return nil
	}
	sort.Sort(IntervalSlice(intervals))
	res := []*Interval{intervals[0]}
	cur := intervals[0].End
	for _, v := range intervals {
		if v.Start <= cur && v.End >= cur {
			res[len(res)-1].End = v.End
			cur = v.End
		} else if v.End < cur {
			continue
		} else {
			res = append(res, v)
			cur = v.End
		}
	}
	return res
}
func (a IntervalSlice) Len() int           { return len(a) }
func (a IntervalSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a IntervalSlice) Less(i, j int) bool { return a[i].Start < a[j].Start }
