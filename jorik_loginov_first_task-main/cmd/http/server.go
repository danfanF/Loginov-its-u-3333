package httpserver

import (
	"log"
	"net/http"
)

func InitHttpServer(port string) {
	routes := GetRoutes()

	if err := http.ListenAndServe(":"+port, routes); err != nil {
		log.Fatal(err)
	}
}
