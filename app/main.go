package main

import (
	"net/http"
	"os"

	"github.com/matt-paul/triggerRemedy/app/handler"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", handler.HomeHandler)
	http.HandleFunc("/new", handler.NewTrigger)
	http.HandleFunc("/triggers", handler.GetAllTriggers)
	http.HandleFunc("/insert", handler.InsertTrigger)

	http.ListenAndServe(":3001", nil)
}
