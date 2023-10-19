package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:@tcp(localhost:3306)/bwastartup?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	authService := auth.NewService()
	userHandler := handler.NewUserHandler(userService, authService)

	fmt.Println(authService.GenerateToken(1001))
	/*
		input := user.LoginInput{
			Email:    "Pos2w@gmail.com",
			Password: "password123",
		}

		user, err := userService.Login(input)




		if err != nil {
			fmt.Println("Terjadi Kesalahan")
			fmt.Println(err.Error())
		}

		fmt.Println("Name : " + user.Name)
		fmt.Println("Email :" + user.Email)
	*/

	userService.SaveAvatar(1, "images/test.jpg")
	router := gin.Default()
	api := router.Group("/api/v1")
	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailibility)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()

}
