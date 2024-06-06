package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kingztech2019/awsimaging/awsclients"
	"github.com/kingztech2019/awsimaging/pkg/awsimagetools"
)

func main() {
	region := "us-west-2"
	bucketName := "awsimaging"
	imageName := "your-image-name.jpg"
	filePath := "image.jpeg"

	// Initialize AWS clients
	awsClients, err := awsclients.NewAWSClients(region)
	if err != nil {
		log.Fatalf("failed to create AWS clients: %v", err)
	}

	base64Image := "Base64 image"

	// Textract
	textractClient := awsimagetools.NewTextractClient(awsClients)
	text, err := textractClient.ExtractText(base64Image)
	if err != nil {
		log.Fatalf("failed to extract text: %v", err)
	}
	fmt.Printf("Extracted Text: %s\n", text)

	// S3 Upload
	s3Uploader := awsimagetools.NewS3Uploader(awsClients)

	s3Url, err := s3Uploader.UploadToS3(base64Image, bucketName, imageName)
	if err != nil {
		log.Fatalf("failed to upload image to S3: %v", err)
	}

	s3UrlJSON, err := json.Marshal(s3Url)
	if err != nil {
		log.Fatalf("failed to encode S3 URL to JSON: %v", err)
	}
	fmt.Printf("S3 Upload URL JSON: %s\n", string(s3UrlJSON))

	// Rekognition
	rekognitionClient := awsimagetools.NewRekognitionClient(awsClients)
	rekResult, err := rekognitionClient.DetectLabels(filePath)
	if err != nil {
		log.Fatalf("failed to detect labels: %v", err)
	}

	rekResultJSON, err := json.Marshal(rekResult)
	if err != nil {
		log.Fatalf("failed to encode Rekognition result to JSON: %v", err)
	}
	fmt.Printf("Rekognition Result JSON: %s\n", string(rekResultJSON))

}
