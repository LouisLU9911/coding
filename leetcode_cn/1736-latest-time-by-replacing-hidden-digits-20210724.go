package leetcode_cn

func maximumTime(time string) string {
	timeByte := []byte(time)
	if timeByte[0] == '?' {
		if timeByte[1] == '?' {
			timeByte[0] = '2'
		} else if timeByte[1] >= '4' {
			timeByte[0] = '1'
		} else {
			timeByte[0] = '2'
		}
	}
	if timeByte[1] == '?' {
		if timeByte[0] <= '1' {
			timeByte[1] = '9'
		} else {
			timeByte[1] = '3'
		}
	}
	if timeByte[3] == '?' {
		timeByte[3] = '5'
	}
	if timeByte[4] == '?' {
		timeByte[4] = '9'
	}
	return string(timeByte)
}
