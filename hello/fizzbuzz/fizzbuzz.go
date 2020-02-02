package fizzbuzz

import "strconv"

func New(number int) Object {
	return Object{
		number: number,
	}
}

type Object struct {
	number int
}

func (o Object) String() string {
	return FizzBuzz(o.number)
}

func FizzBuzz(n int) string {
	if n%3 == 0 && n%5 == 0 {
		return "FizzBuzz"
	}

	if n%3 == 0 {
		return "Fizz"
	}

	if n%5 == 0 {
		return "Buzz"
	}

	return strconv.Itoa(n)
}
