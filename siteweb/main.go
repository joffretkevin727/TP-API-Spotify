package main

import (
	"TP-API-Spotify/router"
	"fmt"
	"net/http"
)

func main() {
	r := router.New()
	fmt.Println("http://localhost:8080/home")
	http.ListenAndServe(":8080", r)
}
