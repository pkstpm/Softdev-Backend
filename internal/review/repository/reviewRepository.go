package repository

import "github.com/pkstpm/Softdev-Backend/internal/review/model"

type ReviewRepository interface {
	CreateReview(review *model.Review) error
	FindReviewById(reviewId string) (*model.Review, error)
	FindReviewByUserId(userId string) ([]model.Review, error)
	FindReviewByUserIdAndReservationId(userId string, reservationId string) (*model.Review, error)
}
