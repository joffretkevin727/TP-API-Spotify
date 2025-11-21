package main

import (
	api "TP-API-Spotify/API"
	"TP-API-Spotify/router"
	"fmt"
	"net/http"
)

func main() {
	api.GetToken()
	//"https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"

	r := router.New()
	fmt.Println("http://localhost:8080/home")
	http.ListenAndServe(":8080", r)

}
