package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/silverscat-3/hostmeta"
)

const sampleXML = "<XRD xmlns=\"http://docs.oasis-open.org/ns/xri/xrd-1.0\"><link rel=\"http://example.org\" type=\"html\"></link><link rel=\"http://toyota.jp/yaris\" type=\"html\"></link></XRD>"

func TestHostMetaHandle(t *testing.T) {
	hmh := HostMetaHandler{
		Links: []hostmeta.Link{
			hostmeta.Link{
				Rel:  "http://example.org",
				Type: "html",
			},
			hostmeta.Link{
				Rel:  "http://toyota.jp/yaris",
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

	body, err := ioutil.ReadAll(r.Body)
	if nil != err {
		t.Fatalf("Failed test by ioutil.ReadAll(). %v", err)
	}

	if string(body) != sampleXML {
		t.Fatalf("Failed test by Data error. %v", string(body))
	}
}
