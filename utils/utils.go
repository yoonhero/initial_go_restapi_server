package utils

import (
	"log"
	"net/http"
)

func HandleErr(err error) {
	if err != nil {
		log.Panic(err)
	}
}

func AllowConnection(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}
