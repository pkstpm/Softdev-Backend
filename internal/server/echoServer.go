package server

import (
	"fmt"

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

	s.app.GET("/", func(c echo.Context) error {
		return c.String(200, "OK")
	})

	s.app.Group(s.conf.Server.Prefix)
	s.initAuthRoute()
	s.initUserRoute()
	s.initRestaurantRoute()
	s.initReservationRoute()

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
	userRouters.PUT("/edit-profile", userController.UpdateProfile)
}

func (s *echoServer) initRestaurantRoute() {
	restaurantRepository := restaurantRepository.NewRestaurantRepository(s.db)
	restaurantService := restaurantService.NewRestaurantService(restaurantRepository)
	restaurantController := restaurantController.NewRestaurantController(restaurantService)

	searchrouter := s.app.Group("/search")
	searchrouter.GET("/:name", restaurantController.FindByName)
	searchrouter.GET("/restaurant/:category", restaurantController.FindByCategory)

	restaurantRouters := s.app.Group("/restaurant")
	restaurantRouters.GET("/get-table/:restaurant_id", restaurantController.GetTable)
	restaurantRouters.Use(middlewares.JWTMiddleware())
	restaurantRouters.GET("/get-time-slot", restaurantController.GetTimeSlot)
	restaurantRouters.POST("/create-time-slot", restaurantController.UpdateTimeSlot)
	restaurantRouters.POST("/create-dish", restaurantController.CreateDish)
	restaurantRouters.POST("/create-table", restaurantController.CreateTable)
	restaurantRouters.PUT("/update-dish", restaurantController.UpdateDish)
}

func (s *echoServer) initReservationRoute() {
	reservationRepository := reservationRepository.NewReservationRepository(s.db)
	restaurantRepository := restaurantRepository.NewRestaurantRepository(s.db)
	reservationService := reservationService.NewReservationService(reservationRepository, restaurantRepository)
	reservationController := reservationController.NewReservationController(reservationService)

	reservationRouters := s.app.Group("/reservation")
	reservationRouters.Use(middlewares.JWTMiddleware())
	reservationRouters.POST("/create-reservation", reservationController.CreateReservation)
}
