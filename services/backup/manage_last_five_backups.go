package backup

import (
	"database/sql"
	"fmt"

	_ "modernc.org/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

const dbPath = "backup_files.db"

type S3ObjectFile struct {
	Id int
	Key string
}

func (db *DatabaseBackup) ManageLastFiveBackups(backup *S3Backup) error {
	// Open SQLite database
	dbSqlite, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return fmt.Errorf("Failed to open SQLite database: %v", err)
	}
	defer dbSqlite.Close()

	// Create table if not exists
	_, err = dbSqlite.Exec(`CREATE TABLE IF NOT EXISTS backups (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		key TEXT NOT NULL,
		path TEXT NOT NULL,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	)`)

	if err != nil {
		return fmt.Errorf("Failed to create table: %v", err)
	}

	// Insert new backup record
	_, err = dbSqlite.Exec(`INSERT INTO backups (key, path) VALUES (?, ?)`, backup.key, backup.url)
	if err != nil {
		return fmt.Errorf("Failed to insert backup record: %v", err)
	}

	// Delete old backups if more than 5
	rows, err := dbSqlite.Query(`SELECT id, key FROM backups ORDER BY created_at DESC LIMIT -1 OFFSET 5`)
	if err != nil {
		return fmt.Errorf("Failed to query old backups: %v", err)
	}
	defer rows.Close()

	var objectsToDelete []S3ObjectFile
	for rows.Next() {
		var obj S3ObjectFile
		if err := rows.Scan(&obj.Id, &obj.Key); err != nil {
			return fmt.Errorf("Failed to scan row: %v", err)
		}
		objectsToDelete = append(objectsToDelete, obj)
	}

	for _, obj := range objectsToDelete {
		_, err := dbSqlite.Exec(`DELETE FROM backups WHERE id = ?`, obj.Id)
		if err != nil {
			return fmt.Errorf("Failed to delete old backup record: %v", err)
		}

		if err := db.DeleteFromS3(obj.Key); err != nil {
			return fmt.Errorf("Failed to delete old backup file: %v", err)
		}
	}
	
	return nil
}