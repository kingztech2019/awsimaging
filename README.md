# awsimaging

`awsimaging` is a Go package that provides a convenient interface for interacting with AWS services such as Rekognition, S3, and Textract. This package allows developers to easily detect labels in images, upload images to S3, and extract text from images using Textract.

## Features

- **Rekognition**: Detect labels in images.
- **S3**: Upload base64 encoded images to S3 and get the URL of the uploaded image.
- **Textract**: Extract text from base64 encoded images.

## Installation

This README covers the setup, usage, and functionality of the package, including examples.

```markdown
# awsimaging

`awsimaging` is a Go package that provides a convenient interface for interacting with AWS services such as Rekognition, S3, and Textract. This package allows developers to easily detect labels in images, upload images to S3, and extract text from images using Textract.

## Features

- **Rekognition**: Detect labels in images.
- **S3**: Upload base64 encoded images to S3 and get the URL of the uploaded image.
- **Textract**: Extract text from base64 encoded images.

## Installation

To install the `awsimaging` package, run:

```sh
go get github.com/kingztech2019/awsimaging
```

## Configuration

Before using the package, you need to set up AWS credentials and specify the AWS region.

## AWS Credentials

You can provide AWS credentials via environment variables or by directly passing them when creating the `AWSClients`.

## Folder Structure

The package is organized as follows:

```
awsimagetools/
├──awsclients/
│       └── clients.go
├── awsimagetools.go
├── rekognition.go
├── s3upload.go
├── textract.go
└── main.go
```

- `internal/awsclients/clients.go`: Contains the code to initialize AWS service clients.
- `rekognition.go`: Contains functions to interact with AWS Rekognition.
- `s3upload.go`: Contains functions to upload images to AWS S3.
- `textract.go`: Contains functions to extract text from images using AWS Textract.
- `main.go`: Example usage of the package.

## Usage

### Initialize AWS Clients

To use the package, you first need to initialize the AWS clients. You can do this by providing the AWS region and credentials.

```go
package main

import (
    "log"
   "github.com/kingztech2019/awsimaging/awsclients"
	"github.com/kingztech2019/awsimaging/pkg/awsimagetools"
)

func main() {
    awsClients, err := awsclients.NewAWSClients("us-west-2", "your-access-key-id", "your-secret-access-key")
    if err != nil {
        log.Fatalf("failed to create AWS clients: %v", err)
    }

    // Use the AWS clients...
}
```

### Detect Labels with Rekognition

Use the Rekognition client to detect labels in an image.

```go
package main

import (
    "fmt"
    "log"
    "github.com/kingztech2019/awsimaging/awsclients"
	"github.com/kingztech2019/awsimaging/pkg/awsimagetools"
)

func main() {
    awsClients, err := awsclients.NewAWSClients("us-west-2", "your-access-key-id", "your-secret-access-key")
    if err != nil {
        log.Fatalf("failed to create AWS clients: %v", err)
    }

    rekClient := awsimagetools.NewRekognitionClient(awsClients)
    result, err := rekClient.DetectLabels("path/to/your/image.jpg")
    if err != nil {
        log.Fatalf("failed to detect labels: %v", err)
    }
    for _, label := range result.Labels {
        fmt.Printf("Label: %s, Confidence: %.2f\n", *label.Name, *label.Confidence)
    }
}
```

### Upload Image to S3

Use the S3 uploader to upload a base64 encoded image to an S3 bucket and get the URL of the uploaded image.

```go
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
	accessKeyID := "your-access-key-id"
	secretAccessKey := "your-secret-access-key"
	bucketName := "your-s3-bucket-name"
	imageName := "your-image-name.jpg"
	base64Image := "your-base64-encoded-image-string"

	awsClients, err := awsclients.NewAWSClients(region, accessKeyID, secretAccessKey)
	if err != nil {
		log.Fatalf("failed to create AWS clients: %v", err)
	}

	s3Uploader := awsimagetools.NewS3Uploader(awsClients)
	s3Url, err := s3Uploader.UploadToS3(base64Image, bucketName, imageName, region, accessKeyID, secretAccessKey)
	if err != nil {
		log.Fatalf("failed to upload image to S3: %v", err)
	}

	s3UrlJSON, err := json.Marshal(s3Url)
	if err != nil {
		log.Fatalf("failed to encode S3 URL to JSON: %v", err)
	}
	fmt.Printf("S3 Upload URL JSON: %s\n", string(s3UrlJSON))
}

```

### Extract Text with Textract

Use the Textract client to extract text from a base64 encoded image.

```go
package main

import (
    "fmt"
    "log"
   "github.com/kingztech2019/awsimaging/awsclients"
	"github.com/kingztech2019/awsimaging/pkg/awsimagetools"
)

func main() {
    awsClients, err := awsclients.NewAWSClients("us-west-2", "your-access-key-id", "your-secret-access-key")
    if err != nil {
        log.Fatalf("failed to create AWS clients: %v", err)
    }

    textractClient := awsimagetools.NewTextractClient(awsClients)
    base64Image := "your-base64-encoded-image-string"
    extractedText, err := textractClient.ExtractTextFromImage(base64Image)
    if err != nil {
        log.Fatalf("failed to extract text from image: %v", err)
    }
    fmt.Printf("Extracted Text:\n%s\n", extractedText)
}
```

## Functions

### Rekognition Functions

- `NewRekognitionClient(clients *awsclients.AWSClients) *RekognitionClient`: Creates a new Rekognition client.
- `(*RekognitionClient) DetectLabels(imagePath string) (*rekognition.DetectLabelsOutput, error)`: Detects labels in the given image.

### S3 Functions

- `NewS3Uploader(clients *awsclients.AWSClients) *S3Uploader`: Creates a new S3 uploader.
- `(*S3Uploader) UploadToS3(base64Image, bucketName, imageName string) (string, error)`: Uploads a base64 encoded image to S3 and returns the URL.

### Textract Functions

- `NewTextractClient(clients *awsclients.AWSClients) *TextractClient`: Creates a new Textract client.
- `(*TextractClient) ExtractTextFromImage(base64Image string) (string, error)`: Extracts text from a base64 encoded image.

## Contributing

Contributions are welcome! Please submit a pull request or open an issue to discuss any changes.

## License

This project is licensed under the MIT License.
```

This README provides comprehensive details about the package, including installation, usage, and function descriptions. It also includes example code to help users understand how to use the package effectively.