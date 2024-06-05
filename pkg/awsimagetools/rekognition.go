package awsimagetools

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/aws/aws-sdk-go-v2/service/rekognition"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	"github.com/kingztech2019/awsimaging/internal/awsclients"
)

// RekognitionClient wraps the AWS Rekognition client
type RekognitionClient struct {
	client *rekognition.Client
}

// NewRekognitionClient creates a new RekognitionClient
func NewRekognitionClient(clients *awsclients.AWSClients) *RekognitionClient {
	return &RekognitionClient{client: clients.RekognitionClient}
}

// DetectLabels detects labels in the given image bytes
func (r *RekognitionClient) DetectLabels(imagePath string) (*rekognition.DetectLabelsOutput, error) {
	imageBytes, err := ioutil.ReadFile(imagePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read image file: %w", err)
	}

	input := &rekognition.DetectLabelsInput{
		Image: &types.Image{
			Bytes: imageBytes,
		},
		MaxLabels:     int32Ptr(10),
		MinConfidence: float32Ptr(75.0),
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
