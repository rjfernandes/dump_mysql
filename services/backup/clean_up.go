package backup

import (
	"log"
	"os"
)

func (db *DatabaseBackup) Cleanup(backupPath, zipPath string) error {
	if err := os.Remove(backupPath); err != nil {
		log.Printf("Failed to remove SQL backup: %v", err)
	}
	if err := os.Remove(zipPath); err != nil {
		log.Printf("Failed to remove ZIP file: %v", err)
	}
	return nil
}
