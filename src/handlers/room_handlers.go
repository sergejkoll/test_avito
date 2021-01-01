package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"test_avito/src/models"
)

func (env *Env) CreateRoomHandler(w http.ResponseWriter, r*http.Request) {
	room := models.Room{}
	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	roomWithId, err := env.Rooms.CreateRoom(room)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(roomWithId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (env *Env) DeleteRoomHandler(w http.ResponseWriter, r*http.Request) {
	paramFromURL := mux.Vars(r)
	id, err := strconv.Atoi(paramFromURL["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = env.Rooms.DeleteRoom(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (env *Env) GetRoomsHandler(w http.ResponseWriter, r*http.Request) {
	sortByPrice, sortByDate, asc, err := parseQuery(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	rooms, err := env.Rooms.GetRooms(sortByPrice, sortByDate, asc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(rooms)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func parseQuery(r*http.Request) (bool, bool, bool, error) {
	sortByPriceStr := r.URL.Query().Get("price")
	sortByDateStr := r.URL.Query().Get("date")
	ascStr := r.URL.Query().Get("asc")

	sortByPrice, sortByDate, asc := false, false, false
	if sortByPriceStr == "true" {
		sortByPrice = true
	} else if sortByPriceStr == "false" || sortByPriceStr == "" {
		sortByPrice = false
	} else {
		return false, false, false, errors.New("Bad Request ")
	}

	if sortByDateStr == "true" {
		sortByDate = true
	} else if sortByDateStr == "false" || sortByDateStr == "" {
		sortByDate = false
	} else {
		return false, false, false, errors.New("Bad Request ")
	}

	if ascStr == "true" {
		asc = true
	} else if ascStr == "false" || ascStr == "" {
		asc = false
	} else {
		return false, false, false, errors.New("Bad Request ")
	}

	return sortByPrice, sortByDate, asc, nil
}