package service

import (
	"errors"
	"log"
	"time"

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
	reservationYear, reservationMonth, reservationDay := dto.StartTime.Date()
	reservationEndYear, reservationEndMonth, reservationEndDay := dto.EndTime.Date()

	today := time.Now().Truncate(24 * time.Hour)

	// Truncate the reservation date to midnight to ignore the time portion
	reservationDate := dto.StartTime.Truncate(24 * time.Hour)

	// Check if the reservation date is in the future
	if !reservationDate.Equal(today) && !reservationDate.After(today) {
		return "", errors.New("reservation date must be today or in the future")
	}

	// Ensure that the reservation is within the correct date range
	if reservationYear != reservationEndYear || reservationMonth != reservationEndMonth || reservationDay != reservationEndDay {
		return "", errors.New("reservation must start and end on the same day")
	}

	currentHour := time.Now().Hour()  // Get the current hour
	startHour := dto.StartTime.Hour() // Get the hour from dto.StartTime

	if startHour < currentHour {
		return "", errors.New("the start time must be in the future")
	}

	if timeSlot.HourStart > dto.StartTime.Hour() || timeSlot.HourEnd < dto.EndTime.Hour() {
		return "", errors.New("reservation time is not within restaurant working hours")
	}

	table, err := r.restaurantRepository.GetTableById(dto.TableID.String())
	if err != nil {
		return "", err
	}

	reservations := table.Reservations
	log.Printf("Reservations: %v", reservations)
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
		Status:       "Pending",
		TotalPrice:   0,
	}

	reservation, err = r.reservationRepository.CreateReservation(reservation)
	if err != nil {
		return "", err
	}

	return reservation.ID.String(), nil
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

	log.Println(reservation)

	// Ensure that the reservation is properly initialized
	if reservation.ID == uuid.Nil {
		return errors.New("reservation not found or invalid")
	}

	if reservation.UserID.String() != userId {
		return errors.New("reservation does not belong to user")
	}

	dishes, err := r.restaurantRepository.GetDishesByRestaurantId(reservation.RestaurantID.String())
	if err != nil {
		return err
	}

	totalPrice := 0

	for _, dish := range dto.DishItems {
		for _, existingDish := range dishes {
			if dish.DishID == existingDish.ID {
				// Parse the UUIDs and handle any errors
				reservationID, err := uuid.Parse(reservation.ID.String())
				if err != nil {
					log.Printf("Error parsing reservation ID: %v", err)
					continue // Skip this iteration if parsing fails
				}

				log.Println(reservationID)

				dishID, err := uuid.Parse(existingDish.ID.String())
				if err != nil {
					log.Printf("Error parsing dish ID: %v", err)
					continue // Skip this iteration if parsing fails
				}

				log.Println(dishID)

				dishItem := &model.DishItem{
					ReservationID: reservationID,
					DishID:        dishID,
					Quantity:      dish.Quantity,
					Price:         existingDish.Price,
					Option:        dish.Option,
					Comment:       dish.Comment,
				}
				totalPrice += existingDish.Price * dish.Quantity

				// Attempt to create the dish item
				if _, err := r.reservationRepository.CreateDishItem(dishItem); err != nil {
					log.Printf("Error creating dish item for dish ID %s: %v", dish.DishID, err)
					continue // Continue processing the rest of the items
				}
			}
		}
	}

	// Update the total price of the reservation
	reservation.TotalPrice = totalPrice

	if err := r.reservationRepository.UpdateReservation(reservation); err != nil {
		return err
	}

	return nil
}

func (s *reservationServiceImpl) StartReservationUpdateRoutine() {
	ticker := time.NewTicker(5 * time.Second)

	go func() {
		for {
			<-ticker.C
			err := s.reservationRepository.UpdateExpiredReservations()
			if err != nil {
				log.Printf("Error updating expired reservations: %v", err)
			} else {
				log.Println("Successfully updated expired reservations to Completed.")
			}
		}
	}()
}
