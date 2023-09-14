package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"text/template"
)

type Data struct {
	DETAIL string
}

func Home(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "ERREUR 405", "Méthode non autorisée")
		return
	}
	if r.URL.Path == "/" {
		posts, err := GetArtists("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			ErrorPage(w, r, http.StatusInternalServerError, "ERREUR 500", "Erreur interne du serveur")
			return
		}
		//ici un template dont le serveur va s'exécuter est déclaré
		t, err := template.ParseFiles("./templates/index.html")
		if err != nil {
			ErrorPage(w, r, http.StatusInternalServerError, "ERREUR 500", "Erreur interne du serveur")
			return
		}
		t.Execute(w, posts)
	} else {
		ErrorPage(w, r, http.StatusNotFound, "ERREUR 404", "Page non trouvée")
		return
	}
}

func ArtistPage(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "ERREUR 405", "Méthode non autorisée")
		return
	}
	// Get the artist ID from the URL
	artistID := r.URL.Path[len("/artists/"):]
	// Get the artist from the database
	artist, err := DetailArt(artistID)
	if err != nil {
		ErrorPage(w, r, http.StatusNotFound, "ERREUR 404", "Page non trouvée")
		return
	}
	// If the artist is not found, return a 404 error
	if artist == nil {
		ErrorPage(w, r, http.StatusNotFound, "ERREUR 404", "Page non trouvée")
		return
	}
	// Render the artist page with the artist's information
	tmpl, err := template.ParseFiles("templates/artists.html")
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "ERREUR 500", "Erreur interne du serveur")
		return
	}
	err = tmpl.Execute(w, artist)
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "ERREUR 500", "Erreur interne du serveur")
		return
	}
}

func ErrorPage(w http.ResponseWriter, r *http.Request, status int, title, message string) {
	w.WriteHeader(status)
	tmpl, err := template.ParseFiles("templates/error.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := struct {
		Title   string
		Message string
	}{
		Title:   title,
		Message: message,
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//************************************************************************
//************************************************************************

// // Fonction pour gérer la page des dates de concert
// func HandleDates(w http.ResponseWriter, r *http.Request) {
// 	dateToArtistsMap, err := GetDateToIDMap()
// 	if err != nil {
// 		http.Error(w, "Erreur lors de la récupération des dates de concert", http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl := template.Must(template.ParseFiles("templates/dates.html"))
// 	tmpl.Execute(w, dateToArtistsMap)
// }

// // Function to handle the locations API
// func HandleLocations(w http.ResponseWriter, r *http.Request) {
// 	locationToIDsMap, err := GetLocationToIDMap()
// 	if err != nil {
// 		http.Error(w, "Erreur lors de la récupération des emplacements", http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl := template.Must(template.ParseFiles("templates/locations.html"))
// 	tmpl.Execute(w, locationToIDsMap)
// }

// Créez une variable globale pour stocker les données des dates de concert en cache
var cachedDatesData map[string][]string

// Créez une variable globale pour stocker les données des emplacements en cache
var cachedLocationsData map[string][]string

// Fonction pour gérer la page des dates de concert
func HandleDates(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "ERREUR 405", "Méthode non autorisée")
		return
	}

	// Vérifier si les données sont en cache
	if cachedDatesData != nil {
		tmpl := template.Must(template.ParseFiles("templates/dates.html"))
		tmpl.Execute(w, cachedDatesData)
		return
	}

	// Si les données ne sont pas en cache, récupérer les données de l'API
	dateToArtistsMap, err := GetDateToIDMap()
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "ERREUR 500", "Erreur lors de la récupération des dates de concert")
		return
	}

	// Stocker les données en cache
	cachedDatesData = dateToArtistsMap

	tmpl := template.Must(template.ParseFiles("templates/dates.html"))
	tmpl.Execute(w, dateToArtistsMap)
}

// Fonction pour gérer la page des emplacements
func HandleLocations(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorPage(w, r, http.StatusMethodNotAllowed, "ERREUR 405", "Méthode non autorisée")
		return
	}

	// Vérifier si les données sont en cache
	if cachedLocationsData != nil {
		tmpl := template.Must(template.ParseFiles("templates/locations.html"))
		tmpl.Execute(w, cachedLocationsData)
		return
	}

	// Si les données ne sont pas en cache, récupérer les données de l'API
	locationToIDsMap, err := GetLocationToIDMap()
	if err != nil {
		ErrorPage(w, r, http.StatusInternalServerError, "ERREUR 500", "Erreur lors de la récupération des emplacements")
		return
	}

	// Stocker les données en cache
	cachedLocationsData = locationToIDsMap

	tmpl := template.Must(template.ParseFiles("templates/locations.html"))
	tmpl.Execute(w, locationToIDsMap)
}

//************************************************************************
//************************************************************************

func GetArtistNameByID(artistID int) string {
	// Get all artists
	artists, err := GetArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return ""
	}

	// Find the artist with the matching ID
	for _, artist := range artists {
		if artist.ID == artistID {
			return artist.Name
		}
	}

	return ""
}

// Fonction pour récupérer les données de l'API et les stocker dans un map associant les dates aux noms d'artistes
func GetDateToIDMap() (map[string][]string, error) {
	// Faites une requête HTTP GET pour récupérer les données de l'API
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Vérifiez si la requête a réussi (code de statut 200)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("La requête a échoué avec le code de statut %d", response.StatusCode)
	}

	// Décodez la réponse JSON dans l'objet ConcertDateResponse
	var concertDateResponse ConcertDateResponse
	err = json.NewDecoder(response.Body).Decode(&concertDateResponse)
	if err != nil {
		return nil, err
	}

	// Créer un map pour stocker les dates associées à leurs noms d'artistes
	dateToArtistsMap := make(map[string][]string)

	// Parcourir les dates de concert et les ajouter au map
	for _, concertDate := range concertDateResponse.Dates {
		date := concertDate.Dates[0:]
		// Ajouter le nom de l'artiste associé à la date dans le map
		artistName := GetArtistNameByID(concertDate.ID)
		if artistName != "" {
			dateToArtistsMap[date[0]] = append(dateToArtistsMap[date[0]], artistName)
		}

	}

	return dateToArtistsMap, nil
}

// Fonction pour récupérer les données de l'API et les stocker dans un map associant les locations aux noms d'artistes
func GetLocationToIDMap() (map[string][]string, error) {
	// Obtenez les données de l'API à partir de la fonction GetLocations
	locationData, err := GetLocations()
	if err != nil {
		return nil, err
	}

	// Créer un map pour stocker les locations associées à leurs noms d'artistes
	locationToNameMap := make(map[string][]string)

	// Parcourir les données de location et les ajouter au map
	for _, location := range locationData {
		locationName := location.Locations[0:]
		// Ajouter le nom de l'artiste associé à la location dans le map
		artistName := GetArtistNameByID(location.ID)
		if artistName != "" {
			locationToNameMap[locationName[0]] = append(locationToNameMap[locationName[0]], artistName)
		}

	}
	return locationToNameMap, nil
}
