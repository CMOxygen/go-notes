package main

import (
	"fmt"
	"go-notes-webapp/main-module/dbmanager"
	"go-notes-webapp/main-module/filemanager"
	"go-notes-webapp/main-module/handlers"
	"log"
	"net/http"
	"strings"
	"time"
)

func main() {
	fmt.Println("Running server...")

	//r := mux.NewRouter()
	//r.HandleFunc("/", handleLoginGet).Methods("GET")
	//r.Handle("/", http.FileServer(http.Dir("../view/"))).Methods("GET")
	//r.HandleFunc("/", handleLoginPost).Methods("POST")

	url, err := filemanager.ReadFile("/etc/server/c/r")
	if err != nil {
		panic(err)
	}

	err = handlers.DBM.Connect(strings.Trim(string(url), "\n"))
	if err != nil {
		panic(err)
	}
	defer func(DBM *dbmanager.DatabaseManager) {
		err := DBM.Close()
		if err != nil {
			panic(err)
		}
	}(&handlers.DBM)

	//var usrs []go_user.User
	//var notes []go_note.Note
	//
	//err = dbm.Select(&usrs, "", 1)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("SELECT USER ", usrs)
	//err = dbm.Select(&usrs, "username = ?", 123)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("SELECT USER WITH PARAMS ", usrs)
	//
	//err = dbm.Select(&notes, "")
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("NOTES = ", notes)

	//insertUser := go_user.User{Username: "go_username2", Password: "go_pass2"}
	//err = dbm.Insert(&insertUser)
	//if err != nil {
	//	panic(err)
	//}

	http.Handle("/", http.FileServer(http.Dir("../view/")))
	//http.Handle("/home", http.FileServer(http.Dir("../view/home/")))

	http.HandleFunc("/login", handlers.HandleLoginRequest)
	http.HandleFunc("/reg", handlers.HandleRegRequest)
	http.HandleFunc("/create-note", handlers.HandleCreateNote)
	http.HandleFunc("/remove-note", handlers.HandleRemoveNote)
	http.HandleFunc("/edit-note", handlers.HandleEditNote)

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
