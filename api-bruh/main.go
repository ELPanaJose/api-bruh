package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Meme struct {
	Title string `json:"Title"`
	Link  string `json:"link"`
}

type Memes []Meme

func allMemes(w http.ResponseWriter, r *http.Request) {
	for i := 0; i <= 1; i++ {
		incrementar()
		fmt.Println("request number :  ", contador-1)
	}
	data, err := ioutil.ReadFile("./memes.json")
	if err != nil {
		fmt.Print(err)
	}

	var memes Memes
	err = json.Unmarshal(data, &memes)
	if err != nil {
		fmt.Print(err)
	}

	json.NewEncoder(w).Encode(memes)
}

var contador = 0

func homePage(w http.ResponseWriter, f *http.Request) {
	fmt.Fprintf(w, "just for fun, a little api with memes👌, api good, go to api-bruh.elpanajose.repl.co/memes")
}

func incrementar() {
	contador++
}

func setUpRoutes() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/memes", allMemes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	fmt.Println("server on port 8080")
	setUpRoutes()
}
