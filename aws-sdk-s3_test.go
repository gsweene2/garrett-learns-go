package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
)

type mockS3Svc struct{}

func (m *mockS3Svc) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	name := "someBucket"
	return &s3.ListObjectsV2Output{
		Name: &name,
	}, nil
}

func TestAwsSdkCountObjectsInBucket(t *testing.T) {
	// Arrange
	// Create mockS3Svc & ListObjectsV2
	want := 1

	// Act
	mockSvc := mockS3Svc{}
	got, err := AwsSdkCountObjectsInBucket(&mockSvc, "someBucket")

	// Assert
	if err != nil {
		t.Error("Error returned")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
