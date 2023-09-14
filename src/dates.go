package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Structure pour stocker l'objet JSON renvoyé par l'API
type ConcertDateResponse struct {
	Dates []ConcertDate `json:"index"`
}

type ConcertDate struct {
	ID      int      `json:"id"`
	Dates   []string `json:"dates"`
	Artists []string `json:"artists"` // Ajouter un champ pour stocker les noms des artistes
}

// Fonction pour récupérer les données de l'API et les stocker dans un tableau de ConcertDate
func GetDates() ([]ConcertDate, error) {
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

	// Obtenir tous les artistes pour mapper leurs noms par ID
	artists, err := GetArtists("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		return nil, err
	}

	// Créer un map pour stocker les noms d'artistes par ID
	artistIDToNameMap := make(map[int]string)
	for _, artist := range artists {
		artistIDToNameMap[artist.ID] = artist.Name
	}

	// Mettre à jour les ConcertDates avec les noms d'artistes correspondants
	for i, concertDate := range concertDateResponse.Dates {
		for _, artistID := range concertDate.Artists {
			id, err := strconv.Atoi(artistID)
			if err != nil {
				return nil, fmt.Errorf("erreur lors de la conversion de l'ID d'artiste en entier: %v", err)
			}
			if name, ok := artistIDToNameMap[id]; ok {
				concertDateResponse.Dates[i].Artists = append(concertDateResponse.Dates[i].Artists, name)
			}
		}
	}

	return concertDateResponse.Dates, nil
}
