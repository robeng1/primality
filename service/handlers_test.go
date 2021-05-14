// This package contains the http handler for getting the highest prime number
// less than the user's input
package service

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func Test_handlePrimeToLeft(t *testing.T) {
	okReq, err := http.NewRequest("GET", "/55", nil)
    if err != nil {
        t.Fatal(err)
    }

    badInputReq, err := http.NewRequest("GET", `/"yes"`, nil)
    if err != nil {
        t.Fatal(err)
    }
    healthReq, err := http.NewRequest("GET", `/`, nil)
    if err != nil {
        t.Fatal(err)
    }

    // We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
    okRecorder := httptest.NewRecorder()
    badInputRecorder := httptest.NewRecorder()
    healthRecorder := httptest.NewRecorder()
	
    router := mux.NewRouter()
    router.HandleFunc("/{num}", handlePrimeToLeft)
	router.HandleFunc("/", handleHealthCheck)
  
    // Our handlers satisfy http.Handler, so we can call their ServeHTTP method 
    // directly and pass in our Request and ResponseRecorder.
    router.ServeHTTP(okRecorder, okReq)
    router.ServeHTTP(badInputRecorder, badInputReq)
    router.ServeHTTP(healthRecorder, healthReq)

    // OK check
    // Check the status code is what we expect.
    if status := okRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected := `{"result":"53"}`
    if okRecorder.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            okRecorder.Body.String(), expected)
    }

    // Bad input check
    // Check the status code is what we expect.
    if status := badInputRecorder.Code; status != http.StatusBadRequest {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusBadRequest)
    }

    // Check the response body is what we expect.
    expected = `{"error":"Bad Input"}`
    if badInputRecorder.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            badInputRecorder.Body.String(), expected)
    }

    // Empty check
    // Check the status code is what we expect.
    if status := healthRecorder.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v",
            status, http.StatusOK)
    }

    // Check the response body is what we expect.
    expected = `{"alive": true}`
    if healthRecorder.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v",
            healthRecorder.Body.String(), expected)
    }
}
