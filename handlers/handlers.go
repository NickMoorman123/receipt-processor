package handlers

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/NickMoorman123/receipt-processor/errors"
	"github.com/NickMoorman123/receipt-processor/objects"
	"github.com/NickMoorman123/receipt-processor/store"
)

type IReceiptHandler interface {
	Process(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
}

type handler struct {
	store store.IReceiptStore
}

func NewReceiptHandler(store store.IReceiptStore) IReceiptHandler {
	return &handler{store: store}
}

func (h *handler) Process(w http.ResponseWriter, r *http.Request) {
	data, err := ioutil.ReadAll(r.Body)
	if err != nil {
		WriteError(w, errors.ErrUnprocessableEntity)
		return
	}
	receipt := &objects.Receipt{}
	if Unmarshal(w, data, receipt) != nil {
		return
	}
	if err := calculateReceiptPoints(receipt); err != nil {
		WriteError(w, err)
		return
	}
	if err = h.store.Process(r.Context(), &objects.ProcessRequest{Receipt: receipt}); err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &objects.ProcessResponseWrapper{UUID: receipt.UUID})
}

func (h *handler) Get(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")
	uuid := pathParts[len(pathParts) - 2]
	receipt, err := h.store.Get(r.Context(), &objects.GetRequest{UUID: uuid})
	if err != nil {
		WriteError(w, err)
		return
	}
	WriteResponse(w, &objects.GetResponseWrapper{Points: receipt.Points})
}