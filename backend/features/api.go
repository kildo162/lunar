package features

import (
	"log"
	"net/http"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	log.Println("Received API request: /api")
	w.Write([]byte("Hello, World!"))
}
