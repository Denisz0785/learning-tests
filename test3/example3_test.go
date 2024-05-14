package test3

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

const succeed = "\u2713"
const failed = "\u2717"

// feed is mocking the XML document we expect to receive.
var feed = ` <?xml version="1.0" encoding="UTF-8"?>
	<rss>
	<channel>
		<title>Going go programming</title>
		<description>Golang : https ://github.goinggo</description>
		<link>http://www.goinggo.net/</link>
		<item>
			<pubDate>Sun 15 03 2015 15:04:00</pubDate>
			<title>OM A HUM</title>
		</item>
	</channel>
	</rss>`

type Item struct {
	XMLName     xml.Name `xml:"item"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
}

// Item defynes the fields accociated with the item tag
type Channel struct {
	XMLName     xml.Name `xml:"channel"`
	Title       string   `xml:"title"`
	Description string   `xml:"description"`
	Link        string   `xml:"link"`
	PubDate     string   `xml:"pubDate"`
	Items       []Item   `xml:"item"`
}

type Document struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
	URI     string
}

// mockServer return a pointer to a server
func mockServer() *httptest.Server {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/xml")
		fmt.Fprintln(w, feed)
	}

	return httptest.NewServer(http.HandlerFunc(f))
}

// TestDownload validates the http Get function can download content
// and the content be unmarshaled and clean.
func TestDownload(t *testing.T) {
	statusCode := http.StatusNoContent

	server := mockServer()
	defer server.Close()

	t.Log("Given the need to test downloading content.")
	{
		t.Logf("\tTest 0:\tWhen checking %q for status code %d", server.URL, statusCode)
		{
			resp, err := http.Get(server.URL)
			if err != nil {
				t.Fatalf("\t%s\tShould be able to make the Get call : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to make the Get call.", succeed)
			defer resp.Body.Close()

			if resp.StatusCode != statusCode {
				t.Fatalf("\t%s\tShould receive a %d status code : %v", failed, statusCode, resp.StatusCode)
			}
			t.Logf("\t%s\tShould receive a %d status code.", succeed, statusCode)

			var d Document
			if err := xml.NewDecoder(resp.Body).Decode(&d); err != nil {
				t.Fatalf("\t%s\tShould be able to unmasrachal the response : %v", failed, err)
			}
			t.Logf("\t%s\tShould be able to unmarshal the response.", succeed)

			if len(d.Channel.Items) == 1 {
				t.Logf("\t%s\tShould have 1 item in the feed", succeed)
			} else {
				t.Errorf("\t%s\tShould have 1 item in the feed : %d", failed, len(d.Channel.Items))
			}
		}

	}
}
