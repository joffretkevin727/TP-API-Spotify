package router

import (
	"TP-API-Spotify/controller"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()

	//--- ROUTES ---
	mux.HandleFunc("/home", controller.Home)
	mux.HandleFunc("/album", controller.Album)
	mux.HandleFunc("/album/damso", controller.AlbumDamso)
	mux.HandleFunc("/track", controller.Track)
	mux.HandleFunc("/track/laylow", controller.TrackLaylow)
	//--- ROUTES ---
	return mux
}
