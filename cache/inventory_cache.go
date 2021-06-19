package cache

import (
	"context"
	"github.com/pupi94/madara/components/xtypes"
	"github.com/pupi94/madara/tools/redis"
)

type InventoryCache struct {
	redis *redis.Client
}

func NewInventoryCache(redisClient *redis.Client) *InventoryCache {
	return &InventoryCache{redis: redisClient}
}

func (pc *InventoryCache) Save(ctx context.Context, storeId, sourceId xtypes.Uuid, inventory int64) error {
	return nil
}

func (pc *InventoryCache) BatchSave(ctx context.Context, storeId xtypes.Uuid, data map[xtypes.Uuid]int64) error {
	return nil
}
