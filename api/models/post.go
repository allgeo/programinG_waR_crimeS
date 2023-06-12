package models

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	ID       primitive.ObjectID `json:"id" bson:"_id"`
	Caption  string             `json:"caption" bson:"caption"`
	ImageURL string             `json:"imageURL" bson:"imageURL"`
}

const (
	POST_COLLECTION = "post"
)

func (r *DBRepository) CreatePost(post Post) (Post, error) {
	collection := r.DB.Collection(POST_COLLECTION)
	post.ID = primitive.NewObjectID()
	_, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (r *DBRepository) GetAllPost() ([]Post, error) {
	posts := []Post{}
	collection := r.DB.Collection(POST_COLLECTION)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var elem Post
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		posts = append(posts, elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
		return []Post{}, err
	}
	return posts, nil
}
