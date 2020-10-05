package main

import (
	"encoding/json"
	"github.com/hunkeelin/server"
	"net/http"
	"time"
)

type Output struct {
	Message   string    `json:"message"`   // Helloworld message
	TimeStamp time.Time `json:"timestamp"` // timestamp in pst
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	output := Output{
		Message:   "hello world",
		TimeStamp: time.Now(),
	}
	json.NewEncoder(w).Encode(output)
	return
}

func startServer() {
	con := http.NewServeMux()
	con.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		mainHandler(w, r)
	})
	j := &server.ServerConfig{
		BindPort: "8080",
		BindAddr: "",
		ServeMux: con,
	}
	server.Server(j)
}
func main() {
	startServer()
}
