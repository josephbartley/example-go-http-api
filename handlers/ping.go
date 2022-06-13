package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PongResponse struct {
	Message string `json:"message"`
}

func PingHandler(w http.ResponseWriter, req *http.Request) {
	response := PongResponse{
		Message: "pong",
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(bytes))
}
