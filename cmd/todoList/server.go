package main

import (
	"fmt"
	"net/http"

	"github.com/advaittrivedi1122/todolist/database"
	"github.com/advaittrivedi1122/todolist/env"
	"github.com/advaittrivedi1122/todolist/internal/router"
)

func main() {
	env := env.GetEnv()

	// Initialise Database
	database.Initialise()

	// Initialise server
	router := router.NewRouter()
	port := env.PORT
	addr := fmt.Sprintf(":%d", port)
	fmt.Printf("Server listening on http://localhost%s\n", addr)

	err := http.ListenAndServe(addr,router)
	if err != nil {
		panic(err)
	}

}