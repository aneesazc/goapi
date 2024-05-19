package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	// "github.com/aneesazc/goapi/internal/handlers"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.SetReportCaller(true)
	var r *chi.Mux = chi.NewRouter()
	handlers.Handler(r)

	fmt.Println("Server GO API service...")
	fmt.Println(`
	________  ________          ________  ________  ___     
	|\   ____\|\   __  \        |\   __  \|\   __  \|\  \    
	\ \  \___|\ \  \|\  \       \ \  \|\  \ \  \|\  \ \  \   
	 \ \  \  __\ \  \\\  \       \ \   __  \ \   ____\ \  \  
	  \ \  \|\  \ \  \\\  \       \ \  \ \  \ \  \___|\ \  \ 
	   \ \_______\ \_______\       \ \__\ \__\ \__\    \ \__\
		\|_______|\|_______|        \|__|\|__|\|__|     \|__|`)

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Error(err)
	}
}
