package models

import "time"

type Booking struct {
	BookingId int       `json:"booking_id"`
	RoomId    int       `json:"room_id"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}

type BookingRequest struct {
	BookingId int    `json:"booking_id"`
	RoomId    int    `json:"room_id"`
	DateStart string `json:"date_start"`
	DateEnd   string `json:"date_end"`
}
