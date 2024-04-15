package src

type Post struct {
	ID             int
	Name           string
	ImageURL       string
	Members        []string
	YearCreation   int
	FirstAlbum     string
	DatesLocations map[string][]string
}

type LocationData struct {
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Post2 struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	ImageURL     string   `json:"image"`
	Members      []string `json:"members"`
	YearCreation int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Emplacement  string   `json:"locations"`
	ConcertDate  string   `json:"concertDates"`
	Relations    string   `json:"relation"`
}

type Loc struct {
	ID   int    `json:"id"`
	LIEU string `json:"locations"`
}

type TempArtist struct {
	ID           int      `json:"id"`
	Name         string   `json:"name"`
	ImageURL     string   `json:"image"`
	Members      []string `json:"members"`
	YearCreation int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	Emplacement  string   `json:"locations"`
	ConcertDate  string   `json:"concertDates"`
	Relations    string   `json:"relation"`
}

type Date struct {
	DATE []string `json:"dates"`
}

type Artist struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type Location struct {
	Id        uint64   `json:"id"`
	Locations []string `json:"locations"`
	DatesUrl  string   `json:"dates"`
}
