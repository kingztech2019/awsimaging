package awsimagetools

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/kingztech2019/awsimaging/awsclients"
)

// RekognitionClient wraps the AWS Rekognition client
type RekognitionClient struct {
	client *rekognition.Client
}

// NewRekognitionClient creates a new RekognitionClient
func NewRekognitionClient(clients *awsclients.AWSClients) *RekognitionClient {
	return &RekognitionClient{client: clients.RekognitionClient}
}

// The `DetectLabels` method in the `RekognitionClient` struct is a function that uses the AWS
// Rekognition client to detect labels in an image file. Here's a breakdown of what it does:
func (r *RekognitionClient) DetectLabels(filePath string) (*rekognition.DetectLabelsOutput, error) {
	imageBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	input := &rekognition.DetectLabelsInput{
		Image: &types.Image{
			Bytes: imageBytes,
		},
		MaxLabels:     int32Ptr(10),
		MinConfidence: float32Ptr(10.0),
	}

	result, err := r.client.DetectLabels(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to detect labels: %w", err)
	}

	return result, nil
}

func int32Ptr(i int) *int32 {
	i32 := int32(i)
	return &i32
}

func float32Ptr(f float64) *float32 {
	f32 := float32(f)
	return &f32
}
