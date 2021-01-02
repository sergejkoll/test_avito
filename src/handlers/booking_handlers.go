package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"

	"test_avito/src/models"
)

func (env *Env) CreateBookingHandler(w http.ResponseWriter, r *http.Request) {
	paramFromURL := mux.Vars(r)
	roomId, err := strconv.Atoi(paramFromURL["room_id"])
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	booking, err := validateBooking(r)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	booking.RoomId = roomId

	bookingWithId, err := env.Booking.CreateBooking(booking)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(bookingWithId)
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func validateBooking(r *http.Request) (booking models.Booking, err error) {
	reqBooking := models.BookingRequest{}
	err = json.NewDecoder(r.Body).Decode(&reqBooking)
	if err != nil {
		return models.Booking{}, err
	}

	dateStart, err := time.Parse("2006-01-02", reqBooking.DateStart)
	if err != nil {
		return models.Booking{}, err
	}
	booking.DateStart = dateStart

	dateEnd, err := time.Parse("2006-01-02", reqBooking.DateEnd)
	if err != nil {
		return models.Booking{}, err
	}
	booking.DateEnd = dateEnd

	return booking, nil
}

func (env *Env) DeleteBookingHandler(w http.ResponseWriter, r*http.Request) {
	paramFromURL := mux.Vars(r)
	id, err := strconv.Atoi(paramFromURL["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = env.Booking.DeleteBooking(id)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (env *Env) GetBookingsHandler(w http.ResponseWriter, r*http.Request) {
	paramFromURL := mux.Vars(r)
	roomId, err := strconv.Atoi(paramFromURL["room_id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	bookings, err := env.Booking.GetBookings(roomId)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(bookings)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}