package api

import (
	"net/http"
	"testing"
)

type MockRoundTriper struct {
	RoundTripperOutput *http.Response
}

func (m MockRoundTriper) RoundTrip(*http.Request) (*http.Response, error) {
	return m.RoundTripperOutput, nil
}

func TestRoundTrip(t *testing.T) {
	myJWTTrasnport := JWTTrasnport{
		Transport: MockRoundTriper{
			RoundTripperOutput: &http.Response{
				StatusCode: 200,
			},
		},
	}
	req := &http.Request{}
	res, err := myJWTTrasnport.RoundTrip(req)
	if err != nil {
		t.Fatalf("RoundTrip error %s\n", err.Error())
	}
	if res.StatusCode != 200 {
		t.Fatalf("Statuscode is not 200, got %d\n", res.StatusCode)
	}
}
