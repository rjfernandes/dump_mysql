package backup

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

func (db *DatabaseBackup) CreateBackup() (string, error) {
	// Create a directory for backups
	backupDir := filepath.Join(os.TempDir(), "mysql_backups")
	err := os.MkdirAll(backupDir, os.ModePerm)
	if err != nil {
		return "", fmt.Errorf("failed to create backup directory: %v", err)
	}

	// Generate a unique filename with timestamp
	filename := fmt.Sprintf("%s_%s.sql", db.Database, time.Now().Format("2006-01-02_15-04-05"))
	backupPath := filepath.Join(backupDir, filename)

	// Construct mysqldump command
	cmd := exec.Command("mysqldump",
		fmt.Sprintf("-h%s", db.Host),
		fmt.Sprintf("-P%s", db.Port),
		fmt.Sprintf("-u%s", db.User),
		fmt.Sprintf("-p%s", db.Password),
		db.Database)

	// Create the output file
	outputFile, err := os.Create(backupPath)
	if err != nil {
		return "", fmt.Errorf("failed to create backup file: %v", err)
	}
	defer outputFile.Close()

	// Execute mysqldump and write to file
	cmd.Stdout = outputFile
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("mysqldump failed: %v", err)
	}

	return backupPath, nil
}