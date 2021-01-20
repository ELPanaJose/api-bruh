package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

type Meme struct {
	Title string `json:"Title"`
	Link string `json:"link"`
}

type Memes []Meme

func allMemes(w http.ResponseWriter, r *http.Request){
    data, err := ioutil.ReadFile("./memes.json")
    if err != nil {
        fmt.Print(err)
    }

    var memes Memes
    err = json.Unmarshal(data, &memes)
    if err != nil {
        fmt.Print(err)
    }

	fmt.Println("all memes here")
	json.NewEncoder(w).Encode(memes)
}

func homePage(w http.ResponseWriter, f *http.Request){
	fmt.Fprintf(w, "just for fun, a little api with memes👌, api good, go to api-bruh.elpanajose.repl.co/memes")
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
