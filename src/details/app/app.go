package app

import (
	"details/types"
	"fmt"
	"log"
	"net/http"

	"github.com/BurntSushi/toml"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/joe-davidson1802/hotwirehandler"
)

func Run(env string) {
	var c types.Config

	config := "local.toml"

	if env != "" {
		config = env + ".toml"
	}

	fmt.Println("reading config...")

	if _, err := toml.DecodeFile(config, &c); err != nil {
		log.Fatal(err)
	}

	fmt.Println("config loaded.")

	r := mux.NewRouter().
		PathPrefix("/apps/details").
		Subrouter()

	r.Handle("/restaurant/{id}", hotwirehandler.New(RestaurantHandler{Config: c})).Methods("GET", "OPTIONS")

	corsObj := handlers.AllowedOrigins([]string{"*"})
	methods := handlers.AllowedMethods([]string{"GET", "OPTIONS"})
	headers := handlers.AllowedHeaders([]string{"turbo-frame"})

	err := http.ListenAndServe(c.Listen, handlers.CORS(corsObj, headers, methods)(r))

	log.Fatal(err)
}
