package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

var (
	errNilBody = errors.New("nil body")

	msgNilBody     = "expected body, got nil"
	msgUserCreated = func(id int) string {
		return fmt.Sprintf("user created with id %d", id)
	}
)

type response struct {
	Error   error       `json:"error"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

func respond(w http.ResponseWriter, code int, resp response) {
	b, _ := json.Marshal(resp)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	_, err := w.Write(b)
	if err != nil {
		log.Println(err)
	}
}
