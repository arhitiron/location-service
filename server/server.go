package server

import (
	"net/http"
	"encoding/json"
	"log"
	"gitlab.antyron.com/ITStudWay2017/location-service/service"
)

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
	http.HandleFunc("/location", h.updateLocation)
	http.ListenAndServe(h.address, nil)
}

func (h *simpleServer) updateLocation(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var req service.LocationRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
	}
	defer r.Body.Close()
	h.messageBroker.Send(req)
}
