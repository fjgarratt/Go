package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", RootIndexHandler)
	router.HandleFunc("/ads", ListClassifiedAdsHandler)
	router.HandleFunc("/ad/{adId}", ViewClassifiedAdHandler)

	port := "8080"

	log.Println("Serving on port:", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func RootIndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Le Good Spot - Classified Ads!\n")
}

func ListClassifiedAdsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ClassifiedAd list\n")
}

func ViewClassifiedAdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	adId := vars["adId"]
	fmt.Fprintf(w, "ClassifiedAd view: %s\n", adId)
}
