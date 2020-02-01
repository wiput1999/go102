package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/wiput1999/go102/hello/fizzbuzz"
)

func FizzBuzzHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	number, err := strconv.Atoi(vars["number"])

	if err != nil {
		s := fmt.Sprintf(`{"message": %q}`, err)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		io.WriteString(w, s)
		return
	}

	s := fmt.Sprintf(`{"message": "%s"}`, fizzbuzz.FizzBuzz(number))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, s)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/fizzbuzz/{number}", FizzBuzzHandler)

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
