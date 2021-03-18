package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/pupi94/madara/config"
	"github.com/pupi94/madara/models"
	"github.com/pupi94/madara/tools/redis"
)

type ProductCache struct {
	redis *redis.Client
}

func NewProductCache() *ProductCache {
	return &ProductCache{redis: config.RedisClient}
}

func (pc *ProductCache) SaveFullProduct(ctx context.Context, product *models.FullProduct) {

}

func (pc *ProductCache) BatchSaveFullProducts(ctx context.Context, product *models.FullProduct) {

}

func (pc *ProductCache) DeleteFullProduct(ctx context.Context, storeId string) {

}

func (pc *ProductCache) GetFullProduct(ctx context.Context, product *models.FullProduct) {

}

func (pc *ProductCache) ListFullProduct(ctx context.Context, product *models.FullProduct) {

}

func (pc *ProductCache) BatchDeleteFullProducts(ctx context.Context, product *models.FullProduct) {

}

func (pc *ProductCache) SaveProduct(ctx context.Context, product *models.Product) error {
	temp := product.Description
	defer func() {
		product.Description = temp
	}()

	product.Description = ""
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	_, err = pc.redis.Set(pc.cacheKey(product.StoreID, product.ID), data)
	return err
}

func (pc *ProductCache) cacheKey(id, storeId int64) string {
	return fmt.Sprintf("%d:product:%d", id, storeId)
}
