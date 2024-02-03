package s3

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"template-manager/pkg/uploader"
)

type S3 struct {
	// contains filtered or unexported fields
	bucket string
	env    string
	region string
	folder string
}

/*
Please ensure that you have the following environment variables set:
- AWS_ACCESS_KEY_ID
- AWS_SECRET_ACCESS_KEY
- AWS_REGION
*/

type Option func(*S3)

func New(bucket string, options ...Option) *S3 {
	s3 := &S3{
		bucket: bucket,
		env:    "dev",
		region: os.Getenv("AWS_REGION"),
		folder: "template-manager",
	}
	for _, option := range options {
		option(s3)
	}
	return s3
}

func (s S3) UploadFile(ctx context.Context, filename string, buf *bytes.Buffer) (*uploader.UploadOutput, error) {
	sess := session.Must(session.NewSession(
		&aws.Config{
			Region: aws.String(s.region),
		},
	)) // The session the S3 Uploader will use

	key := fmt.Sprintf("%s/%s", s.folder, filename)
	s3Uploader := s3manager.NewUploader(sess)
	response, err := s3Uploader.UploadWithContext(ctx, &s3manager.UploadInput{ // Upload the file to S3.
		Bucket:          aws.String(s.bucket),
		ContentLanguage: aws.String("en"),
		ContentType:     aws.String("text/html"),
		Tagging:         aws.String(fmt.Sprintf("env=%s", s.env)),
		StorageClass:    aws.String("STANDARD"),
		Key:             aws.String(key),
		Body:            buf,
	})
	if err != nil {
		return nil, err
	}

	return &uploader.UploadOutput{
		URL:     response.Location,
		Version: aws.StringValue(response.VersionID),
		Key:     key,
	}, nil
}
