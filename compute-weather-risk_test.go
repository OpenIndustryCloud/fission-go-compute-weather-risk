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
	userJson60     = `{"history":{"dailysummary":[{"fog":"0","maxpressurem":"1025","maxtempm":"7","maxwspdm":"40","minpressurem":"1014","mintempm":"0","minwspdm":"7","rain":"1","tornado":"0"}]},"response":{"version":"0.1"}}`
	userJson80     = `{"history":{"dailysummary":[{"fog":"0","maxpressurem":"1025","maxtempm":"7","maxwspdm":"70","minpressurem":"1014","mintempm":"0","minwspdm":"7","rain":"1","tornado":"0"}]},"response":{"version":"0.1"}}`
	userJson20     = `{"history":{"dailysummary":[{"fog":"0","maxpressurem":"1025","maxtempm":"7","maxwspdm":"10","minpressurem":"1014","mintempm":"0","minwspdm":"7","rain":"1","tornado":"0"}]},"response":{"version":"0.1"}}`
	outputJSON60   = `{"riskScore":60,"description":"Possibly stormy weather"}`
	outputJSON80   = `{"riskScore":80,"description":"Stormy Weather identified"}`
	outputJSON20   = `{"riskScore":20,"description":"Very less likelyhood of Storm"}`
	emptyReqOutput = `{"status":"400","message":"EOF"}`
	// ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr = httptest.NewRecorder()
)

func TestHandler(t *testing.T) {
	//Convert string to reader and
	//Create request with JSON body
	req20, err := http.NewRequest("POST", "", strings.NewReader(userJson20))
	req60, err := http.NewRequest("POST", "", strings.NewReader(userJson60))
	req80, err := http.NewRequest("POST", "", strings.NewReader(userJson80))
	reqEmpty, err := http.NewRequest("POST", "", strings.NewReader(""))
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
		{"Test Data", args{rr, req20}, outputJSON20},
		{"Test Data", args{rr, req60}, outputJSON60},
		{"Test Data", args{rr, req80}, outputJSON80},
		{"Test Data", args{rr, reqEmpty}, emptyReqOutput},
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
