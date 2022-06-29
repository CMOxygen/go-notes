package main

import (
	"encoding/json"
	"fmt"
	"go-notes-webapp/main-module/dbmanager"
	"go-notes-webapp/main-module/go_note"
	"go-notes-webapp/main-module/go_user"
	"html"
	"log"
	"net/http"
	"time"
)

//type User struct {
//	ID       uint32 `json:"id"`
//	Username string `json:"username"`
//	Password string `json:"password"`
//}

//type Note struct {
//	NoteID    uint32 `json:"noteID"`
//	UserID    uint32 `json:"userID"`
//	NoteTitle string `json:"noteTitle"`
//	NoteText  string `json:"noteText"`
//}

func processLoginRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(html.EscapeString(r.URL.Path), r.Method)

	switch r.Method {
	case "GET":

	case "POST":
		fmt.Println("LOGIN REQUEST POST")
		var usr go_user.User
		fmt.Println("REQUEST = ", r.Body)
		err := json.NewDecoder(r.Body).Decode(&usr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err)
		}
		fmt.Println(usr)
	}
}

func handleLoginPost(w http.ResponseWriter, r *http.Request) {
	//var usr go_user.User
	var usr go_user.User
	err := json.NewDecoder(r.Body).Decode(&usr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println("usr", usr)

	out, err := json.Marshal(usr)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	_, err = w.Write(out)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}

func main() {
	fmt.Println("Running server...")

	//r := mux.NewRouter()
	//r.HandleFunc("/", handleLoginGet).Methods("GET")
	//r.Handle("/", http.FileServer(http.Dir("../view/"))).Methods("GET")
	//r.HandleFunc("/", handleLoginPost).Methods("POST")

	var dbm dbmanager.DatabaseManager
	err := dbm.Connect("go_notes:AlSkDjFhG_2@/go_notes")
	if err != nil {
		panic(err)
	}

	defer func(dbm *dbmanager.DatabaseManager) {
		err := dbm.Close()
		if err != nil {
			panic(err)
		}
	}(&dbm)

	var usrs []go_user.User
	var notes []go_note.Note

	err = dbm.Select(&usrs, "", 1)
	if err != nil {
		panic(err)
	}
	fmt.Println("SELECT USER ", usrs)
	err = dbm.Select(&usrs, "username = ?", 123)
	if err != nil {
		panic(err)
	}
	fmt.Println("SELECT USER WITH PARAMS ", usrs)

	err = dbm.Select(&notes, "")
	if err != nil {
		panic(err)
	}
	fmt.Println("NOTES = ", notes)

	//insertUser := go_user.User{Username: "go_username2", Password: "go_pass2"}
	//err = dbm.Insert(&insertUser)
	//if err != nil {
	//	panic(err)
	//}

	http.Handle("/", http.FileServer(http.Dir("../view/")))
	http.HandleFunc("/login", processLoginRequest)

	s := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Fatal(s.ListenAndServe())
}
