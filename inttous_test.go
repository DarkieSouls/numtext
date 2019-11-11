package numtext

import (
	"testing"
)

func TestIntToWordsUS(t *testing.T) {
	tcs := []struct{
		number int64
		result string
	}{
		{1, "one"},
		{10, "ten"},
		{100,"one hundred"},
		{1000, "one thousand"},
		{1000000, "one million"},
		{1000000000, "one billion"},
		{1000000000000, "one trillion"},
		{1000000000000000, "one quadrillion"},
		{1000000000000000000, "one quintillion"},
		{1000056772000, "one trillion fifty six million seven hundred seventy two thousand"},
		{190037711212, "one hundred ninety billion thirty seven million seven hundred eleven thousand two hundred twelve"},
		{-5, "minus five"},
		{-515982772, "minus five hundred fifteen million nine hundred eighty two thousand seven hundred seventy two"},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(string(tc.number), func(t *testing.T) {
			num := IntToWordsUS(tc.number)
			if num != tc.result {
				t.Errorf("expected: %s, but got: %s", tc.result, num)
			}
		})
	}
}

func TestIntToWordsUSFmt(t *testing.T) {
	tcs := []struct{
		number int64
		result string
		fmt rune
	}{
		{1, "one", 'c'},
		{5, "five", 'h'},
		{10, "ten", 's'},
		{100,"one-hundred", 'H'},
		{1000, "one thousand", 's'},
		{1000000, "one million", 'c'},
		{1000000000000000, "one quadrillion", 'c'},
		{1000000000000000000, "one quintillion", 'S'},
		{1000056772000, "one-trillion-fifty-six-million-seven-hundred-seventy-two-thousand", 'h'},
		{190037711212, "one hundred ninety billion, thirty seven million, seven hundred eleven thousand, two hundred twelve", 'C'},
		{-515982772, "minus five-hundred-fifteen-million-nine-hundred-eighty-two-thousand-seven-hundred-seventy-two", 'h'},
		{-812800, "minus eight hundred twelve thousand, eight hundred", 'c'},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(string(tc.number), func(t *testing.T) {
			num := IntToWordsUSFmt(tc.number, tc.fmt)
			if num != tc.result {
				t.Errorf("expected: %s, but got: %s", tc.result, num)
			}
		})
	}
}
