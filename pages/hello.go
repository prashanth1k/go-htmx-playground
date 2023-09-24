package pages

import (
	"io"
	"net/http"
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World")
}
