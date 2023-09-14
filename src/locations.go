package src

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"text/template"
)

// Structure pour stocker l'objet JSON renvoyé par l'API
type LocationResponse struct {
	Locations []LocationDATA `json:"index"`
}

type LocationDATA struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
	DatesURL  string   `json:"dates"`
}

// Fonction pour récupérer les données de l'API et les stocker dans un tableau de LocationDATA
func GetLocations() ([]LocationDATA, error) {
	// Faites une requête HTTP GET pour récupérer les données de l'API
	response, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Vérifiez si la requête a réussi (code de statut 200)
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("La requête a échoué avec le code de statut %d", response.StatusCode)
	}

	// Décodez la réponse JSON dans l'objet LocationResponse
	var locationResponse LocationResponse
	err = json.NewDecoder(response.Body).Decode(&locationResponse)
	if err != nil {
		return nil, err
	}

	return locationResponse.Locations, nil
}

// Fonction pour inverser les données de locations et les sauvegarder dans un fichier HTML
func ReverseLocationsAndSaveHTML(locations []LocationDATA, outputFilePath string) error {
	// Inverser les données de locations
	locationsMap := make(map[string][]int)
	for _, locationData := range locations {
		for _, location := range locationData.Locations {
			locationsMap[location] = append(locationsMap[location], locationData.ID)
		}
	}

	// Charger le modèle à partir du fichier HTML
	tmpl, err := template.ParseFiles("templates/locations.html")
	if err != nil {
		return err
	}

	// Ouvrir le fichier de sortie
	file, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// Exécuter le modèle et écrire le résultat dans le fichier HTML
	return tmpl.Execute(file, locationsMap)
}
