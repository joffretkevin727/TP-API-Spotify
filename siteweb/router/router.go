package router

import (
	"TP-API-Spotify/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	//--- ROUTES ---
	mux.HandleFunc("/", controller.Home)
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/album/damso", controller.AlbumDamso)
	mux.HandleFunc("/tracks/laylow", controller.TracksLaylow)

	// --- STATIC FILES ---
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static/", fileServer))

	//--- ROUTES ---
	return mux
}
