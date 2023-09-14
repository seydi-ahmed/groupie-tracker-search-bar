package src

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func getDataFromAPI(apiURL string, result interface{}) error {
	// Effectuer une requête HTTP GET pour récupérer les données.
	response, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Lire le corps de la réponse en tant que données JSON.
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	// Analyser les données JSON dans la structure correspondante.
	err = json.Unmarshal(data, result)
	if err != nil {
		return err
	}

	return nil
}
func getRelationDataAPI(apiURL string, result interface{}) error {
	// Effectuer une requête HTTP GET pour récupérer les données.
	response, err := http.Get(apiURL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	// Lire le corps de la réponse en tant que données JSON.
	err = json.NewDecoder(response.Body).Decode(result)
	if err != nil {
		return err
	}

	return nil
}
