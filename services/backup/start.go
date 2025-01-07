package backup

import (
	"fmt"
	"os"
)

func NewInstance() (*DatabaseBackup, error) {
	// Validate required environment variables
	requiredVars := []string{
		"MYSQL_HOST", "MYSQL_PORT", "MYSQL_USER", "MYSQL_PASSWORD", "MYSQL_DATABASE",
		"AWS_S3_BUCKET", "AWS_REGION", 
		"AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY",
	}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return nil, fmt.Errorf("missing required environment variable: %s", v)
		}
	}

	// Create DatabaseBackup struct from environment variables

	basePath := os.Getenv("BASE_PATH")
	if basePath == "" {
		basePath = "backup_database"
	}

	return &DatabaseBackup{
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     os.Getenv("MYSQL_PORT"),
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
		S3Bucket: os.Getenv("AWS_S3_BUCKET"),
		S3Region: os.Getenv("AWS_REGION"),
		BasePath: basePath,
	}, nil
}