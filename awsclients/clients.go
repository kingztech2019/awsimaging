package awsclients

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/textract"
)

// AWSClients holds the AWS service clients
type AWSClients struct {
	RekognitionClient *rekognition.Client
	S3Client          *s3.Client
	TextractClient    *textract.Client
}

// The `NewAWSClients` function creates and returns AWS clients for Rekognition, S3, and Textract
// services using the provided region, access key ID, and secret access key.
func NewAWSClients(region, accessKeyID, secretAccessKey string) (*AWSClients, error) {
	staticCredentials := aws.NewCredentialsCache(credentials.NewStaticCredentialsProvider(accessKeyID, secretAccessKey, ""))

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithCredentialsProvider(staticCredentials),
	)
	if err != nil {
		return nil, fmt.Errorf("unable to load SDK config, %v", err)
	}

	rekognitionSvc := rekognition.NewFromConfig(cfg)
	s3Svc := s3.NewFromConfig(cfg)
	textractSvc := textract.NewFromConfig(cfg)

	return &AWSClients{
		RekognitionClient: rekognitionSvc,
		S3Client:          s3Svc,
		TextractClient:    textractSvc,
	}, nil
}
