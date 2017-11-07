package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

var (
	server *httptest.Server
	//Test Data TV{"description":"Possible Stormy weather","riskScore":50}
	userJson   = `{"history":{"dailysummary":[{"fog":"0","maxpressurem":"1025","maxtempm":"7","maxwspdm":"28","minpressurem":"1014","mintempm":"0","minwspdm":"7","rain":"1","tornado":"0"}]},"response":{"version":"0.1"}}`
	outputJSON = `{"riskScore":60,"description":"Possibly stormy weather"}`
	// ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr = httptest.NewRecorder()
)

func TestHandler(t *testing.T) {
	//Convert string to reader and
	//Create request with JSON body
	req, err := http.NewRequest("POST", "", strings.NewReader(userJson))
	if err != nil {
		t.Error(err) //Something is wrong while sending request
	}

	//TEST CASES
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test Data", args{rr, req}, outputJSON},
	}
	for _, tt := range tests {
		// call ServeHTTP method
		// directly and pass Request and ResponseRecorder.
		handler := http.HandlerFunc(Handler)
		handler.ServeHTTP(tt.args.w, tt.args.r)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
		//check content type
		if ctype := rr.Header().Get("Content-Type"); ctype != "application/json" {
			t.Errorf("content type header does not match: got %v want %v",
				ctype, "application/json")
		}
		// check the output
		res, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Error(err) //Something is wrong while read res
		}
		got := string(res)
		if got != tt.want {
			t.Errorf("%q. compute weather risk() = %v, want %v", tt.name, got, tt.want)
		}
	}
}