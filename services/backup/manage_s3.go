package backup

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (db *DatabaseBackup) UploadToS3(zipPath string) (*S3Backup, error) {
	// Open the zip file
	file, err := os.Open(zipPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open zip file: %v", err)
	}

	defer file.Close()

	svc := db.createS3Client()

	// Prepare S3 upload input

	filename := filepath.Join(db.BasePath, filepath.Base(zipPath))
	
	input := &s3.PutObjectInput{
		Bucket: aws.String(db.S3Bucket),
		Key:    aws.String(filename),
		Body:   file,
		ACL:    aws.String("public-read"), // Set the file as public
	}

	// Upload to S3
	_, err = svc.PutObject(input)
	if err != nil {
		return nil, fmt.Errorf("failed to upload to S3: %v", err)
	}

	// Construct the full URL of the uploaded file
	url := fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", db.S3Bucket, db.S3Region, filename)

	return &S3Backup{
		url: url,
		key: filename,
	}, nil
}

func  (db *DatabaseBackup) DeleteFromS3(key string) error {
	svc := db.createS3Client()

	_, err := svc.DeleteObject(&s3.DeleteObjectInput{
		Bucket: aws.String(db.S3Bucket),
		Key:    aws.String(key),
	})

	return err
}