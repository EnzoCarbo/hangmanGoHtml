package main

import "hangmanweb/hangman"

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

func main() {

	hangman.Start()
	hangman.Player.CheckInput("guess")
	/*
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
	   		if logs.lvl == "1" {
	   			hangman.Start("......")
	   			http.Redirect(w, r, "/game", 301)
	   		}

	   	})
	   	type PageGame struct {
	   	}
	   	http.HandleFunc("/game", func(w http.ResponseWriter, r *http.Request) {

	   		temp.ExecuteTemplate(w, "easy", logs)
	   	})

	   	http.HandleFunc("/game/treatment", func(w http.ResponseWriter, r *http.Request) {
	   	})

	   	rootDoc, _ := os.Getwd()
	   	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	   	http.Handle("/static/", http.StripPrefix("/static/", fileserver))
	   	//Init serv
	   	http.ListenAndServe("localhost:8080", nil)
	   }
	*/

}
