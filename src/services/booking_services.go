package services

import (
	"database/sql"
	"log"

	"test_avito/src/models"
)

type BookingModel struct {
	DB *sql.DB
}

func (m BookingModel) CreateBooking(booking models.Booking) (models.Booking, error) {
	err := m.DB.QueryRow("INSERT INTO booking (room_id, date_start, date_end) values ($1, $2, $3) RETURNING booking_id",
		booking.RoomId, booking.DateStart, booking.DateEnd).Scan(&booking.BookingId)
	if err != nil {
		log.Println("Add booking in db error: ", err)
		return models.Booking{}, err
	}
	return booking, nil
}

func (m BookingModel) DeleteBooking(bookingId int) error {
	_, err := m.DB.Exec("DELETE FROM booking WHERE booking_id = $1", bookingId)
	if err != nil {
		log.Println("Delete booking from db error: ", err)
		return err
	}
	return nil
}

func (m BookingModel) GetBookings(roomId int) ([]models.Booking, error) {
	rows, err := m.DB.Query("SELECT * FROM booking WHERE room_id = $1 ORDER BY date_start", roomId)
	if err != nil {
		log.Println("Error get bookings from database: ", err)
		return nil, err
	}

	var bookings []models.Booking
	for rows.Next() {
		booking := models.Booking{}

		err = rows.Scan(&booking.BookingId, &booking.RoomId, &booking.DateStart, &booking.DateEnd)
		if err != nil {
			log.Println("Scan from db error: ", err)
			return nil, err
		}

		bookings = append(bookings, booking)
	}

	return bookings, nil
}
