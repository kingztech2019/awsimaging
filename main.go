package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/kingztech2019/awsimaging/awsclients"
	"github.com/kingztech2019/awsimaging/pkg/awsimagetools"
)

func main() {
	awsClients, err := awsclients.NewAWSClients("us-west-2", "Your AWS-Access key", "Your AWS-secret-key")
	if err != nil {
		log.Fatalf("failed to create AWS clients: %v", err)
	}

	// Rekognition
	rekClient := awsimagetools.NewRekognitionClient(awsClients)
	rekResult, err := rekClient.DetectLabels("image.jpeg")
	if err != nil {
		log.Fatalf("failed to detect labels: %v", err)
	}

	rekResultJSON, err := json.Marshal(rekResult)
	if err != nil {
		log.Fatalf("failed to encode Rekognition result to JSON: %v", err)
	}
	fmt.Printf("Rekognition Result JSON: %s\n", string(rekResultJSON))

	// S3 Upload
	s3Uploader := awsimagetools.NewS3Uploader(awsClients)

	s3Url, err := s3Uploader.UploadToS3("base64 image", "bucket_name", "image_name", "region-name", "access-key", "secret-key")
	if err != nil {
		log.Fatalf("failed to upload image to S3: %v", err)
	}

	s3UrlJSON, err := json.Marshal(s3Url)
	if err != nil {
		log.Fatalf("failed to encode S3 URL to JSON: %v", err)
	}
	fmt.Printf("S3 Upload URL JSON: %s\n", string(s3UrlJSON))

	// Textract
	textractClient := awsimagetools.NewTextractClient(awsClients)
	extractedText, err := textractClient.ExtractTextFromImage("your base64 image")
	if err != nil {
		log.Fatalf("failed to extract text from image: %v", err)
	}

	extractedTextJSON, err := json.Marshal(extractedText)
	if err != nil {
		log.Fatalf("failed to encode extracted text to JSON: %v", err)
	}
	fmt.Printf("Extracted Text JSON: %s\n", string(extractedTextJSON))
}
