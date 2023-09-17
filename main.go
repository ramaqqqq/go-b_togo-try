package main

import (
	"bookingoto-try/config"
	"bookingoto-try/helpers"
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		helpers.Logger("error", "Error getting env")
	}

	db := config.Init()
	defer db.Close()

	r := mux.NewRouter()

	presenter := config.Route(r, db)
	fmt.Println(presenter)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"POST", "GET", "OPTIONS", "PUT", "DELETE"},
		AllowedHeaders: []string{"Accept", "content-type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})

	p := os.Getenv("PORT")
	h := c.Handler(r)
	s := new(http.Server)
	s.Handler = h
	s.Addr = ":" + p
	appASCIIArt := fmt.Sprintf(`

		▄▄▄▄·             ▄ •▄ ▪   ▐ ▄  ▄▄ • ▄▄▄▄▄       ▄▄ •       
		▐█ ▀█▪▪     ▪     █▌▄▌▪██ •█▌▐█▐█ ▀ ▪•██  ▪     ▐█ ▀ ▪▪     
		▐█▀▀█▄ ▄█▀▄  ▄█▀▄ ▐▀▀▄·▐█·▐█▐▐▌▄█ ▀█▄ ▐█.▪ ▄█▀▄ ▄█ ▀█▄ ▄█▀▄ 
		██▄▪▐█▐█▌.▐▌▐█▌.▐▌▐█.█▌▐█▌██▐█▌▐█▄▪▐█ ▐█▌·▐█▌.▐▌▐█▄▪▐█▐█▌.▐▌
		·▀▀▀▀  ▀█▄▀▪ ▀█▄▀▪·▀  ▀▀▀▀▀▀ █▪·▀▀▀▀  ▀▀▀  ▀█▄▀▪·▀▀▀▀  ▀█▄▀▪
				
		Developed by Lutfi M 
		Server run in port %s
	`, s.Addr)
	fmt.Println(appASCIIArt)
	s.ListenAndServe()
}
