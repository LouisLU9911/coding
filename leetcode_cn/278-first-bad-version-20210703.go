package leetcode_cn

// a fake function
func isBadVersion(version int) bool {
	return true
}

/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {
	l, r := 1, n
	var mid, res int
	for l <= r && l >= 1 && r <= n {
		mid = (l + r) >> 1
		if isBadVersion(mid) {
			res = mid
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return res
}
