package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

var keyLocation = "key.txt"

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "<a href=\"https://github.com/Thor77/nginx-rtmp-keyauth\">nginx-rtmp-keyauth</a>")
}

func authHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "", http.StatusNotImplemented)
		return
	}
	keyfile, err := ioutil.ReadFile(keyLocation)
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		log.Print(err)
		return
	}
	key := strings.TrimRight(strings.TrimRight(string(keyfile), "\n"), "\r")
	if givenKey := r.FormValue("key"); givenKey != key {
		http.Error(w, "", http.StatusForbidden)
	}

}

func main() {
	if len(os.Args) >= 2 {
		keyLocation = os.Args[1]
	}
	r := mux.NewRouter()
	r.HandleFunc("/", indexHandler)
	r.HandleFunc("/auth", authHandler)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
