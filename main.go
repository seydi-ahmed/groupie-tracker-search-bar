package main

import (
	"fmt"
	"groupie/src"
	"log"
	"net/http"
)

func main() {

	a, _ := src.FetchLocations()
	fmt.Println(a)

	http.HandleFunc("/", src.Home)
	http.HandleFunc("/artists/", src.ArtistPage)
	http.HandleFunc("/dates", src.HandleDates)

	// Redirection permanente depuis "/locations/" vers "/locations"
	http.HandleFunc("/locations/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/locations", http.StatusMovedPermanently)
	})

	http.HandleFunc("/locations", src.HandleLocations)

	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	fmt.Printf("Le serveur a démarré sous http://localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
