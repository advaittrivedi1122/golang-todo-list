package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/advaittrivedi1122/todolist/database"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	session := database.Session
	var clusterName string
	if err := session.Query(`SELECT cluster_name FROM system.local`).Scan(&clusterName); err != nil {
		log.Fatalf("unable to query system.local: %v", err)
	}
	fmt.Printf("Connected to cluster: %s\n", clusterName)
	w.Write([]byte("Welcome to Todos List!!"))
}
