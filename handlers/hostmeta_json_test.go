package handlers

import (
	"encoding/json"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/silverscat-3/hostmeta"
	"github.com/stretchr/testify/assert"
)

const expectedJSON = `
{
  "links": [
    {
      "type": "html",
      "href": "http://example.org"
    },
    {
      "type": "html",
      "href": "http://toyota.jp/yaris"
    }
  ]
}
`

func TestHostMetaJSONHandle(t *testing.T) {
	hmh := HostMetaJSONHandler{
		Links: []*hostmeta.Link{
			&hostmeta.Link{
				Href: "http://example.org",
				Type: "html",
			},
			&hostmeta.Link{
				Href: "http://toyota.jp/yaris",
				Type: "html",
			},
		},
	}

	ts := httptest.NewServer(http.HandlerFunc(hmh.ServeHTTP))
	defer ts.Close()

	r, err := http.Get(ts.URL)
	if nil != err {
		t.Fatalf("Failed test by http.Get(). %v", err)
	}

	ct, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if "application/json" != ct {
		t.Fatalf("Failed test! Invalid Content-Type. %v", r.Header.Get("Content-Type"))
	}

	if http.StatusOK != r.StatusCode {
		t.Fatalf("Failed test by HTTP error. %v", r.StatusCode)
	}

	body, err := ioutil.ReadAll(r.Body)
	if nil != err {
		t.Fatalf("Failed test by ioutil.ReadAll(). %v", err)
	}

	var expected interface{}
	json.Unmarshal([]byte(expectedJSON), &expected)

	var actual interface{}
	if err = json.Unmarshal(body, &actual); nil != err {
		t.Fatalf("Failed test! Invalid response. %v", string(body))
	}

	assert.Equal(t, expected, actual)
}
