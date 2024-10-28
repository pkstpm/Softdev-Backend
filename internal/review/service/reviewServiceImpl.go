package service

import (
	"errors"

	"github.com/google/uuid"
	reservationRepository "github.com/pkstpm/Softdev-Backend/internal/reservation/repository"
	"github.com/pkstpm/Softdev-Backend/internal/review/dto"
	"github.com/pkstpm/Softdev-Backend/internal/review/model"
	"github.com/pkstpm/Softdev-Backend/internal/review/repository"
)

type reviewServiceImpl struct {
	reviewRepository      repository.ReviewRepository
	reservationRepository reservationRepository.ReservationRepository
}

func NewReviewService(reviewRepository repository.ReviewRepository) ReviewService {
	return &reviewServiceImpl{reviewRepository: reviewRepository}
}

func (s *reviewServiceImpl) CreateReview(userId string, reservationId string, review *dto.ReviewDTO) error {
	// Step 1: Find the reservation by ID
	reservation, err := s.reservationRepository.GetReservationById(reservationId)
	if err != nil {
		return err
	}

	// Step 2: Check if the reservation exists and belongs to the user
	if reservation == nil {
		return errors.New("reservation not found")
	}

	if reservation.UserID.String() != userId {
		return errors.New("user does not own this reservation")
	}

	existReview, _ := s.reviewRepository.FindReviewByUserIdAndReservationId(userId, reservationId)
	if existReview != nil {
		return errors.New("review already exists")
	}

	// Step 3: Create the review
	newReview := &model.Review{
		UserID:         uuid.MustParse(userId), // Convert string to UUID
		ReservationID:  reservation.ID,         // Use the reservation's UUID
		Content:        review.Content,
		FoodRating:     review.FoodRating,
		ServiceRating:  review.ServiceRating,
		AbbienceRating: review.AmbienceRating,
	}

	// Step 4: Save the review to the database
	err = s.reviewRepository.CreateReview(newReview)
	if err != nil {
		return err
	}

	reservation.ID = uuid.MustParse(reservationId)
	err = s.reservationRepository.UpdateReservation(reservation)

	if err != nil {
		return err
	}

	return nil
}
