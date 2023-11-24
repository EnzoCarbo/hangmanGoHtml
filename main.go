package main

import (
	"fmt"
	"hangmanweb/hangman"
	"html/template"
	"net/http"
	"os"
)

type PageInit struct {
	Username string
	lvl      string
}

type Game struct {
	State        string
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnsLeft    int
}

var logs PageInit
var MesUser string

func main() {
	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Printf(fmt.Sprintf("ERREUR => %s", err.Error()))
		return
	}

	http.HandleFunc("/intro", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "intro", nil)
	})

	http.HandleFunc("/whatsapp", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "whatsapp", nil)
	})

	http.HandleFunc("/init", func(w http.ResponseWriter, r *http.Request) {
		temp.ExecuteTemplate(w, "username", nil)

	})

	http.HandleFunc("/init/treatment", func(w http.ResponseWriter, r *http.Request) {
		logs = PageInit{r.FormValue("pseudo"), r.FormValue("lvl")}
		fmt.Println(logs)
		hangman.Start(logs.lvl)
		http.Redirect(w, r, "/game", 301)
	})
	type PageGame struct {
		Hiddenword []string
		Listletter []string
		Leftpv     int
		MesUser    string
	}

	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

		data := PageGame{hangman.Player.FoundLetters, hangman.Player.UsedLetters, hangman.Player.TurnsLeft, MesUser}
		if hangman.HasWon(hangman.Player.FoundLetters, hangman.Player.Word) || hangman.Player.TurnsLeft <= 0 {
			http.Redirect(w, r, "/end", 301)
		}
		temp.ExecuteTemplate(w, "easy", data)
	})

	http.HandleFunc("/game/treatment", func(w http.ResponseWriter, r *http.Request) {
		value := r.FormValue("value")
		MesUser = hangman.Player.CheckInput(value)
		http.Redirect(w, r, "/game", 301)
	})

	http.HandleFunc("/end", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("end"))
	})

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	//Init serv
	http.ListenAndServe("localhost:8080", nil)
}
