package numtext

import (
	"strconv"
	"strings"
)

// IntToWordsGB converts a specified number to its GB verbal representation.
func IntToWordsGB(number int64) string {
	if number == 0 {
		return "zero"
	}
	if number > 9223372036854775807 || number < -9223372036854775808 {
		panic("number value too large")
	}

	var nums [][]string
	if number < 0 {
		nums = splitNums(strconv.FormatInt(-number, 10))
	} else {
		nums = splitNums(strconv.FormatInt(number, 10))
	}
	val := numsToGB(nums)
	if number < 0 {
		return "minus " + val
	}
	return val
}

// UIntToWordsGB converts a specified number to its GB verbal representation.
func UIntToWordsGB(number uint64) string {
	if number == 0 {
		return "zero"
	}
	if number > 18446744073709551615 {
		panic("number value too large")
	}

	nums := splitNums(strconv.FormatUint(number, 10))
	return numsToGB(nums)
}

// IntToWordsGBFmt converts a specified number to its GB verbal representation.
//
// s formats the number with spaces,
// h formats the number with hyphens,
// c formats the number with comma separation.
func IntToWordsGBFmt(number int64, format rune) string {
	num := IntToWordsGB(number)
	return fmtEngNum(num, format)
}

// UIntToWordsGBFmt converts a specified number to its GB verbal representation.
//
// s formats the number with spaces,
// h formats the number with hyphens,
// c formats the number with comma separation.
func UIntToWordsGBFmt(number uint64, format rune) string {
	num := UIntToWordsGB(number)
	return fmtEngNum(num, format)
}

func numsToGB(nums [][]string) string {
	var result string
	for k, n := range nums {
		if n[0] == "" && n[1] == "" && n[2] == "" && k != 4 && !(k == 2 && len(nums) == 4) {
			continue
		}
		switch k {
		case 1:
			result = orders[1] + result
		case 2:
			result = orders[2] + result
		case 3:
			result = orders[1] + result
		case 4:
			result = orders[3] + result
		case 5:
			result = orders[1] + result
		case 6:
			result = orders[2] + result
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
				if n[0] == "" {
					result = n[1] + result
				} else {
					result = n[0] + orders[0] + n[1] + result
				}
			} else {
				if n[0] == "" {
					result = n[1] + n[2] + result
				} else {
					result = n[0] + orders[0] + n[1] + n[2] + result
				}
			}
		}
	}

	return strings.TrimSpace(result)
}
