package main

import (
	api "TP-API-Spotify/API"
	"TP-API-Spotify/router"
	"fmt"
	"net/http"
)

// FONCTION PRINCIPAL
func main() {
	api.GetToken()
	//liens utile au TP
	//"https://api.spotify.com/v1/artists/2UwqpfQtNuhBwviIC0f2ie/albums"
	//"https://api.spotify.com/v1/tracks/0nAHBAlzkyaQXUp7qTULqv"

	r := router.New()
	fmt.Println("http://localhost:8080/home")
	http.ListenAndServe(":8080", r)

}
