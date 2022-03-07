package routers

import (
	"github.com/gorilla/mux"
	"github.com/iyashjayesh/Go-Lang-Resources/projects/ToDo_React_and_Go/server/middleware"
)

func Router() *mux.Router {

	router := mux.NewRouter() // create a new router

	//Get
	router.HandleFunc("/api/task", middleware.GetAllTasks).Methods("GET", "OPTIONS") // get all tasks

	//Post
	router.HandleFunc("/api/tasks", middleware.CreateTask).Methods("POST", "OPTIONS") // create a new task

	//Put
	router.HandleFunc("/api/tasks/{id}", middleware.TaskComplete).Methods("PUT", "OPTIONS") // mark task as complete
	router.HandleFunc("/api/undoTask/{id}", middleware.UndoTask).Methods("PUT", "OPTIONS")  // Undo the deleted task

	//Delete
	router.HandleFunc("/api/deleteTask/{id}", middleware.DeleteTask).Methods("DELETE", "OPTIONS")                      // delete a task
	router.HandleFunc("/api/deleteAllTasks", middleware.DeleteAllTasks).Methods("DELETE", "OPTIONS")                   // delete all tasks
	router.HandleFunc("/api/deleteAllCompletedTasks", middleware.DeleteAllCompletedTasks).Methods("DELETE", "OPTIONS") // delete all completed tasks

	return router
}
