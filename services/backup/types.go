package backup

type DatabaseBackup struct {
	Host     string
	Port     string
	User     string
	Password string
	Database string
	S3Bucket string
	S3Region string
	BasePath string
}

type S3Backup struct {
	url string
	key string
}
