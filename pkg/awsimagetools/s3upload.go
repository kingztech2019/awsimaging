package awsimagetools

import (
	"context"
	"encoding/base64"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/kingztech2019/awsimaging/awsclients"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"

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

// This `UploadToS3` function in the `S3Uploader` struct is responsible for uploading an image to an
// AWS S3 bucket. Here is a breakdown of what the function does:
func (u *S3Uploader) UploadToS3(base64Image, bucketName, imageName, region, accessKeyID, secretAccessKey string) (*UploadResponse, error) {
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

	// Get the bucket's region
	regionOutput, err := u.client.GetBucketLocation(context.TODO(), &s3.GetBucketLocationInput{
		Bucket: aws.String(bucketName),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get bucket location: %w", err)
	}

	// Determine the endpoint based on the bucket's region
	bucketRegion := string(regionOutput.LocationConstraint)
	if bucketRegion == "" {
		bucketRegion = "us-east-1" // Default region
	}
	log.Println(bucketRegion, accessKeyID, secretAccessKey)

	// Create a new S3 client with the correct region
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(bucketRegion),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, "")),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create S3 client with correct region: %w", err)
	}
	u.client = s3.NewFromConfig(cfg)

	// Perform the upload
	_, err = u.client.PutObject(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to upload image to S3: %w", err)
	}

	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", bucketName, bucketRegion, imageName)
	return &UploadResponse{URL: url}, nil
}
