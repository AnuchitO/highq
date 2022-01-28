package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/anuchito/ursho/config"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"response"`
}


var prefix string

func init() {
	c, err := config.ReadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if c.Options.Prefix == "" {
		prefix = ""
	} else {
		prefix = c.Options.Prefix
	}
}


func createResponse(w http.ResponseWriter, r Response) {
	d, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	w.Write(d)
}
