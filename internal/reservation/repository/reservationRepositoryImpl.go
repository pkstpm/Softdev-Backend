package repository

import (
	"github.com/google/uuid"
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

func (r *reservationRepository) CreateReservation(reservation *model.Reservation) (uuid.UUID, error) {
	err := r.db.Create(reservation).Error
	if err != nil {
		return uuid.Nil, err
	}
	return reservation.ID, nil
}

func (r *reservationRepository) CreateDishItem(dishItem *model.DishItem) error {
	err := r.db.Create(dishItem).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *reservationRepository) GetReservationById(reservationId string) (*model.Reservation, error) {
	var reservation model.Reservation
	err := r.db.Where("id = ?", reservationId).First(&reservation).Error
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
