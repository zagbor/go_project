package services

import (
	"context"
	"log"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// IntegrationService handles interaction with external services like MinIO.
type IntegrationService struct {
	minioClient *minio.Client
}

// NewIntegrationService creates a new integration service.
func NewIntegrationService(endpoint, accessKeyID, secretAccessKey string) *IntegrationService {
	// Initialize MinIO client object.
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: false, // Set to true if using HTTPS
	})
	if err != nil {
		log.Printf("Error initializing MinIO client: %v", err)
		// We don't panic here to allow the service to run even if MinIO isn't configured for the homework
		return &IntegrationService{minioClient: nil}
	}

	return &IntegrationService{minioClient: minioClient}
}

// UploadFile uploads a file to the configured bucket (example method).
func (s *IntegrationService) UploadFile(ctx context.Context, bucketName, objectName, filePath string) error {
	if s.minioClient == nil {
		log.Println("MinIO client not initialized")
		return nil
	}
	
	// Create bucket if it doesn't exist
	exists, err := s.minioClient.BucketExists(ctx, bucketName)
	if err != nil {
		return err
	}
	if !exists {
		err = s.minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{})
		if err != nil {
			return err
		}
	}

	// Upload the file
	info, err := s.minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}

	log.Printf("Successfully uploaded %s of size %d\n", objectName, info.Size)
	return nil
}
