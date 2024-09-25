package cache

import (
	"context"
	"time"
)

// ICache interface working with cache implementation. Should be implemented as in a description.
// namespace is a group that aggregate keys.
// So you can clean whole namespace or a separate key.
// This represented as a map tree as assoc array
//
//	map[namespace][key] = value
//
// so you can remove specific key of namespace with Delete method which will delete only specific key at namespace
// or you can remove whole namespace.
// NOTICE: Developers are responsible to not cross namespace they using.
type ICache interface {
	Put(ctx context.Context, namespace string, key string, value interface{}, period time.Duration)
	PutList(ctx context.Context, namespace string, key string, value interface{}, total int64, period time.Duration)
	Load(ctx context.Context, namespace string, key string, dest interface{}) error
	LoadList(ctx context.Context, namespace string, key string, dest interface{}, total *int64) error
	ClearNamespace(ctx context.Context, namespace string)
	Delete(ctx context.Context, namespace string, key string)
}
