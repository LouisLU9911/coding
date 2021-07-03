package leetcode_cn

import "sort"

var byteNum = make(map[byte]int)

type byteList []byte

func (l byteList) Len() int {
	return len(l)
}

func (l byteList) Less(i, j int) bool {
	return byteNum[l[i]] > byteNum[l[j]]
}

func (l byteList) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func frequencySort(s string) string {
	mylist := byteList{}
	for _, v := range []byte(s) {
		if _, ok := byteNum[v]; ok {
			byteNum[v] += 1
		} else {
			byteNum[v] = 1
			mylist = append(mylist, v)
		}
	}
	sort.Sort(mylist)
	ret := []byte{}
	for i := 0; i < mylist.Len(); i++ {
		for j := 0; j < byteNum[mylist[i]]; j++ {
			ret = append(ret, mylist[i])
		}
	}
	return string(ret)
}
