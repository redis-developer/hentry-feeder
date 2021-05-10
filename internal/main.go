package main

import (
	"fmt"
	"net/http"

	"github.com/YashKumarVerma/hentry-feeder/internal/config"
	"github.com/YashKumarVerma/hentry-feeder/internal/redis"
	"github.com/YashKumarVerma/hentry-feeder/internal/server"
)

func main() {
	config.Init()
	redis.Init()

	fmt.Println("Listening on port :" + config.Load.PORT)
	http.HandleFunc("/timeseries", server.TimeSeriesFormHandler)
	http.HandleFunc("/snapshot", server.SnapshotFormHandler)
	http.ListenAndServe(":"+config.Load.PORT, nil)
}
