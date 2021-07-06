package leetcode_cn

import (
	"fmt"
	"sort"
	"strconv"
)

func displayTable(orders [][]string) [][]string {
	tables := make(map[string]bool)
	orderMap := make(map[string]map[string]int)
	for i := 0; i < len(orders); i++ {
		if _, ok := tables[orders[i][1]]; !ok {
			tables[orders[i][1]] = true
		}
		if _, ok := orderMap[orders[i][2]]; ok {
			orderMap[orders[i][2]][orders[i][1]] += 1
		} else {
			orderMap[orders[i][2]] = make(map[string]int)
			orderMap[orders[i][2]][orders[i][1]] = 1
		}
	}

	ret := make([][]string, 0, len(tables)+1)
	for i := 0; i < len(tables)+1; i++ {
		ret = append(ret, make([]string, 0, len(orderMap)+1))
	}
	tablestr := make([]int, 0, len(tables))
	for k := range tables {
		kInt, _ := strconv.Atoi(k)
		tablestr = append(tablestr, kInt)
	}
	sort.Ints(tablestr)

	ret[0] = append(ret[0], "Table")
	for i := 0; i < len(tablestr); i++ {
		ret[i+1] = append(ret[i+1], fmt.Sprintf("%d", tablestr[i]))
	}

	orderstr := make([]string, 0, len(orderMap))
	for k := range orderMap {

		orderstr = append(orderstr, k)
	}
	sort.Strings(orderstr)

	for i := 0; i < len(orderMap); i++ {
		ret[0] = append(ret[0], orderstr[i])
	}

	for i := 1; i < len(ret); i++ {
		for j := 1; j < len(ret[0]); j++ {
			if v, ok := orderMap[ret[0][j]][ret[i][0]]; ok {
				ret[i] = append(ret[i], fmt.Sprintf("%d", v))
			} else {
				ret[i] = append(ret[i], "0")
			}
		}
	}
	return ret
}
