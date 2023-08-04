package main

import (
	"app/database"
	"app/handler"
	"app/middleware"
	"app/repository"
	"app/service"
	"app/transaction"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db := database.Mariadb()
	defer db.Close()
	conn := database.MongoDB()
	defer conn.Client().Disconnect(context.Background())

	r := repository.NewRepository(db)
	t := transaction.NewTransaction(conn)
	s := service.NewService(r, t)
	h := handler.NewHandler(s)

	router := gin.Default()

	user := router.Group("/", middleware.UserAuth)
	user.GET("/get", h.Get)

	admin := router.Group("/", middleware.AdminAuth)
	admin.POST("/add", h.Add)
	admin.PUT("/update/:id", h.Update)
	admin.DELETE("/delete/:id", h.Delete)

	if err := router.Run(":9000"); err != nil {
		log.Fatal(err.Error())
	}
}
