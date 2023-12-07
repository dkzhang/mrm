package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"log"
	"mrm/ent"
	"mrm/handlers"
	"os"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")

	client, err := ent.Open("postgres",
		fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase))
	//"host=192.168.128.27 port=5432 user=postgres dbname=mydatabase password=mysecretpassword sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	defer client.Close()

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	handler := handlers.NewHandler(client)

	r := gin.Default()

	r.GET("/rooms", handler.GetRooms)
	r.POST("/room", handler.CreateRoom)
	r.PUT("/room/:id", handler.UpdateRoom)
	r.DELETE("/room/:id", handler.DeleteRoom)

	r.POST("/allocate", handler.Allocate)
	r.GET("/meetings/date/:date", handler.QueryAllocateByDate)
	r.GET("/meetings/room/:id", handler.QueryAllocateByRoom)
	r.GET("meeting/:id", handler.QueryMeeting)
	r.GET("/svg/:date", handler.SVG)

	r.DELETE("/meeting/:id", handler.DeleteMeeting)

	r.Run(":8080")
}
