package features

import (
	"net/http"
)

func HandleAPI(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World!"))
}
