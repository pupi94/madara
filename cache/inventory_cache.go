package cache

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pupi94/madara/component/redis"
	"github.com/pupi94/madara/component/syncx"
	"github.com/pupi94/madara/component/xtype"
	"github.com/pupi94/madara/model"
	"gorm.io/gorm"
)

type InventoryCache struct {
	redis      *redis.Client
	sharedCall syncx.SharedCalls
	db         *gorm.DB
}

func NewInventoryCache(redisClient *redis.Client, db *gorm.DB) *InventoryCache {
	return &InventoryCache{
		redis:      redisClient,
		db:         db,
		sharedCall: syncx.NewSharedCalls(),
	}
}

func (ic *InventoryCache) Save(ctx context.Context, storeId, sourceId xtype.Uuid, inventory int64) error {
	res := ic.redis.Set(ic.cacheKey(storeId, sourceId), inventory)
	return res.Error()
}

func (ic *InventoryCache) BatchSave(ctx context.Context, storeId xtype.Uuid, data map[xtype.Uuid]int64) error {
	var cacheData = make(map[string]interface{})

	for k, v := range data {
		cacheData[ic.cacheKey(storeId, k)] = v
	}
	res := ic.redis.MultiSet(cacheData)
	return res.Error()
}

func (ic *InventoryCache) Delete(ctx context.Context, storeId xtype.Uuid, sourceIds ...xtype.Uuid) error {
	var keys []string
	for _, id := range sourceIds {
		keys = append(keys, ic.cacheKey(storeId, id))
	}
	res := ic.redis.Delete(keys...)
	return res.Error()
}

func (ic *InventoryCache) Get(ctx context.Context, storeId, sourceId xtype.Uuid) (int64, error) {
	return ic.doSharedGet(ctx, storeId, sourceId)
}

func (ic *InventoryCache) MultiGet(ctx context.Context, storeId xtype.Uuid, sourceIds []xtype.Uuid) (map[xtype.Uuid]int64, error) {
	var (
		result = make(map[xtype.Uuid]int64)
		missed = make([]xtype.Uuid, 0)
	)

	for _, id := range sourceIds {
		resp := ic.redis.Get(ic.cacheKey(storeId, id))

		if resp.Error() != nil {
			return nil, resp.Error()
		}

		if resp.Hit() {
			inventory, err := resp.Int64()
			if err != nil {
				return nil, err
			}
			result[id] = inventory
		} else {
			missed = append(missed, id)
		}
	}

	missedResult, err := ic.batchBackfill(ctx, storeId, missed)
	if err != nil {
		return nil, err
	}

	for k, v := range missedResult {
		result[k] = v
	}
	return result, nil
}

func (ic *InventoryCache) cacheKey(storeId, id xtype.Uuid) string {
	return fmt.Sprintf("%s:inventory_quantity:%s", storeId, id)
}

func (ic *InventoryCache) doSharedGet(ctx context.Context, storeId, sourceId xtype.Uuid) (int64, error) {
	val, err := ic.sharedCall.Do(ic.cacheKey(storeId, sourceId), func() (interface{}, error) {
		return ic.doGet(ctx, storeId, sourceId)
	})
	return val.(int64), err
}

func (ic *InventoryCache) doGet(ctx context.Context, storeId, sourceId xtype.Uuid) (int64, error) {
	resp := ic.redis.Get(ic.cacheKey(storeId, sourceId))

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

func (ic *InventoryCache) backfill(ctx context.Context, storeId, sourceId xtype.Uuid) (int64, error) {
	inventory, err := model.GetInventory(ctx, ic.db, sourceId)
	if err != nil {
		return 0, err
	}
	if err := ic.Save(ctx, storeId, sourceId, inventory); err != nil {
		return 0, err
	}
	return inventory, nil
}

func (ic *InventoryCache) batchBackfill(ctx context.Context, storeId xtype.Uuid, sourceIds []xtype.Uuid) (map[xtype.Uuid]int64, error) {
	inventories, err := model.SelectInventories(ctx, ic.db, sourceIds)
	if err != nil {
		return nil, err
	}
	var result = make(map[xtype.Uuid]int64)
	if len(inventories) == 0 {
		return result, err
	}
	for _, inv := range inventories {
		result[inv.SourceID] = aws.Int64Value(inv.Value)
	}

	if err := ic.BatchSave(ctx, storeId, result); err != nil {
		return nil, err
	}
	return result, nil
}
