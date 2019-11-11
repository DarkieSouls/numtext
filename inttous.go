package numtext

import "strings"

// IntToWordsUS converts a specified number to its US verbal representation.
func IntToWordsUS(number int64) string {
	if number == 0 {
		return "zero"
	}
	if number > 9223372036854775807 || number < -9223372036854775808 {
		panic("number value too large")
	}

	var nums [][]string
	if number < 0 {
		nums = splitNums(string(-number))
	} else {
		nums = splitNums(string(number))
	}
	val := numsToUS(nums)
	if number < 0 {
		return "minus " + val
	}
	return val
}

// UIntToWordsUS converts a specified number to its US verbal representation.
func UIntToWordsUS(number uint64) string {
	if number == 0 {
		return "zero"
	}
	if number > 18446744073709551615 {
		panic("number value too large")
	}

	nums := splitNums(string(number))
	return numsToUS(nums)
}

// IntToWordsUSFmt converts a specified number to its US verbal representation.
// s formats the number with spaces
// h formats the number with hyphens
// c formats the number with comma separation
func IntToWordsUSFmt(number int64, format rune) string {
	num := IntToWordsUS(number)
	return fmtEngNum(num, format)
}

// UIntToWordsUSFmt converts a specified number to its US verbal representation.
// s formats the number with spaces
// h formats the number with hyphens
// c formats the number with comma separation
func UIntToWordsUSFmt(number uint64, format rune) string {
	num := UIntToWordsUS(number)
	return fmtEngNum(num, format)
}

func numsToUS(nums [][]string) string {
	var result string
	for k, n := range nums {
		if k > 1 {
			result = orders[k] + result
		}
		if k == len(nums)-1 {
			switch len(n) {
			case 1:
				result = n[0] + result
			case 2:
				hundreds := false
				for _, x := range units {
					if x == n[0] {
						result = n[0] + orders[0] + n[1] + result
						hundreds = true
						break
					}
				}
				if !hundreds {
					result = n[0] + n[1] + result
				}
			case 3:
				result = n[0] + orders[0] + n[1] + n[2] + result
			}
		} else {
			if len(n) == 2 {
				result = n[0] + orders[0] + n[1] + result
			} else {
				result = n[0] + orders[0] + n[1] + n[2] + result
			}
		}
	}

	return strings.TrimSpace(result)
}