package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/endpoint_b", handleRequestB).Methods("GET")
	r.HandleFunc("/endpoint_b/getA", getA).Methods("GET")

	log.Println(http.ListenAndServe("127.0.0.1:8081", r))
}

func handleRequestB(w http.ResponseWriter, r *http.Request) {
	messageToSend := "Привет от сервиса B!"
	fmt.Fprintln(w, messageToSend)
}

func getA(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8080/endpoint_a")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	messageToGet, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Сейчас мы в сервисе В и Ответ от сервиса A:", string(messageToGet))
}
