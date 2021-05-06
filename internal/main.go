package main

import (
	"net/http"

	"github.com/YashKumarVerma/hentry-feeder/internal/server"
)

func main() {
	// redis.Init()

	http.HandleFunc("/data", server.FormHandler)
	http.ListenAndServe(":9898", nil)
}
