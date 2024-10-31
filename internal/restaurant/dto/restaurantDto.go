package dto

type UpdateRestaurantDTO struct {
	Description string  `json:"description" validate:"required,min=3,max=200"`
	Category    string  `json:"category" validate:"required,min=3,max=30"`
	Latitude    float64 `json:"latitude" validate:"required"`
	Longitude   float64 `json:"longitude" validate:"required"`
}

type UpdateTimeSlotDTO struct {
	Weekday   int    `json:"weekday" validate:"gte=0,lte=6"`
	HourStart int    `json:"hour_start" validate:"required,gte=0,lte=23"`
	HourEnd   int    `json:"hour_end" validate:"required,gte=0,lte=23,gtfield=HourStart"`
	Slots     string `json:"slots"`
	IsClosed  bool   `json:"is_closed"`
}

type UpdateTimeDTO struct {
	TimeSlots []UpdateTimeSlotDTO `json:"time_slot" validate:"required,dive"`
}

type CreateDishDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Description string `json:"description" validate:"required,min=3,max=200"`
	Price       int    `json:"price" validate:"required"`
	Slots       string `json:"slots"`
}

type UpdateDishDTO struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Description string `json:"description" validate:"required,min=3,max=200"`
	Price       int    `json:"price" validate:"required"`
}

type CreateTableDTO struct {
	TableNumber string `json:"table_number" validate:"required"`
	Capacity    int    `json:"capacity" validate:"required"`
}
