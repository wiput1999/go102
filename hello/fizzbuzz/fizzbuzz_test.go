package fizzbuzz

import "testing"

func TestGivenOneSayOne(t *testing.T) {
	var want string = "1"
	var given int = 1

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenTwoSayTwo(t *testing.T) {
	var want string = "2"
	var given int = 2

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenThreeSayFizz(t *testing.T) {
	var want string = "Fizz"
	var given int = 3

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenFourSayFour(t *testing.T) {
	var want string = "4"
	var given int = 4

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenFiveSayBuzz(t *testing.T) {
	var want string = "Buzz"
	var given int = 5

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenSixSayFizz(t *testing.T) {
	var want string = "Fizz"
	var given int = 6

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenSevenSaySeven(t *testing.T) {
	var want string = "7"
	var given int = 7

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenEightSayEight(t *testing.T) {
	var want string = "8"
	var given int = 8

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenNineSayFizz(t *testing.T) {
	var want string = "Fizz"
	var given int = 6

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenTenSayBuzz(t *testing.T) {
	var want string = "Buzz"
	var given int = 10

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}

func TestGivenFifteenSayFizzBuzz(t *testing.T) {
	var want string = "FizzBuzz"
	var given int = 15

	var get string = FizzBuzz(given)

	if want != get {
		t.Errorf("given %d given %q but get %q", given, want, get)
	}
}
