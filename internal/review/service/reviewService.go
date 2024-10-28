package service

import "github.com/pkstpm/Softdev-Backend/internal/review/dto"

type ReviewService interface {
	CreateReview(userId string, reservationId string, review *dto.ReviewDTO) error
}
