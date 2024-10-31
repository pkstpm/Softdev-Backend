package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/dto"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
	"github.com/pkstpm/Softdev-Backend/internal/restaurant/repository"
)

type restaurantServiceImpl struct {
	restaurantRepository repository.RestaurantRepository
}

func NewRestaurantService(restaurantRepository repository.RestaurantRepository) RestaurantService {
	return &restaurantServiceImpl{restaurantRepository: restaurantRepository}
}

func (r *restaurantServiceImpl) FindRestaurantByName(name string) ([]model.Restaurant, error) {
	restaurants, err := r.restaurantRepository.FindRestaurantByName(name)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantServiceImpl) FindRestaurantByCategory(category string) ([]model.Restaurant, error) {
	restaurants, err := r.restaurantRepository.FindRestaurantByCategory(category)
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantServiceImpl) CreateDish(userId string, dto *dto.CreateDishDTO, imgPath string) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	_, err = r.restaurantRepository.FindDishByName(dto.Name, restaurant.ID.String())
	if err == nil {
		return errors.New("dish name already exists")
	}

	var dish = &model.Dish{
		RestaurantID: restaurant.ID,
		Name:         dto.Name,
		Description:  dto.Description,
		Price:        dto.Price,
		ImgPath:      imgPath,
	}

	err = r.restaurantRepository.CreateDish(dish)
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantServiceImpl) UpdateDish(userId string, dishId string, dto *dto.CreateDishDTO, imgPath string) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	dish, err := r.restaurantRepository.FindDishById(dishId)
	if err != nil {
		return err
	}

	if restaurant.ID != dish.RestaurantID {
		return errors.New("dish does not belong to restaurant")
	}

	_, err = r.restaurantRepository.FindDishByName(dto.Name, restaurant.ID.String())
	if err == nil {
		return errors.New("dish name already exists")
	}

	timeSlots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(restaurant.ID.String())
	if err != nil {
		return err
	}

	if err := checkCurrentTimeInTimeSlots(timeSlots); err != nil {
		return err
	}

	dish.Name = dto.Name
	dish.Description = dto.Description
	dish.Price = dto.Price
	dish.ImgPath = imgPath

	err = r.restaurantRepository.UpdateDish(dish)
	if err != nil {
		return err
	}
	return nil
}

func (r *restaurantServiceImpl) GetAllRestaurants() ([]model.Restaurant, error) {
	restaurants, err := r.restaurantRepository.GetAllRestaurants()
	if err != nil {
		return nil, err
	}
	return restaurants, nil
}

func (r *restaurantServiceImpl) CreateTimeSlot(userId string) error {

	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	for i := 0; i <= 6; i++ {
		timeslot := &model.TimeSlot{
			RestaurantID: restaurant.ID,
			Weekday:      i,
			HourStart:    9,
			HourEnd:      21,
			IsClosed:     false,
			Slots:        "09:00-21:00",
		}
		err = r.restaurantRepository.CreateTimeSlot(timeslot)
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *restaurantServiceImpl) GetTimeSlot(userId string) ([]model.TimeSlot, error) {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return nil, err
	}
	timeSlots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(restaurant.ID.String())
	if err != nil {
		return nil, err
	}
	fmt.Println(timeSlots)
	return timeSlots, err
}

func (r *restaurantServiceImpl) GetTimeSlotByRestaurantId(restaurantId string) ([]model.TimeSlot, error) {
	timeSlots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(restaurantId)
	if err != nil {
		return nil, err
	}
	fmt.Println(timeSlots)
	return timeSlots, nil
}

