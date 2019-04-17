package handler

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"
)

// Trigger is our struct is it not
type Trigger struct {
	ID       int
	Name     string
	Category string
	Created  time.Time
}

var tmpl = template.Must(template.ParseGlob("form/*"))

// HomeHandler handlers the "/" endpoint
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	x := "wow"
	fmt.Fprintf(w, "Fantastic stuff hello %v\n", x)
}

// NewTrigger handles "/new"
func NewTrigger(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

// GetAllTriggers handlers the "/triggers" endpoint
func GetAllTriggers(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM `Trigger` ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	trg := Trigger{}
	res := []Trigger{}
	for selDB.Next() {
		var id int
		var name, category string
		err = selDB.Scan(&id, &name, &category)
		if err != nil {
			panic(err.Error())
		}
		trg.ID = id
		trg.Name = name
		trg.Category = category
		res = append(res, trg)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()

}

// InsertTrigger handlers the "/insert" endpoint
func InsertTrigger(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		name := r.FormValue("name")
		category := r.FormValue("category")
		insForm, err := db.Prepare("INSERT INTO `Trigger`(name, category) VALUES(?,?)")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(name, category)
		log.Println("INSERT: Name: " + name + " | Category: " + category)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "username"
	dbPassword := os.Getenv("MYSQL_TRIGGER_PASSWORD")
	dbName := "triggers"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPassword+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}
