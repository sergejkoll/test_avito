package handlers

import "test_avito/src/models"

type Env struct {
	Rooms interface{
		CreateRoom(room models.Room) (models.Room, error)
		DeleteRoom(roomId int) error
		GetRooms(bool, bool, bool) ([]models.Room, error)
	}
	Booking interface{
		CreateBooking(booking models.Booking) (models.Booking, error)
		DeleteBooking(bookingId int) error
		GetBookings(roomId int) ([]models.Booking, error)
	}
}
