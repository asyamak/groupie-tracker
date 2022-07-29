package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"groupie-tracker/internal/logic"
	"groupie-tracker/models"
)

type Error struct {
	Status  int
	Message string
}
type Artists struct {
	Artistss   []models.ArtistResult
	Relationss []models.Relations
}

type Index struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

type Handler struct {
	count int
}

func NewHandler() *Handler {
	return &Handler{
		count: 0,
	}
}

// GETHandler func receives only GET request and displays main page.
func (h *Handler) GETHandler(w http.ResponseWriter, r *http.Request) {
	status := checkMethodAndPath(r, "/", http.MethodGet)
	if status != 200 {
		w.WriteHeader(status)
		// erroHandler(w, http.StatusBadRequest int){
		// 	Status = statusbadrequesth
		// 	message = http.textmessage(staus)
		// 	t1.execute(w, html, error)
		// }

		templateExecution(w, "./ui/html/"+fmt.Sprint(status)+".html", status)
		return
	}
	var artists []models.ArtistResult
	if h.count == 0 {
		artists = logic.GetArtistss()
		logic.GetRelations()
		h.count++
	} else {
		artists = logic.Array
	}
	data := Artists{
		Artistss: artists,
	}
	templateExecution(w, "./ui/html/home.html", data)
}

// Post handler responces only post request and processes data we receive through FromValue
// checks if text is correct and do not contain cyrilic alphabet, correct new lines,checks if font name is correct
// and if file that contains format font haven't been modified through HashSum & ConverFont func.
func (h *Handler) POSTHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		id_str := strings.TrimPrefix(r.URL.Path, "/artist/")
		id, err := strconv.Atoi(id_str)
		if err != nil {
			return
		}
		artist := logic.MapArtists[id]
		artist.Relation = logic.MapRelations[id]
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(artist)
	}
}
