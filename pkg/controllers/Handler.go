package controllers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	//"./pkg/controllers"
	"github.com/KASYANRoman/go-restapi-jokes/pkg/models"

	//"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

var jokes []models.Joke

func GetJoks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(jokes)
}

func SearchJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range jokes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&models.Joke{})
}

func CreateJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book models.Joke
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	jokes = append(jokes, book)
	json.NewEncoder(w).Encode(book)

	rawDataOut, err := json.MarshalIndent(jokes, "", "   ")
	if err != nil {
		log.Fatal("JSON marshaling failed: ", err)
	}
	err = ioutil.WriteFile("reddit_jokes.json", rawDataOut, 0)
	if err != nil {
		log.Fatal("Cannot write:", err)
	}
}

func UpdateJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range jokes {
		if item.ID == params["id"] {
			jokes = append(jokes[:index], jokes[index+1:]...)
			var book models.Joke
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			jokes = append(jokes, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func DeleteJoke(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range jokes {
		if item.ID == params["id"] {
			jokes = append(jokes[:index], jokes[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(jokes)
}
