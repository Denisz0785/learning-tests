package handlers_test

import (
	"ard_lab_test_learn/test4/handlers"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

func init() {
	handlers.Routes()
}

func TestSendJSON(t *testing.T) {
	url := "/sendjson"
	statusCode := 200

	t.Log("Given the need to test the SendJSON endpoint.")
	{
		r := httptest.NewRequest("GET", url, nil)
		w := httptest.NewRecorder()
		http.DefaultServeMux.ServeHTTP(w, r)

		testID := 0
		t.Logf("\tTest %d\tWhen checking %q for status code %d", testID, url, statusCode)
		{
			if w.Code != 200 {
				t.Fatalf("\t%s\tTest %d\tShould receive a status code of %d for the reponse. Received[%d]", failed, testID, statusCode, w.Code)
			}
			t.Logf("\t%s\tTest %d\tShould receive a status code of %d for the responce.Succeed", succeed, testID, statusCode)
		}
		var u struct {
			Name  string
			Email string
		}

		err := json.NewDecoder(w.Body).Decode(&u)

		if err != nil {
			t.Fatalf("\t%s\tTest %d\tShould be able to decode the response", failed, testID)
		}
		t.Logf("\t%s\tTest %d\tShould be able to decode the response.", succeed, testID)

		if u.Name == "Roma" {
			t.Logf("\t%s\tTest %d\tShould have \"Roma\" for name", succeed, testID)
		} else {
			t.Errorf("\t%s\tTest %d\tShould have \"Roma\" for name. Received[%s]", failed, testID, u.Name)
		}

		if u.Email == "omahung@svaha.ru" {
			t.Logf("\t%s\tTest %d\tShould have \"omahung\" for email", succeed, testID)
		} else {
			t.Errorf("\t%s\tTest %d\tShould have \"omahung\" for email", failed, testID)
		}
	}
}
