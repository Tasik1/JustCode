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
	r.HandleFunc("/endpoint_a", handleRequestA).Methods("GET")
	r.HandleFunc("/endpoint_a/getB", getB).Methods("GET")

	log.Println(http.ListenAndServe("127.0.0.1:8080", r))
}

func handleRequestA(w http.ResponseWriter, r *http.Request) {
	messageToSend := "Привет от сервиса A!"
	fmt.Fprintln(w, messageToSend)
}

func getB(w http.ResponseWriter, r *http.Request) {
	response, err := http.Get("http://localhost:8081/endpoint_b")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	message, err := io.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Сейчас мы в сервисе А и Ответ от сервиса B:", string(message))
}
