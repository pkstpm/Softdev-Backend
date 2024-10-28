package repository

import (
	"github.com/pkstpm/Softdev-Backend/internal/database"
	"github.com/pkstpm/Softdev-Backend/internal/review/model"
	"gorm.io/gorm"
)

type reviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db database.Database) ReviewRepository {
	return &reviewRepository{db: db.GetDb()}
}

func (r *reviewRepository) CreateReview(review *model.Review) error {
	err := r.db.Create(review).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *reviewRepository) FindReviewById(reviewId string) (*model.Review, error) {
	var review model.Review
	err := r.db.Where("id = ?", reviewId).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (r *reviewRepository) FindReviewByUserId(userId string) ([]model.Review, error) {
	var reviews []model.Review
	err := r.db.Where("user_id = ?", userId).Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (r *reviewRepository) FindReviewByUserIdAndReservationId(userId string, reservationId string) (*model.Review, error) {
	var review model.Review
	err := r.db.Where("user_id = ? AND reservation_id = ?", userId, reservationId).First(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}
