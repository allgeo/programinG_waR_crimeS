package models

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//DBRepository ...
type DBRepository struct {
	client *mongo.Client
	DB     *mongo.Database
}

//NewDBRepository ...
func NewDBRepository(client *mongo.Client, db *mongo.Database) DBRepositoryInterface {
	return &DBRepository{client, db}
}

type DBRepositoryInterface interface {
	// Post
	CreatePost(post Post) (Post, error)
	GetAllPost() ([]Post, error)
}
