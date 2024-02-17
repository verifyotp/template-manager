package syncer

import "context"

type Syncer interface {
	Retrieve(ctx context.Context, version, id string) (map[string]any, error)
	Upsert(ctx context.Context, version, id string, data map[string]any) error
}
