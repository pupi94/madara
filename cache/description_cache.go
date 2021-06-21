package cache

import (
	"context"
	"fmt"
	"github.com/pupi94/madara/components/xtypes"
	"github.com/pupi94/madara/tools/redis"
)

type DescriptionCache struct {
	redis *redis.Client
}

func NewDescCache(redisClient *redis.Client) *DescriptionCache {
	return &DescriptionCache{redis: redisClient}
}

func (dc *DescriptionCache) Save(ctx context.Context, storeId, productId xtypes.Uuid, desc string) error {
	res := dc.redis.Set(dc.cacheKey(storeId, productId), desc)
	return res.Error()
}

func (dc *DescriptionCache) Delete(ctx context.Context, storeId xtypes.Uuid, productIds ...xtypes.Uuid) error {
	var keys []string
	for _, id := range productIds {
		keys = append(keys, dc.cacheKey(storeId, id))
	}
	res := dc.redis.Delete(keys...)
	return res.Error()
}

func (dc *DescriptionCache) BatchSave(ctx context.Context, storeId xtypes.Uuid, data map[xtypes.Uuid]string) error {
	var cacheData = make(map[string]interface{})

	for k, v := range data {
		cacheData[dc.cacheKey(storeId, k)] = v
	}
	res := dc.redis.MultiSet(cacheData)
	return res.Error()
}

func (dc *DescriptionCache) Get(ctx context.Context, storeId, productId xtypes.Uuid) (string, error) {
	res := dc.redis.Get(dc.cacheKey(storeId, productId))
	return res.String()
}

func (dc *DescriptionCache) cacheKey(storeId, id xtypes.Uuid) string {
	return fmt.Sprintf("%s:product_description:%s", storeId, id)
}
