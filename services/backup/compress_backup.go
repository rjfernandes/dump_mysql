package backup

import (
	"fmt"
	"os/exec"
)

func (db *DatabaseBackup) CompressBackup(backupPath string) (string, error) {
	zipPath := backupPath + ".zip"
	cmd := exec.Command("zip", "-j", zipPath, backupPath) // Use -j to junk the path and place the file in the root
	
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("failed to compress backup: %v", err)
	}

	return zipPath, nil
}