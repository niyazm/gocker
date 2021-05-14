package main
import (
	"fmt"
	"net/http"
	"html"
	"log"
	)


func main() {
	http.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello world, from badonkadocker %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "you might be looking for /heartbeat %q", html.EscapeString(r.URL.Path))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
