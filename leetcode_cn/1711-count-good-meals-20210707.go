package leetcode_cn

func countPairs(deliciousness []int) int {
	foodNum := make(map[int]int)
	for _, v := range deliciousness {
		foodNum[v] += 1
	}
	foodName := []int{}
	for k := range foodNum {
		foodName = append(foodName, k)
	}
	ret := 0
	for i := 0; i < len(foodName); i++ {
		for j := i; j < len(foodName); j++ {
			if i == j {
				if fn := foodNum[foodName[i]]; fn > 1 && isPowerOfTwo(foodName[i]+foodName[j]) {
					ret = (ret + fn*(fn-1)/2) % 1000000007
				}
			} else if i != j && isPowerOfTwo(foodName[i]+foodName[j]) {
				ret = (ret + foodNum[foodName[i]]*foodNum[foodName[j]]) % 1000000007
			}
		}
	}
	return ret
}

func isPowerOfTwo(number int) bool {
	if number == 0 {
		return false
	}
	return (number & (number - 1)) == 0
}

func countPairs_v2(deliciousness []int) (ans int) {
	maxVal := deliciousness[0]
	for _, val := range deliciousness[1:] {
		if val > maxVal {
			maxVal = val
		}
	}
	maxSum := maxVal * 2
	cnt := map[int]int{}
	for _, val := range deliciousness {
		for sum := 1; sum <= maxSum; sum <<= 1 {
			ans += cnt[sum-val]
		}
		cnt[val]++
	}
	return ans % (1e9 + 7)
}
