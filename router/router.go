package router

import (
	"TP-API-Spotify/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	//--- ROUTES ---
	mux.HandleFunc("/home", controller.Home)
	return mux
}
