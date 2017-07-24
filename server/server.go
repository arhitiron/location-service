package server

import (
	"encoding/json"
	"log"
	"fmt"
	"net/http"

	"github.com/arhitiron/location-service/service"
)

const (
	OK_MESSAGE          = "location added to queue"
	CONTENT_TYPE_HEADER = "Content-Type"
	APPLICATION_JSON    = "application/json"
)

var OkResponse string

func init() {
	res, err := json.Marshal(struct {
		Message string
	}{Message: OK_MESSAGE})

	if err != nil {
		panic(err)
	}
	OkResponse = string(res)
}

type Server interface {
	Handle()
}

type simpleServer struct {
	address       string
	messageBroker service.MessageBroker
}

func NewServer(addr string, mb service.MessageBroker) Server {
	return &simpleServer{addr, mb}
}

func (h *simpleServer) Handle() {
	http.HandleFunc("/", h.main)
	http.HandleFunc("/location", h.updateLocation)
	http.ListenAndServe(h.address, nil)
}

func (h *simpleServer) main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string("Welcome to location service"))
}

func (h *simpleServer) updateLocation(w http.ResponseWriter, r *http.Request) {
	w.Header().Add(CONTENT_TYPE_HEADER, APPLICATION_JSON)
	fmt.Fprint(w, OkResponse)
	go func(r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var req service.LocationRequest
		err := decoder.Decode(&req)
		if err != nil {
			log.Println(err)
		}
		defer r.Body.Close()
		h.messageBroker.Send(req)
	}(r)
}
