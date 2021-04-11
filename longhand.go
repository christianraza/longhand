package longhand

import (
	"errors"
	"strings"
)

const maxInt32 = 1<<31 - 1
const minInt32 = -1 << 31

var numerals = map[string]int64{
	"zero":      0,
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"ten":       10,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eighteen":  18,
	"nineteen":  19,
	"twenty":    20,
	"thirty":    30,
	"forty":     40,
	"fifty":     50,
	"sixty":     60,
	"seventy":   70,
	"eighty":    80,
	"ninety":    90,
}

// Exclusive of 10
var ordersOfMagnitude = map[string]int64{
	"hundred":     100,
	"thousand":    1000,
	"k":           1000,
	"million":     1000000,
	"mil":         1000000,
	"m":           1000000,
	"billion":     1000000000,
	"bil":         1000000000,
	"b":           1000000000,
	"trillion":    1000000000000,
	"tril":        1000000000000,
	"t":           1000000000000,
	"quadrillion": 1000000000000000,
	"quad":        1000000000000000,
	"quintillion": 1000000000000000000,
	"quin":        1000000000000000000,
}

// LonghandError stores relevant information associated with a longhand error.
type LonghandError struct {
	Field   string // The infringing field
	Err     error  // The corresponding error
	FunName string // The associated function name
}

func (err *LonghandError) Error() string {
	return err.Err.Error() + ": \"" + err.Field + "\" in " + "longhand." + err.FunName
}

func rangeError(funName, s string) *LonghandError {
	return &LonghandError{s, errors.New("value out of range"), funName}
}

func syntaxError(funName, s string) *LonghandError {
	return &LonghandError{s, errors.New("invalid syntax"), funName}
}

func mul64(x int64, y int64) (int64, error) {
	const fnMul64 = "Mul64"
	z := x * y
	if y != 0 && x != z/y {
		return 0, errors.New("integer overflow")
	}
	return z, nil
}

// ParseLonghand64 interprets a longhand string lh and
// returns the corresponding 64-bit signed integer i64.
//
// If lh is empty or not in valid longhand, an error is
// returned and i64 = 0.
//
// The range for ParseLonghand64 is between -9,223,372,036,854,775,808
// and 9,223,372,036,854,775,807 (inclusive). If lh is not in this range
// an error is returned and i64 = 0.
//
// The errors returned by ParseLonghand64 are of type *LonghandError.
func ParseLonghand64(lh string) (i64 int64, err error) {
	const funName = "ParseLonghand64"

	if lh == "" {
		return 0, syntaxError(funName, lh)
	}

	var neg int64 = 1
	var preN, sum int64
	var in bool
	var b strings.Builder

	size := len(lh)
	for i := 0; i <= size; i++ {
		var c uint8
		if i < size {
			c = lh[i]
		}
		if c != ' ' && c != '\t' && c != '\r' && c != '\n' && c != '-' && i < size {
			in = true
			if 'A' <= c && c <= 'Z' {
				c += 32
			}
			b.WriteString(string(c))
		} else if in {
			// Handle current string in buffer
			s := b.String()
			// Validate s
			n, isN := numerals[s]
			m, isM := ordersOfMagnitude[s]
			if !isN && !isM {
				if s == "negative" {
					neg = -1
				} else if s != "and" {
					return 0, syntaxError(funName, s)
				}
			}
			// Process s
			if isN {
				if 0 < preN/10 && 0 < n/10 {
					sum *= 100
				} else if preN == n {
					sum *= 10
				}
				sum += n
				preN = n
			}
			if isM {
				if sum == 0 {
					sum += m
				} else {
					sum, err = mul64(sum, m)
					if err != nil {
						return 0, rangeError(funName, lh)
					}
				}
				if m != 100 {
					i64 += sum
					sum = 0
				}
				preN = 0
			}
			b.Reset()
			in = false
		}
	}

	i64 += sum
	i64 *= neg
	if neg == -1 && i64 > 0 || neg == 1 && i64 < 0 {
		return 0, rangeError(funName, lh)
	}
	return i64, nil
}

// ParseLonghand interprets a longhand string lh and
// returns the corresponding 32-bit signed integer i.
//
// If lh is empty or not in valid longhand, an error is
// returned and i = 0.
//
// The range for ParseLonghand is between -2,147,483,648
// and 2,147,483,647 (inclusive). If lh is not in this range
// an error is returned and i = 0.
//
// The errors returned by ParseLonghand are of type *LonghandError.
//
// ParseLonghand is equivalent to calling ParseLonghand64,
// enforcing 32-bit signed integer range, and performing an
// int type conversion.
func ParseLonghand(lh string) (i int, err error) {
	const funName = "ParseLonghand"

	i64, err := ParseLonghand64(lh)
	if e, ok := err.(*LonghandError); ok {
		e.FunName = funName
	}

	if minInt32 > i64 || i64 > maxInt32 {
		return 0, rangeError(funName, lh)
	}
	return int(i64), err
}
