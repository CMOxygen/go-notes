package handlers

import (
	"encoding/json"
	"fmt"
	"go-notes-webapp/main-module/dbmanager"
	"go-notes-webapp/main-module/encryption"
	"go-notes-webapp/main-module/go_note"
	"go-notes-webapp/main-module/go_user"
	"go-notes-webapp/main-module/sessions"
	"html"
	"net/http"
)

var DBM dbmanager.DatabaseManager
var sm sessions.SessionManager

func HandleLoginRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(html.EscapeString(r.URL.Path), r.Method)

	switch r.Method {
	case "GET":

	case "POST":

		fmt.Println("LOGIN REQUEST POST")
		var requestUser go_user.UserStringPass
		fmt.Println("REQUEST = ", r.Body)

		err := json.NewDecoder(r.Body).Decode(&requestUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			panic(err)
		}

		fmt.Println(requestUser)
		responseUser := go_user.User{}

		pass, err := encryption.EncryptSHA([]byte(requestUser.Password))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		err = DBM.Select(&responseUser, "username=? AND password=?", requestUser.Username, pass)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			panic(err)
		}

		if responseUser.ID <= 0 {
			http.Error(w, "no user with such username and password", http.StatusNotFound)
			fmt.Printf("NO USER FOUND. request user = %v \n response user = %v", requestUser, responseUser)
		} else {
			err = DBM.Select(&responseUser.Notes, "userId=?", responseUser.ID)
			resp, err := json.Marshal(responseUser)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}
			c, err := r.Cookie("sessionID")
			if err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				//panic(err)
				fmt.Println(err.Error())
			}
			if c == nil {

				sm.CreateSession(responseUser)
				sessionCookie := http.Cookie{Name: "sessionID", Value: sm.Sessions[len(sm.Sessions)-1].SessionID}
				http.SetCookie(w, &sessionCookie)

			} else if sm.SessionExists(c.Value) {

			}
			fmt.Println(sm.Sessions)

			_, err = w.Write(resp)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				panic(err)
			}
		}
	}
}

func HandleRegRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Println(html.EscapeString(r.URL.Path), r.Method)

	switch r.Method {
	case "GET":

	case "POST":
		fmt.Println("REGISTRATION REQUEST POST")

		var requestUser go_user.UserStringPass
		fmt.Printf("REQUEST = %v\n", r.Body)

		err := json.NewDecoder(r.Body).Decode(&requestUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		fmt.Printf("REQUEST USER %v\n", requestUser)

		p, err := encryption.EncryptSHA([]byte(requestUser.Password))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		responseUser := go_user.User{Username: requestUser.Username, Password: p}

		err = DBM.Insert(&responseUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		err = DBM.Select(&responseUser, "")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		fmt.Printf("SELECT USER FROM DB = %v\n", responseUser)

		resp, err := json.Marshal(responseUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	}
}

func HandleCreateNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println(html.EscapeString(r.URL.Path), r.Method)

	switch r.Method {

	case "POST":
		fmt.Println("CREATE NOTE REQUEST POST")
		fmt.Printf("REQUEST = %v\n", r.Body)

		var requestNote go_note.Note
		err := json.NewDecoder(r.Body).Decode(&requestNote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		err = DBM.Insert(&requestNote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		var responseUser go_user.User
		err = DBM.Select(&responseUser, "userId=?", requestNote.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		err = DBM.Select(&responseUser.Notes, "userId=?", requestNote.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		resp, err := json.Marshal(responseUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	}
}

func HandleRemoveNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println(html.EscapeString(r.URL.Path), r.Method)
	fmt.Println(r.Body)

	switch r.Method {

	case "POST":
		fmt.Println("HANDLE REMOVE NOTE POST")

		var requestNote go_note.Note
		err := json.NewDecoder(r.Body).Decode(&requestNote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		err = DBM.Delete(&requestNote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		var responseUser go_user.User
		err = DBM.Select(&responseUser, "userId=?", requestNote.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		err = DBM.Select(&responseUser.Notes, "userId=?", responseUser.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		resp, err := json.Marshal(responseUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	}
}

func HandleEditNote(w http.ResponseWriter, r *http.Request) {
	fmt.Println(html.EscapeString(r.URL.Path), r.Method)
	fmt.Println(r.Body)

	switch r.Method {
	case "POST":
		fmt.Println("HANDLE EDIT NOTE POST")

		var requestNote go_note.Note
		err := json.NewDecoder(r.Body).Decode(&requestNote)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		var noteToUpdate go_note.Note
		err = DBM.Select(&noteToUpdate, "noteId=?", requestNote.NoteID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		err = DBM.Update(&noteToUpdate, "noteTitle=?, noteText=?", requestNote.NoteTitle, requestNote.NoteText)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		var responseUser go_user.User
		err = DBM.Select(&responseUser, "userId=?", requestNote.UserID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		err = DBM.Select(&responseUser.Notes, "userId=?", responseUser.ID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		resp, err := json.Marshal(responseUser)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}

		_, err = w.Write(resp)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			panic(err)
		}
	}
}
