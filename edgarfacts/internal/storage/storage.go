// Declare Package
package storage

// Make Imports
import (
	"bytes"
	"context"
	"io"
	"time"

	"cloud.google.com/go/storage"
)

// Upload Bytes
func UploadBytes(data []byte, bucket, path string) error {
	// Create Buffer
	buffer := bytes.NewBuffer(data)

	// Create Client
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	// Create Writer
	ctx, cancel := context.WithTimeout(ctx, 120*time.Second)
	defer cancel()
	writer := client.Bucket(bucket).Object(path).NewWriter(ctx)

	// Copy Data from Buffer to Google Cloud Storage
	_, err = io.Copy(writer, buffer)
	if err != nil {
		return err
	}

	// Close Writer
	err = writer.Close()
	if err != nil {
		return err
	}

	// Return Result
	return nil

}
