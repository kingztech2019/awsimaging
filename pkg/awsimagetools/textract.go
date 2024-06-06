package awsimagetools

import (
	"context"
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/textract"
	"github.com/aws/aws-sdk-go-v2/service/textract/types"
	"github.com/kingztech2019/awsimaging/awsclients"
)

// TextractClient wraps the AWS Textract client
type TextractClient struct {
	client *textract.Client
}

// NewTextractClient creates a new TextractClient
func NewTextractClient(clients *awsclients.AWSClients) *TextractClient {
	return &TextractClient{client: clients.TextractClient}
}

// ExtractText extracts text from a base64 encoded image
func (t *TextractClient) ExtractText(base64Image string) (string, error) {
	imageBytes, err := base64.StdEncoding.DecodeString(base64Image)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 image: %w", err)
	}

	input := &textract.DetectDocumentTextInput{
		Document: &types.Document{
			Bytes: imageBytes,
		},
	}

	result, err := t.client.DetectDocumentText(context.TODO(), input)
	if err != nil {
		return "", err
	}

	var extractedText strings.Builder
	for _, block := range result.Blocks {
		if block.BlockType == types.BlockTypeLine {
			extractedText.WriteString(*block.Text)
			extractedText.WriteString("\n")
		}
	}

	return extractedText.String(), nil
}
