package s3

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type Uploader interface {
	UploadWithContext(ctx aws.Context, input *s3manager.UploadInput, options ...func(*s3manager.Uploader)) (*s3manager.UploadOutput, error)
}

type S3Actions interface {
	PutObjectRequest(input *s3.PutObjectInput) (*request.Request, *s3.PutObjectOutput)
	HeadObjectWithContext(ctx aws.Context, input *s3.HeadObjectInput, options ...request.Option) (*s3.HeadObjectOutput, error)
	ListObjectVersionsWithContext(ctx aws.Context, input *s3.ListObjectVersionsInput, options ...request.Option) (*s3.ListObjectVersionsOutput, error)
}

type S3 struct {
	ENV         string
	Bucket      string
	Region      string
	ContentType string // e.g text/html, image/png, image/jpeg, application/pdf, video/mp4, etc
	Folder      string
	Session     client.ConfigProvider
	Uploader    Uploader
	SVC         S3Actions
}

func NewS3(Bucket string, Region string, ContentType string, Folder string) (*S3, error) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(Region),
	})
	if err != nil {
		return nil, err
	}
	return &S3{
		Bucket:      Bucket,
		Region:      Region,
		ContentType: ContentType,
		Folder:      Folder,
		Session:     sess,
		Uploader:    s3manager.NewUploader(sess),
		SVC:         s3.New(sess),
	}, nil
}

func (s S3) UploadFile(ctx context.Context, filename string, buf *bytes.Buffer) (*s3manager.UploadOutput, error) {
	// Upload the file to S3.
	return s.Uploader.UploadWithContext(ctx, &s3manager.UploadInput{
		Bucket:          aws.String(s.Bucket),
		ContentLanguage: aws.String("en"),
		ContentType:     aws.String(s.ContentType),
		Tagging:         aws.String(fmt.Sprintf("env=%s", s.ENV)),
		StorageClass:    aws.String("STANDARD"),
		Key:             aws.String(fmt.Sprintf("%s/%s", s.Folder, filename)),
		Body:            buf,
		ACL:             aws.String("public-read"),
	})
}

func (s S3) getFolder() string {
	return fmt.Sprintf("%s/%s", s.ENV, s.Folder)
}

func (s S3) UploadPresignedURL(ctx context.Context, filename string, expiry time.Duration) (string, error) {
	req, _ := s.SVC.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", s.getFolder(), filename)),
	})

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), expiry)
	defer cancel()

	// Presign the request using the context
	urlStr, _, err := req.PresignRequest(expiry)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func (s S3) GetFileDetails(ctx context.Context, filename string) (*s3.HeadObjectOutput, error) {
	return s.SVC.HeadObjectWithContext(ctx, &s3.HeadObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(fmt.Sprintf("%s/%s", s.getFolder(), filename)),
	})
}

func (s S3) VerifyFileExists(ctx context.Context, filename string) (bool, error) {
	_, err := s.GetFileDetails(ctx, filename)

	if err != nil {
		return false, err
	}

	return true, nil
}

func (s S3) GetPublicURl(filename string) string {
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s/%s", s.Bucket, s.Region, s.getFolder(), filename)
}

func (s S3) GetFileVersions(ctx context.Context, filename string) ([]*s3.ObjectVersion, error) {
	result, err := s.SVC.ListObjectVersionsWithContext(ctx, &s3.ListObjectVersionsInput{
		Bucket: aws.String(s.Bucket),
		Prefix: aws.String(fmt.Sprintf("%s/%s", s.getFolder(), filename)),
	})
	if err != nil {
		return nil, err
	}
	if len(result.Versions) == 0 {
		return nil, fmt.Errorf("no versions found for %s", filename)
	}

	return result.Versions, nil
}
