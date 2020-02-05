package handlers

import (
	"encoding/xml"
	"net/http"

	"github.com/silverscat-3/hostmeta"
)

type HostMetaHandler struct {
	Links []*hostmeta.Link
}

func (h *HostMetaHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	type alias []*hostmeta.Link
	l := &struct {
		*alias  `xml:"link"`
		XMLName xml.Name `xml:"XRD"`
		Xmlns   string   `xml:"xmlns,attr"`
	}{
		alias: (*alias)(&h.Links),
		Xmlns: "http://docs.oasis-open.org/ns/xri/xrd-1.0",
	}

	body, err := xml.Marshal(l)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/xrd+xml")
	w.Write(body)
}
