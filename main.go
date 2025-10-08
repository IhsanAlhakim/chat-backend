package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	PORT := ":8080"
	fmt.Println("React Go ChatApp v1")
	fmt.Println("Server started at Port" + PORT)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	if err := http.ListenAndServe(PORT, nil); err != nil {
		log.Println("Shutting down server...")
	}
}
