package storage

import (
    "context"
    "fmt"
    "os"

    "github.com/minio/minio-go/v7"
    "github.com/minio/minio-go/v7/pkg/credentials"
)

var MinioClient *minio.Client

func InitMinio() error {
    endpoint := os.Getenv("MINIO_ENDPOINT")
    accessKey := os.Getenv("MINIO_ROOT_USER")
    secretKey := os.Getenv("MINIO_ROOT_PASSWORD")

    client, err := minio.New(endpoint, &minio.Options{
        Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
        Secure: false,
    })

    if err != nil {
        return fmt.Errorf("failed to init MinIO client: %w", err)
    }

    MinioClient = client

    // Just a test call
    _, err = client.ListBuckets(context.Background())
    if err != nil {
        return fmt.Errorf("failed to list buckets (MinIO not ready?): %w", err)
    }

    fmt.Println("✅ Connected to MinIO")
    return nil
}

