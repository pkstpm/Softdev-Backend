package service

import (
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

func (r *reservationServiceImpl) CreateReservation(userId uuid.UUID, dto dto.CreateReservationDTO) error {
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
		return err
	}

	for _, dishItem := range dto.DishItems {
		price, _ := r.restaurantRepository.GetDishPrice(dishItem.DishID)

		dishItem := &model.DishItem{
			ReservationID: reservationId,
			DishID:        dishItem.DishID,
			Quantity:      dishItem.Quantity,
			Price:         price * dishItem.Quantity,
			Comment:       dishItem.Comment,
		}

		err := r.reservationRepository.CreateDishItem(dishItem)
		if err != nil {
			return err
		}
	}

	return nil
}
