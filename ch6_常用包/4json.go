package main

import (
	"log"
	"fmt"
	"encoding/json"
)

type Movie struct {
	Title  string
	Year   int `json:"year"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942,
		Actors: []string{"Hum", "In"}},
	{Title: "Cool", Year: 1967,
		Actors: []string{"asdf", "asdf"}},
}

func main() {
	//普通json
	data,err:=json.Marshal(movies)
	if err != nil {
		log.Fatalf("JSON marshaling:%s",err)
	}
	fmt.Printf("%s\n",data)

	//格式化json
	data,err=json.MarshalIndent(movies,"","		")
	if err != nil {
		log.Fatalf("JSON marshaling:%s",err)
	}
	fmt.Printf("%s\n",data)

	//解析json
	var titles []Movie
	json.Unmarshal(data,&titles)
	fmt.Println(titles)

	var titles2 []struct{Title string }
	json.Unmarshal(data,&titles2)
	fmt.Println(titles2)
}
