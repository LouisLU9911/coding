package leetcode_cn

func checkInclusion(s1, s2 string) bool {
	s1Sum := make(map[byte]int)
	for i := 0; i < len(s1); i++ {
		s1Sum[s1[i]]++
	}
	s2Sum := newS2Sum()
	// [l, r)
	l, r, total := 0, 0, 0
	for r < len(s2) {
		if v1, ok := s1Sum[s2[r]]; ok {
			s2Sum[s2[r]]++
			if v2, _ := s2Sum[s2[r]]; v2 > v1 {
				for s2Sum[s2[r]] > v1 {
					s2Sum[s2[l]]--
					if s2Sum[s2[l]] == s1Sum[s2[l]]-1 {
						total--
					}
					l++
				}
			} else if v2 == v1 {
				total++
				if total == len(s1Sum) {
					return true
				}
			}
			r++
		} else {
			r++
			l = r
			total = 0
			s2Sum = newS2Sum()
		}
	}
	return false
}

func newS2Sum() map[byte]int {
	return make(map[byte]int)
}
