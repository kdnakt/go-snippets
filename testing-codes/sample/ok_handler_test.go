package sample_test

import (
	"io/ioutil"
	"net/http/httptest"
	"testing"

	"github.com/kdnakt/go-snippets/testing-codes/sample"
)

func TestOkHandler(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/sample", nil)

	sample.OkHandler(w, r)
	res := w.Result()
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("ioutil.ReadAll() returns unexpected error `%#v`", err)
	}

	const expected = "{\"status\":\"OK\"}\n"
	if got := string(b); got != expected {
		t.Errorf("OkHandler response = `%#v`, want `%#v`", got, expected)
	}
}
