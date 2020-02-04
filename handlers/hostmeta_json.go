package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/silverscat-3/hostmeta"
)

type HostMetaJSONHandler struct {
	Links []hostmeta.Link
}

func (h *HostMetaJSONHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if http.MethodGet != r.Method {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	l := &struct {
		Links *[]hostmeta.Link `json:"links"`
	}{
		Links: &h.Links,
	}

	body, err := json.Marshal(l)
	if nil != err {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("content-type", "application/json")
	w.Write(body)
}
