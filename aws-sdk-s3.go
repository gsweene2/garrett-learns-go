package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func GetS3Client() *s3.S3 {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := s3.New(sess)

	return svc
}

// Created soely for the purpose of testing
// During the test, I need to mock the object passed into AwsSdkCountObjectsInBucket
// If it is *s3.S3, all functions would need to be mocket
// With a defined interface, only the method used needs to be mocked
// As long as the object of type *s3.S3 is passed implements ListObjectsV2, this works
type S3ServiceInterface interface {
	ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
}

func AwsSdkCountObjectsInBucket(svc S3ServiceInterface, bucket string) (int, error) {

	resp, err := svc.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: &bucket})

	if err != nil {
		fmt.Printf("Error Listing Objects: %v\n", err)
		return -1, err
	}

	fmt.Printf("len(resp.Contents): %v\n", len(resp.Contents))
	return len(resp.Contents), nil
}

func main() {
	svc := GetS3Client()
	bucket := "garretts-s3-bucket-name"
	AwsSdkCountObjectsInBucket(svc, bucket)
}
