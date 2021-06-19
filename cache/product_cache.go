package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/pupi94/madara/components/xtypes"
	"github.com/pupi94/madara/models"
	"github.com/pupi94/madara/tools/redis"
	"golang.org/x/sync/errgroup"
)

type ProductCache struct {
	redis          *redis.Client
	descCache      *DescriptionCache
	inventoryCache *InventoryCache
}

func NewProductCache(redisClient *redis.Client) *ProductCache {
	return &ProductCache{
		redis:          redisClient,
		descCache:      NewDescCache(redisClient),
		inventoryCache: NewInventoryCache(redisClient),
	}
}

func (pc *ProductCache) Save(ctx context.Context, product *models.FullProduct) error {
	eg, ctx := errgroup.WithContext(ctx)

	eg.Go(func() error {
		if err := pc.saveMainInfo(ctx, &product.Product); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		if err := pc.descCache.Save(ctx, product.StoreID, product.ID, aws.StringValue(product.Description)); err != nil {
			return err
		}
		return nil
	})

	eg.Go(func() error {
		var inventoryDate = make(map[xtypes.Uuid]int64)
		inventoryDate[product.ID] = aws.Int64Value(product.InventoryQuantity)
		for _, v := range product.Variants {
			inventoryDate[v.ID] = aws.Int64Value(v.InventoryQuantity)
		}

		if err := pc.inventoryCache.BatchSave(ctx, product.StoreID, inventoryDate); err != nil {
			return err
		}
		return nil
	})
	return eg.Wait()
}

func (pc *ProductCache) BatchSave(ctx context.Context, product *models.Product) {

}

func (pc *ProductCache) Delete(ctx context.Context, storeId, productId xtypes.Uuid) error {
	pc.redis.Delete(pc.cacheKey(storeId, productId))
	return
}

func (pc *ProductCache) BatchDelete(ctx context.Context, product *models.Product) {

}

func (pc *ProductCache) Get(ctx context.Context, product *models.Product) {

}

func (pc *ProductCache) Select(ctx context.Context, productIds []uint) {

}

func (pc *ProductCache) saveMainInfo(ctx context.Context, product *models.Product) error {
	temp := product.Description
	defer func() {
		product.Description = temp
	}()

	product.Description = nil
	data, err := json.Marshal(product)
	if err != nil {
		return err
	}
	_, err = pc.redis.Set(pc.cacheKey(product.StoreID, product.ID), data)
	return err
}

func (pc *ProductCache) cacheKey(storeId, id xtypes.Uuid) string {
	return fmt.Sprintf("%s:product:%s", storeId, id)
}
