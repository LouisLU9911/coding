# Leetcode_cn 34 find-first-and-last-position-of-element-in-sorted-array

## Description

Given an array of integers `nums` sorted in ascending order, find the starting and ending position of a given `target` value.

If `target` is not found in the array, return `[-1, -1]`.

You must write an algorithm with `O(log n)` runtime complexity.

Reference：[34. Find First and Last Position of Element in Sorted Array](https://leetcode-cn.com/problems/find-first-and-last-position-of-element-in-sorted-array)

Example 1:

```shell
Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
```

Example 2:

```shell
Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]
```

Example 3:

```shell
Input: nums = [], target = 0
Output: [-1,-1]
```

Constraints:

* `0 <= nums.length <= 10^5`
* `-10^9 <= nums[i] <= 10^9`
* `nums` is a non-decreasing array.
* `-10^9 <= target <= 10^9`

## Solution

```go
package leetcode_cn

func searchRange(nums []int, target int) []int {
    if len(nums) == 0 {
        return []int{-1, -1}
    }
    left := search(nums, target, true)
    right := search(nums, target, false)
    return []int{left, right}
}

func search(nums []int, target int, isLeft bool) int {
    // `target` in [l, r]
    l, r := 0, len(nums) - 1
    for l <= r {
        mid := (l + r) >> 1
        if nums[mid] == target {
            if isLeft && mid > 0 && nums[mid-1] == target {
                r = mid - 1
                continue
            } else if !isLeft && mid < len(nums)-1 && nums[mid+1] == target {
                l = mid + 1
                continue
            }
            return mid
        } else if nums[mid] > target {
            r = mid - 1
        } else {
            l = mid + 1
        }
    }
    return -1
}
```
