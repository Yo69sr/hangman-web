package main

import (
	"hangman"
	"html/template"
	"net/http"
	"strings"
)

type Game struct {
	Level    string
	Historic string
}

var currentGame = Game{}
var GameData = hangman.DataHang{}

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("_templates_/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		panic(err)
	}
}

func hangmant(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	if r.URL.Query().Get("level") != "" {
		GameData = hangman.Play(r.URL.Query().Get("level"))
		currentGame.Level = r.URL.Query().Get("level")
		currentGame.Historic = ""
	}

	input := r.Form.Get("letter")
	if input != "" {
		if len(input) > 1 {
			if strings.EqualFold(input, GameData.SearchedW) {
				GameData.Word = GameData.SearchedW
			} else {
				GameData.Pv -= 2
			}
			currentGame.Historic += input + " "
		} else {
			if !strings.Contains(currentGame.Historic, input) {
				currentGame.Historic += input + " "
				if strings.Contains(GameData.SearchedW, input) {
					for i, c := range GameData.SearchedW {
						if strings.EqualFold(string(c), input) {
							GameData.Word = GameData.Word[:i] + input + GameData.Word[i+1:]
						}
					}
				} else {
					GameData.Pv--
				}
			}
		}

		if GameData.Pv <= 0 {
			currentGame.Historic += "(Défaite)"
		} else if !strings.Contains(GameData.Word, "_") {
			currentGame.Historic += "(Victoire)"
		}
	}

	tmpl := template.Must(template.ParseFiles("_templates_/page-2.html"))
	i := struct {
		Level    string
		Historic string
		Pv       int
		Word     string
	}{
		Level:    currentGame.Level,
		Historic: currentGame.Historic,
		Pv:       GameData.Pv,
		Word:     GameData.Word,
	}
	err := tmpl.Execute(w, i)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.Handle("/css/", http.StripPrefix("/css", http.FileServer(http.Dir("_templates_/"))))
	http.HandleFunc("/", index)
	http.HandleFunc("/hangman", hangmant)
	http.ListenAndServe(":8080", nil)
}
