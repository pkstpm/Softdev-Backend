package dto

type ReviewDTO struct {
	Content        string `json:"content"`
	FoodRating     int    `json:"food_rating" valide:"required"`
	ServiceRating  int    `json:"service_rating" valide:"required"`
	AmbienceRating int    `json:"ambience_rating" valide:"required"`
}
