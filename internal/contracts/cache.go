package contracts

import (
	"context"
	"time"
)
type CacheI interface{
	Set(ctx context.Context,key string ,value interface{},dur time.Duration)error
	Get(ctx context.Context,key string)(string ,error)
}