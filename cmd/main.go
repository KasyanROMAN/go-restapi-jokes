package main

import (
	"log"
	"net/http"

	"github.com/KASYANRoman/go-restapi-jokes/pkg/controllers"
	"github.com/KASYANRoman/go-restapi-jokes/pkg/models"
	utils "github.com/KASYANRoman/go-restapi-jokes/pkg/untils"

	//"github.com/gofiber/fiber/v2"
	"github.com/gorilla/mux"
)

func main() {
	var jokes []models.Joke
	utils.ParseJSON("reddit_jokes.json", &jokes)

	r := mux.NewRouter()
	r.HandleFunc("/Jokes", controllers.GetJoks).Methods("GET")
	r.HandleFunc("/Jokes/Search/{id}", controllers.SearchJoke).Methods("GET")
	r.HandleFunc("/Jokes/Add", controllers.CreateJoke).Methods("POST")
	r.HandleFunc("/Jokes/UpdataJoke/{id}", controllers.UpdateJoke).Methods("PUT")
	r.HandleFunc("/Jokes/Delete/{id}", controllers.DeleteJoke).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8001", r))
}

/*func Setup(app *fiber.App) {

	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

}
*/
