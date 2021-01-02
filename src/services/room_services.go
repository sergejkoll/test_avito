package services

import (
	"database/sql"
	"log"
	"strings"
	"time"

	"test_avito/src/models"
)

type RoomModel struct {
	DB *sql.DB
}

func (m RoomModel) CreateRoom(room models.Room) (models.Room, error) {
	room.DateAdded = time.Now().Unix()
	err := m.DB.QueryRow("INSERT INTO room (description, price, date_added) values ($1, $2, $3) RETURNING id",
		room.Description, room.Price, room.DateAdded).Scan(&room.Id)
	if err != nil {
		log.Println("Add room in db error: ", err)
		return models.Room{}, err
	}
	return room, nil
}

func (m RoomModel) DeleteRoom(roomId int) error {
	_, err := m.DB.Exec("DELETE FROM room WHERE id = $1", roomId)
	if err != nil {
		log.Println("Delete room from db error: ", err)
		return err
	}
	return nil
}

func (m RoomModel) GetRooms(sortByPrice bool, sortByDate bool, asc bool) ([]models.Room, error) {
	query := "SELECT * FROM room"

	// TODO: remove this shit
	var queryParams []string
	if sortByPrice || sortByDate {
		query += " ORDER BY "
		if sortByPrice {
			if asc {
				queryParams = append(queryParams, "price ASC")
			} else {
				queryParams = append(queryParams, "price DESC")
			}
		}
		if sortByDate {
			if asc {
				queryParams = append(queryParams, "date_added ASC")
			} else {
				queryParams = append(queryParams, "date_added DESC")
			}
		}
	}
	query += strings.Join(queryParams, ", ")

	rows, err := m.DB.Query(query)
	if err != nil {
		log.Println("Error get rooms from database: ", err)
		return nil, err
	}

	var rooms []models.Room
	for rows.Next() {
		room := models.Room{}

		err = rows.Scan(&room.Id, &room.Description, &room.Price, &room.DateAdded)
		if err != nil {
			log.Println("Scan from db error: ", err)
			return nil, err
		}

		rooms = append(rooms, room)
	}
	return rooms, nil
}
