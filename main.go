package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	tencentCls "github.com/tencentcloud/tencentcloud-cls-sdk-go"
	"log"
	"mrm/ent"
	"mrm/handlers"
	"os"
	"time"
)

func main() {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbDatabase := os.Getenv("DB_NAME")

	producerConfig := tencentCls.GetDefaultAsyncProducerClientConfig()
	producerConfig.Endpoint = "ap-guangzhou.cls.tencentcs.com"
	producerConfig.AccessKeyID = ""
	producerConfig.AccessKeySecret = ""

	producerInstance, err := tencentCls.NewAsyncProducerClient(producerConfig)
	if err != nil {
		fmt.Println(err)
		return
	}

	// 异步发送程序，需要启动
	producerInstance.Start()
	defer producerInstance.Close(60000)

	// try to open database 3 times
	var client *ent.Client
	//var err error
	count := 0
	for {
		client, err = ent.Open("postgres",
			fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbUser, dbPassword, dbDatabase))
		//"host=192.168.128.27 port=5432 user=postgres dbname=mydatabase password=mysecretpassword sslmode=disable")
		if err != nil {
			if count >= 3 {
				log.Fatalf("failed opening connection to postgres after 3 times try: %v", err)
			} else {
				count += 1
				log.Printf("failed opening connection to postgres, trying %d", count)
				time.Sleep(time.Second * 3)
			}
		} else {
			break
		}
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

type LogInstance struct {
	producerInstance *tencentCls.AsyncProducerClient
	TopicId          string
}

type Callback struct {
}

func (callback *Callback) Success(result *tencentCls.Result) {
	attemptList := result.GetReservedAttempts()
	for _, attempt := range attemptList {
		fmt.Printf("%+v \n", attempt)
	}
}

func (callback *Callback) Fail(result *tencentCls.Result) {
	fmt.Println(result.IsSuccessful())
	fmt.Println(result.GetErrorCode())
	fmt.Println(result.GetErrorMessage())
	fmt.Println(result.GetReservedAttempts())
	fmt.Println(result.GetRequestId())
	fmt.Println(result.GetTimeStampMs())
}
