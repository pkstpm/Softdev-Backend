package repository

import (
	"errors"
	"log"

	"github.com/pkstpm/Softdev-Backend/internal/database"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/model"
	"gorm.io/gorm"
)

type reservationRepository struct {
	db *gorm.DB
}

func NewReservationRepository(db database.Database) ReservationRepository {
	return &reservationRepository{db: db.GetDb()}
}

func (r *reservationRepository) CreateReservation(reservation *model.Reservation) (*model.Reservation, error) {
	err := r.db.Create(reservation).Error
	if err != nil {
		return reservation, err
	}
	return reservation, nil
}

func (r *reservationRepository) CreateDishItem(dishItem *model.DishItem) (*model.DishItem, error) {
	err := r.db.Create(dishItem).Error
	if err != nil {
		return nil, err
	}
	return dishItem, nil
}

func (r *reservationRepository) GetReservationById(reservationId string) (*model.Reservation, error) {
	var reservation model.Reservation
	log.Printf("reservationId: %s", reservationId)
	err := r.db.Preload("DishItems").Where("id = ?", reservationId).First(&reservation).Error

	// Check if the record was not found
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Explicitly return nil for reservation if not found
	}

	// Return any other errors
	if err != nil {
		return nil, err
	}

	return &reservation, nil
}

func (r *reservationRepository) GetResvationByUserId(userId string) ([]model.Reservation, error) {
	var reservations []model.Reservation
	err := r.db.Where("user_id = ?", userId).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *reservationRepository) GetReservationByRestaurantId(restaurantId string) ([]model.Reservation, error) {
	var reservations []model.Reservation
	err := r.db.Where("restaurant_id = ?", restaurantId).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *reservationRepository) GetReservationByUserId(userId string) ([]model.Reservation, error) {
	var reservations []model.Reservation
	err := r.db.Preload("DishItems").Where("user_id = ?", userId).Find(&reservations).Error
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *reservationRepository) UpdateReservation(reservation *model.Reservation) error {
	err := r.db.Save(reservation).Error
	if err != nil {
		return err
	}
	return nil
}
