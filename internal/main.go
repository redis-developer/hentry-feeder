package main

import (
	"net/http"
    "fmt"

	"github.com/YashKumarVerma/hentry-feeder/internal/redis"
	"github.com/YashKumarVerma/hentry-feeder/internal/server"
)

func main() {
	redis.Init()
    fmt.Println("Listening on port 9898")
	http.HandleFunc("/data", server.FormHandler)
	http.ListenAndServe(":9898", nil)
}
