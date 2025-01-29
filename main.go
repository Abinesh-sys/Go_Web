package main

import (
    "log"
    "net/http"
	"go_web/routes"
)

func main() {

	router := routes.Registerroutes()

    log.Println("Server started at http://localhost:9090")
    log.Fatal(http.ListenAndServe(":9090", router))
}


