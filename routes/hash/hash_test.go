package hash

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestHashHandler1_ServeHTTP(t *testing.T) {
	h := HashHandler{}

	form := url.Values{}
	form.Add("password", "angryMonkey")

	req, err := http.NewRequest("POST", "/hash", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Form = form

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Body.String() != "ZEHhWB65gUlzdVwtDQArEyx+KVLzp/aTaRaPlBzYRIFj6vjFdqEb0Q5B8zVKCZ0vKbZPZklJz0Fd7su2A+gf7Q==" {
		t.Errorf("unexpected response: %s", rec.Body.String())
	}
}

func TestHashHandler2_ServeHTTP(t *testing.T) {
	h := HashHandler{}

	form := url.Values{}
	// Missing password parameter

	req, err := http.NewRequest("POST", "/hash", strings.NewReader(form.Encode()))
	if err != nil {
		t.Fatal(err)
	}
	req.Form = form

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Body.String() != "Missing required field 'password'" {
		t.Errorf("unexpected response: %s", rec.Body.String())
	}
}

func TestHashHandler3_ServeHTTP(t *testing.T) {
	h := HashHandler{}

	req, err := http.NewRequest("GET", "/hash", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)

	if rec.Body.String() != "Missing required field 'password'" {
		t.Errorf("unexpected response: %s", rec.Body.String())
	}
}

func TestStatsHandler_ServeHTTP(t *testing.T) {
	h := StatsHandler{}

	req, err := http.NewRequest("GET", "/stats", nil)
	if err != nil {
		t.Fatal(err)
	}

	rec := httptest.NewRecorder()

	h.ServeHTTP(rec, req)
	//fmt.Println(rec.Body)

	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("content type header does not match: got %v want %v",
			ct, "application/json")
	}

	jsonRes := []byte(rec.Body.String())
	var resObj map[string]interface{}
	if err := json.Unmarshal(jsonRes, &resObj); err != nil {
		t.Fatal(err)
	}

	if _, ok := resObj["total"]; ok {
		total := resObj["total"].(float64) // interface{} is float64
		// THIS WILL NEED TO CHANGE IF YOU ADD MORE TESTS ABOVE
		if int(total) != 1 {
			t.Errorf("unexpected response: %x", resObj)
		}
	} else {
		t.Error("expected 'total' as a key in stats response.")
	}
}
