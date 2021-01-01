package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"test_avito/src/handlers"
	"test_avito/src/services"
)

func main() {
	db, err := services.OpenDB()
	if err != nil {
		log.Panic()
	}
	defer db.Close()

	err = services.Setup("src/database/database.sql", db)
	if err != nil {
		log.Panic()
	}
	log.Println("Database is ready!")

	env := &handlers.Env{
		Rooms: services.RoomModel{DB: db},
	}

	r := mux.NewRouter()
	r.HandleFunc("/room/create", env.CreateRoomHandler).Methods("POST")
	r.HandleFunc("/room/{id}", env.DeleteRoomHandler).Methods("DELETE")
	r.HandleFunc("/room", env.GetRoomsHandler).Methods("GET")

	r.HandleFunc("/booking/create", env.CreateBookingHandler).Methods("POST")
	r.HandleFunc("/booking/{id}", env.DeleteBookingHandler).Methods("DELETE")
	r.HandleFunc("/booking/{room_id}", env.GetBookingsHandler).Methods("GET")

	err = http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatal(err)
	}
}
