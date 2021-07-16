package cache

import (
	"context"
	"fmt"
	"github.com/pupi94/madara/components/xtypes"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/models"
	"github.com/pupi94/madara/tools/redis"
	"github.com/pupi94/madara/tools/syncx"
	"gorm.io/gorm"
)

type InventoryCache struct {
	redis      *redis.Client
	sharedCall syncx.SharedCalls
	db         *gorm.DB
}

func NewInventoryCache(redisClient *redis.Client) *InventoryCache {
	return &InventoryCache{redis: redisClient, db: config.DB}
}

func (ic *InventoryCache) Save(ctx context.Context, storeId, sourceId xtypes.Uuid, inventory int64) error {
	res := ic.redis.Set(ic.cacheKey(storeId, sourceId), inventory)
	return res.Error()
}

func (ic *InventoryCache) BatchSave(ctx context.Context, storeId xtypes.Uuid, data map[xtypes.Uuid]int64) error {
	var cacheData = make(map[string]interface{})

	for k, v := range data {
		cacheData[ic.cacheKey(storeId, k)] = v
	}
	res := ic.redis.MultiSet(cacheData)
	return res.Error()
}

func (ic *InventoryCache) Delete(ctx context.Context, storeId xtypes.Uuid, sourceIds ...xtypes.Uuid) error {
	var keys []string
	for _, id := range sourceIds {
		keys = append(keys, ic.cacheKey(storeId, id))
	}
	res := ic.redis.Delete(keys...)
	return res.Error()
}

func (ic *InventoryCache) Get(ctx context.Context, storeId, sourceId xtypes.Uuid) (int64, error) {
	return ic.sharedGet(ctx, storeId, sourceId)
}

func (ic *InventoryCache) cacheKey(storeId, id xtypes.Uuid) string {
	return fmt.Sprintf("%s:inventory_quantity:%s", storeId, id)
}

func (ic *InventoryCache) sharedGet(ctx context.Context, storeId, sourceId xtypes.Uuid) (int64, error) {
	val, err := ic.sharedCall.Do(ic.cacheKey(storeId, sourceId), func() (interface{}, error) {
		return ic.doGet(ctx, storeId, sourceId)
	})
	return val.(int64), err
}

func (ic *InventoryCache) doGet(ctx context.Context, storeId, sourceId xtypes.Uuid) (int64, error) {
	cacheKey := ic.cacheKey(storeId, sourceId)
	resp := ic.redis.Get(cacheKey)

	if resp.Error() != nil {
		return 0, resp.Error()
	}

	if resp.Hit() {
		inventory, err := resp.Int64()
		if err != nil {
			return 0, err
		}
		return inventory, err
	}

	return ic.backfill(ctx, storeId, sourceId)
}

func (ic *InventoryCache) backfill(ctx context.Context, storeId, sourceId xtypes.Uuid) (int64, error) {
	inventory, err := models.GetProductInventory(ctx, ic.db, sourceId)
	if err != nil {
		return 0, err
	}
	if err := ic.Save(ctx, storeId, sourceId, inventory); err != nil {
		return 0, err
	}
	return inventory, nil
}
