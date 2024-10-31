package server

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	authController "github.com/pkstpm/Softdev-Backend/internal/auth/controller"
	authService "github.com/pkstpm/Softdev-Backend/internal/auth/service"
	"github.com/pkstpm/Softdev-Backend/internal/config"
	"github.com/pkstpm/Softdev-Backend/internal/database"
	"github.com/pkstpm/Softdev-Backend/internal/middlewares"
	reservationController "github.com/pkstpm/Softdev-Backend/internal/reservation/controller"
	reservationRepository "github.com/pkstpm/Softdev-Backend/internal/reservation/repository"
	reservationService "github.com/pkstpm/Softdev-Backend/internal/reservation/service"
	restaurantController "github.com/pkstpm/Softdev-Backend/internal/restaurant/controller"
	restaurantRepository "github.com/pkstpm/Softdev-Backend/internal/restaurant/repository"
	restaurantService "github.com/pkstpm/Softdev-Backend/internal/restaurant/service"
	reviewController "github.com/pkstpm/Softdev-Backend/internal/review/controller"
	reviewRepository "github.com/pkstpm/Softdev-Backend/internal/review/repository"
	reviewService "github.com/pkstpm/Softdev-Backend/internal/review/service"

	notificationController "github.com/pkstpm/Softdev-Backend/internal/notification/controller"
	notificationRepository "github.com/pkstpm/Softdev-Backend/internal/notification/repository"
	notificationService "github.com/pkstpm/Softdev-Backend/internal/notification/service"

	userController "github.com/pkstpm/Softdev-Backend/internal/users/controller"
	userRepository "github.com/pkstpm/Softdev-Backend/internal/users/repository"
	userService "github.com/pkstpm/Softdev-Backend/internal/users/service"
)

type echoServer struct {
	app  *echo.Echo
	db   database.Database
	conf *config.Config
}

func NewEchoServer(conf *config.Config, db database.Database) Server {
	echoApp := echo.New()
	echoApp.Logger.SetLevel(log.DEBUG)

	return &echoServer{
		app:  echoApp,
		db:   db,
		conf: conf,
	}
}

func (s *echoServer) Start() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())

	s.app.Use(middleware.CORS())
	s.app.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"}, // Allow all origins, or specify your origins here
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{"Content-Type", "Authorization"},
	}))

	s.app.GET("/", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.app.Static("/uploads", "uploads")

	s.app.Group(s.conf.Server.Prefix)
	s.initAuthRoute()
	s.initUserRoute()
	s.initRestaurantRoute()
	s.initReservationRoute()
	s.initReviewRoute()
	s.initImageRoute()
	s.initNotificationRoute()

	serverUrl := fmt.Sprintf(":%d", s.conf.Server.Port)
	s.app.Logger.Fatal(s.app.Start(serverUrl))
}

func (s *echoServer) initAuthRoute() {
	userRepository := userRepository.NewUserRepository(s.db)
	restaurantRepository := restaurantRepository.NewRestaurantRepository(s.db)
	authService := authService.NewAuthService(userRepository, restaurantRepository)
	restaurantService := restaurantService.NewRestaurantService(restaurantRepository)
	authController := authController.NewAuthController(authService, restaurantService)

	// Routers
	authRouters := s.app.Group("/auth")
	authRouters.POST("/register", authController.Register)
	authRouters.POST("/login", authController.Login)

	authRouters.Use(middlewares.JWTMiddleware())
	authRouters.GET("/me", authController.Me)
	authRouters.POST("/register-restaurant", authController.RegisterRestaurant)
}

func (s *echoServer) initUserRoute() {
	userRepository := userRepository.NewUserRepository(s.db)
	userService := userService.NewUserService(userRepository)
	userController := userController.NewUserController(userService)

	userRouters := s.app.Group("/users")
	userRouters.Use(middlewares.JWTMiddleware())
	userRouters.GET("/profile", userController.ViewProfile)
	userRouters.POST("/upload-profile-picture", userController.UploadUserProfilePicture)
	userRouters.PUT("/edit-profile", userController.UpdateProfile)
	userRouters.PUT("/change-password", userController.ChangePassword)
	userRouters.POST("/add-favourite-restaurant/:restaurant_id", userController.AddFavouriteRestaurant)
	userRouters.DELETE("/remove-favourite-restaurant/:restaurant_id", userController.RemoveFavouriteRestaurant)
}

