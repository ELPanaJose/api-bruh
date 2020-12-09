package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Meme struct {
	Tittle string `json:"Tittle"`
	Link string `json:"link"`
	
}

type Memes []Meme

func allMemes(w http.ResponseWriter, r *http.Request){
    memes := Memes{
    	Meme{Tittle:"bruh que buen momo", Link:"https://images3.memedroid.com/images/UPLOADED230/5ac47bad8c3da.jpeg"},
    	{Tittle:"eso brad kkkk", Link:"https://images3.memedroid.com/images/UPLOADED230/5ac47bad8c3da.jpeg"},
    	{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},{Tittle:"", Link:""},

    }

	fmt.Println("all memes here")
	json.NewEncoder(w).Encode(memes)
}

func homePage(w http.ResponseWriter, f *http.Request){
	fmt.Fprintf(w, "BRUH")
}

func setUpRoutes(){
	http.HandleFunc("/", homePage)
	http.HandleFunc("/memes", allMemes)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main(){
	fmt.Println("BRUH")
    setUpRoutes()
}