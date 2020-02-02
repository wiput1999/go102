package main

import (
	"reflect"
	"testing"
)

type fakeRandomer string

func (r fakeRandomer) Intn(int) int {
	return 3
}

func TestFizzBuzzByRandom(t *testing.T) {
	r := fakeRandomer("")

	want := response{
		Message: "Fizz",
		Number:  "3",
	}

	get := fizzBuzzByRandom(r)

	if !reflect.DeepEqual(want, get) {
		t.Error("not equal")
	}

}
