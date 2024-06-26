package test2

import (
	"net/http"
	"testing"
)

const (
	succeed = "\u2713"
	failed  = "\u2717"
)

// TestDownload validates the http Get function can download content
// and handles different status conditions properly.
func TestDownload(t *testing.T) {
	tests := []struct {
		url        string
		statusCode int
	}{
		{"http://www.baidu.com", http.StatusOK},
		{"http://www.baidu.com", http.StatusNotFound},
	}

	t.Log("Given the need to test the download function.")
	{
		for i, tt := range tests {
			t.Logf("\tTest %d:\tWhen checking %q for status codes %d", i, tt.url, tt.statusCode)
			{
				resp, err := http.Get(tt.url)
				if err != nil {
					t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
				}
				t.Logf("\t%s\tShould be able to make the Get call.", succeed)

				defer resp.Body.Close()

				if resp.StatusCode == tt.statusCode {
					t.Logf("\t%s\tShould receive a %d status code", succeed, tt.statusCode)
				} else {
					t.Errorf("\t%s\tShould receive a %d statuc code", failed, tt.statusCode)
				}
			}
		}
	}
}
