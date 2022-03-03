package main

import (
	"fmt"
	"log"
	"net/http"

	"ToDo_React+Go/server/routers"

	"github.com/iyashjayesh/Go-Lang-Resources/projects/ToDo_React+Go/server/routers"
)

func main() {

	fmt.Println("To Do Application")

	r := routers.Router()

	fmt.Println("Startinmg the server on the 8000...")
	log.Fatal(http.ListenAndServe(":9000", r))

}
