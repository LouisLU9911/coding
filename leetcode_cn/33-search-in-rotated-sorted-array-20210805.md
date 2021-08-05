# Leetcode_cn 33 Search in Rotated Sorted Array

## Description

There is an integer array `nums` sorted in ascending order (with **distinct** values).

Prior to being passed to your function, `nums` is **rotated** at an unknown pivot index `k` (`0 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (**0-indexed**). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array `nums` **after** the rotation and an integer `target`, return *the index of* `target` *if it is in* `nums`*, or* `-1` *if it is not in* `nums`.

You must write an algorithm with `O(log n)` runtime complexity.

Reference：[33. Search in Rotated Sorted Array](https://leetcode-cn.com/problems/search-in-rotated-sorted-array/)

Example 1:

```shell
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

Example 2:

```shell
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

Example 3:

```shell
Input: nums = [1], target = 0
Output: -1
```

Constraints:

* `1 <= nums.length <= 5000`
* `-10^4 <= nums[i] <= 10^4`
* All values of `nums` are **unique**.
* `nums` is guaranteed to be rotated at some pivot.
* `-10^4 <= target <= 10^4`

## Solution

```go
package leetcode_cn

func search(nums []int, target int) int {
    l, r := 0, len(nums)-1
    for l <= r {
        mid := (l + r) >> 1
        if nums[mid] == target {
            return mid
        }
        if nums[l] == nums[r] {
            return -1
        } else if nums[l] > nums[r] {
            if nums[mid] == nums[l] {
                l = mid + 1
            } else if nums[mid] > nums[l] {
                if target > nums[mid] || target < nums[l] {
                    l = mid + 1
                } else {
                    r = mid - 1
                }
            } else {
                if target < nums[mid] || target > nums[r] {
                    r = mid - 1
                } else {
                    l = mid + 1
                }
            }
        } else {
            if nums[mid] < target {
                l = mid + 1
            } else {
                r = mid - 1
            }
        }
    }
    return -1
}
```
