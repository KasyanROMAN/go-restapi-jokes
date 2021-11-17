package utils

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/KASYANRoman/go-restapi-jokes/pkg/models"
)

func ParseJSON(path string, list *[]models.Joke) {
	file, _ := os.Open(path)
	decoder := json.NewDecoder(file)

	err := decoder.Decode(&list)
	if err != nil {
		fmt.Println("Decode error")
	}
}
