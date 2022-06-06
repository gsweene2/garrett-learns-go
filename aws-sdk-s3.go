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

/*
	S3ServiceInterface created for the purpose of testing
	When testing AwsSdkCountObjectsInBucket, I need to mock the `svc` parameter
	If I defined `svc` of type *s3.S3, all *s3.S3 functions need to be mocked
	By creating an interface, I can only define functions I use
	Thus, only the functions I use have to be mocked
	And an argument of type *s3.S3 would satisfy the interface
*/
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
