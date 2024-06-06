package awsimagetools

import (
	"context"
	"encoding/base64"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/kingztech2019/awsimaging/awsclients"

	"strings"
)

// S3Uploader wraps the AWS S3 client
type S3Uploader struct {
	client *s3.Client
}

// NewS3Uploader creates a new S3Uploader
func NewS3Uploader(clients *awsclients.AWSClients) *S3Uploader {
	return &S3Uploader{client: clients.S3Client}
}

// UploadResponse is the structured response for the S3 upload function
type UploadResponse struct {
	URL string `json:"url"`
}

// UploadToS3 uploads a base64 encoded image to an S3 bucket and returns the URL
func (u *S3Uploader) UploadToS3(base64Image, bucketName, imageName string) (*UploadResponse, error) {
	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 image: %w", err)
	}

	input := &s3.PutObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(imageName),
		Body:   strings.NewReader(string(imageBytes)),
		ACL:    types.ObjectCannedACLPublicRead,
	}

	// Perform the upload using the existing client
	_, err = u.client.PutObject(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image to S3: %w", err)
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, u.client.Options().Region, imageName)
	return &UploadResponse{URL: url}, nil
}
