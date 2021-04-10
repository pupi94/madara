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

func (pc *ProductCache) Save(ctx context.Context, product *models.Product) {
	pc.redis.Set()
}

func (pc *ProductCache) BatchSave(ctx context.Context, product *models.Product) {

}

func (pc *ProductCache) Delete(ctx context.Context, storeId string) {

}

func (pc *ProductCache) Get(ctx context.Context, product *models.Product) {

}

func (pc *ProductCache) Select(ctx context.Context, productIds []uint) {

}

func (pc *ProductCache) BatchDelete(ctx context.Context, product *models.Product) {

}

func (pc *ProductCache) saveProduct(ctx context.Context, product *models.Product) error {
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

func (pc *ProductCache) cacheKey(id, storeId uint64) string {
	return fmt.Sprintf("%d:product:%d", id, storeId)
}
