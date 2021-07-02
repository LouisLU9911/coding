package leetcode_cn

func maxIceCream(costs []int, coins int) int {
    sort.Ints(costs)
    ret := 0
    for i := 0; i < len(costs); i++ {
        if coins >= costs[i] {
            ret += 1
            coins -= costs[i]
        } else {
            break
        }
    }
    return ret
}
