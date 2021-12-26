# LeetCode 56

1. Sort intervals
2. Merge one interval with all overlapping intervals after it

## Solution

```go
package main

import (
    "fmt"
    "sort"
)

func merge(intervals [][]int) [][]int {
    sort.Slice(intervals, func(i, j int) bool {
        if intervals[i][0] < intervals[j][0] {
            return true
        } else {
            return false
        }
    })
    var i, j int
    res := [][]int{}
    for i = 0; i < len(intervals); {
        left, right := intervals[i][0], intervals[i][1]
        j = i + 1
        for j < len(intervals) {
            if right < intervals[j][0] {
                break
            } else if right <= intervals[j][1] {
                right = intervals[j][1]
                j += 1
            } else {
                j += 1
            }
        }
        item := []int{left, right}
        res = append(res, item)
        i = j
    }
    return res
}

func main() {
    intervals := [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}
    fmt.Println(intervals)
    res := merge(intervals)
    fmt.Println(res)
}
```
