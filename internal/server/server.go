package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// FormHandler handles the incoming data
func FormHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)

	var body PostBody
	err := decoder.Decode(&body)
	if err != nil {
		fmt.Fprintf(w, "BAD REQUEST")
	}

	fmt.Println(body)
	fmt.Fprintf(w, "OK")
}
