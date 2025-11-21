package router

import (
	"TP-API-Spotify/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	//--- ROUTES ---
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/album/damso", controller.AlbumDamso)
	mux.HandleFunc("/tracks/laylow", controller.TracksLaylow)
	//--- ROUTES ---
	return mux
}
