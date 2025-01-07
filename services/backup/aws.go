package backup

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)


func (db *DatabaseBackup) createAwsSession() *session.Session {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(db.S3Region),
		Credentials: credentials.NewStaticCredentials(
			os.Getenv("AWS_ACCESS_KEY_ID"),
			os.Getenv("AWS_SECRET_ACCESS_KEY"),
			"", // session token (optional)
		),
	})

	if err != nil {
		log.Fatalf("failed to create AWS session: %v", err)
	}

	return sess
}

func (db *DatabaseBackup) createS3Client() *s3.S3 {
	sess := db.createAwsSession()
	return s3.New(sess)
}