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
	mux.HandleFunc("/albumdamso", controller.AlbumDamso)
	mux.HandleFunc("/track", controller.Track)
	mux.HandleFunc("/tracklaylow", controller.TrackLaylow)
	//--- ROUTES ---
	return mux
}
