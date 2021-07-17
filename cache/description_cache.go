package cache

import (
	"context"
	"fmt"
	"github.com/pupi94/madara/component/redis"
	"github.com/pupi94/madara/component/xtype"
	"github.com/pupi94/madara/model"
	"gorm.io/gorm"
)

type DescriptionCache struct {
	redis *redis.Client
	db    *gorm.DB
}

func NewDescCache(redisClient *redis.Client, db *gorm.DB) *DescriptionCache {
	return &DescriptionCache{redis: redisClient, db: db}
}

func (dc *DescriptionCache) Save(ctx context.Context, storeId, productId xtype.Uuid, desc string) error {
	res := dc.redis.Set(dc.cacheKey(storeId, productId), desc)
	return res.Error()
}

func (dc *DescriptionCache) Delete(ctx context.Context, storeId xtype.Uuid, productIds ...xtype.Uuid) error {
	var keys []string
	for _, id := range productIds {
		keys = append(keys, dc.cacheKey(storeId, id))
	}
	res := dc.redis.Delete(keys...)
	return res.Error()
}

func (dc *DescriptionCache) BatchSave(ctx context.Context, storeId xtype.Uuid, data map[xtype.Uuid]string) error {
	var cacheData = make(map[string]interface{})

	for k, v := range data {
		cacheData[dc.cacheKey(storeId, k)] = v
	}
	res := dc.redis.MultiSet(cacheData)
	return res.Error()
}

func (dc *DescriptionCache) Get(ctx context.Context, storeId, productId xtype.Uuid) (string, error) {
	res := dc.redis.Get(dc.cacheKey(storeId, productId))
	return res.String()
}

func (dc *DescriptionCache) cacheKey(storeId, id xtype.Uuid) string {
	return fmt.Sprintf("%s:product_description:%s", storeId, id)
}

func (dc *DescriptionCache) doGet(ctx context.Context, storeId, id xtype.Uuid) (string, error) {
	resp := dc.redis.Get(dc.cacheKey(storeId, id))

	if resp.Error() != nil {
		return "", resp.Error()
	}

	if resp.Hit() {
		desc, err := resp.String()
		if err != nil {
			return "", err
		}
		return desc, err
	}

	return dc.backfill(ctx, storeId, id)
}

func (dc *DescriptionCache) backfill(ctx context.Context, storeId, productId xtype.Uuid) (string, error) {
	desc, err := model.GetProductDescription(ctx, dc.db, storeId, productId)
	if err != nil {
		return "", err
	}
	if err := dc.Save(ctx, storeId, productId, desc); err != nil {
		return "", err
	}
	return desc, nil
}
