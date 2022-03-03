package routers

import (
	"fmt"
	"github.com/iyashjayesh/Go-Lang-Resources/projects/ToDo_React+Go/server/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router{

	router := mux.NewRouter()
	
	router.HandleFunc("/api/task", middleware.GetAllTasks).Method("GET",)

	retrun router
}