package controllers

import (
	"blog-apis/api/models"
	"blog-apis/api/utils"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) CreatePost(r *gin.Context) {

	var newPost models.Post

	err := json.NewDecoder(r.Request.Body).Decode(&newPost)
	if err != nil {
		utils.ErrorResponse(r, http.StatusBadRequest, "Unable to decode the body")
		return
	}
	post, err := s.service.Repo.CreatePost(newPost)
	if err != nil {
		log.Println("Error while creating new post: ", err)
		utils.ErrorResponse(r, http.StatusBadRequest, "Unable to create new post")
		return
	}
	utils.OkResponse(r, post)
}

func (s *HTTPHandler) GetAllPost(r *gin.Context) {
	allPosts, err := s.service.Repo.GetAllPost()
	if err != nil {
		log.Println("Error while getting all posts: ", err)
		utils.ErrorResponse(r, http.StatusBadRequest, "Unable to get all posts")
		return
	}
	utils.OkResponse(r, allPosts)
}

func (s *HTTPHandler) UploadPostFile(r *gin.Context) {
	file, header, err := r.Request.FormFile("files")
	if err != nil {
		log.Println("Error while uploading files: ", err)
		utils.ErrorResponse(r, http.StatusBadRequest, "Error while uploading files")
		return
	}
	filename := header.Filename

	filepath, err := s.service.AwsS3Service.Upload(filename, file)
	if err != nil {
		log.Println("Error while uploading files: ", err)
		utils.ErrorResponse(r, http.StatusBadRequest, "Error while uploading files")
		return
	}

	type response struct {
		ImageURL string `json:"imageURL"`
	}

	utils.OkResponse(r, response{ImageURL: filepath})
}
