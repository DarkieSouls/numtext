package numtext

import "strings"

var (
	units = map[string]string{"1":"one ", "2":"two ", "3":"three ", "4":"four ", "5":"five ", "6":"six ", "7":"seven ", "8":"eight ", "9":"nine "}
	teens = map[string]string{"1":"eleven ", "2":"twelve ", "3":"thirteen ", "4":"fourteen ", "5":"fifteen ", "6":"sixteen ", "7":"seventeen ", "8":"eighteen ", "9":"nineteen "}
	tens = map[string]string{"1":"ten ", "2":"twenty ", "3":"thirty ", "4":"forty ", "5":"fifty ", "6":"sixty ", "7":"seventy ", "8":"eighty ", "9":"ninety "}
	orders = []string{"hundred ", "thousand ", "million ", "billion ", "trillion ", "quadrillion ", "quintillion "}
)

func splitNums(number string) [][]string {
	var nums []string
	for len(number) > 3 {
		nums = append(nums, number[len(number)-3:])
		number = number[0:len(number)-3]
	}
	nums = append(nums, number)

	var numStrings [][]string
	for _, n := range nums {
		switch len(n) {
		case 1:
			numStrings = append(numStrings, []string{units[n]})
		case 2:
			if string(n[0]) == "1" && string(n[1]) != "0" {
				numStrings = append(numStrings, []string{teens[string(n[0])]})
			} else {
				numStrings = append(numStrings, []string{tens[string(n[0])], units[string(n[1])]})
			}
		case 3:
			var s []string
			s = append(s, units[string(n[0])])
			if string(n[1]) == "1" && string(n[2]) != "0" {
				s = append(s, teens[string(n[2])])
			} else {
				s = append(s, tens[string(n[1])], units[string(n[2])])
			}
			numStrings = append(numStrings, s)
		}
	}

	return numStrings
}

func fmtEngNum(number string, format rune) string {
	var s string
	switch format {
	case 's','S':
		s = number
	case 'h','H':
		s = strings.ReplaceAll(number, " ", "-")
	case 'c','C':
		s = strings.ReplaceAll(number, "ion ", "ion, ")
		s = strings.ReplaceAll(s, "and ", "and, ")
	}

	return s
}
