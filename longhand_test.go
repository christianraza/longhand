package longhand

import (
	"testing"
)

// ParseLonghand tests
func TestLonghandEmpty(t *testing.T) {
	i, err := ParseLonghand("")
	if err == nil {
		t.Fatal(err)
	}
	var expecting int = 0
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandSimple(t *testing.T) {
	i, err := ParseLonghand("two")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 2
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandComplex(t *testing.T) {
	i, err := ParseLonghand("seven thousand two hundred thirty one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 7231
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandRepetition(t *testing.T) {
	i, err := ParseLonghand("thirty thirty twenty-one one one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 30302111
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandYear(t *testing.T) {
	i, err := ParseLonghand("nineteen twenty")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 1920
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandAlternativeSpellingDash(t *testing.T) {
	i, err := ParseLonghand("seven thousand two hundred thirty-one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 7231
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandAlternativeSpellingAnd(t *testing.T) {
	i, err := ParseLonghand("seven thousand two hundred and thirty one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 7231
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandAlternativeSpellingDashAnd(t *testing.T) {
	i, err := ParseLonghand("seven thousand two hundred and thirty one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 7231
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandNegative(t *testing.T) {
	i, err := ParseLonghand("negative seven thousand two hundred thirty one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = -7231
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandNegativeEdgeCase(t *testing.T) {
	i, err := ParseLonghand(`negative two billion one hundred forty seven million
						     four hundred eighty three thousand six hundred forty eight`)
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = -2147483648
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandPositiveEdgeCase(t *testing.T) {
	i, err := ParseLonghand(`two billion one hundred forty seven million
						     four hundred eighty three thousand six hundred forty seven`)
	if err != nil {
		t.Fatal(err)
	}
	var expecting int = 2147483647
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandTypo(t *testing.T) {
	i, err := ParseLonghand("Three hunded twenty four thousand")
	if err == nil {
		t.Fatal(err)
	}
	var expecting int = 0
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandMix(t *testing.T) {
	i, err := ParseLonghand("five 1000 one hundred eighty seven")
	if err == nil {
		t.Fatal(err)
	}
	var expecting int = 0
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandMixLarge(t *testing.T) {
	i, err := ParseLonghand(`two billion one hundred forty seven million
						     four hundred eighty three thousand 647`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int = 0
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandNegativeOverflow(t *testing.T) {
	i, err := ParseLonghand(`negative two billion one hundred forty seven million
							 four hundred eighty three thousand six hundred forty nine`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int = 0
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

func TestLonghandPositiveOverflow(t *testing.T) {
	i, err := ParseLonghand(`two billion one hundred forty seven million
						     four hundred eighty three thousand six hundred forty eight`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int = 0
	if i != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i)
	}
}

// ParseLonghand64 tests
func TestLonghand64Empty(t *testing.T) {
	i64, err := ParseLonghand64("")
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64Complex(t *testing.T) {
	i64, err := ParseLonghand64("seven thousand two hundred thirty one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = 7231
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64Repetition(t *testing.T) {
	i64, err := ParseLonghand64("thirty thirty twenty-one one one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = 30302111
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64Year(t *testing.T) {
	i64, err := ParseLonghand64("nineteen twenty")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = 1920
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64AlternativeSpellingDash(t *testing.T) {
	i64, err := ParseLonghand64("seven thousand two hundred thirty-one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = 7231
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64AlternativeSpellingAnd(t *testing.T) {
	i64, err := ParseLonghand64("seven thousand two hundred and thirty one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = 7231
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64AlternativeSpellingDashAnd(t *testing.T) {
	i64, err := ParseLonghand64("seven thousand two hundred and thirty-one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = 7231
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64Negative(t *testing.T) {
	i64, err := ParseLonghand64("negative seven thousand two hundred thirty one")
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = -7231
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64NegativeEdgeCase(t *testing.T) {
	i64, err := ParseLonghand64(`negative nine quintillion two hundred twenty three quadrillion
								 three hundred seventy two trillion thirty six billion
								 eight hundred fifty four million seven hundred seventy
								 five thousand eight hundred eight`)
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = -9223372036854775808
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64PositiveEdgeCase(t *testing.T) {
	i64, err := ParseLonghand64(`nine quintillion two hundred twenty three quadrillion
								 three hundred seventy two trillion thirty six billion
								 eight hundred fifty four million seven hundred seventy
								 five thousand eight hundred seven`)
	if err != nil {
		t.Fatal(err)
	}
	var expecting int64 = 9223372036854775807
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64Typo(t *testing.T) {
	i64, err := ParseLonghand64("Three hunded twenty four thousand")
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64Mix(t *testing.T) {
	i64, err := ParseLonghand64("five 1000 one hundred eighty seven")
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64MixLarge(t *testing.T) {
	i64, err := ParseLonghand64(`nine quintillion two hundred twenty three quadrillion
								 three hundred seventy two trillion thirty six billion
								 eight hundred fifty four million seven hundred seventy
								 five 1807`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64NegativeOverflow(t *testing.T) {
	i64, err := ParseLonghand64(`negative nine quintillion two hundred twenty three quadrillion
							   	 three hundred seventy two trillion thirty six billion
							   	 eight hundred fifty four million seven hundred seventy
							     five thousand eight hundred nine`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64PositiveOverflow(t *testing.T) {
	i64, err := ParseLonghand64(`nine quintillion two hundred twenty three quadrillion
							   	 three hundred seventy two trillion thirty six billion
							     eight hundred fifty four million seven hundred seventy
							     five thousand eight hundred eight`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64NegativeWraparoundOverflow(t *testing.T) {
	i64, err := ParseLonghand64(`negative twenty quintillion two hundred twenty three quadrillion
							   	 three hundred seventy two trillion thirty six billion
							   	 eight hundred fifty four million seven hundred seventy
							   	 five thousand eight hundred nine`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}

func TestLonghand64PositiveWraparoundOverflow(t *testing.T) {
	i64, err := ParseLonghand64(`twenty quintillion two hundred twenty three quadrillion
							   	 three hundred seventy two trillion thirty six billion
							   	 eight hundred fifty four million seven hundred seventy
							   	 five thousand eight hundred eight`)
	if err == nil {
		t.Fatal(err)
	}
	var expecting int64 = 0
	if i64 != expecting {
		t.Fatalf("Expected: %d Got: %d", expecting, i64)
	}
}
