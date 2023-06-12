package api

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	s3 "blog-apis/api/clients/s3service"
	"blog-apis/api/controllers"
	"blog-apis/api/middlewares"
	"blog-apis/api/models"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("sad .env file found")
	}
}

func GetMongoSession() (*mongo.Client, *mongo.Database) {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if os.Getenv("DB_NAME") == "" {
		log.Panic("db name is empty")
	}
	dburl := os.Getenv("MONGODB_URL")
	if dburl == "" {
		log.Panic("mongo db url is empty")
	}
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(dburl))

	DB := client.Database(os.Getenv("DB_NAME"))
	if err != nil {
		fmt.Printf("Cannot connect to monog database")
		log.Fatal("This is the error connecting to mongo:", err)
	} else {
		fmt.Printf("We are connected to the monog database")
	}
	return client, DB
}

func getServicer() *controllers.Service {

	bucketName := os.Getenv("AWS_BUCKET_NAME")
	regionName := os.Getenv("AWS_REGION_NAME")
	accessKey := os.Getenv("AWS_ACCESS_KEY")
	secretKey := os.Getenv("AWS_SECRET_KEY")

	if bucketName == "" || regionName == "" || accessKey == "" || secretKey == "" {
		log.Panic("Either AWS_BUCKET_NAME/AWS_REGION_NAME/AWS_ACCESS_KEY/AWS_SECRET_KEY is empty")

	}

	awss3client, err := s3.NewAwsS3ServiceClient(bucketName, regionName, accessKey, secretKey)

	if err != nil {
		log.Panic("Error while create aws s3 client")
	}

	weatherBaseURL := os.Getenv("WEATHER_BASE_URL")
	weatherApiKey := os.Getenv("WEATHER_API_KEY")

	if weatherBaseURL == "" || weatherApiKey == "" {
		log.Panic("Either WEATHER_BASE_URL/WEATHER_API_KEY is empty")

	}

	client, db := GetMongoSession()
	r := models.NewDBRepository(client, db)

	router := gin.Default()
	router.Use(middlewares.CORSMiddleware())

	return controllers.NewService(awss3client, r, router)

}

func Run() {
	server := getServicer()
	controllers.NewHttpHandler(server)
	apiPort := fmt.Sprintf(":%s", os.Getenv("API_PORT"))
	fmt.Printf("Listening to port %s", apiPort)

	server.Run(apiPort)
}
