// This package contains the http handler for getting the highest prime number
// less than the user's input
package service

import (
	"net/http"
	"testing"
)

func Test_handlePrimeToLeft(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			handlePrimeToLeft(tt.args.w, tt.args.r)
		})
	}
}
