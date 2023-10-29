package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Location2 struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}

type Artist2 struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ArtistLocation struct {
	ArtistID        int
	ArtistName      string
	ArtistLocations []string
}

func fetchLocations() ([]Location2, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var locations []Location2
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &locations); err != nil {
		return nil, err
	}

	return locations, nil
}

func fetchArtists() ([]Artist2, error) {
	resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var artists []Artist2
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(body, &artists); err != nil {
		return nil, err
	}

	return artists, nil
}

func createArtistLocations(locations []Location2, artists []Artist2) []ArtistLocation {
	artistLocations := make([]ArtistLocation, 0)

	for _, artist := range artists {
		var artistLocation ArtistLocation
		artistLocation.ArtistID = artist.ID
		artistLocation.ArtistName = artist.Name

		for _, location := range locations {
			if contains(location.ID, artist.Locations) {
				artistLocation.ArtistLocations = append(artistLocation.ArtistLocations, location.Locations...)
			}
		}

		artistLocations = append(artistLocations, artistLocation)
	}

	return artistLocations
}

func contains(id int, ids []int) bool {
	for _, value := range ids {
		if value == id {
			return true
		}
	}
	return false
}

func Search() {
	// Récupérer les données depuis les API
	locations, err := fetchLocations()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des locations:", err)
		return
	}

	artists, err := fetchArtists()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des artistes:", err)
		return
	}

	// Créer la structure ArtistLocation
	artistLocations := createArtistLocations(locations, artists)

	// À ce stade, vous avez une liste de structures ArtistLocation qui contiennent "ArtistID", "ArtistName" et "ArtistLocations" liés.

	// Vous pouvez maintenant intégrer ces "locations" dans votre datalist HTML.
	// Cela dépendra de la manière dont vous souhaitez générer le HTML.
}
