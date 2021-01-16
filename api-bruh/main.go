package main

import(
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Meme struct {
	Title string `json:"Title"`
	Link string `json:"link"`
	
}

type Memes []Meme

func allMemes(w http.ResponseWriter, r *http.Request){
    memes := Memes{
    	Meme{Title:"bruh que buen momo", Link:"https://i.pinimg.com/originals/a7/ce/0d/a7ce0d33f4de01f6b81f9d5dc89828e8.jpg"},


      {Title:"que monda", Link:"https://images3.memedroid.com/images/UPLOADED115/6001f7b90e8df.jpeg"},


      {Title:"Pa casita", Link:"https://images7.memedroid.com/images/UPLOADED528/6001d05f63cca.jpeg"},


      {Title:"LOL", Link:"https://images7.memedroid.com/images/UPLOADED959/600213661813c.jpeg"},


      {Title:"LMAO", Link:"https://i.pinimg.com/236x/e7/22/1c/e7221ce620389f8b4abe03cb8c5866ce.jpg"},


      {Title:"BRUH", Link:"https://i.pinimg.com/236x/c3/3d/82/c33d825bfde06a3b0368d00943c02c1c.jpg"},


      {Title:"No gracias ya comi", Link:"https://images3.memedroid.com/images/UPLOADED78/6002c2fb8951c.jpeg"},


      {Title:"Asies", Link:"https://images3.memedroid.com/images/UPLOADED16/6002714b8b6ba.jpeg"},


      {Title:"Que hermoso aww", Link:"https://images7.memedroid.com/images/UPLOADED529/60027b6498af8.jpeg"},

      {Title: "Es una serie coreana de mierda de dos Larvas pedorreandose, eructando, peleando y cogiendo. Puedes ver esa mierda en Netflix", Link:"https://images7.memedroid.com/images/UPLOADED619/6001efb67699e.jpeg"},

      {Title: "Muy random todo", Link:"https://images7.memedroid.com/images/UPLOADED701/6001aea562772.jpeg"},

      {Title: "Y si.", Link:"https://images3.memedroid.com/images/UPLOADED849/6000fe4594b77.jpeg"},

      {Title: "dame un chin", Link:"https://images3.memedroid.com/images/UPLOADED121/6000f899a4309.jpeg"},

      {Title: "si", Link:"https://images7.memedroid.com/images/UPLOADED211/6000a5d59f93d.jpeg"},

      {Title: "VAMO A CURARTE :D", Link:"https://images3.memedroid.com/images/UPLOADED109/6000a13790548.jpeg"},

      {Title: "XD", Link:"https://images3.memedroid.com/images/UPLOADED202/6000a5f201f1b.jpeg"},

      {Title: "Humor Negro bruh", Link:"https://images3.memedroid.com/images/UPLOADED668/6000a19a10e2b.jpeg"},

      {Title: "Yo perdí un familiar hace meses y aun duele", Link:"https://images3.memedroid.com/images/UPLOADED350/600008c582273.jpeg"},

      {Title: "gotica culona ", Link:"https://i.redd.it/84wd7zwajpb61.jpg"},

      {Title: "", Link:"https://preview.redd.it/w49am7itumb61.jpg?width=640&height=606&crop=smart&auto=webp&s=e9d644a6acada2b8c47049f854646f7bd179acef"},

      {Title: "No te voy a mentir, me rei ._.", Link:"https://preview.redd.it/c7e1525tbkb61.jpg?width=640&height=235&crop=smart&auto=webp&s=23b91fc4b11ba61a40a1e3d3f3b476c05e095dc9"},

      {Title: "10010", Link:"https://i.redd.it/ilz5decs5nb61.jpg"},

      {Title: "Me interesa", Link:"https://preview.redd.it/tuvqkdcn1qb61.jpg?width=640&height=640&crop=smart&auto=webp&s=ce401e1791cf939245cd91e9e5c5f0194f4c8a48"},

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
