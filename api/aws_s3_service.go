package api

import (
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"net/http"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"os"
)

func Respond(w http.ResponseWriter, data ApiResponse) {
	w.Header().Add("Content-Type", "application/json") // key is added, which tells that Content-Type of the response is application/json
	json.NewEncoder(w).Encode(data)                    // response is appended after the conversion of data from struct to JSON
}

func AwsBucketList(ClientId string, ClientSecret string, Region string) (*s3.ListBucketsOutput, error) {

	err := os.Setenv("AWS_ACCESS_KEY_ID", ClientId)
	if err != nil {
		log.Printf("Error while setting env variable, %v ", err)
		return nil, err
	}
	err = os.Setenv("AWS_SECRET_KEY_ID", ClientSecret)
	if err != nil {
		log.Printf("Error while setting env variable, %v ", err)
		return nil, err
	}
	err = os.Setenv("REGION_NAME", Region)
	if err != nil {
		log.Printf("Error while setting env variable, %v ", err)
		return nil, err
	}

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(ClientId, ClientSecret, ""),
	})
	svc := s3.New(sess)

	input := &s3.ListBucketsInput{}

	result, err := svc.ListBuckets(input)

	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return result, err
}

func AwsCreateBucket(ClientId string, ClientSecret string, Region string) (*s3.CreateBucketOutput, error) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(ClientId, ClientSecret, ""),
	})
	svc := s3.New(sess)

	input := &s3.CreateBucketInput{
		Bucket: aws.String("ansh-bucket"),
		CreateBucketConfiguration: &s3.CreateBucketConfiguration{
			LocationConstraint: aws.String("ap-south-1"),
		},
	}
	result, err := svc.CreateBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return result, err
}

func AwsDeleteBucket(ClientId string, ClientSecret string, Region string) (*s3.DeleteBucketOutput, error) {

	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String(Region),
		Credentials: credentials.NewStaticCredentials(ClientId, ClientSecret, ""),
	})
	svc := s3.New(sess)

	input := &s3.DeleteBucketInput{
		Bucket: aws.String("ansh-bucket"),
	}
	result, err := svc.DeleteBucket(input)
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case s3.ErrCodeBucketAlreadyExists:
				fmt.Println(s3.ErrCodeBucketAlreadyExists, aerr.Error())
			case s3.ErrCodeBucketAlreadyOwnedByYou:
				fmt.Println(s3.ErrCodeBucketAlreadyOwnedByYou, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			fmt.Println(err.Error())
		}
	}
	return result, err
}
