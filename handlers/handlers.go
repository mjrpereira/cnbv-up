package handlers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB     *gorm.DB
	Router *gin.Engine
}

var server = Server{}

func Run() {
	server.initializeDB()
	server.initializeRouter()
	server.Run(":8080")

}

func (s *Server) Run(addr string) {
	log.Fatal(
		http.ListenAndServe(addr, server.Router),
	)
}

func (s *Server) initializeDB() {

	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	s.DB = db

}

func (s *Server) initializeRouter() {

	s.Router = gin.Default()

	s.Router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Pong",
		})
	})

	v1 := s.Router.Group("/api/v1")
	{
		v1.POST("/users", s.CreateUser)
	}
}
