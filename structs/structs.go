package structs

type Artists struct {
	ID             int      `json:"id"`
	Image          string   `json:"image,omitempty"`
	Name           string   `json:"name"`
	Members        []string `json:"members"`
	CreationDate   int      `json:"creationDate"`
	FirstAlbum     string   `json:"firstAlbum"`
	Locations      string   `json:"locations"`
	ConcertDates   string   `json:"concertDates"`
	DatesLocations string   `json:"datesLocations"`
}

type Relations struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type RelationID struct {
	Index []Relations `json:"index"`
}

type Locations struct {
	ID        int      `json:"id"`
	Locations []string `json:"locations"`
}
type LocationIndex struct {
	Index []Locations `json:"index"`
}

type ConcertDates struct {
	ID    int      `json:"id"`
	Dates []string `json:"dates"`
}
type ConcertIndex struct {
	Index []ConcertDates `json:"index"`
}
type Search []struct {
	ID             int                 `json:"id"`
	Image          string              `json:"image,omitempty"`
	Name           string              `json:"name"`
	Members        []string            `json:"members"`
	CreationDate   int                 `json:"creationDate"`
	FirstAlbum     string              `json:"firstAlbum"`
	Locations      []string            `json:"locations"`
	ConcertDates   []string            `json:"concertDates"`
	DatesLocations map[string][]string `json:"datesLocations"`
}
