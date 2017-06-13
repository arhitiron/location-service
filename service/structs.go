package service

type LocationRequest struct {
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	User      string `json:"user"`
	Created   int64  `json:"created"`
}
