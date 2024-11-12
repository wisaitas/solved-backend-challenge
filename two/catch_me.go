package two

func CatchMe(input string) string {
	result := make([]byte, len(input)+1)

	for i := range result {
		result[i] = '0'
	}

	for i := 0; i < len(input); i++ {
		if input[i] == 'R' {
			result[i+1] = result[i] + 1
		} else if input[i] == '=' {
			result[i+1] = result[i]
		}
	}

	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == 'L' {
			if result[i] <= result[i+1] {
				result[i] = result[i+1] + 1
			}
		} else if input[i] == '=' {
			if result[i] != result[i+1] {
				if result[i] > result[i+1] {
					result[i+1] = result[i]
				} else {
					result[i] = result[i+1]
				}
			}
		}
	}

	return string(result)
}
