package s3service

import (
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

//ClientInterface ...
type ClientInterface interface {
	Upload(filename string, file io.Reader) (string, error)
}

type AwsS3ServiceClient struct {
	bucketName string
	regionName string
	accessKey  string
	secretKey  string
	client     *http.Client
}

func NewAwsS3ServiceClient(bucketName, regionName, accessKey, secretKey string) (*AwsS3ServiceClient, error) {
	client := &http.Client{
		Transport: &http.Transport{
			MaxIdleConns:    10,
			IdleConnTimeout: 30 * time.Second,
		},
		Timeout: 3 * time.Second,
	}

	return &AwsS3ServiceClient{bucketName: bucketName, regionName: regionName, accessKey: accessKey, secretKey: secretKey, client: client}, nil
}

func (awss3 *AwsS3ServiceClient) Upload(filename string, file io.Reader) (string, error) {

	sess, err := session.NewSession(
		&aws.Config{
			Region: aws.String(awss3.regionName),
			Credentials: credentials.NewStaticCredentials(
				awss3.accessKey,
				awss3.secretKey,
				"", // a token will be created when the session it's used.
			),
		})

	if err != nil {
		panic(err)
	}

	uploader := s3manager.NewUploader(sess)
	//upload to the s3 bucket
	uploadResult, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(awss3.bucketName),
		ACL:    aws.String("private"),
		Key:    aws.String(filename),
		Body:   file,
	})
	if err != nil {
		log.Println("Failed to upload file: ", err)
		return "", errors.New("FAILED TO UPLOAD FILESs")
	}
	return uploadResult.Location, nil
}