func (s *echoServer) initRestaurantRoute() {
	restaurantRepository := restaurantRepository.NewRestaurantRepository(s.db)
	restaurantService := restaurantService.NewRestaurantService(restaurantRepository)
	restaurantController := restaurantController.NewRestaurantController(restaurantService)

	searchrouter := s.app.Group("/search")
	searchrouter.GET("/:name", restaurantController.FindByName)
	searchrouter.GET("/category/:category", restaurantController.FindByCategory)

	restaurantRouters := s.app.Group("/restaurant")
	restaurantRouters.GET("/get-all", restaurantController.GetAllRestaurants)
	restaurantRouters.GET("/get-table/:restaurant_id", restaurantController.GetTable)
	restaurantRouters.GET("/:restaurant_id", restaurantController.GetRestaurantByID)
	restaurantRouters.GET("/get-time-slot/:restaurant_id", restaurantController.GetTimeSlotById)
	restaurantRouters.GET("/get-dish/:restaurant_id", restaurantController.GetDishesByRestaurantId)

	restaurantRouters.Use(middlewares.JWTMiddleware())
	restaurantRouters.GET("/get-my-restaurant", restaurantController.GetMyRestaurant)
	restaurantRouters.GET("/get-dish", restaurantController.GetDishesById)
	restaurantRouters.GET("/get-time-slot", restaurantController.GetTimeSlot)
	restaurantRouters.POST("/update-time-slot", restaurantController.UpdateTimeSlot)
	restaurantRouters.POST("/create-dish", restaurantController.CreateDish)
	restaurantRouters.POST("/create-table", restaurantController.CreateTable)
	restaurantRouters.PUT("/update-dish", restaurantController.UpdateDish)
	restaurantRouters.POST("/upload-restaurant-picture", restaurantController.UploadRestaurantPictures)
	restaurantRouters.DELETE("/delete-restaurant-picture:image_id", restaurantController.DeleteRestauranPictures)
}

func (s *echoServer) initReservationRoute() {
	reservationRepository := reservationRepository.NewReservationRepository(s.db)
	restaurantRepository := restaurantRepository.NewRestaurantRepository(s.db)
	reservationService := reservationService.NewReservationService(reservationRepository, restaurantRepository)
	reservationController := reservationController.NewReservationController(reservationService)

	reservationRouters := s.app.Group("/reservation")
	reservationRouters.Use(middlewares.JWTMiddleware())
	reservationRouters.POST("/create-reservation", reservationController.CreateReservation)
	reservationRouters.GET("/get-reservation/:reservation_id", reservationController.GetReservationById)
	reservationRouters.GET("/get-my-reservation", reservationController.GetReservationByUserId)
	reservationRouters.POST("/add-dish/:reservation_id", reservationController.AddDishItem)
}

func (s *echoServer) initReviewRoute() {
	reviewRepository := reviewRepository.NewReviewRepository(s.db)
	reservationRepository := reservationRepository.NewReservationRepository(s.db)
	reviewService := reviewService.NewReviewService(reviewRepository, reservationRepository)
	reviewController := reviewController.NewReviewController(reviewService)

	reviewRouters := s.app.Group("/review")
	reviewRouters.Use(middlewares.JWTMiddleware())
	reviewRouters.POST("/create-review/:reservation-id", reviewController.CreateReview)
}

func (s *echoServer) initImageRoute() {
	s.app.GET("/images/:filename", func(c echo.Context) error {
		filename := c.Param("filename")
		filePath := filepath.Join("uploads", filename)

		// Check if the file exists
		if _, err := os.Stat(filePath); os.IsNotExist(err) {
			return c.JSON(http.StatusNotFound, echo.Map{
				"message": "File not found",
			})
		}

		// Serve the file for download
		return c.File(filePath)
	})
}

func (s *echoServer) initNotificationRoute() {
	notificationRepository := notificationRepository.NewNotificationRepository(s.db)
	restaurantRepository := restaurantRepository.NewRestaurantRepository(s.db)
	notificationService := notificationService.NewNotificationService(notificationRepository, restaurantRepository)
	notificationController := notificationController.NewNotificationController(notificationService)

	notificaitonRounter := s.app.Group("/notification")
	notificaitonRounter.Use(middlewares.JWTMiddleware())
	notificaitonRounter.GET("/get-user-not-read-notification", notificationController.GetUserNotReadNotification)
	notificaitonRounter.GET("/get-restaurant-not-read-notification", notificationController.GetRestaurantNotReadNotification)
	notificaitonRounter.GET("/get-all-user-notification", notificationController.GetAllUserNotification)
	notificaitonRounter.GET("/get-all-restaurant-notification", notificationController.GetAllRestaurantNotification)
}
