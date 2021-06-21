package cache

import (
	"context"
	"fmt"
	"github.com/pupi94/madara/components/xtypes"
	"github.com/pupi94/madara/tools/redis"
)

type InventoryCache struct {
	redis *redis.Client
}

func NewInventoryCache(redisClient *redis.Client) *InventoryCache {
	return &InventoryCache{redis: redisClient}
}

func (ic *InventoryCache) Save(ctx context.Context, storeId, sourceId xtypes.Uuid, inventory int64) error {
	res := ic.redis.Set(ic.cacheKey(storeId, sourceId), inventory)
	return res.Error
}

func (ic *InventoryCache) BatchSave(ctx context.Context, storeId xtypes.Uuid, data map[xtypes.Uuid]int64) error {
	var cacheData = make(map[string]interface{})

	for k, v := range data {
		cacheData[ic.cacheKey(storeId, k)] = v
	}
	res := ic.redis.MultiSet(cacheData)
	return res.Error
}

func (ic *InventoryCache) Delete(ctx context.Context, storeId xtypes.Uuid, sourceIds ...xtypes.Uuid) error {
	var keys []string
	for _, id := range sourceIds {
		keys = append(keys, ic.cacheKey(storeId, id))
	}
	res := ic.redis.Delete(keys...)
	return res.Error
}

func (ic *InventoryCache) cacheKey(storeId, id xtypes.Uuid) string {
	return fmt.Sprintf("%s:inventory_quantity:%s", storeId, id)
}
