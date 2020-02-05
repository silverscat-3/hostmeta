package handlers

import (
	"encoding/xml"
	"io/ioutil"
	"mime"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/silverscat-3/hostmeta"
	"github.com/stretchr/testify/assert"
)

const expectedXML = `
<?xml version="1.0"?>
<XRD xmlns="http://docs.oasis-open.org/ns/xri/xrd-1.0">
  <link type="html" href="http://example.org"/>
  <link type="html" href="http://toyota.jp/yaris"/>
</XRD>
`

func TestHostMetaHandle(t *testing.T) {

	hmh := HostMetaHandler{
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
		t.Fatalf("Failed test! http.Get(). %v", err)
	}

	ct, _, _ := mime.ParseMediaType(r.Header.Get("Content-Type"))
	if "application/xrd+xml" != ct {
		t.Fatalf("Failed test! Invalid Content-Type. %v", r.Header.Get("Content-Type"))
	}

	if http.StatusOK != r.StatusCode {
		t.Fatalf("Failed test! HTTP error. %v", r.StatusCode)
	}

	body, err := ioutil.ReadAll(r.Body)
	if nil != err {
		t.Fatalf("Failed test! ioutil.ReadAll(). %v", err)
	}

	var expected interface{}
	xml.Unmarshal([]byte(expectedXML), &expected)

	var actual interface{}
	if err = xml.Unmarshal(body, &actual); nil != err {
		t.Fatalf("Failed test! Invalid response type. %v", string(body))
	}

	assert.Equal(t, expected, actual)
}
