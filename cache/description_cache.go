package cache

import (
	"context"
	"github.com/pupi94/madara/components/xtypes"
	"github.com/pupi94/madara/tools/redis"
)

type DescriptionCache struct {
	redis *redis.Client
}

func NewDescCache(redisClient *redis.Client) *DescriptionCache {
	return &DescriptionCache{redis: redisClient}
}

func (pc *DescriptionCache) Save(ctx context.Context, storeId, productId xtypes.Uuid, desc string) error {

}

func (pc *DescriptionCache) BatchSave(ctx context.Context, storeId xtypes.Uuid, data map[xtypes.Uuid]string) {

}
