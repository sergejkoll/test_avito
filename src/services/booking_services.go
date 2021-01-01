package services

import (
	"database/sql"
	"test_avito/src/models"
)

type BookingModel struct {
	DB *sql.DB
}

func (m BookingModel) CreateBooking(booking models.Booking) (models.Booking, error) {
	return booking, nil
}

func (m BookingModel) DeleteBooking(bookingId int) error {
	return nil
}

func (m BookingModel) GetBookings(roomId int) ([]models.Booking, error) {
	return []models.Booking{}, nil
}
