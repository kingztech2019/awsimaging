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

### Setting Up AWS Credentials

The AWS SDK for Go automatically searches for credentials using the default credential provider chain, which includes environment variables, shared credentials files, and EC2 instance metadata.

Using Environment Variables
Set the environment variables in your shell:

```sh
export AWS_ACCESS_KEY_ID=your-access-key-id
export AWS_SECRET_ACCESS_KEY=your-secret-access-key
export AWS_REGION=your-region

```
### Using Shared Credentials File

Edit or create the ~/.aws/credentials file with the following content:

```ini
[default]
aws_access_key_id = your-access-key-id
aws_secret_access_key = your-secret-access-key

```
Edit or create the ~/.aws/config file with the following content:

```ini
[default]
region = your-region

```
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

     region := "us-west-2"
    // Initialize AWS clients
	awsClients, err := awsclients.NewAWSClients(region)
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
    region := "us-west-2"
	filePath := "path/to/your/image.jpg"

	// Initialize AWS clients
	awsClients, err := awsclients.NewAWSClients(region)
	if err != nil {
		log.Fatalf("failed to create AWS clients: %v", err)
	}

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
```

### Upload Image to S3

Use the S3 uploader to upload a base64 encoded image to an S3 bucket and get the URL of the uploaded image.

```go
package main

import (
    "encoding/base64"
    "log"
    "github.com/kingztech2019/awsimaging/awsclients"
	"github.com/kingztech2019/awsimaging/pkg/awsimagetools"
)

func main() {
    region := "us-west-2"
	bucketName := "awsimaging"
	imageName := "your-image-name.jpg"
	base64Image := "Your base64 Image"

	// Initialize AWS clients
	awsClients, err := awsclients.NewAWSClients(region)
	if err != nil {
		log.Fatalf("failed to create AWS clients: %v", err)
	}

    s3Uploader := awsimagetools.NewS3Uploader(awsClients)
    s3Url, err := s3Uploader.UploadToS3(base64Image, bucketName, imageName)
    if err != nil {
        log.Fatalf("failed to upload image to S3: %v", err)
    }

    log.Printf("S3 Upload URL: %s", s3Url)
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

   region := "us-west-2"
   base64Image := "your-base64-encoded-image-string"

   // Initialize AWS clients
	awsClients, err := awsclients.NewAWSClients(region)
	if err != nil {
		log.Fatalf("failed to create AWS clients: %v", err)
	}

    textractClient := awsimagetools.NewTextractClient(awsClients)
    text, err := textractClient.ExtractText(base64Image)
    if err != nil {
        log.Fatalf("failed to extract text: %v", err)
    }

    log.Printf("Extracted Text: %s", text)
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