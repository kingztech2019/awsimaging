package awsimagetools

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/kingztech2019/awsimaging/internal/awsclients"
)

// TextractClient wraps the AWS Textract client
type TextractClient struct {
	client *textract.Client
}

// NewTextractClient creates a new TextractClient
func NewTextractClient(clients *awsclients.AWSClients) *TextractClient {
	return &TextractClient{client: clients.TextractClient}
}

// ExtractTextResponse is the structured response for the Textract function
type ExtractTextResponse struct {
	Text string `json:"text"`
}

// This function `ExtractTextFromImage` is a method of the `TextractClient` struct. It takes a base64
// encoded image as input, decodes it, and then uses the AWS Textract service to extract text from the
// image.
func (t *TextractClient) ExtractTextFromImage(base64Image string) (*ExtractTextResponse, error) {
	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return nil, fmt.Errorf("failed to decode base64 image: %w", err)
	}

	input := &textract.DetectDocumentTextInput{
		Document: &types.Document{
			Bytes: imageBytes,
		},
	}

	result, err := t.client.DetectDocumentText(context.TODO(), input)
	if err != nil {
		return nil, fmt.Errorf("failed to extract text: %w", err)
	}

	var extractedText strings.Builder
	for _, block := range result.Blocks {
		if block.BlockType == types.BlockTypeLine {
			extractedText.WriteString(*block.Text)
			extractedText.WriteString("\n")
		}
	}

	return &ExtractTextResponse{Text: extractedText.String()}, nil
}
