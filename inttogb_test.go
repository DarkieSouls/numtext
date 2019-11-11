package numtext

import (
	"testing"
)

func TestIntToWordsGB(t *testing.T) {
	tcs := []struct{
		number int64
		result string
	}{
		{1, "one"},
		{10, "ten"},
		{100,"one hundred"},
		{1000, "one thousand"},
		{1000000, "one million"},
		{1000000000, "one thousand million"},
		{1000000000000, "one billion"},
		{1000000000000000, "one thousand billion"},
		{1000000000000000000, "one million billion"},
		{1000056772000, "one billion fifty six million seven hundred seventy two thousand"},
		{190037711212, "one hundred ninety thousand thirty seven million seven hundred eleven thousand two hundred twelve"},
		{-5, "minus five"},
		{-515982772, "minus five hundred fifteen million nine hundred eighty two thousand seven hundred seventy two"},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(string(tc.number), func(t *testing.T) {
			num := IntToWordsGB(tc.number)
			if num != tc.result {
				t.Errorf("expected: %s, but got: %s", tc.result, num)
			}
		})
	}
}

func TestIntToWordsGBFmt(t *testing.T) {
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
		{1000000000000000, "one thousand billion", 'c'},
		{1000000000000000000, "one million billion", 'S'},
		{1000056772000, "one-billion-fifty-six-million-seven-hundred-seventy-two-thousand", 'h'},
		{190037711212, "one hundred ninety thousand, thirty seven million, seven hundred eleven thousand, two hundred twelve", 'C'},
		{-515982772, "minus five-hundred-fifteen-million-nine-hundred-eighty-two-thousand-seven-hundred-seventy-two", 'h'},
		{-812800, "minus eight hundred twelve thousand, eight hundred", 'c'},
	}

	for _, tc := range tcs {
		tc := tc
		t.Run(string(tc.number), func(t *testing.T) {
			num := IntToWordsGBFmt(tc.number, tc.fmt)
			if num != tc.result {
				t.Errorf("expected: %s, but got: %s", tc.result, num)
			}
		})
	}
}
