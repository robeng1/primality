// This package contains the http handler for getting the highest prime number
// less than the user's input
package service

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/stably/prime"
)

// this is the request handler for finding the highest prime number less than a given input
func handlePrimeToLeft(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	// grab the input value
	input := vars["num"]

	// if the input is empty the user did not enter anything
	if input == "" {
		WriteError(w, http.StatusBadRequest, errors.New("Input cannot be empty"))
	} else {
		// check that the input value is valid
		val, err := strconv.ParseInt(input, 10, 64)
		if err != nil {
			WriteError(w, http.StatusBadRequest, errors.New("Bad Input"))
		} else {

			if val <= 0 {
				// negative values are not accepted return an error to the user
				WriteError(w, http.StatusBadRequest, errors.New("Negative numbers or zero is/are not allowed"))
			} else if val <= 2 {
				// 2 is the lowest prime number so we report that there's no prime number lower than 2
				WriteError(w, http.StatusNotFound, errors.New("Prime not found"))
			} else {
				// find the highest prime number less than Val
				primeNumber := prime.FindPrimeToLeft(uint64(val))

				if primeNumber == 0 {
					// report if no prime number is found
					// we're unlikely to reach here since a prime number will always be found
					// for values > 2
					WriteError(w, http.StatusNotFound, errors.New("Prime not found"))
				} else {
					WriteData(w, http.StatusOK, primeNumber)
				}
			}
		}

	}

}

func HandleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	// this takes the input number
	router.HandleFunc("/{num}", handlePrimeToLeft)

	// the port number is hard code to make it simple
	// for a more complex up, it could be provided through env variables
	log.Println("Serving on port 10000")
	log.Fatal(http.ListenAndServe(":10000",
		handlers.CORS(handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}
