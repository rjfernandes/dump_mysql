package backup

import (
	"log"
	"os"
)

func Backup() {
	// Load configuration from .env file
	backup, err := NewInstance()
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	// Create database backup
	backupPath, err := backup.CreateBackup()
	if err != nil {
		log.Fatalf("Backup creation failed: %v", err)
	}
	defer func() {
		if backupPath != "" {
			os.Remove(backupPath)
		}
	}()

	// Compress backup
	zipPath, err := backup.CompressBackup(backupPath)
	if err != nil {
		log.Fatalf("Compression failed: %v", err)
	}
	defer func() {
		if zipPath != "" {
			os.Remove(zipPath)
		}
	}()

	// Upload to S3
	s3backup, err := backup.UploadToS3(zipPath)
	if err != nil {
		log.Fatalf("S3 upload failed: %v", err)
	}

	// Optional: Clean up local files
	if err := backup.Cleanup(backupPath, zipPath); err != nil {
		log.Printf("Cleanup failed: %v", err)
	}

	// Manage last 5 backups
	if err := backup.ManageLastFiveBackups(s3backup); err != nil {
		log.Printf("Manage backups failed: %v", err)
	}

	if err := SendEmailWithURL(s3backup.url); err != nil {
		log.Printf("Send email failed: %v", err)
	}
}