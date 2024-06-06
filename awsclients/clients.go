package awsclients

import (
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/textract"
)

type AWSClients struct {
	RekognitionClient *rekognition.Client
	S3Client          *s3.Client
	TextractClient    *textract.Client
}

// The NewAWSClients function creates and returns AWS clients for Rekognition, S3, and Textract
// services in the specified region.
func NewAWSClients(region string) (*AWSClients, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
	)
	log.Println(cfg.Credentials)
	if err != nil {
		return nil, fmt.Errorf("failed to load configuration: %w", err)
	}

	return &AWSClients{
		RekognitionClient: rekognition.NewFromConfig(cfg),
		S3Client:          s3.NewFromConfig(cfg),
		TextractClient:    textract.NewFromConfig(cfg),
	}, nil
}
