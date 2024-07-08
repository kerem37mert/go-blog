package server

import (
	"log"
	"net/http"
)

func Start(port string) {
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
