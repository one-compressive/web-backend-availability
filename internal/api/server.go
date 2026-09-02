package api

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/one-compressive/web-backend-availability/internal/app/handler"
	"github.com/one-compressive/web-backend-availability/internal/app/repository"
	"github.com/sirupsen/logrus"
)

func StartServer() {
	log.Println("Starting server")

	repo, err := repository.NewRepository()
	if err != nil {
		logrus.Error("ошибка инициализации репозитория")
		return
	}

	h := handler.NewHandler(repo)

	r := gin.Default()
	r.LoadHTMLGlob("./templates/*")
	r.Static("/static", "./static")

	r.GET("/", h.GetComponents)
	r.GET("/components", h.GetComponents)
	r.GET("/component", h.GetComponent)
	r.GET("/component/:id", h.GetComponent)
	r.GET("/add_component", h.AddComponent)

	addr := ":8081"
	log.Printf("Listening on http://localhost%s", addr)
	if err := r.Run(addr); err != nil {
		logrus.Error(err)
	}

	log.Println("Server down")
}
