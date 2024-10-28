package database

import (
	"fmt"
	"log"
	"sync"

	"github.com/pkstpm/Softdev-Backend/internal/config"
	notificationModel "github.com/pkstpm/Softdev-Backend/internal/notification/model"
	reservationModel "github.com/pkstpm/Softdev-Backend/internal/reservation/model"
	restaurantModel "github.com/pkstpm/Softdev-Backend/internal/restaurant/model"
	reviewModel "github.com/pkstpm/Softdev-Backend/internal/review/model"
	userModels "github.com/pkstpm/Softdev-Backend/internal/users/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type postgresDatabase struct {
	Db *gorm.DB
}

var (
	once       sync.Once
	dbInstance *postgresDatabase
)

func NewPostgresDatabase(conf *config.Config) Database {
	log.Print("testtestset", conf.Database.Host,
		conf.Database.User,
		conf.Database.Password,
		conf.Database.Name,
		conf.Database.Port,
		conf.Database.SSLMode,
		conf.Database.TimeZone)
	once.Do(func() {

		// Construct the DSN string
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=%s",
			conf.Database.Host,
			conf.Database.User,
			conf.Database.Password,
			conf.Database.Name,
			conf.Database.Port,
			conf.Database.SSLMode,
			conf.Database.TimeZone,
		)

		// Open the database connection
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("failed to connect database")
		}

		// Set the dbInstance if connection is successful
		dbInstance = &postgresDatabase{Db: db}
	})

	// Return the initialized dbInstance
	return dbInstance
}

func (p *postgresDatabase) GetDb() *gorm.DB {
	// Return the GORM DB instance
	return dbInstance.Db
}

func (p *postgresDatabase) Migrate() {
	var err error
	err = p.GetDb().Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"").Error
	if err != nil {
		log.Fatal("failed to create extension:", err)
	}

	// p.GetDb().Exec("CREATE TYPE user_type AS ENUM ('Customer', 'Restaurant');")
	// p.GetDb().Exec("CREATE TYPE status_type AS ENUM ('Accepted', 'Denied', 'Completed', 'Cancelled', 'Pending);")
	// Migrate the schema
	// p.GetDb().Migrator().DropTable(&userModels.User{}, &restaurantModel.Restaurant{}, &restaurantModel.Dish{}, &restaurantModel.Table{}, &restaurantModel.TimeSlot{}, &restaurantModel.Dish{}, &reservationModel.Reservation{}, &reservationModel.DishItem{}, &restaurantModel.Image{}, &userModels.Favourite{}, &reviewModel.Review{}, &notificationModel.Notification{})
	// p.GetDb().Migrator().DropTable(&reservationModel.DishItem{})
	err = p.GetDb().AutoMigrate(&userModels.User{}, &restaurantModel.Restaurant{}, &restaurantModel.Dish{}, &restaurantModel.Table{}, &restaurantModel.TimeSlot{}, &restaurantModel.Dish{}, &reservationModel.Reservation{}, &reservationModel.DishItem{}, &restaurantModel.Image{}, &userModels.Favourite{}, &reviewModel.Review{}, &notificationModel.Notification{})
	if err != nil {
		panic("failed to migrate database")
	}
}
