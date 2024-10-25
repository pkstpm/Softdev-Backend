package dto

type UpdateTimeSlotDTO struct {
	Weekday   int `json:"weekday" validate:"required,gte=0,lte=6"`
	HourStart int `json:"hour_start" validate:"required,gte=0,lte=23"`
	HourEnd   int `json:"hour_end" validate:"required,gte=0,lte=23,gtfield=HourStart"`
}

type UpdateTimeDTO struct {
	TimeSlots []UpdateTimeSlotDTO `json:"time_slot" validate:"required,dive"`
}

type CreateDishDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Description string `json:"description" validate:"required,min=3,max=200"`
	Price       int    `json:"price" validate:"required"`
}

type UpdateDishDTO struct {
	ID          string `json:"id" validate:"required"`
	Name        string `json:"name" validate:"required,min=3,max=30"`
	Description string `json:"description" validate:"required,min=3,max=200"`
	Price       int    `json:"price" validate:"required"`
}

type CreateTableDTO struct {
	TableNumber int `json:"table_number" validate:"required"`
	Capacity    int `json:"capacity" validate:"required"`
}
