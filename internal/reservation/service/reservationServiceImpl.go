package service

import (
	"errors"

	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/dto"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/model"
	"github.com/pkstpm/Softdev-Backend/internal/reservation/repository"
	restaurantRepository "github.com/pkstpm/Softdev-Backend/internal/restaurant/repository"
)

type reservationServiceImpl struct {
	reservationRepository repository.ReservationRepository
	restaurantRepository  restaurantRepository.RestaurantRepository
}

func NewReservationService(reservationRepository repository.ReservationRepository, restaurantRepository restaurantRepository.RestaurantRepository) ReservationService {
	return &reservationServiceImpl{reservationRepository: reservationRepository, restaurantRepository: restaurantRepository}
}

func (r *reservationServiceImpl) CreateReservation(userId uuid.UUID, dto dto.CreateReservationDTO) (string, error) {

	timeSlots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(dto.RestaurantID.String())

	if err != nil {
		return "", err
	}

	timeSlot := timeSlots[(int(dto.StartTime.Weekday()))]

	if timeSlot.HourStart > dto.StartTime.Hour() || timeSlot.HourEnd < dto.EndTime.Hour() {
		return "", errors.New("reservation time is not within restaurant working hours")
	}

	table, err := r.restaurantRepository.GetTableById(dto.TableID.String())
	if err != nil {
		return "", err
	}

	reservations := table.Reservations
	for _, reservation := range reservations {
		if reservation.StartTime.Before(dto.EndTime) && reservation.EndTime.After(dto.StartTime) {
			return "", errors.New("table is already reserved")
		}
	}

	reservation := &model.Reservation{
		UserID:       userId,
		TableID:      dto.TableID,
		RestaurantID: dto.RestaurantID,
		StartTime:    dto.StartTime,
		EndTime:      dto.EndTime,
		TotalPrice:   0,
	}

	reservationId, err := r.reservationRepository.CreateReservation(reservation)
	if err != nil {
		return "", err
	}

	return reservationId.String(), nil
}

func (r *reservationServiceImpl) GetReservationById(reservationId string) (*model.Reservation, error) {
	reservation, err := r.reservationRepository.GetReservationById(reservationId)
	if err != nil {
		return nil, err
	}
	return reservation, nil
}

func (r *reservationServiceImpl) GetReservationsByUserId(userId string) ([]model.Reservation, error) {
	reservations, err := r.reservationRepository.GetReservationByUserId(userId)
	if err != nil {
		return nil, err
	}
	return reservations, nil
}

func (r *reservationServiceImpl) AddDishItem(userId string, reservationId string, dto dto.AddDishItemDTO) error {
	reservation, err := r.reservationRepository.GetReservationById(reservationId)
	if err != nil {
		return err
	}

	if reservation.UserID.String() != userId {
		return errors.New("reservation does not belong to user")
	}

	dishes, err := r.restaurantRepository.GetDishesByRestaurantId(reservation.RestaurantID.String())

	if err != nil {
		return err
	}

	prices := 0

	for _, dish := range dto.DishItems {
		for _, existingDish := range dishes {
			if dish.DishID == existingDish.ID {
				dishItem := &model.DishItem{
					ReservationID: reservation.ID,
					DishID:        dish.DishID,
					Quantity:      dish.Quantity,
					Price:         existingDish.Price,
					Option:        dish.Option,
					Comment:       dish.Comment,
				}
				prices += existingDish.Price * dish.Quantity
				err = r.reservationRepository.CreateDishItem(dishItem)
				if err != nil {
					return err
				}
			}
		}
	}

	reservation.TotalPrice = prices

	err = r.reservationRepository.UpdateReservation(reservation)
	if err != nil {
		return err
	}

	return nil
}
