package x_test

import (
	"testing"

	"github.com/xpetit/x/v2"
	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

func TestFormatByte(t *testing.T) {
	expected := map[int]string{
		0:           "0 B",
		1:           "1 B",
		9:           "9 B",
		10:          "10 B",
		11:          "11 B",
		49:          "49 B",
		50:          "50 B",
		51:          "51 B",
		99:          "99 B",
		100:         "100 B",
		101:         "101 B",
		149:         "149 B",
		150:         "150 B",
		151:         "151 B",
		249:         "249 B",
		250:         "250 B",
		251:         "251 B",
		999:         "999 B",
		1000:        "1 kB",
		1001:        "1 kB",
		1049:        "1 kB",
		1050:        "1.1 kB",
		1051:        "1.1 kB",
		1449:        "1.4 kB",
		1450:        "1.4 kB",
		1451:        "1.5 kB",
		1900:        "1.9 kB",
		1949:        "1.9 kB",
		1950:        "2 kB",
		1951:        "2 kB",
		1999:        "2 kB",
		2000:        "2 kB",
		2001:        "2 kB",
		4999:        "5 kB",
		5000:        "5 kB",
		5001:        "5 kB",
		9499:        "9.5 kB",
		9500:        "9.5 kB",
		9501:        "9.5 kB",
		9901:        "9.9 kB",
		9949:        "9.9 kB",
		9950:        "10 kB",
		9951:        "10 kB",
		9999:        "10 kB",
		10_000:      "10 kB",
		19_999:      "20 kB",
		99_999:      "100 kB",
		100_000:     "100 kB",
		900_000:     "900 kB",
		999_000:     "999 kB",
		999_499:     "999 kB",
		999_500:     "1 MB",
		999_501:     "1 MB",
		999_999:     "1 MB",
		1_000_000:   "1 MB",
		1_000_001:   "1 MB",
		31_350_000:  "31 MB",
		31_450_000:  "31 MB",
		999_499_999: "999 MB",
		999_500_000: "1 GB",
	}
	keys := maps.Keys(expected)
	slices.Sort(keys)
	for _, i := range keys {
		want := expected[i]
		if got := x.FormatByte(i); got != want {
			t.Errorf("for i=%d, got:%q, want:%q", i, got, want)
		}
	}
}
