package uploader

import (
	"bytes"
	"context"
)

// UploadOutput represents a response from the Upload() call.
type UploadOutput struct {
	URL     string
	Version string
	Key     string
}

type Uploader interface {
	UploadFile(ctx context.Context, filename string, buf *bytes.Buffer) (*UploadOutput, error)
}
