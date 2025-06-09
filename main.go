package main

import (
	"fmt"
	"log"
	"net/http"
)

func moodHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/v1/mood/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "male male")
}

func main() {
	http.HandleFunc("/v1/mood/", moodHandler)
	log.Println("Server in ascolto su http://localhost:8080/v1/mood/")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
