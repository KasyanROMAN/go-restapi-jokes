package models

type Joke struct {
	ID    string `json:"id"`
	Score int `json:"Score"`
	Title string `json:"title"`
	Body  string `json:"body"`
}
