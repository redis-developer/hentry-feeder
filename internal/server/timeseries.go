package server

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/YashKumarVerma/hentry-feeder/internal/redis"
	"github.com/YashKumarVerma/hentry-feeder/internal/structure"
)

// FormHandler handles the incoming data
func TimeSeriesFormHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var body structure.PostBody
	err := decoder.Decode(&body)
	if err != nil {
		fmt.Fprintf(w, "BAD REQUEST")
		fmt.Println(err)
		return
	}

	go redis.Save("time", body)
	fmt.Fprintf(w, "OK")
}
