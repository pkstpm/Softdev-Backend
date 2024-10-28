package dto

type ReviewDTO struct {
	Content        string `json:"content" valide:"required"`
	FoodRating     int    `json:"foodRating" valide:"required"`
	ServiceRating  int    `json:"serviceRating" valide:"required"`
	AmbienceRating int    `json:"ambienceRating" valide:"required"`
}