func (r *restaurantServiceImpl) UpdateTimeSlot(userId string, dto *dto.UpdateTimeDTO) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	timeslots, err := r.restaurantRepository.GetTimeSlotsByRestaurantId(restaurant.ID.String())
	if err != nil {
		return err
	}

	for _, timeslot := range timeslots {
		for _, dto := range dto.TimeSlots {
			if timeslot.Weekday == dto.Weekday {
				timeslot.HourStart = dto.HourStart
				timeslot.HourEnd = dto.HourEnd
				timeslot.IsClosed = dto.IsClosed
				timeslot.Slots = dto.Slots
				err = r.restaurantRepository.UpdateTimeSlot(&timeslot)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

func (r *restaurantServiceImpl) GetTablesByRestaurantId(restaurantId string) ([]model.Table, error) {
	tables, err := r.restaurantRepository.GetTablesByRestaurantId(restaurantId)
	if err != nil {
		return nil, err
	}
	return tables, nil
}

func (r *restaurantServiceImpl) GetAllDishesByRestaurantId(restaurantId string) ([]model.Dish, error) {
	dishes, err := r.restaurantRepository.GetAllDishesByRestaurantId(restaurantId)
	if err != nil {
		return nil, err
	}
	return dishes, nil
}

func (r *restaurantServiceImpl) CreateTable(userId string, dto *dto.CreateTableDTO) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	table := &model.Table{
		RestaurantID: restaurant.ID,
		TableNumber:  dto.TableNumber,
		Capacity:     dto.Capacity,
	}

	err = r.restaurantRepository.CreateTable(table)
	if err != nil {
		return err
	}

	return nil
}

func (r *restaurantServiceImpl) UploadRestaurantPictures(userId string, uploadedFiles []string) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	for _, file := range uploadedFiles {
		err = r.restaurantRepository.CreateImages(&model.Image{
			RestaurantID: restaurant.ID,
			ImgPath:      file,
		})
		if err != nil {
			return err
		}
	}

	return nil
}

func (r *restaurantServiceImpl) GetRestaurantByID(restaurantId string) (*model.Restaurant, error) {
	restaurant, err := r.restaurantRepository.FindRestaurantByID(restaurantId)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func (r *restaurantServiceImpl) DeletetRestaurantPicture(userId string, pictureId string) error {
	_, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	parsedID, err := uuid.Parse(pictureId)
	if err != nil {
		return err
	}

	err = r.restaurantRepository.DeleteImage(parsedID)
	if err != nil {
		return err
	}

	return nil
}

func (r *restaurantServiceImpl) GetRestaurantByUserId(userId string) (*model.Restaurant, error) {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return nil, err
	}
	return restaurant, nil
}

func checkCurrentTimeInTimeSlots(timeSlots []model.TimeSlot) error {
	currentTime := time.Now()
	currentWeekday := int(currentTime.Weekday())
	currentHour := currentTime.Hour()

	for _, slot := range timeSlots {
		if slot.Weekday == currentWeekday && currentHour >= slot.HourStart && currentHour < slot.HourEnd {
			return errors.New("failed the restaurant is currently open")
		}
	}
	return nil
}

func (r *restaurantServiceImpl) GetDishByID(dishId string) (*model.Dish, error) {
	dish, err := r.restaurantRepository.GetDishByID(dishId)
	if err != nil {
		return nil, err
	}
	return dish, nil
}

func (r *restaurantServiceImpl) GetTableByID(tableId string) (*model.Table, error) {
	table, err := r.restaurantRepository.GetTableByID(tableId)
	if err != nil {
		return nil, err
	}
	return table, nil
}

func (r *restaurantServiceImpl) UpdateRestaurant(userId string, dto *dto.UpdateRestaurantDTO) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	restaurant.Description = dto.Description
	restaurant.Category = dto.Category
	restaurant.Latitude = dto.Latitude
	restaurant.Longitude = dto.Longitude

	err = r.restaurantRepository.UpdateRestaurant(restaurant)
	if err != nil {
		return err
	}

	return nil
}

func (r *restaurantServiceImpl) UploadTablePicture(userId string, dstPath string) error {
	restaurant, err := r.restaurantRepository.FindRestaurantByUserID(userId)
	if err != nil {
		return err
	}

	restaurant.ImgPath = dstPath

	err = r.restaurantRepository.UpdateRestaurant(restaurant)
	if err != nil {
		return err
	}

	return nil
}
