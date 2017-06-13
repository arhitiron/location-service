package server

import (
	"net/http"
	"encoding/json"
	"log"
	"gitlab.antyron.com/ITStudWay2017/location-service/service"
	"fmt"
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
	http.HandleFunc("/", h.main)
	http.HandleFunc("/location", h.updateLocation)
	http.ListenAndServe(h.address, nil)
}

func (h *simpleServer) main(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, string("Welcome to location service"))
}

func (h *simpleServer) updateLocation(w http.ResponseWriter, r *http.Request) {
	log.Printf("DEBUG: request updateLocation")
	decoder := json.NewDecoder(r.Body)
	var req service.LocationRequest
	err := decoder.Decode(&req)
	if err != nil {
		log.Println(err)
	}
	log.Printf("DEBUG: request %v", req)
	defer r.Body.Close()
	go h.messageBroker.Send(req)

	sendJsonResponse(struct {
		Message string
	}{Message: "location added to queue"}, w)
}

func sendJsonResponse(obj interface{}, w http.ResponseWriter) {
	res, err := json.Marshal(obj)
	if err != nil {
		fmt.Fprint(w, "error: %v", res)
	}
	w.Header().Add("Content-Type", "application/json")
	fmt.Fprint(w, string(res))
}
