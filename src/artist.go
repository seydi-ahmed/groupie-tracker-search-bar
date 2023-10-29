package src

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

func GetArtists(srt string) ([]Post2, error) {
	// Lire le contenu du fichier JSON
	dat, err := http.Get(srt)
	if err != nil {
		// log.Fal(err)
		return nil, err
	}
	data, err := io.ReadAll(dat.Body)

	if err != nil {
		// log.Fal(err)
		return nil, err
	}

	// Parser le fichier JSON
	var tempArtists []TempArtist
	err = json.Unmarshal(data, &tempArtists)
	if err != nil {
		// log.Fatal(err)
		return nil, err
	}

	// Convertir la structure temporaire en la structure Artist
	var artists []Post2
	for _, temp := range tempArtists {
		artist := Post2(temp) // Convert TempArtist to Artist
		artists = append(artists, artist)
	}

	return artists, nil
}

func DetailArt(id string) (*Post, error) {
	dat := "https://groupietrackers.herokuapp.com/api/artists"
	dat3 := "https://groupietrackers.herokuapp.com/api/relation"

	// Parse the ID string to an integer
	artistID, err := strconv.Atoi(id)
	if err != nil {
		// Return an error if parsing the ID fails
		return nil, fmt.Errorf("unable to parse artist ID: %v", err)
	}

	// Get all artists
	artists, err := GetArtists(dat)
	if err != nil {
		return nil, fmt.Errorf("impossible d'aller chercher des artistes: %v", err)
	}

	// Find the artist with the matching ID
	for _, artist := range artists {
		if artist.ID == artistID {
			artistID2 := strconv.Itoa(artist.ID)

			relationString := dat3 + "/" + artistID2
			var RELLOCDAT LocationData

			err := getRelationDataAPI(relationString, &RELLOCDAT)
			if err != nil {
				return nil, fmt.Errorf("erreur lors de la récupération des données: %v", err)
			}
			var Detail Post
			Detail.ID = artist.ID
			Detail.Name = artist.Name
			Detail.ImageURL = artist.ImageURL
			Detail.Members = artist.Members
			Detail.FirstAlbum = artist.FirstAlbum
			Detail.YearCreation = artist.YearCreation
			Detail.DatesLocations = RELLOCDAT.DatesLocations

			return &Detail, nil
		}
	}
	return nil, fmt.Errorf("artiste non trouvé pour l'ID %s", id)
}
